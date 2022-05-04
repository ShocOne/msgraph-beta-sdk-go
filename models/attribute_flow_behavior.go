package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of application entities.
type AttributeFlowBehavior int

const (
    FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR AttributeFlowBehavior = iota
    FLOWALWAYS_ATTRIBUTEFLOWBEHAVIOR
)

func (i AttributeFlowBehavior) String() string {
    return []string{"FLOWWHENCHANGED", "FLOWALWAYS"}[i]
}
func ParseAttributeFlowBehavior(v string) (interface{}, error) {
    result := FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR
    switch strings.ToUpper(v) {
        case "FLOWWHENCHANGED":
            result = FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR
        case "FLOWALWAYS":
            result = FLOWALWAYS_ATTRIBUTEFLOWBEHAVIOR
        default:
            return 0, errors.New("Unknown AttributeFlowBehavior value: " + v)
    }
    return &result, nil
}
func SerializeAttributeFlowBehavior(values []AttributeFlowBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}