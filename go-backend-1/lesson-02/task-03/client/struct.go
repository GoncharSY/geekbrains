package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Структура чат-клиента.
type Structure struct {
	Host string
	Port string

	Dialer     net.Dialer
	Connection net.Conn

	Sending   chan string
	Receiving chan string
	Stopping  chan struct{}
	Clossing  chan struct{}
}

//
//

// Получить адрес чат-сервера, к которому присоединяется клиент.
func (cli *Structure) GetAddress() string {
	return fmt.Sprintf("%s:%s", cli.Host, cli.Port)
}

//
//

// Остановить работу клиента.
func (cli *Structure) Stop() {
	cli.Stopping <- struct{}{}
}

//
//

// Начать работу клиента с чат-сервером.
func (cli *Structure) Start() error {
	var adr = cli.GetAddress()
	var err error

	cli.Connection, err = cli.Dialer.Dial("tcp", adr)

	if err == nil {
		go cli.StartSending()
		go cli.StartReceiving()
		go cli.StartProcessing()
	}

	return err
}

//
//

// Начать обработку процессов взамодействия.
// Клиент может:
//   - Получать сообщения;
//   - Отправлять сообщеия;
//   - Покинуть чат.
func (cli *Structure) StartProcessing() {
	defer close(cli.Stopping)

	for {
		select {
		case <-cli.Stopping:
			cli.Connection.Close()
			return
		case msg, opn := <-cli.Receiving:
			if opn {
				log.Println(msg)
			}
		case msg, opn := <-cli.Sending:
			if opn {
				fmt.Fprintln(cli.Connection, msg)
			}
		}
	}
}

//
//

// Начать процесс получения сообщений из чата.
func (cli *Structure) StartReceiving() {
	var con = cli.Connection
	var chn = cli.Receiving
	var inp = bufio.NewScanner(con)

	defer close(cli.Clossing)
	defer close(cli.Receiving)

	for inp.Scan() {
		chn <- inp.Text()
	}
}

//
//

// Начать процесс отпраления сообщений в чат.
func (cli *Structure) StartSending() {
	var inp = bufio.NewScanner(os.Stdin)

	defer close(cli.Sending)

	for inp.Scan() {
		cli.Sending <- inp.Text()
	}
}
