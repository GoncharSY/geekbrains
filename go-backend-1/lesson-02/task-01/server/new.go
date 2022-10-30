package server

import (
	"context"
	"net"
	"sync"
	"time"
)

// New создает новый сервер и возвращает указатель на него.
func New(host, port string, ctx context.Context) *Structure {
	if host == "" {
		host = DefaultHost
	}

	if port == "" {
		port = DefaultPort
	}

	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithCancel(ctx)

	return &Structure{
		Context:      ctx,
		Cancel:       cancel,
		Config:       net.ListenConfig{KeepAlive: time.Minute},
		ConnectGroup: sync.WaitGroup{},

		Host: host,
		Port: port,
	}
}
