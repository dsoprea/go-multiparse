package parse

import (
    "fmt"
    "reflect"
    "go/types"
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

    KindNameZeroType = map[string]reflect.Type {
        "string": reflect.TypeOf(types.Bool),
        "int8": reflect.TypeOf(types.Int8),
        "int16": reflect.TypeOf(types.Int16),
        "int32": reflect.TypeOf(types.Int32),
        "int64": reflect.TypeOf(types.Int64),
        "uint8": reflect.TypeOf(types.Uint8),
        "uint16": reflect.TypeOf(types.Uint16),
        "uint32": reflect.TypeOf(types.Uint32),
        "uint64": reflect.TypeOf(types.Uint64),
        "hex8": reflect.TypeOf(types.String),
        "hex16": reflect.TypeOf(types.String),
        "hex32": reflect.TypeOf(types.String),
        "hex64": reflect.TypeOf(types.String),
        "float32": reflect.TypeOf(types.Float32),
        "float64": reflect.TypeOf(types.Float64),
        "bool": reflect.TypeOf(types.Bool),
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
    Uint64(value interface{}) uint64

    Hex8(value interface{}) uint8
    Hex16(value interface{}) uint16
    Hex32(value interface{}) uint32
    Hex64(value interface{}) uint64

    Float32(value interface{}) float32
    Float64(value interface{}) float64

    Bool(value interface{}) bool
}

func AddParser(fromKind reflect.Kind, p Parser) {
    parsers[fromKind] = p
}

func GetParser(k reflect.Kind) Parser {
    p, found := parsers[k]
    if found == false {
        panic(fmt.Errorf("no parser register for kind [%s]", k))
    }

    return p
}

func Parse(valueRaw interface{}, toKindName string) interface{} {
    if valueRaw == nil {
        if t, found := KindNameZeroType[toKindName]; found == false {
            panic(fmt.Errorf("kind [%s] does not have a zero-type defined", toKindName))
        } else {
            return reflect.Zero(t)
        }
    }

    fromKind := reflect.TypeOf(valueRaw).Kind()
    p := GetParser(fromKind)

    mn, found := NameMethodMap[toKindName]
    if found == false {
        panic(fmt.Errorf("no operation from kind (%d) to kind [%s]", fromKind, toKindName))
    }

    pValue := reflect.ValueOf(p)

    m := pValue.MethodByName(mn)
    if m.IsValid() == false {
        panic(fmt.Errorf("parser [%s] method [%s] not valid", pValue.Type(), mn))
    }

    vV := reflect.ValueOf(valueRaw)
    parsed := m.Call([]reflect.Value { vV })
    return parsed[0].Interface()
}
