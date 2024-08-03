package types

import "net"

type PacketHeader struct {
	address *net.UDPAddr
	size    int
}

func NewPacketHeader(address *net.UDPAddr, size int) PacketHeader {
	return PacketHeader{
		address: address,
		size:    size,
	}
}

type Packet struct {
	header PacketHeader
	data   []byte
}

func NewPacket(header PacketHeader, data []byte) Packet {
	return Packet{}
}
