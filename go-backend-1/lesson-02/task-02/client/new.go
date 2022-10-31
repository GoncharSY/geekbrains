package client

import (
	"net"
	"time"
)

func New(host, port string) *Structure {
	return &Structure{
		Host: host,
		Port: port,

		Dialer: net.Dialer{
			KeepAlive: 20 * time.Second,
			Timeout:   time.Second,
		},

		Sending:   make(chan string),
		Receiving: make(chan string, 10),
		Stopping:  make(chan struct{}),
		Clossing:  make(chan struct{}),
	}
}
