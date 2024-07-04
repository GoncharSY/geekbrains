package server

import (
	"bufio"
	"fmt"
	"net"
)

// Создать нового клиента и вернуть указатель на него.
func NewClient(conn net.Conn) *Client {
	var cli = Client{
		Name:       conn.RemoteAddr().String(),
		Connection: conn,
		Messaging:  make(chan string),
	}

	go cli.Start()
	return &cli
}

//
//

// Структура клиента в чате.
type Client struct {
	Name       string
	Connection net.Conn
	Messaging  chan string
}

//
//

// Отправить сообщение клиенту.
func (cli *Client) Send(msg string) {
	fmt.Fprintln(cli.Connection, msg)
}

//
//

// Остановить работу с клиентом.
// При этом соедиеннение будет закрыто.
func (cli *Client) Stop() {
	cli.Connection.Close()
}

//
//

// Начать работу с клиентом.
// При этом активизируется канал обмена сообщений с клиентом.
func (cli *Client) Start() {
	var con = cli.Connection
	var inp = bufio.NewScanner(con)
	var chn = cli.Messaging

	for inp.Scan() {
		chn <- inp.Text()
	}

	close(chn)
}
