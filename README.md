## Introduction

`go-multiparse` allows you to easily convert one type of value into a built-in type (e.g. strings from query values in web requests to integers or floats). The original purpose of this project is to reduce the repetition in parsing from string values in web-requests though it is not limited to just parsing strings. Parse errors will simply panic.


## Example

```go
// import "github.com/dsoprea/go-multiparse"

valueInterface := parse.Parse("123.456", "float64")
value := valueInterface.(float64)
```


## Implementation

This is a general parsing framework. Parser implementations must satisfy the `Parser` interface and be registered for specific input types. You will have to register multiple times if you are trying to parse both a base type *and* one or more aliases of it. The `StringParser` is included and parses over string values. It is automatically registered.


### Parser Interface

This is the interface:

```go
type Parser interface {
    String(value interface{}) string

    Int8(value interface{}) int8
    Int16(value interface{}) int16
    Int32(value interface{}) int32
    Int64(value interface{}) int64

    Uint8(value interface{}) uint8
    Uint16(value interface{}) uint16
    Uint32(value interface{}) uint32
    Uint64(value interface{}) uint64

    Hex8(value interface{}) uint8
    Hex16(value interface{}) uint16
    Hex32(value interface{}) uint32
    Hex64(value interface{}) uint64

    Float32(value interface{}) float32
    Float64(value interface{}) float64

    Bool(value interface{}) bool
}
```

To register a new implementation, use `AddParser()`. An example based on how we automatically register `StringParser`:

```go
type := reflect.TypeOf("")
parse.AddParser(type, parserInstance)
```

Of course, it would not make any sense to register another string parser. However, if you have a more exotic type that you often need to convert, implement the interface methods that you require, panic on the others, and use `AddParser()` to register it.


### Alternate Usage

The `Parse()` function is a convenience wrapper (which uses a dictionary to find the method name). You may acquire and use the `Parser` instance directly:

```go
// import "reflect"

vRaw := "123.456"
t := reflect.TypeOf(vRaw)
p := parse.GetParser(t)
v := p.Float64(vRaw)
```

Note that using a parser this way returns the desired type directly (not as a `interface{}`).


## Utilities

We also include a convenience function to parse values from an HTTP request 
(query arguments or body data):

```go
// func ParseArg(r *http.Request, name string, kindName string, required bool) (value interface{}) {
parse.ParseArg(r, "x", "float64", true)
```
