package server

import (
	"log"
	"net"
)

func (srv *Structure) Start() error {
	var ctx = srv.Context
	var ntw = "tcp"
	var adr = srv.GetAddress()
	var err error

	srv.Listener, err = srv.Config.Listen(ctx, ntw, adr)

	if err == nil {
		go srv.StartProcessing()
		go srv.StartAccepting()
	}

	return err
}

func (srv *Structure) StartProcessing() {
	for {
		select {
		case cli, ok := <-srv.Entering:
			if ok {
				msg := cli.Name + ": has arrived"
				srv.Announce(msg)
				srv.AddClient(cli)
				log.Println(msg)
			}
		case cli, ok := <-srv.Leaving:
			if ok {
				msg := cli.Name + ": has left"
				srv.RemoveClient(cli)
				srv.Announce(msg)
				log.Println(msg)
			}
		case msg, ok := <-srv.Messaging:
			if ok {
				srv.Announce(msg)
			}
		}
	}
}

func (srv *Structure) StartAccepting() {
	for {
		if con, err := srv.Listener.Accept(); err != nil {
			log.Println(err)
		} else {
			go srv.StartClient(con)
		}
	}
}

func (srv *Structure) StartClient(con net.Conn) {
	var cli = NewClient(con)
	var nme = cli.Name

	cli.Send("You are " + nme)
	srv.Entering <- cli

	for {
		select {
		case <-srv.Context.Done():
			cli.Stop()
		case msg, ok := <-cli.Messaging:
			if ok {
				srv.Messaging <- nme + ": " + msg
			} else {
				srv.Leaving <- cli
				return
			}
		}
	}
}
