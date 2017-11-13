package server

import (
	"fmt"
	"strconv"
	"goraklib/protocol"
)

type Session struct {
	address string
	port int
	opened bool
	packets []protocol.IPacket
}

func NewSession(address string, port int) *Session {
	fmt.Println("Session created for ip: " + address + ":" + strconv.Itoa(port))
	return &Session{address: address, port: port, opened: false}
}

func (session *Session) Open() {
	session.opened = true
}

func (session *Session) Forward(packet protocol.IPacket) {
	packet.Decode()

	session.packets = append(session.packets, packet)
}

func (session *Session) IsStackEmpty() bool {
	return len(session.packets) == 0
}

func (session *Session) FetchFromStack() protocol.IPacket {
	var packet = session.packets[0]
	session.packets = session.packets[1:]
	return packet
}
