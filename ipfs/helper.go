package ipfs

import (
	"encoding/binary"
	"errors"
	"literatecarnival/logger"
	"literatecarnival/types"
)

func Distance(x types.NodeId, y types.NodeId) (int, error) {
	distanceBytes, err := xorId(x, y)
	if err != nil {
		logger.DefaultLogger.Fatalln(err)
		return -1, err
	}
	return int(binary.BigEndian.Uint64(distanceBytes)), nil
}

func xorId(x types.NodeId, y types.NodeId) ([]byte, error) {
	if len(x) != len(y) {
		return nil, errors.New("Length of the node id are not the same")
	}
	result := make([]byte, len(x))
	for i := range len(x) {
		xor := x[i] ^ y[i]
		result[i] = xor
	}
	return result, nil
}
