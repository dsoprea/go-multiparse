package parse

import (
    "net/http"

    "fmt"
    "reflect"
)

// Maps
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

    stringType = reflect.TypeOf("")

    KindNameZeroType = map[string]reflect.Type {
        "string": stringType,
        "int8": reflect.TypeOf(int8(0)),
        "int16": reflect.TypeOf(int16(0)),
        "int32": reflect.TypeOf(int32(0)),
        "int64": reflect.TypeOf(int64(0)),
        "uint8": reflect.TypeOf(uint8(0)),
        "uint16": reflect.TypeOf(uint16(0)),
        "uint32": reflect.TypeOf(uint32(0)),
        "uint64": reflect.TypeOf(uint64(0)),
        "hex8": stringType,
        "hex16": stringType,
        "hex32": stringType,
        "hex64": stringType,
        "float32": reflect.TypeOf(float32(0)),
        "float64": reflect.TypeOf(float64(0)),
        "bool": reflect.TypeOf(true),
    }

    parsers = map[string]Parser {}
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

func AddParser(fromType reflect.Type, p Parser) {
    parsers[fromType.Name()] = p
}

func GetParser(fromType reflect.Type) Parser {
    p, found := parsers[fromType.Name()]
    if found == false {
        panic(fmt.Errorf("no parser registered for type [%s]", fromType))
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

    fromType := reflect.TypeOf(valueRaw)
    p := GetParser(fromType)

    mn, found := NameMethodMap[toKindName]
    if found == false {
        panic(fmt.Errorf("no operation from type [%s] to kind [%s]", fromType, toKindName))
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

// ParseRequestArg A convenience function to parse values from an incoming 
// request.
func ParseRequestArg(r *http.Request, name string, kindName string, required bool) (value interface{}) {
    valueRaw := r.FormValue(name)
    if valueRaw == "" {
        if required == true {
            panic(fmt.Errorf("query argument empty or omitted: [%s]", name))
        } else {
            return nil
        }
    }

    return Parse(valueRaw, kindName)
}
