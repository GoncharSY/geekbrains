package server

import (
	"context"
	"fmt"
	"net"
	"sync"
)

// Структура данных для сервера.
type Structure struct {
	Context context.Context
	Cancel  context.CancelFunc

	Config       net.ListenConfig
	Listener     net.Listener
	ConnectGroup sync.WaitGroup

	Host string
	Port string
}

// Получить адрес на котором запускается сервер.
func (srv *Structure) GetAddress() string {
	return fmt.Sprintf("%s:%s", srv.Host, srv.Port)
}
