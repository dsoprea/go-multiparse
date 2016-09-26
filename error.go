package parse

type ParseError struct {
    value interface{}
    message string
}

func (pe *ParseError) Value() interface{} {
    return pe.value
}

func (pe *ParseError) Error() string {
    return pe.message
}

func NewParseError(value interface{}, message string) error {
    return &ParseError{
        value: value,
        message: message,
    }
}
