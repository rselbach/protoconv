package protoconv

import (
	"errors"

	structpb "github.com/golang/protobuf/ptypes/struct"
)

// NewNumber returns a new Value of type NumberValue
func NewNumber() *Number {
	return &Number{}
}

// Number holds a numeric Value
type Number struct {
	structpb.Value_NumberValue
}

// Value returns the Number as a Value
func (n Number) Value() *structpb.Value {
	return &structpb.Value{
		Kind: &n.Value_NumberValue,
	}
}

// SetValue sets a new value
func (n *Number) SetValue(value *structpb.Value) error {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return errors.New("value is not a numeric value")
	}
	n.Value_NumberValue = *nv
	return nil
}

// Int32Val creates a new Value with val
func Int32Val(val int32) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(val),
		},
	}
}

// Int32 returns value as an int32
func Int32(value *structpb.Value) (int32, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return int32(nv.NumberValue), nil
}

// Int32 returns the value of Number as an int32
func (n Number) Int32() int32 {
	return int32(n.NumberValue)
}

// SetInt32 sets the value of the Number to the int32 value
func (n *Number) SetInt32(val int32) *Number {
	n.NumberValue = float64(val)
	return n
}

// Int64Val creates a new Value with val
func Int64Val(val int64) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(val),
		},
	}
}

// Int64 returns value as an int64
func Int64(value *structpb.Value) (int64, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return int64(nv.NumberValue), nil
}

// Int64 returns the value of Number as an int64
func (n Number) Int64() int64 {
	return int64(n.NumberValue)
}

// SetInt64 sets the value of the Number to the int64 value
func (n *Number) SetInt64(val int64) *Number {
	n.NumberValue = float64(val)
	return n
}

// IntVal creates a new Value with val
func IntVal(val int) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(val),
		},
	}
}

// Int returns value as an int
func Int(value *structpb.Value) (int, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return int(nv.NumberValue), nil
}

// Int returns the value of Number as an int
func (n Number) Int() int {
	return int(n.NumberValue)
}

// SetInt sets the value of the Number to the int value
func (n *Number) SetInt(val int) *Number {
	n.NumberValue = float64(val)
	return n
}

// Uint32Val creates a new Value with val
func Uint32Val(val uint32) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(val),
		},
	}
}

// Uint32 returns value as an uint32
func Uint32(value *structpb.Value) (uint32, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return uint32(nv.NumberValue), nil
}

// Uint32 returns the value of Number as an uint32
func (n Number) Uint32() uint32 {
	return uint32(n.NumberValue)
}

// SetUint32 sets the value of the Number to the uint32 value
func (n *Number) SetUint32(val uint32) *Number {
	n.NumberValue = float64(val)
	return n
}

// Uint64Val creates a new Value with val
func Uint64Val(val uint64) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(val),
		},
	}
}

// Uint64 returns value as an uint64
func Uint64(value *structpb.Value) (uint64, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return uint64(nv.NumberValue), nil
}

// Uint64 returns the value of Number as an uint64
func (n Number) Uint64() uint64 {
	return uint64(n.NumberValue)
}

// SetUint64 sets the value of the Number to the uint64 value
func (n *Number) SetUint64(val uint64) *Number {
	n.NumberValue = float64(val)
	return n
}

// UintVal creates a new Value with val
func UintVal(val uint) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(val),
		},
	}
}

// Uint returns value as an uint
func Uint(value *structpb.Value) (uint, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return uint(nv.NumberValue), nil
}

// Uint returns the value of Number as an uint
func (n Number) Uint() uint {
	return uint(n.NumberValue)
}

// SetUint sets the value of the Number to the uint value
func (n *Number) SetUint(val uint) *Number {
	n.NumberValue = float64(val)
	return n
}

// Float32Val creates a new Value with val
func Float32Val(val float32) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(val),
		},
	}
}

// Float32 returns value as an float32
func Float32(value *structpb.Value) (float32, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return float32(nv.NumberValue), nil
}

// Float32 returns the value of Number as an float32
func (n Number) Float32() float32 {
	return float32(n.NumberValue)
}

// SetFloat32 sets the value of the Number to the float32 value
func (n *Number) SetFloat32(val float32) *Number {
	n.NumberValue = float64(val)
	return n
}

// Float64Val creates a new Value with val
func Float64Val(val float64) *structpb.Value {
	return &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: val,
		},
	}
}

// Float64 returns value as an float64
func Float64(value *structpb.Value) (float64, error) {
	nv, ok := value.Kind.(*structpb.Value_NumberValue)
	if !ok {
		return 0, errors.New("value is not numeric")
	}
	return nv.NumberValue, nil
}

// Float64 returns the value of Number as an float64
func (n Number) Float64() float64 {
	return n.NumberValue
}

// SetFloat64 sets the value of the Number to the float64 value
func (n *Number) SetFloat64(val float64) *Number {
	n.NumberValue = val
	return n
}
