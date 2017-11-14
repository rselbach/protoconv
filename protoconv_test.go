package protoconv_test

import (
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/robteix/protoconv"
)

func TestTraverseListValue(t *testing.T) {
	list := protoconv.NewList()
	for i := 0; i < 10; i++ {
		list.Append(protoconv.NewNumber().SetInt(i).Value())
	}

	count := 0
	sum := 0
	list.Traverse(func(val *structpb.Value) error {
		count++
		nn := protoconv.NewNumber()
		nn.SetValue(val)
		sum += nn.Int()

		return nil
	})
	if count != 10 {
		t.Errorf("TraverseListValue got count %d, want %d", count, 10)
	}

	if sum != 45 { // 0+1+2+...+9 = 45
		t.Errorf("TraverseListValue got sum %d, want %d", sum, 45)
	}
}

func TestNumber(t *testing.T) {
	n := protoconv.NewNumber()
	if n.SetInt32(32).Int32() != 32 {
		t.Errorf("SetInt32 failed: got %d, want %d", 32, 32)
	}
}

func TestString(t *testing.T) {
	s := protoconv.NewString()
	if s.String() != "" {
		t.Error("TestString: empty String should returns empty string")
	}
	s = protoconv.NewString().SetString("foobar")
	if res := s.String(); res != "foobar" {
		t.Errorf("TestString got %s, want foobar", res)
	}
}

func ExampleList_Traverse() {
	// populate a list with numbers
	list := protoconv.NewList()
	for i := 0; i < 10; i++ {
		list.Append(protoconv.IntVal(i))
	}

	// traverse the list and calculate the sum of all
	// numbers in it
	sum := 0
	err := list.Traverse(func(val *structpb.Value) error {
		// convert the value to an int
		n, err := protoconv.Int(val)
		if err != nil {
			// if there's an error, the traverse will stop here
			return err
		}

		sum += n

		return nil
	})

	if err != nil {
		panic(err)
	}
}
