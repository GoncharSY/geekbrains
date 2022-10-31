package server

import (
	"bufio"
	"fmt"
	"net"
)

func NewClient(conn net.Conn) *Client {
	var cli = Client{
		Name:       conn.RemoteAddr().String(),
		Connection: conn,
		Messaging:  make(chan string),
	}

	go cli.Start()
	return &cli
}

type Client struct {
	Name       string
	Connection net.Conn
	Messaging  chan string
}

func (cli *Client) Send(msg string) {
	fmt.Fprintln(cli.Connection, msg)
}

func (cli *Client) Stop() {
	cli.Connection.Close()
}

func (cli *Client) Start() {
	var con = cli.Connection
	var inp = bufio.NewScanner(con)
	var chn = cli.Messaging

	for inp.Scan() {
		chn <- inp.Text()
	}

	close(chn)
}
