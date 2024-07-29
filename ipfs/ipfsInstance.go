package ipfs

import (
	"literatecarnival/logger"
	"net"
)

type IPFSInstance struct {
	node   IPFSRouting
	socket net.Listener
}

func NewIPFSInstance(kBucketSize int, address string) (*IPFSInstance, error) {
	ln, err := net.Listen("udp", address)
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
		conn, err := ipfsInstance.socket.Accept()
		if err != nil {
			logger.DefaultLogger.Fatalln(err)
			continue
		}

		// Handel of the connection

	}
}

func (ipfsInstance *IPFSInstance) handler(conn net.Conn) {

}
