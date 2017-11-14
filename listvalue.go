package protoconv

import (
	"errors"

	structpb "github.com/golang/protobuf/ptypes/struct"
)

// List is a Value holding a list of Values
type List struct {
	values []*structpb.Value
}

// NewList creates a new value of type list
func NewList() *List {
	return &List{values: []*structpb.Value{}}
}

// Value returns the List as a Value
func (lv List) Value() *structpb.Value {
	return &structpb.Value{
		Kind: lv.ListValue(),
	}
}

// SetValue sets the value of the List
func (lv *List) SetValue(value *structpb.Value) error {
	list, ok := value.Kind.(*structpb.Value_ListValue)
	if !ok {
		return errors.New("value is not a list value")
	}

	lv.values = list.ListValue.Values
	return nil
}

// ListValue returns the List as a Value_ListValue
func (lv List) ListValue() *structpb.Value_ListValue {
	return &structpb.Value_ListValue{
		ListValue: &structpb.ListValue{Values: lv.values},
	}
}

// Traverse will execute fn for each Value contained in the List.
// If fn returns an error, the traverse is aborted and the error
// is returned.
func (lv List) Traverse(fn func(val *structpb.Value) error) error {
	for _, v := range lv.values {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// Len returns the number of Values stored in the list
func (lv List) Len() int {
	return len(lv.values)
}

// Append will append val to the List values
func (lv *List) Append(val *structpb.Value) {
	lv.values = append(lv.values, val)
}

// ListValue returns a new Value of type ListValue
func ListValue() *structpb.Value {
	return NewList().Value()
}

// TraverseListValue traverses a ListValue calling fn for each value within. If fn
// returns an error, the traverse is stopped and the errors is returned
func TraverseListValue(lv *structpb.Value, fn func(*structpb.Value) error) error {
	list, ok := lv.Kind.(*structpb.Value_ListValue)
	if !ok {
		return errors.New("lv is not a listvalue")
	}

	if list.ListValue == nil {
		return nil
	}

	for _, v := range list.ListValue.Values {
		if err := fn(v); err != nil {
			return err
		}
	}

	return nil
}

// AppendValue appends a value to the a list
func AppendValue(lv *structpb.Value, val *structpb.Value) error {
	if lv == nil {
		return errors.New("lv cannot be nil")
	}

	list, ok := lv.Kind.(*structpb.Value_ListValue)
	if !ok {
		return errors.New("cannot append to non-list value")
	}

	if list.ListValue == nil {
		list.ListValue = &structpb.ListValue{}
	}

	list.ListValue.Values = append(list.ListValue.Values, val)
	return nil
}
