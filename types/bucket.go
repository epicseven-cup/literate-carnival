package types

import "literatecarnival/proto"

type Bucket struct {
	buckets [NODEID_SIZE]*proto.NODE
	size    int
}

func (bucket *Bucket) GetBuckets() [NODEID_SIZE]*proto.NODE {
	return bucket.buckets
}

func (bucket *Bucket) GetSize() int {
	return bucket.size
}
