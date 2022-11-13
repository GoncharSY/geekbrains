package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
)

// Структура чат-сервера.
type Structure struct {
	Context context.Context
	Cancel  context.CancelFunc

	Config   net.ListenConfig
	Listener net.Listener

	Host string
	Port string

	Clients     map[string]*Client
	ClientGroup sync.WaitGroup

	Entering  chan *Client
	Leaving   chan *Client
	Messaging chan string
}

//
//

// Проверить, наличие клиента под указанным именем.
func (srv *Structure) HasClient(name string) bool {
	_, ok := srv.Clients[name]
	return ok
}

//
//

// Добавить клиента в чат.
func (srv *Structure) AddClient(cli *Client) error {
	if srv.HasClient(cli.Name) {
		return errors.New("already exists")
	}

	srv.ClientGroup.Add(1)
	srv.Clients[cli.Name] = cli
	return nil
}

//
//

// Удалить клиента из чата.
func (srv *Structure) RemoveClient(cli *Client) {
	if srv.HasClient(cli.Name) {
		delete(srv.Clients, cli.Name)
		srv.ClientGroup.Done()
	}
}

//
//

// Получить адрес чат-сервера.
func (srv *Structure) GetAddress() string {
	return fmt.Sprintf("%s:%s", srv.Host, srv.Port)
}

//
//

// Отправить сообщение всем клиентам в чате.
func (srv *Structure) Announce(msg string) {
	for _, cli := range srv.Clients {
		go cli.Send(msg)
	}
}