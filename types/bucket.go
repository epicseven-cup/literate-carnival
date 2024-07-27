package types

import "literatecarnival/proto"

const (
	NODEID_SIZE = 128
)

type Bucket struct {
	buckets [NODEID_SIZE]*proto.NODE
	size    int
}

func NewBucket(size int) *Bucket {
	return &Bucket{
		buckets: [NODEID_SIZE]*proto.NODE{},
		size:    size,
	}
}

func (bucket *Bucket) GetBuckets() [NODEID_SIZE]*proto.NODE {
	return bucket.buckets
}

func (bucket *Bucket) GetSize() int {
	return bucket.size
}
