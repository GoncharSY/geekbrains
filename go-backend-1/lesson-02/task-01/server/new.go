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
		host = "localhost"
	}

	if port == "" {
		port = "9000"
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

		Host: "localhost",
		Port: "9000",
	}
}
