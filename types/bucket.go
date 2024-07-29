package types

import "literatecarnival/proto"

const (
	NODEID_SIZE = 128
)

type Bucket struct {
	buckets []*proto.NODE
	size    int
}

func NewBucket(size int) *Bucket {
	return &Bucket{
		buckets: make([]*proto.NODE, size),
		size:    size,
	}
}

func (bucket *Bucket) GetBuckets() []*proto.NODE {
	return bucket.buckets
}

func (bucket *Bucket) GetSize() int {
	return bucket.size
}
