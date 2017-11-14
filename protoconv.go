package protoconv

import structpb "github.com/golang/protobuf/ptypes/struct"

// ProtoValuer is an interface that provides values
type ProtoValuer interface {
	Value() *structpb.Value
	SetValue(*structpb.Value) error
}
