package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

// Запустить сервер.
func (srv *Structure) Start() error {
	if srv.Listener != nil {
		return errors.New("already started")
	}

	lis, err := srv.Config.Listen(srv.Context, "tcp", srv.GetAddress())

	if err != nil {
		return err
	} else {
		srv.Listener = lis
		go srv.StartAccepting()
	}

	return nil
}

// Начать прием новых соединений к серверу.
func (srv *Structure) StartAccepting() {
	var con net.Conn
	var err error

	for {
		con, err = srv.Listener.Accept()

		if err != nil {
			log.Println(err)
			return
		} else {
			log.Println("new connection with address:", con.RemoteAddr())
			go srv.StartConnect(con)
		}
	}
}

// Начать обработку соединения.
func (srv *Structure) StartConnect(con net.Conn) {
	srv.ConnectGroup.Add(1)
	defer srv.StopConnect(con)

	tck := time.NewTicker(time.Second)
	defer tck.Stop()

	for {
		select {
		case time := <-tck.C:
			if _, err := fmt.Fprintf(con, "time: %s\n", time); err != nil {
				return
			}
		case <-srv.Context.Done():
			return
		}
	}
}
