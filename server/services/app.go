package services

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"srv/db"
	"srv/structs"
	"time"

	"github.com/nats-io/nats.go"
)

type app struct {
	Server  http.Server
	ctx     chan os.Signal
	cfg     structs.Config
	subChan chan *nats.Msg
}

func CreateApp() *app {
	return &app{}
}

func CreateWorker(ctx chan os.Signal, sub chan *nats.Msg) {
	for {
		select {
		case <-ctx:
			break
		case msg := <-sub:
			DoJob(msg)
		}
	}
}

func (a *app) Init() error {
	a.ctx = make(chan os.Signal, 1)
    cfgApp, errors := ParseJsonConfig()
    if len(errors) != 0 {
        return fmt.Errorf("missing values in config %v", errors) 
    }
    a.cfg = cfgApp
    db.ConnectDB(a.cfg.Postgres)
	return nil
}

func (a app) Start(router http.Handler) {
	a.Server = http.Server{
		Addr:    a.cfg.Port,
		Handler: router,
	}
    
	for i := 1; i <= a.cfg.MaxWorkers; i++ {
		fmt.Println(i)
	}

	if err := a.Server.ListenAndServe(); err != nil {
		return
	}

}

func (a app) ShutDown() error {
	signal.Notify(a.ctx, os.Interrupt)
	<-a.ctx

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := a.Server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
