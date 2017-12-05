package protocol

import (
	"goraklib/protocol/identifiers"
	"strconv"
)

type UnconnectedPong struct {
	*UnconnectedMessage
	PingTime int64
	ServerId int64
	ServerName string
	ServerProtocol uint
	ServerVersion string
	OnlinePlayers uint
	MaximumPlayers uint
	Motd string
	DefaultGameMode string

	// Raw data from pong.
	ServerData string
}

func NewUnconnectedPong() *UnconnectedPong {
	return &UnconnectedPong{NewUnconnectedMessage(NewPacket(
		identifiers.UnconnectedPong,
	)), 0, 0, "", 0, "", 0, 20, "", "", ""}
}

func (pong *UnconnectedPong) Encode() {
	pong.EncodeId()
	pong.PutLong(pong.PingTime)
	pong.PutLong(pong.ServerId)
	pong.PutMagic()
	pong.PutString(
		"MCPE;" +
		pong.ServerName + ";" +
		strconv.Itoa(int(pong.ServerProtocol)) + ";" +
		pong.ServerVersion + ";" +
		strconv.Itoa(int(pong.OnlinePlayers)) + ";" +
		strconv.Itoa(int(pong.MaximumPlayers)) + ";" +
		strconv.Itoa(int(pong.ServerId)) + ";" +
		pong.Motd + ";" +
		pong.DefaultGameMode + ";",
	)
}

func (pong *UnconnectedPong) Decode() {
	pong.DecodeStep()
	pong.PingTime = pong.GetLong()
	pong.ServerId = pong.GetLong()
	pong.ReadMagic()
	pong.ServerData = pong.GetString()
}
