package server

import (
	"fmt"
	"net"
)

// Client - описывает клиента, участника чата.
type Client struct {
	name       string
	connection *net.Conn
}

// NewClient - создаст новый объект Client и вернет указатель на него.
func NewClient(name string, conn *net.Conn) *Client {
	var client = Client{}

	client.name = name
	client.connection = conn

	return &client
}

// GetName - вернет имя клиента.
func (c *Client) GetName() string {
	return c.name
}

// GetConnection - вернет соединение с клиентом.
func (c *Client) GetConnection() *net.Conn {
	return c.connection
}

// SendMessage - отправит сообщение клиенту.
func (c *Client) SendMessage(msg string) {
	fmt.Fprintln(*c.connection, msg)
}
