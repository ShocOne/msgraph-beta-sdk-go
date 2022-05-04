package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessPackageCustomExtensionStage int

const (
    ASSIGNMENTREQUESTCREATED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE AccessPackageCustomExtensionStage = iota
    ASSIGNMENTREQUESTAPPROVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTREQUESTGRANTED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTREQUESTREMOVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTFOURTEENDAYSBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    ASSIGNMENTONEDAYBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
)

func (i AccessPackageCustomExtensionStage) String() string {
    return []string{"ASSIGNMENTREQUESTCREATED", "ASSIGNMENTREQUESTAPPROVED", "ASSIGNMENTREQUESTGRANTED", "ASSIGNMENTREQUESTREMOVED", "ASSIGNMENTFOURTEENDAYSBEFOREEXPIRATION", "ASSIGNMENTONEDAYBEFOREEXPIRATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageCustomExtensionStage(v string) (interface{}, error) {
    result := ASSIGNMENTREQUESTCREATED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
    switch strings.ToUpper(v) {
        case "ASSIGNMENTREQUESTCREATED":
            result = ASSIGNMENTREQUESTCREATED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "ASSIGNMENTREQUESTAPPROVED":
            result = ASSIGNMENTREQUESTAPPROVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "ASSIGNMENTREQUESTGRANTED":
            result = ASSIGNMENTREQUESTGRANTED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "ASSIGNMENTREQUESTREMOVED":
            result = ASSIGNMENTREQUESTREMOVED_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "ASSIGNMENTFOURTEENDAYSBEFOREEXPIRATION":
            result = ASSIGNMENTFOURTEENDAYSBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "ASSIGNMENTONEDAYBEFOREEXPIRATION":
            result = ASSIGNMENTONEDAYBEFOREEXPIRATION_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONSTAGE
        default:
            return 0, errors.New("Unknown AccessPackageCustomExtensionStage value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageCustomExtensionStage(values []AccessPackageCustomExtensionStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}