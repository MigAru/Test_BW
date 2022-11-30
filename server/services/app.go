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
)

type app struct {
	Server  http.Server
	cfg     structs.Config
}

func CreateApp() *app {
	return &app{}
}

func (a *app) Init() error {
    cfgApp, errors := ParseJsonConfig()
    if len(errors) != 0 {
        return fmt.Errorf("missing values in config %v", errors) 
    }
    a.cfg = cfgApp
    if err := db.ConnectDB(a.cfg.Postgres); err != nil {
        return err
    }
	return nil
}

func (a app) Start(router http.Handler) {
	a.Server = http.Server{
		Addr:    a.cfg.Port,
		Handler: router,
	}
    
	if err := a.Server.ListenAndServe(); err != nil {
		return
	}

}

func (a app) ShutDown() error {
	var quit = make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := a.Server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
