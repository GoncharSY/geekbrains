package server

import (
	"context"
	"go-backend-1/lesson-02/task-03/server/player"
	"go-backend-1/lesson-02/task-03/server/quest"
	"net"
	"sync"
	"time"
)

func New(host, port string) *Structure {
	if host == "" {
		host = DefaultHost
	}

	if port == "" {
		port = DefaultPort
	}

	var cont context.Context
	var canc context.CancelFunc

	cont = context.Background()
	cont, canc = context.WithCancel(cont)

	return &Structure{
		Context: cont,
		Cancel:  canc,

		Config:   net.ListenConfig{KeepAlive: time.Minute},
		Listener: nil,

		Host: host,
		Port: port,

		Players:  make(map[string]*player.Structure),
		Question: quest.New(),
		QuestMtx: sync.Mutex{},

		Broadcast: make(chan string, 10),
	}
}
