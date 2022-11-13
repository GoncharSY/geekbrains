package server

import (
	"context"
	"net"
	"sync"
	"time"
)

// New создает новый сервер и возвращает указатель на него.
func New(host, port string) *Structure {
	if host == "" {
		host = DefaultHost
	}

	if port == "" {
		port = DefaultPort
	}

	var ctx context.Context
	var ccl context.CancelFunc

	ctx = context.Background()
	ctx, ccl = context.WithCancel(ctx)

	return &Structure{
		Context: ctx,
		Cancel:  ccl,
		Config:  net.ListenConfig{KeepAlive: time.Minute},

		Connections:  make(map[string]net.Conn),
		ConnectGroup: sync.WaitGroup{},

		Host: host,
		Port: port,

		Entering:  make(chan net.Conn, 10),
		Leaving:   make(chan string, 10),
		Ticking:   make(chan string, 10),
		Messaging: make(chan string, 10),
	}
}
