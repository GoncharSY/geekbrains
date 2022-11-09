package player

import "net"

func New(conn net.Conn) *Structure {
	var plr = &Structure{
		Name:       conn.RemoteAddr().String(),
		Connection: conn,
		Answering:  make(chan string, 10),
	}

	go plr.startAnswering()

	return plr
}
