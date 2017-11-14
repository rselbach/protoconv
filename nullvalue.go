package protoconv

import structpb "github.com/golang/protobuf/ptypes/struct"

// Null holds a struct.Value of type NullValue
type Null struct {
	n structpb.Value_NullValue
}

// NewNull creates a new Null
func NewNull() *Null {
	return &Null{}
}

// Value returns Null as a struct.Value
func (e *Null) Value() *structpb.Value {
	return &structpb.Value{Kind: &e.n}
}

// SetValue sets the value of the nullvalue. This function
// does nothing.
func (e *Null) SetValue(value *structpb.Value) error {
	return nil
}

// NullVal returns a new Value of type NullValue
func NullVal() *structpb.Value {
	return NewNull().Value()
}

// IsNull tests whether value is a NullValue
func IsNull(value *structpb.Value) bool {
	_, ok := value.Kind.(*structpb.Value_NullValue)
	return ok
}
