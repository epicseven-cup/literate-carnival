package types

type Ledger struct {
	owner      NodeId
	partner    NodeId
	bytes_sent int
	bytes_recv int
	timestamp  Timestamp
}

type Peer struct {
	nodeid    NodeId
	ledger    Ledger
	last_seen Timestamp
	want_list []Multihash
}

type BitSwap struct {
	// Can not use NodeId directly as key since it is a slice of bytes
	ledgers map[string]Ledger
	// currently open connections to other nodes
	active map[string]Peer
	//checksums of blocks this node needs
	need_list []Multihash
	// checksums of blocks this node has
	have_list []Multihash
}
