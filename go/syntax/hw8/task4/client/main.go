package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Start - запустит новый клиент.
func Start(addr string) {
	var conn net.Conn
	var err error

	// Установим соединение.
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Копируем в консоль все, что приходит от сервера.
	go func() {
		io.Copy(os.Stdout, conn)
	}()

	// Отправляем на сервер все, что вводим в консоль.
	io.Copy(conn, os.Stdin)
	fmt.Printf("%s: exit", conn.LocalAddr())
}
