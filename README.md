protoconv
=========


`protoconv` is a Go package providing helper functions to convert between
Go types and `struct.Value` and its related types.

I created this package to help me when I need to work with [ptypes/struct](https://godoc.org/github.com/golang/protobuf/ptypes/struct) values as they
often need quite a lot of boilerplate for simple operations.

## Documentation

You can find the full documentation on
[godoc](https://godoc.org/github.com/rselbach/protoconv). 

The purpose of this package is to turn something like this --

```
val := &structpb.Value{
    Kind: &structpb.Value_NumberValue{
        NumberValue: float64(12),
    },
}
```

into something like --

```
val := protoconv.IntVal(12)
```

And getting an `int` from a `Value` from this --

```
v, ok := val.Kind.(*structpb.Value_NumberValue)
if !ok {
    // do something
}
i := int(v.NumberValue)
```

To this --

```
i := protoconv.Int(val)
```

In short, helper functions to help converting from `struct.Value` to
standard Go types.

## Status

`protoconv` is ready for production use but it may lack functionality you need
as I add functions whenever I need them at work.

Feel free to contribute pull requests or file an issue if you need something
that isn't there yet.


## License

Copyright (c) 2017 Roberto Selbach Teixeira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.