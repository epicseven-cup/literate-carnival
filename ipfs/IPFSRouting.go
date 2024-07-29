package ipfs

import (
	"literatecarnival/proto"
	"literatecarnival/types"
)

type IPFSRouting interface {
	// Ping Node
	Ping(node types.NodeId) (*types.NodeId, error)
	// Get a particular peer network address
	FindPeer(node types.NodeId) (*proto.NODE, error)
	// // stores a small metadata value in DHT
	// SetValue(key []byte, value []byte)
	// // retrieves small meta data value
	// GetValue(key []byte)
	// // announces this node can serve a large value
	// ProvideValue(key types.Multihash)
	// //Get a number of peers serving a large value
	// FindValuePeers(key types.Multihash, min int)
}
