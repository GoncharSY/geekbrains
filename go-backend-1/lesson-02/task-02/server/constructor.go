package server

import (
	"context"
	"net"
	"sync"
	"time"
)

// Создать новый чат-сервер и вернуть указатель на него.
func New(host, port string) *Structure {
	var ctx context.Context
	var ccl context.CancelFunc

	ctx, ccl = context.WithCancel(context.Background())

	return &Structure{
		Context: ctx,
		Cancel:  ccl,

		Config:   net.ListenConfig{KeepAlive: 20 * time.Second},
		Listener: nil,

		Host: host,
		Port: port,

		Clients:     map[string]*Client{},
		ClientGroup: sync.WaitGroup{},

		Entering:  make(chan *Client, 10),
		Leaving:   make(chan *Client, 10),
		Messaging: make(chan string, 10),
	}
}
