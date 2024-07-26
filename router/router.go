package router

import (
	"literatecarnival/types"
)

const (
	NODEID_SIZE = 128
)

type Router struct {
	values map[string]types.IPFSObject
	bucket types.Bucket
}
