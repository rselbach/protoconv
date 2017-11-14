package protoconv

import (
	"errors"

	structpb "github.com/golang/protobuf/ptypes/struct"
)

// NewString returns a new Value of type StringValue
func NewString() *String {
	return new(String)
}

// String holds a string Value
type String string

// Value returns the String as a Value
func (s String) Value() *structpb.Value {
	return &structpb.Value{Kind: s.StringValue()}
}

// SetValue sets a new value
func (s *String) SetValue(value *structpb.Value) error {
	sv, ok := value.Kind.(*structpb.Value_StringValue)
	if !ok {
		return errors.New("value is not a string value")
	}
	*s = String(sv.StringValue)
	return nil
}

// StringVal creates a new Value with val
func StringVal(val string) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_StringValue{
			StringValue: val,
		},
	}
}

// IsString tests if the value is a StringValue
func IsString(value *structpb.Value) bool {
	_, ok := value.Kind.(*structpb.Value_StringValue)
	return ok
}

// AsString returns value as a string. If value does
// not hold a string, returns the empty string
func AsString(value *structpb.Value) string {
	s, ok := value.Kind.(*structpb.Value_StringValue)
	if !ok {
		return ""
	}

	return s.StringValue
}

// StringValue returns the String as a Value_StringValue
func (s String) StringValue() *structpb.Value_StringValue {
	return &structpb.Value_StringValue{StringValue: string(s)}
}

// String returns the value of Number as an int32
func (s String) String() string {
	return string(s)
}

// SetString sets the value of the Number to the string value
func (s *String) SetString(val string) *String {
	*s = String(val)
	return s
}
