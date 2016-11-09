package parse

import (
    "strconv"
    "time"

    "github.com/dsoprea/go-logging"
)

type StringParser struct {
}

func NewStringParser() Parser {
    return new(StringParser)
}

func (sp StringParser) String(value interface{}) string {
    s := value.(string)
    return s
}

func (sp StringParser) Int8(value interface{}) int8 {
    s := value.(string)

    p, err := strconv.ParseInt(s, 10, 8)
    log.PanicIf(err)

    return int8(p)
}

func (sp StringParser) Int16(value interface{}) int16 {
    s := value.(string)

    p, err := strconv.ParseInt(s, 10, 16)
    log.PanicIf(err)

    return int16(p)
}

func (sp StringParser) Int32(value interface{}) int32 {
    s := value.(string)

    p, err := strconv.ParseInt(s, 10, 32)
    log.PanicIf(err)
    
    return int32(p)
}

func (sp StringParser) Int64(value interface{}) int64 {
    s := value.(string)

    p, err := strconv.ParseInt(s, 10, 64)
    log.PanicIf(err)
    
    return int64(p)
}

func (sp StringParser) Uint8(value interface{}) uint8 {
    s := value.(string)
    
    p, err := strconv.ParseUint(s, 10, 8)
    log.PanicIf(err)
    
    return uint8(p)
}

func (sp StringParser) Uint16(value interface{}) uint16 {
    s := value.(string)

    p, err := strconv.ParseUint(s, 10, 16)
    log.PanicIf(err)

    return uint16(p)
}

func (sp StringParser) Uint32(value interface{}) uint32 {
    s := value.(string)

    p, err := strconv.ParseUint(s, 10, 32)
    log.PanicIf(err)

    return uint32(p)
}

func (sp StringParser) Uint64(value interface{}) uint64 {
    s := value.(string)

    p, err := strconv.ParseUint(s, 10, 64)
    log.PanicIf(err)
    
    return p
}

func (sp StringParser) Hex8(value interface{}) uint8 {
    s := value.(string)

    p, err := strconv.ParseUint(s, 16, 8)
    log.PanicIf(err)
    
    return uint8(p)
}

func (sp StringParser) Hex16(value interface{}) uint16 {
    s := value.(string)

    p, err := strconv.ParseUint(s, 16, 16)
    log.PanicIf(err)
    
    return uint16(p)
}

func (sp StringParser) Hex32(value interface{}) uint32 {
    s := value.(string)

    p, err := strconv.ParseUint(s, 16, 32)
    log.PanicIf(err)

    return uint32(p)
}

func (sp StringParser) Hex64(value interface{}) uint64 {
    s := value.(string)
    
    p, err := strconv.ParseUint(s, 16, 64)
    log.PanicIf(err)
    
    return p
}

func (sp StringParser) Float32(value interface{}) float32 {
    s := value.(string)

    p, err := strconv.ParseFloat(s, 32)
    log.PanicIf(err)
    
    return float32(p)
}

func (sp StringParser) Float64(value interface{}) float64 {
    s := value.(string)
    
    p, err := strconv.ParseFloat(s, 64)
    log.PanicIf(err)
    
    return p
}

func (sp StringParser) Bool(value interface{}) bool {
    s := value.(string)

    p, err := strconv.ParseBool(s)
    log.PanicIf(err)
    
    return p
}

func (sp StringParser) Rfc3339(value interface{}) time.Time {
    s := value.(string)

    t, err := time.Parse(time.RFC3339, s)
    log.PanicIf(err)
    
    return t
}

func init() {
    p := NewStringParser()
    AddParser(stringType, p)
}
