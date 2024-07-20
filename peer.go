package literatecarnival

type PeerProtocol interface {
	open(nodeid NodeId, ledger Ledger)
	send_want_list(want_list []Multihash)
	send_block(block Block) bool
	close(final bool)
}
