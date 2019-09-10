package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"unicode"
)

// Chat - описывает текстовый чат.
type Chat struct {
	BroadcastChannel chan string
	IncomeChannel    chan net.Conn
	Clients          map[string]*Client
	History          []string
}

// NewChat - создаст новый чат и вернет указатель на него.
func NewChat() *Chat {
	var chat = Chat{
		BroadcastChannel: make(chan string, 10),
		IncomeChannel:    make(chan net.Conn, 10),
		Clients:          make(map[string]*Client),
		History:          []string{},
	}

	return &chat
}

// Start - запустит чат.
func (c *Chat) Start() {
	go func() {
		for {
			select {
			case conn := <-c.IncomeChannel:
				if err := c.AddClient(&conn); err != nil {
					fmt.Println("ERROR CONNECTION:", err)
				}
			case message := <-c.BroadcastChannel:
				for i := range c.Clients {
					go c.Clients[i].SendMessage(message)
				}

				if len(c.History) < 10 {
					c.History = append(c.History, message)
				} else {
					c.History = append(c.History[1:], message)
				}
			}
		}
	}()
}

// ================================================================================================
// GETTERS PART ===================================================================================
// ================================================================================================

// ================================================================================================
// CLIENT PART ====================================================================================
// ================================================================================================

// AddClient - добавит нового клиента в чат.
// И сообщит об этом в общий канал.
func (c *Chat) AddClient(conn *net.Conn) error {
	var name = <-c.askName(conn)
	var client *Client

	if name == "" {
		return fmt.Errorf("No name for new client")
	}

	// Новый клиент.
	client = NewClient(name, conn)
	c.Clients[client.GetName()] = client

	// Отправим историю последних сообщений.
	for i := range c.History {
		client.SendMessage(c.History[i])
	}
	c.BroadcastChannel <- fmt.Sprintf("%v joined the chat.", client.GetName())

	// Запустим программу, которая будет работать с новым клиентом.
	// Будет передавать сообщения от клиента в чат. Своеобразный агент.
	go func() {
		var connection = client.GetConnection()
		var scanner = bufio.NewScanner(*connection)

		for scanner.Scan() {
			c.BroadcastChannel <- fmt.Sprintf("%s: %s", client.GetName(), scanner.Text())
		}

		c.RemoveClient(client.GetName())
	}()

	return nil
}

// RemoveClient - удалит клиента из чата.
// И сообщит об этом в общий канал.
func (c *Chat) RemoveClient(name string) error {
	if client, ok := c.Clients[name]; ok {
		// Удалим из списка.
		delete(c.Clients, name)

		// Закроем соединение.
		(*client.GetConnection()).Close()

		// Сообщим об этом в общий канал.
		c.BroadcastChannel <- fmt.Sprintf("%s left the chat.", name)
		return nil
	}

	return fmt.Errorf("No client with name: %s", name)
}

// ================================================================================================
// PRIVATE PART ===================================================================================
// ================================================================================================

// askName - запросит у клиента его имя и вернет канал, по которому придет имя.
// Возвращаемый канал - не буферизированный.
func (c *Chat) askName(conn *net.Conn) <-chan string {
	var chName = make(chan string)

	go func() {
		var question = "Enter your name, please..."
		var input = bufio.NewScanner(*conn)
		var name string

		fmt.Fprintln(*conn, question)
		for input.Scan() {
			// Допустимы только буквы и цифры.
			name = strings.TrimFunc(input.Text(), func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})

			if len(name) == 0 {
				// пустая строка
				fmt.Fprintln(*conn, question)
			} else if _, ok := c.Clients[name]; ok {
				// существует уже
				fmt.Fprintf(*conn, "User with name %s already exists in the chat.\n", name)
				fmt.Fprintln(*conn, question)
			} else {
				chName <- name
				close(chName)
				break
			}
		}
	}()

	return chName
}
