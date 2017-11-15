package server

import (
	"math/rand"
)

type GoRakLibServer struct {
	serverName string
	serverId int64
	port uint16
	udp *UDPServer
	sessionManager *SessionManager
	sessionCount uint
	maxSessionCount uint
	motd string
	defaultGameMode string
	minecraftProtocol uint
	minecraftVersion string
}

func NewGoRakLibServer(serverName string, address string, port uint16) *GoRakLibServer {
	var server = GoRakLibServer{}
	server.serverName = serverName
	var udp = NewUDPServer(address, port)

	server.serverId = int64(rand.Int())

	server.udp = &udp
	server.port = port
	server.sessionManager = NewSessionManager(&server)

	go func() {
		for {
			var packet, ip, port, err = udp.ReadBuffer()
			if err != nil {
				continue
			}
			if !server.sessionManager.SessionExists(ip, port) {
				server.sessionManager.CreateSession(ip, port)
			}
			var session, _ = server.sessionManager.GetSession(ip, port)
			go session.Forward(packet)
		}
	}()

	return &server
}

func (server *GoRakLibServer) GetSessionManager() *SessionManager {
	return server.sessionManager
}

func (server *GoRakLibServer) GetServerName() string {
	return server.serverName
}

func (server *GoRakLibServer) GetServerId() int64 {
	return server.serverId
}

func (server *GoRakLibServer) GetUDP() *UDPServer {
	return server.udp
}

func (server *GoRakLibServer) GetPort() uint16 {
	return server.port
}

func (server *GoRakLibServer) SetServerName(name string) {
	server.serverName = name
}

func (server *GoRakLibServer) GetConnectedSessionCount() uint {
	return server.sessionCount
}

func (server *GoRakLibServer) GetMaxConnectedSessions() uint {
	return server.maxSessionCount
}

func (server *GoRakLibServer) SetConnectedSessionCount(count uint) {
	server.sessionCount = count
}

func (server *GoRakLibServer) SetMaxConnectedSessions(count uint) {
	server.maxSessionCount = count
}

func (server *GoRakLibServer) SetMotd(motd string) {
	server.motd = motd
}

func (server *GoRakLibServer) GetMotd() string {
	return server.motd
}

func (server *GoRakLibServer) GetDefaultGameMode() string {
	return server.defaultGameMode
}

func (server *GoRakLibServer) SetDefaultGameMode(gameMode string) {
	server.defaultGameMode = gameMode
}

func (server *GoRakLibServer) GetMinecraftProtocol() uint {
	return server.minecraftProtocol
}

func (server *GoRakLibServer) SetMinecraftProtocol(protocol uint) {
	server.minecraftProtocol = protocol
}

func (server *GoRakLibServer) GetMinecraftVersion() string {
	return server.minecraftVersion
}

func (server *GoRakLibServer) SetMinecraftVersion(version string) {
	server.minecraftVersion = version
}

func (server *GoRakLibServer) Tick() {
	server.sessionManager.Tick()
}

