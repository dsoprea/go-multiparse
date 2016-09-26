## Introduction

`go-multiparse` allows you to easily convert one type of value into another, e.g. strings from query values in web requests to integers or floats. The original purpose of this project is to reduce the repetition in processing string values though it is not limited to just parsing strings. For more details see below. Parse errors will simply panic.


## Example

```go
// import "github.com/dsoprea/go-multiparse"

valueInterface := parse.Parse("123.456", "float64")
value := valueInterface.(float64)
```


## Implementation

This is a general parsing framework. Parser implementations must satisfy the `Parser` interface and be registered for a specific input type. You might have to register multiple times if you are trying to parse both a base type *and* one or more aliases of it. The `StringParser` is included and parses over string values. It is automatically registered.


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

To register a new implementation, use `AddParser`. An example based on how we automatically register `StringParser`:

```go
type := reflect.TypeOf("")
parse.AddParser(type, parserInstance)
```

Of course, it usually would not make any sense to register another string parser. However, if you have a more exotic type that you often need to convert, implement the interface methods that make sense, panic on the others, and use `AddParser` to register it.

The `Parse()` function is a convenience wrapper. You may acquire and use the parser instance directly:

```go
// import "reflect"

valueRaw := "123.456"
type := reflect.TypeOf(valueRaw)
parser := parse.GetParser(type)
value := parser.Float64(valueRaw)
```

Note that using a parser this way returns the desired type directly (not as a `interface{}`).
