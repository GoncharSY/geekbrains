package client

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

// Запустить клиент.
func Start(addr string, ctx context.Context) error {
	var dlr = net.Dialer{
		KeepAlive: 10 * time.Second,
		Timeout:   time.Second,
	}

	var con net.Conn
	var err error
	var dat = make(chan string, 3)

	if con, err = dlr.DialContext(ctx, "tcp", addr); err != nil {
		return err
	} else {
		defer con.Close()
		go StartReading(con, dat)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case str, ok := <-dat:
			if ok {
				fmt.Print(str)
			} else {
				return nil
			}
		}
	}
}

// Начать читать данные из сетевого соединения.
func StartReading(con net.Conn, dat chan<- string) {
	var buf = make([]byte, 256)

	defer close(dat)

	for {
		if _, err := con.Read(buf); err == io.EOF {
			break
		} else {
			dat <- string(buf)
		}
	}
}
