package models
import (
    "strings"
    "errors"
)
// Provides operations to call the evaluateApplication method.
type ContentFormat int

const (
    DEFAULT_ESCAPED_CONTENTFORMAT ContentFormat = iota
    EMAIL_CONTENTFORMAT
)

func (i ContentFormat) String() string {
    return []string{"DEFAULT_ESCAPED", "EMAIL"}[i]
}
func ParseContentFormat(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_CONTENTFORMAT
    switch strings.ToUpper(v) {
        case "DEFAULT_ESCAPED":
            result = DEFAULT_ESCAPED_CONTENTFORMAT
        case "EMAIL":
            result = EMAIL_CONTENTFORMAT
        default:
            return 0, errors.New("Unknown ContentFormat value: " + v)
    }
    return &result, nil
}
func SerializeContentFormat(values []ContentFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}