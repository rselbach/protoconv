package protoconv

import (
	"errors"

	structpb "github.com/golang/protobuf/ptypes/struct"
)

// Struct holds a struct.Value of type StructValue
type Struct struct {
	value *structpb.Struct
}

// NewStruct creates a new Struct
func NewStruct() *Struct {
	return &Struct{
		value: &structpb.Struct{
			Fields: make(map[string]*structpb.Value),
		},
	}
}

// Value return Struct as a Value
func (s *Struct) Value() *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_StructValue{
			StructValue: s.value,
		},
	}
}

// StructVal returns a new empty Value of type StructValue
func StructVal() *structpb.Value {
	return NewStruct().Value()
}

// IsStruct tests whether value is a StructValue
func IsStruct(value *structpb.Value) bool {
	_, ok := value.Kind.(*structpb.Value_StructValue)
	return ok
}

// SetValue sets the value of the Struct
func (s *Struct) SetValue(value *structpb.Value) error {
	sv, ok := value.Kind.(*structpb.Value_StructValue)
	if !ok {
		return errors.New("value is not a struct")
	}

	s.value = sv.StructValue
	return nil
}

// Set sets the value of key
func (s *Struct) Set(key string, value *structpb.Value) {
	s.value.Fields[key] = value
}

// Get returns the value of key. If key is not found,
// the returned value is a NullValue
func (s *Struct) Get(key string) *structpb.Value {
	v, ok := s.value.Fields[key]
	if !ok {
		return NullVal()
	}

	return v
}

// Remove removes the item indexed by key from the
// Struct. If key does not exists, this does nothing
func (s *Struct) Remove(key string) {
	delete(s.value.Fields, key)
}
