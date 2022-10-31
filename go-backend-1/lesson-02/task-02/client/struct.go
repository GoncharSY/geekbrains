package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

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

func (cli *Structure) GetAddress() string {
	return fmt.Sprintf("%s:%s", cli.Host, cli.Port)
}

func (cli *Structure) Stop() error {
	cli.Connection.Close()
	cli.Stopping <- struct{}{}

	close(cli.Sending)
	close(cli.Receiving)
	close(cli.Stopping)

	return nil
}

func (cli *Structure) Start() error {
	var adr = cli.GetAddress()
	var err error

	cli.Connection, err = cli.Dialer.Dial("tcp", adr)

	if err == nil {
		go cli.StartProcessing()
		go cli.StartReceiving()
		go cli.StartSending()
	}

	return err
}

func (cli *Structure) StartProcessing() {
	for {
		select {
		case <-cli.Stopping:
			return
		case msg := <-cli.Receiving:
			log.Println(msg)
		case msg := <-cli.Sending:
			fmt.Fprintln(cli.Connection, msg)
		}
	}
}

func (cli *Structure) StartReceiving() {
	var con = cli.Connection
	var chn = cli.Receiving
	var inp = bufio.NewScanner(con)

	for inp.Scan() {
		chn <- inp.Text()
	}

	close(cli.Clossing)
}

func (cli *Structure) StartSending() {
	var inp = bufio.NewScanner(os.Stdin)

	for inp.Scan() {
		cli.Sending <- inp.Text()
	}
}
