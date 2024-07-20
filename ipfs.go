package literatecarnival

type IPFSLink struct {
	Name string
	Hash Multihash
	Size int
}

type IPFSObject struct {
	links []IPFSLink
	data  []byte
}
