package server

import (
	"log"
	"net"
)

// Запустить сервер.
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

//
//

// Начать обработку взаимодействия сервера с клиентами.
// Клиенты могут:
//   - Присоединяться к чату
//   - Покидать чать
//   - Отправлять сообщения в чат
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

//
//

// Начать процесс приема входящих соединений.
func (srv *Structure) StartAccepting() {
	for {
		if con, err := srv.Listener.Accept(); err != nil {
			log.Println(err)
		} else {
			go srv.StartClient(con)
		}
	}
}

//
//

// Начать работу с отдельным клиентом.
func (srv *Structure) StartClient(con net.Conn) {
	var cli = NewClient(con)
	srv.StartAuthorization(cli)
	srv.StartMessaging(cli)
}

//
//

// Начать процесс авторизации клиента в чате.
func (srv *Structure) StartAuthorization(cli *Client) {
	for {
		cli.Send("Enter your name")

		select {
		case <-srv.Context.Done():
			cli.Stop()
		case nme, opn := <-cli.Messaging:
			if !opn {
				return
			} else if nme == "" {
				continue
			} else if _, ok := srv.Clients[nme]; ok {
				cli.Send(nme + " already exists")
				continue
			} else {
				cli.Name = nme
				cli.Send("Hi " + nme + "!")
				srv.Entering <- cli
				return
			}
		}
	}
}

//
//

// Начать обмен сообщениями клиента с чатом.
func (srv *Structure) StartMessaging(cli *Client) {
	for {
		select {
		case <-srv.Context.Done():
			cli.Stop()
		case msg, ok := <-cli.Messaging:
			if ok {
				srv.Messaging <- cli.Name + ": " + msg
			} else {
				srv.Leaving <- cli
				return
			}
		}
	}
}
