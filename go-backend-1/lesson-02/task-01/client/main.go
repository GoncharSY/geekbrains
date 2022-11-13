package client

import (
	"bufio"
	"context"
	"log"
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

	if con, err = dlr.DialContext(ctx, "tcp", addr); err != nil {
		return err
	} else {
		defer con.Close()
	}

	var dat = bufio.NewScanner(con)

	for dat.Scan() {
		select {
		case <-ctx.Done():
			return nil
		default:
			log.Print(dat.Text())
		}
	}

	return nil
}
