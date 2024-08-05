package ipfs

// Remove this to node

// type IPFSInstance struct {
// 	node   IPFSRouting
// 	socket *net.UDPConn
// }

// func NewIPFSInstance(kBucketSize int, address *net.UDPAddr) (*IPFSInstance, error) {
// 	ln, err := net.ListenUDP("udp", address)
// 	if err != nil {
// 		logger.DefaultLogger.Fatalln(err)
// 		return nil, err
// 	}
// 	return &IPFSInstance{
// 		node:   NewNode(kBucketSize),
// 		socket: ln,
// 	}, nil
// }

// func (ipfsInstance *IPFSInstance) Start() {
// 	// listends for reqeust
// 	for {
// 		data := make([]byte, BLOCK_SIZE)
// 		bytesRead, address, err := ipfsInstance.socket.ReadFromUDP(data)
// 		if err != nil {
// 			logger.DefaultLogger.Fatalln(err)
// 			continue
// 		}
// 		logger.DefaultLogger.Println("New Message arrived, from %s. Size: %d", address.String(), bytesRead)
// 		// Handel of the connection
// 		packet := proto.Packet{}
// 		protobuf.Unmarshal(data, &packet)

// 		ipfsInstance.node.Incoming <- packet
// 	}
// }
