package router

import (
	"literatecarnival/types"
)

type Router struct {
	values map[string]types.IPFSObject
	bucket *types.Bucket
}

func NewRouter(size int) IPFSRouting {
	return &Router{
		values: map[string]types.IPFSObject{},
		bucket: types.NewBucket(size),
	}
}