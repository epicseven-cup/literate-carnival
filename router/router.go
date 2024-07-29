package router

import (
	"errors"
	"literatecarnival/types"
)

const (
	NODEID_SPACE_SIZE = 128
)

type Router struct {
	values   map[string]types.IPFSObject
	kBuckets []*types.Bucket
}

func NewRouter(size int) *Router {
	kBucketLength := NODEID_SPACE_SIZE / size
	kBuckets := make([]*types.Bucket, kBucketLength)
	for i := range kBucketLength {
		kBuckets[i] = types.NewBucket(size)
	}
	return &Router{
		values:   map[string]types.IPFSObject{},
		kBuckets: kBuckets,
	}
}

func (router *Router) GetKBuckets(index int) (*types.Bucket, error) {
	if index > len(router.kBuckets) {
		return nil, errors.New("The request k bucket index is out of bounce")
	} else {
		return router.kBuckets[index], nil
	}
}
