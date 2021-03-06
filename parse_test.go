package parse

import (
    "testing"
    "os"

    "net/http"
    "net/url"

    "github.com/dsoprea/go-logging"
)

const (
    testEnvironmentKey = "_MULTIPARSE_TEST"
)

func TestFromRequestBody_String_Required_Hit(t *testing.T) {
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Fatalf("Could not fabricate request: [%s]", err)
    }

    req.Form = url.Values{}
    req.Form.Add("SomeKey", "SomeValue")

    actual := FromRequestBody(req, "SomeKey", "string", true)
    expected := "SomeValue"

    if actual != expected {
        t.Fatalf("Actual value does not equal expected value: [%s] != [%s]",
                 actual, expected)
    }
}

func TestFromRequestBody_String_Required_Miss(t *testing.T) {
    defer func() {
        errRaw := recover()
        if errRaw == nil {
            t.Fatalf("Did not get error for missing but required variable.")
        }

        err := errRaw.(error)

        if err.Error() != "regular body argument empty or omitted: [SomeKey]" {
            t.Fatalf("There was an error for a missing but required query argument but it was not the right error: [%v]", err)
        }
    }()

    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Fatalf("Could not fabricate request: [%s]", err)
    }

    FromRequestBody(req, "SomeKey", "string", true)
}

func TestFromRequestBody_String_Optional_Hit(t *testing.T) {
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Fatalf("Could not fabricate request: [%s]", err)
    }

    req.Form = url.Values{}
    req.Form.Add("SomeKey", "SomeValue")

    actual := FromRequestBody(req, "SomeKey", "string", false)
    expected := "SomeValue"

    if actual != expected {
        t.Fatalf("Actual value does not equal expected value: [%s] != [%s]",
                 actual, expected)
    }
}

func TestFromRequestBody_String_Optional_Miss(t *testing.T) {
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Fatalf("Could not fabricate request: [%s]", err)
    }

    recovered := FromRequestBody(req, "SomeKey", "string", false)
    if recovered != nil {
        t.Fatalf("Read value should've been nil: [%s]", recovered)
    }
}

func TestFromRequestBody_Uint64(t *testing.T) {
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Fatalf("Could not fabricate request: [%s]", err)
    }

    valueRaw := "123"

    req.Form = url.Values{}
    req.Form.Add("SomeKey", valueRaw)

    recovered := FromRequestBody(req, "SomeKey", "uint64", false)
    if recovered != uint64(123) {
        t.Fatalf("Read value does not equal written UINT64 value: [%v] != [%s]",
                 recovered, valueRaw)
    }
}

func TestFromRequestQuery(t *testing.T) {
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Fatalf("Could not fabricate request: [%s]", err)
    }

    u, err := url.Parse("?aa=123")
    log.PanicIf(err)

    req.URL = u

    actualRaw := FromRequestQuery(req, "aa", "uint64", true)
    actual := actualRaw.(uint64)

    expected := uint64(123)

    if actual != expected {
        t.Fatalf("Actual value does not equal expected value: [%d] != [%d]",
                 actual, expected)
    }
}

func TestFromRequestHeader(t *testing.T) {
    req, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {
        t.Fatalf("Could not fabricate request: [%s]", err)
    }

    req.Header.Set("AA", "123")

    actualRaw := FromRequestHeader(req, "AA", "uint64", true)
    actual := actualRaw.(uint64)

    expected := uint64(123)

    if actual != expected {
        t.Fatalf("Actual value does not equal expected value: [%d] != [%d]",
                 actual, expected)
    }
}

func TestFromEnviron_String_Required_Hit(t *testing.T) {
    value := "SOME-VALUE"

    os.Unsetenv(testEnvironmentKey)

    err := os.Setenv(testEnvironmentKey, value)
    if err != nil {
        t.Fatalf("Could not set variable: %v", err)
    }

    defer func() {
        os.Unsetenv(testEnvironmentKey)
    }()

    recovered := FromEnviron(testEnvironmentKey, "string", true)
    if recovered != value {
        t.Fatalf("Read value does not equal written value: [%s] != [%s]",
                 recovered, value)
    }
}

func TestFromEnviron_String_Required_Miss(t *testing.T) {
    defer func() {
        errRaw := recover()
        if errRaw == nil {
            t.Fatalf("Did not get error for missing but required variable.")
        }

        err := errRaw.(error)

        if err.Error() != "environment argument empty or omitted: [_MULTIPARSE_TEST]" {
            t.Fatalf("There was an error for a missing but required variable but it was not the right error: [%v]", err)
        }
    }()

    os.Unsetenv(testEnvironmentKey)
    FromEnviron(testEnvironmentKey, "string", true)
}

func TestFromEnviron_String_Optional_Hit(t *testing.T) {
    value := "SOME-VALUE2"

    os.Unsetenv(testEnvironmentKey)

    err := os.Setenv(testEnvironmentKey, value)
    if err != nil {
        t.Fatalf("Could not set variable: %v", err)
    }

    defer func() {
        os.Unsetenv(testEnvironmentKey)
    }()

    recovered := FromEnviron(testEnvironmentKey, "string", false)
    if recovered != value {
        t.Fatalf("Read value does not equal written STRING value: [%s] != [%s]",
                 recovered, value)
    }
}

func TestFromEnviron_String_Optional_Miss(t *testing.T) {
    recovered := FromEnviron(testEnvironmentKey, "string", false)
    if recovered != nil {
        t.Fatalf("Read value should've been nil: [%s]", recovered)
    }
}

func TestFromEnviron_Uint64(t *testing.T) {
    value := "123"

    os.Unsetenv(testEnvironmentKey)

    err := os.Setenv(testEnvironmentKey, value)
    if err != nil {
        t.Fatalf("Could not set variable: %v", err)
    }

    defer func() {
        os.Unsetenv(testEnvironmentKey)
    }()

    recovered := FromEnviron(testEnvironmentKey, "uint64", false)
    if recovered != uint64(123) {
        t.Fatalf("Read value does not equal written UINT64 value: [%v] != [%s]",
                 recovered, value)
    }
}
