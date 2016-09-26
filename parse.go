package parse

import (
    "reflect"
    "fmt"
)

var (
    // Use from standard reflection kinds to internal names.
    KindNameMap = map[reflect.Kind]string {
        reflect.String: "string",
        reflect.Int8: "int8",
        reflect.Int16: "int16",
        reflect.Int32: "int32",
        reflect.Int64: "int64",
        reflect.Uint8: "uint8",
        reflect.Uint16: "uint16",
        reflect.Uint32: "uint32",
        reflect.Uint64: "uint64",
        reflect.Float32: "float32",
        reflect.Float64: "float64",
    }

    // Comprehensive mapping to method names.
    NameMethodMap = map[string]string {
        "string": "String",
        "int8": "Int8",
        "int16": "Int16",
        "int32": "Int32",
        "int64": "Int64",
        "uint8": "Uint8",
        "uint16": "Uint16",
        "uint32": "Uint32",
        "uint64": "Uint64",
        "hex8": "Hex8",
        "hex16": "Hex16",
        "hex32": "Hex32",
        "hex64": "Hex64",
        "float32": "Float32",
        "float64": "Float64",
        "bool": "Bool",
    }

    parsers = map[reflect.Kind]Parser {}
)

type Parser interface {
    String(value interface{}) string

    Int8(value interface{}) int8
    Int16(value interface{}) int16
    Int32(value interface{}) int32
    Int64(value interface{}) int64

    Uint8(value interface{}) uint8
    Uint16(value interface{}) uint16
    Uint32(value interface{}) uint32
    Uint64(value interface{}) uint16

    Hex8(value interface{}) uint8
    Hex16(value interface{}) uint16
    Hex32(value interface{}) uint32
    Hex64(value interface{}) uint64

    Float32(value interface{}) float32
    Float64(value interface{}) float64

    Bool(value interface{}) bool
}

func AddParser(fromKind string, p Parser) {
    parsers[fromKind] = p
}

func GetParser(k string) Parser {
    p, found := parsers[k]
    if found == false {
        panic(fmt.Errorf("no parser register for kind [%s]", k))
    }

    return p
}

func Parse(value interface{}, toKindName string) interface{} {
    k := reflect.Kind(value)
    p := GetParser(k)
    m, err := NameMethodMap[toKindName]
    if err != nil {
        panic(err)
    }

    parsed := reflect.ValueOf(&p).MethodByName(m).Call([]reflect.Value { value })
    return parsed
}

func ParseK(value interface{}, toKind reflect.Kind) {
    if toKindName, found := KindNameMap[toKind]; found == false {
        panic(fmt.Errorf("to-kind (%d) is not supported", toKind))
    } else {
        return Parse(value, toKindName)
    }
}
