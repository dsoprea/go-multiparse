package parse

import (
    "testing"
    "time"

    "github.com/dsoprea/go-logging"
)

func TestParseString(t *testing.T) {
    phrases := [][]interface{} {
        { "some string", "some string" },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.String(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%s] != [%s]", parsed2, pair[1])
        }
    }
}

func TestParseInt8(t *testing.T) {
    phrases := [][]interface{} {
        { "-67", int8(-67) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Int8(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseInt16(t *testing.T) {
    phrases := [][]interface{} {
        { "-7013", int16(-7013) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Int16(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseInt32(t *testing.T) {
    phrases := [][]interface{} {
        { "-33908743", int32(-33908743) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Int32(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseInt64(t *testing.T) {
    phrases := [][]interface{} {
        { "-0945895663560175", int64(-945895663560175) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Int64(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseUint8(t *testing.T) {
    phrases := [][]interface{} {
        { "75", uint8(75) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Uint8(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseUint16(t *testing.T) {
    phrases := [][]interface{} {
        { "15798", uint16(15798) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Uint16(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseUint32(t *testing.T) {
    phrases := [][]interface{} {
        { "07876050", uint32(7876050) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Uint32(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseUint64(t *testing.T) {
    phrases := [][]interface{} {
        { "4101229679899789", uint64(4101229679899789) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Uint64(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseHex8(t *testing.T) {
    phrases := [][]interface{} {
        { "C4", uint8(0xC4) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Hex8(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseHex16(t *testing.T) {
    phrases := [][]interface{} {
        { "9F06", uint16(0x9F06) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Hex16(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseHex32(t *testing.T) {
    phrases := [][]interface{} {
        { "9E311959", uint32(0x9E311959) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Hex32(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseHex64(t *testing.T) {
    phrases := [][]interface{} {
        { "427F2A1E7A23DC3E", uint64(0x427F2A1E7A23DC3E) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Hex64(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%d] != [%d]", parsed2, pair[1])
        }
    }
}

func TestParseFloat32(t *testing.T) {
    phrases := [][]interface{} {
        { "-0.6046603", float32(-0.6046603) },
        { "0.6046603", float32(0.6046603) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Float32(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%f] != [%f]", parsed2, pair[1])
        }
    }
}

func TestParseFloat64(t *testing.T) {
    phrases := [][]interface{} {
        { "-0.9405090880450124", float64(-0.9405090880450124) },
        { "0.9405090880450124", float64(0.9405090880450124) },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Float64(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%f] != [%f]", parsed2, pair[1])
        }
    }
}

func TestParseBool(t *testing.T) {
    phrases := [][]interface{} {
        { "0", false },
        { "1", true },
    }

    for _, pair := range phrases {
        p := NewStringParser()
        parsed2 := p.Bool(pair[0])

        if parsed2 != pair[1] {
            t.Errorf("parse did not produce correct result: [%v] != [%v]", parsed2, pair[1])
        }
    }
}

func TestParseRfc3339(t *testing.T) {
    phrases := []string {
        "2016-11-08T19:32:29Z",
        "2016-11-08T19:32:29+12:34",
    }

    for _, phrase := range phrases {
        parsed1, err := time.Parse(time.RFC3339, phrase)
        if err != nil {
            log.Panic(err)
        }

        p := NewStringParser()
        parsed2 := p.Rfc3339(phrase)

        if parsed2.Equal(parsed1) == false {
            t.Error("parse did not produce correct result")
        }
    }
}
