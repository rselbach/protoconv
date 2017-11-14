package protoconv

import (
	"errors"

	structpb "github.com/golang/protobuf/ptypes/struct"
)

// NewBoolean returns a new Value of type BoolValue
func NewBoolean() *Boolean {
	return &Boolean{}
}

// Boolean holds a boolean Value
type Boolean struct {
	value bool
}

// Value returns the Boolean as a Value
func (b Boolean) Value() *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_BoolValue{
			BoolValue: b.value,
		},
	}
}

// SetValue sets a new value
func (b *Boolean) SetValue(value *structpb.Value) error {
	nv, ok := value.Kind.(*structpb.Value_BoolValue)
	if !ok {
		return errors.New("value is not a boolean value")
	}
	b.value = nv.BoolValue
	return nil
}

// BoolVal creates a new Value with val
func BoolVal(val bool) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_BoolValue{
			BoolValue: val,
		},
	}
}

// Bool returns value as an bool
func Bool(value *structpb.Value) (bool, error) {
	bv, ok := value.Kind.(*structpb.Value_BoolValue)
	if !ok {
		return false, errors.New("value is not boolean")
	}
	return bv.BoolValue, nil
}

// Bool returns the value of Number as an bool
func (b Boolean) Bool() bool {
	return b.value
}

// SetBool sets the value of the Boolean to the int32 value
func (b *Boolean) SetBool(val bool) *Boolean {
	b.value = val
	return b
}
