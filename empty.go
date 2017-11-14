package protoconv

import (
	"github.com/golang/protobuf/ptypes/empty"
)

// Empty returns a new empty.Empty
func Empty() *empty.Empty {
	return &empty.Empty{}
}
