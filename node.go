package literatecarnival

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"literatecarnival/logger"
	"literatecarnival/pki"
)

const (
	// The amount of preceding zero bits in the node id hash
	// The higher the value the harder to generate the node id , this provides sercuity for unknow node to join the network
	DIFFICULTY = 5
)

type Node struct {
	NodeId Multihash
	// Private and public key should be encrypted with passcode
	PubKey PublicKey
	PriKey PrivateKey
	//Routing table, should not be access from other code
	router IPFSRouting
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

func NewNode() *Node {
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
	node := Node{NodeId: nodeId, PubKey: pubKey_bytes, PriKey: privKey_bytes}
	return &node
}
