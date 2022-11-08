package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

// Структура данных для сервера.
type Structure struct {
	Context context.Context
	Cancel  context.CancelFunc

	Config   net.ListenConfig
	Listener net.Listener

	Connections  map[string]net.Conn
	ConnectGroup sync.WaitGroup

	Host string
	Port string

	Entering  chan net.Conn
	Leaving   chan string
	Ticking   chan string
	Messaging chan string
}

// Получить адрес на котором запускается сервер.
func (srv *Structure) GetAddress() string {
	return fmt.Sprintf("%s:%s", srv.Host, srv.Port)
}

// Определить существует ли сооединение под указанным именем (адресом).
func (srv *Structure) HasConnection(name string) bool {
	_, yes := srv.Connections[name]
	return yes
}

// Добавить новое соединение сервера с клиентом.
func (srv *Structure) AddConnection(conn net.Conn) error {
	var addr = conn.RemoteAddr().String()

	if srv.HasConnection(addr) {
		return errors.New("already exists")
	}

	srv.ConnectGroup.Add(1)
	srv.Connections[addr] = conn
	return nil
}

// Удалить соединение сервера с клиентом.
func (srv *Structure) RemoveConnection(name string) {
	if srv.HasConnection(name) {
		delete(srv.Connections, name)
		srv.ConnectGroup.Done()
	}
}

// Отправить сообщение всем клиентам, присоединившимся к серверу.
func (srv *Structure) Announce(mssg string) {
	for addr, conn := range srv.Connections {
		go func(addr string, conn net.Conn) {
			if _, err := fmt.Fprintf(conn, "server: %s", mssg); err != nil {
				log.Printf("error when sending to %v: %v\n", addr, err)
			}
		}(addr, conn)
	}
}
