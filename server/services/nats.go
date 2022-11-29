package services

import (
	"errors"
	"strings"

	"github.com/nats-io/nats.go"
)

var (
	natsConnect *nats.Conn
	mapQueue    = make(map[string]*nats.Subscription)
)

func ConnectNats(urls []string, username, password *string) error {
	if len(urls) == 0 {
		return errors.New("urls doesn't exist")
	}

	url := urls[0]
	if len(urls) > 1 {
		url = strings.Join(urls, ",")
	}

	if username == nil || password == nil {
		conn, err := nats.Connect(url)
		if err != nil {
			return err
		}
		natsConnect = conn
		return nil

	}

	userInfo := nats.UserInfo(*username, *password)
	conn, err := nats.Connect(url, userInfo)
	if err != nil {
		return err
	}
	natsConnect = conn
	return nil
}

func Subscribe(channel, queue string, chanSub chan *nats.Msg) error {
	sub, err := natsConnect.ChanQueueSubscribe(channel, queue, chanSub)
	if err != nil {
		return err
	}
	mapQueue[channel+"/"+queue] = sub
	return nil
}

func Unsubscribe(channel, queue string) error {
	sub, ok := mapQueue[channel+"/"+queue]
	if !ok {
		return errors.New("channel doesn't exist")
	}
	if err := sub.Unsubscribe(); err != nil {
		return err
	}
	return nil
}
