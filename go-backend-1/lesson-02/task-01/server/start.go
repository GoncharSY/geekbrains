package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
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
		go srv.StartProcessing()
		go srv.StartMessaging()
		go srv.StartTicking()
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
		}

		go srv.StartConnect(con)
	}
}

// Начать обработку соединений.
func (srv *Structure) StartProcessing() {
	for {
		select {
		case conn := <-srv.Entering:
			if err := srv.AddConnection(conn); err != nil {
				log.Println("entering error:", err)
			} else {
				log.Println("joined:", conn.RemoteAddr())
			}
		case addr := <-srv.Leaving:
			srv.RemoveConnection(addr)
			log.Println("disconnected:", addr)
		case time := <-srv.Ticking:
			srv.Announce(time)
		case mssg := <-srv.Messaging:
			srv.Announce(mssg)
		case <-srv.Context.Done():
			return
		}
	}
}

// Начать тиканье сервера.
func (srv *Structure) StartTicking() {
	var tick = time.NewTicker(time.Second)

	defer tick.Stop()

	for {
		select {
		case time := <-tick.C:
			srv.Ticking <- fmt.Sprintf("time: %v\n", time)
		case <-srv.Context.Done():
			return
		}
	}
}

// Начать пересылку сообщений.
func (srv *Structure) StartMessaging() {
	var inpt = bufio.NewScanner(os.Stdin)

	for inpt.Scan() {
		select {
		case <-srv.Context.Done():
			return
		default:
			srv.Messaging <- inpt.Text() + "\n"
		}
	}
}

// Начать обработку соединения.
func (srv *Structure) StartConnect(con net.Conn) {
	srv.Entering <- con

	defer func() {
		con.Close()
		srv.Leaving <- con.RemoteAddr().String()
	}()

	for bufio.NewScanner(con).Scan() {
		select {
		case <-srv.Context.Done():
			return
		default:
			time.Sleep(time.Second)
		}
	}
}
