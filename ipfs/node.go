package ipfs

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"literatecarnival/logger"
	"literatecarnival/pki"
	"literatecarnival/proto"
	"literatecarnival/router"
	"literatecarnival/types"
)

const (
	// The amount of preceding zero bits in the node id hash
	// The higher the value the harder to generate the node id , this provides sercuity for unknow node to join the network
	DIFFICULTY = 5
)

type Node struct {
	NodeId types.NodeId
	// Private and public key should be encrypted with passcode
	PubKey types.PublicKey
	PriKey types.PrivateKey
	// Routing table, should not be access from other code
	router *router.Router
	// Message channel
	message chan []byte
}

func count_preceding_zero(hash []byte) int {
	// Number of leading zero
	preceding_zero := 0
	for b := range hash {
		if b == 0 {
			preceding_zero += 1
		} else {
			break
		}
	}
	return preceding_zero
}

func NewNode(size int) IPFSRouting {
	var nodeId []byte
	var pubKey_bytes, privKey_bytes []byte
	for count_preceding_zero(nodeId) < DIFFICULTY {
		pubKey, privKey := pki.GenKeyPair()
		pubKey_bytes = pem.EncodeToMemory(
			&pem.Block{
				Type:    "RSA PRIVATE KEY",
				Headers: map[string]string{},
				Bytes:   x509.MarshalPKCS1PublicKey(pubKey),
			},
		)

		privKey_bytes = pem.EncodeToMemory(
			&pem.Block{
				Type:    "RSA PRIVATE KEY",
				Headers: map[string]string{},
				Bytes:   x509.MarshalPKCS1PrivateKey(privKey),
			},
		)
		hash := sha256.New()
		_, err := hash.Write(pubKey_bytes)
		if err != nil {
			logger.DefaultLogger.Fatalln("Error when tring to write public key hash into the hash", err)
		}
		nodeId = hash.Sum(nil)
	}
	node := Node{
		NodeId: nodeId,
		PubKey: pubKey_bytes,
		PriKey: privKey_bytes,
		router: router.NewRouter(size),
	}
	return &node
}

func SendPacket(packet types.Packet) {

}

func (node *Node) Ping(nodeId types.NodeId) (*types.NodeId, error) {
	return nil, nil
}

func (node *Node) CurrentDistance(nodeId types.NodeId) (int, error) {
	distanceBytes, err := xorId(node.NodeId, nodeId)
	if err != nil {
		logger.DefaultLogger.Fatalln(err)
		return -1, err

	}
	return int(binary.BigEndian.Uint64(distanceBytes)), nil
}

func Distance(x types.NodeId, y types.NodeId) (int, error) {
	distanceBytes, err := xorId(x, y)
	if err != nil {
		logger.DefaultLogger.Fatalln(err)
		return -1, err
	}
	return int(binary.BigEndian.Uint64(distanceBytes)), nil
}

func xorId(x types.NodeId, y types.NodeId) ([]byte, error) {
	if len(x) != len(y) {
		return nil, errors.New("Length of the node id are not the same")
	}
	result := make([]byte, len(x))
	for i := range len(x) {
		xor := x[i] ^ y[i]
		result[i] = xor
	}
	return result, nil
}

func (node *Node) FindPeer(nodeId types.NodeId) (*proto.NODE, error) {
	distance, err := node.CurrentDistance(nodeId)
	if err != nil {
		logger.DefaultLogger.Fatalln(err)
		return nil, err
	}
	kBuckets, err := node.router.GetKBuckets(distance)
	if err != nil {
		logger.DefaultLogger.Fatalln(err)
		return nil, err
	}

	buckets := kBuckets.GetBuckets()
	// Going through the buckets to see if one matchs the idea if not go to the closest one
	for i := range buckets {
		distance, err := Distance(buckets[i].GetNodeId(), nodeId)
		if err != nil {
			logger.DefaultLogger.Fatalln(err)
			return nil, err
		}

		if distance == 0 {
			return buckets[i], nil
		} else {
			// Send UDP request to other nodes and wait on the channel to hear back
		}

	}
	return &proto.NODE{}, nil
}
