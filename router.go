package literatecarnival

type IPFSRouting interface {
	// Get a particular peer network address
	FindPeer(node NodeID)
	// stores a small metadata value in DHT
	SetValue(key []byte, value []byte)
	// retrieves small meta data value
	GetValue(key []byte)
	// announces this node can serve a large value
	ProvideValue(key Multihash)
	//Get a number of peers serving a large value
	FindValuePeers(key Multihash, min int)
}
