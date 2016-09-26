package parse

import (
    "strconv"
)

type StringParser struct {
}

func (sp StringParser) String(value interface{}) string {
    s = value.(string)
    return s
}

func (sp StringParser) Int8(value interface{}) int8 {
    s = value.(string)
    if p, err := strconv.ParseInt(s, 10, 8); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Int16(value interface{}) int16 {
    s = value.(string)
    if p, err := strconv.ParseInt(s, 10, 16); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Int32(value interface{}) int32 {
    s = value.(string)
    if p, err := strconv.ParseInt(s, 10, 32); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Int64(value interface{}) int64 {
    s = value.(string)
    if p, err := strconv.ParseInt(s, 10, 64); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Uint8(value interface{}) uint8 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 10, 8); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Uint16(value interface{}) uint16 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 10, 16); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Uint32(value interface{}) uint32 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 10, 32); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Uint64(value interface{}) uint16 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 10, 64); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Hex8(value interface{}) uint8 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 16, 8); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Hex16(value interface{}) uint16 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 16, 16); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Hex32(value interface{}) uint32 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 16, 32); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Hex64(value interface{}) uint64 {
    s = value.(string)
    if p, err := strconv.ParseUint(s, 16, 64); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Float32(value interface{}) float32 {
    s = value.(string)
    if p, err := strconv.ParseFloat(s, 32); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Float64(value interface{}) float64 {
    s = value.(string)
    if p, err := strconv.ParseFloat(s, 64); if err != nil {
        panic(err)
    } else {
        return p
    }
}

func (sp StringParser) Bool(value interface{}) bool {
    s = value.(string)
    if p, err := strconv.ParseBool(s); if err != nil {
        panic(err)
    } else {
        return p
    }

}

func init() {
    AddParser("string", &StringParser{})
}