/*
Package protoconv provides helper functions to convert between
protobuf values and standard Go types.

Working with protobuf Values often requires a fair amount of
boilerplate and this package hopes to simplify that.

For instance, if one wants to get an int from a struct.Value,
one would do something like this:

	// val is a struct.Value pointer
	v, ok := val.Kind.(*structpb.Value_NumberValue)
	if !ok {
		// deal with the type being wrong
	}
	i := int(v.NumberValue)

protobuf simplifies that to

    i := protoconv.Int(val)

Note that for simplicity sake, all conversion functions return
the empty value for the types if the provided Value is not of
the correct type. This mimicks the behaviour of protobuf's
own functions.

Creating a new value is also simplified. For instance, to
create a value from an int, one normally needs to do this:

	i := 42
	val := &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: float64(i),
		},
	}

And protoconv simplifies this to

    val := protoconv.IntVal(42)

For lists this is even more important. For instance, to create
list with a few integers, you'd need to do this

	listVal := &structpb.ListValue
	for i := 0; i < 2; i++ {
		v := &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(i),
			},
		}
		listVal.Values = append(listVal.Values, v)
	}
	val := &structpb.Value{
		Kind: &structpb.Value_ListValue{
			ListValue: listVal,
		},
	}

But with protoconv

    list := protoconv.NewList()
	for i := 0; i < 2; i++ {
		list.Append(protoconv.IntVal(i))
	}
	val := list.Value()

Iterating through a list of Value containing integers so
we can append them to a slice would be

	listVal := val.Kind.(*structpb.Value_ListValue)
	for _, v := range listVal.ListValue.Values {
		nv, ok := v.Kind.(*structpb.Value_NumberValue)
		if ok {
		    ilist = append(ilist, int(nv.NumberValue))
		}
	}

This becomes either this

    for _, v := range list.Values() {
		ilist = append(ilist, protoconv.Int(v))
	}

Or alternatively

    list.Traverse(func (v *structpb.Value) error {
		ilist = append(ilist, protoconv.Int(v))
		return nil
	})

*/
package protoconv

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
)

// ProtoValuer is an interface that provides values
type ProtoValuer interface {
	Value() *structpb.Value
	SetValue(*structpb.Value) error
}
