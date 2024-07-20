package literatecarnival

type Node struct {
	NodeId    Multihash
	Multihash []byte
	PubKey    PublicKey
	PriKey    PrivateKey
}

func CreateNode() *Node {
	return nil
}
