package ipfs

import (
	"literatecarnival/logger"
	"literatecarnival/types"
	"net"
)

type IPFSInstance struct {
	node   IPFSRouting
	socket *net.UDPConn
}

func NewIPFSInstance(kBucketSize int, address *net.UDPAddr) (*IPFSInstance, error) {
	ln, err := net.ListenUDP("udp", address)
	if err != nil {
		logger.DefaultLogger.Fatalln(err)
		return nil, err
	}
	return &IPFSInstance{
		node:   NewNode(kBucketSize),
		socket: ln,
	}, nil
}

func (ipfsInstance *IPFSInstance) Start() {
	// listends for reqeust
	for {
		data := make([]byte, BLOCK_SIZE)
		bytesRead, address, err := ipfsInstance.socket.ReadFromUDP(data)
		if err != nil {
			logger.DefaultLogger.Fatalln(err)
			continue
		}

		header := types.NewPacketHeader(address, bytesRead)
		packet := types.NewPacket(header, data)
		// Handel of the connection
	}
}
