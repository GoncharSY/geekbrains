package player

import (
	"bufio"
	"fmt"
	"net"
)

type Structure struct {
	Name       string
	Connection net.Conn
	Answering  chan string
}

// ============================================================================
// ============================================================================
// ============================================================================

func (plr *Structure) Send(msg string) {
	fmt.Fprint(plr.Connection, msg)
}

// ============================================================================
// ============================================================================
// ============================================================================

func (plr *Structure) startAnswering() {
	var msg = bufio.NewScanner(plr.Connection)

	defer close(plr.Answering)
	for msg.Scan() {
		plr.Answering <- msg.Text()
	}
}
