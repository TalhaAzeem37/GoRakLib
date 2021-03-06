package protocol

const (
	BitFlagValid = 0x80
	BitFlagIsAck = 0x40
	BitFlagIsNak = 0x20

	BitFlagPacketPair     = 0x10
	BitFlagContinuousSend = 0x08
	BitFlagNeedsBAndAs    = 0x04
)

type Datagram struct {
	*Packet

	PacketPair     bool
	ContinuousSend bool
	NeedsBAndAs    bool

	SequenceNumber uint32

	packets *[]*EncapsulatedPacket
}

func NewDatagram() *Datagram {
	var datagram = &Datagram{NewPacket(0), false, false, false, 0, &[]*EncapsulatedPacket{}}
	datagram.ResetStream()
	return datagram
}

func (datagram *Datagram) GetPackets() *[]*EncapsulatedPacket {
	return datagram.packets
}

func (datagram *Datagram) Encode() {
	datagram.Buffer = []byte{}
	var flags = BitFlagValid
	if datagram.PacketPair {
		flags |= BitFlagPacketPair
	}
	if datagram.ContinuousSend {
		flags |= BitFlagContinuousSend
	}
	if datagram.NeedsBAndAs {
		flags |= BitFlagNeedsBAndAs
	}

	datagram.PutByte(byte(flags))
	datagram.PutLittleTriad(datagram.SequenceNumber)

	for _, packet := range *datagram.GetPackets() {
		packet.Encode()
		datagram.PutBytes(packet.Buffer)
	}
}

func (datagram *Datagram) Decode() {
	var flags = datagram.GetByte()
	datagram.PacketPair = (flags & BitFlagPacketPair) != 0
	datagram.ContinuousSend = (flags & BitFlagContinuousSend) != 0
	datagram.NeedsBAndAs = (flags & BitFlagNeedsBAndAs) != 0

	datagram.SequenceNumber = datagram.GetLittleTriad()

	for !datagram.Feof() {
		packet := NewEncapsulatedPacket()
		packet, err := packet.GetFromBinary(datagram)
		if err == nil {
			var packets = append(*datagram.packets, packet)
			datagram.packets = &packets
		}
	}
}

func (datagram *Datagram) GetLength() int {
	var length = 4
	for _, pk := range *datagram.GetPackets() {
		length += pk.GetLength()
	}
	return length
}

func (datagram *Datagram) AddPacket(packet *EncapsulatedPacket) {
	var packets = append(*datagram.packets, packet)
	datagram.packets = &packets
}
