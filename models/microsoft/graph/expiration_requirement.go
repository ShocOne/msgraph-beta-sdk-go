package graph
import (
    "strings"
    "errors"
)
// 
type ExpirationRequirement int

const (
    REMEMBERMULTIFACTORAUTHENTICATIONONTRUSTEDDEVICES_EXPIRATIONREQUIREMENT ExpirationRequirement = iota
    TENANTTOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT
    AUDIENCETOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT
    SIGNINFREQUENCYPERIODICREAUTHENTICATION_EXPIRATIONREQUIREMENT
    NGCMFA_EXPIRATIONREQUIREMENT
    SIGNINFREQUENCYEVERYTIME_EXPIRATIONREQUIREMENT
    UNKNOWNFUTUREVALUE_EXPIRATIONREQUIREMENT
)

func (i ExpirationRequirement) String() string {
    return []string{"REMEMBERMULTIFACTORAUTHENTICATIONONTRUSTEDDEVICES", "TENANTTOKENLIFETIMEPOLICY", "AUDIENCETOKENLIFETIMEPOLICY", "SIGNINFREQUENCYPERIODICREAUTHENTICATION", "NGCMFA", "SIGNINFREQUENCYEVERYTIME", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseExpirationRequirement(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "REMEMBERMULTIFACTORAUTHENTICATIONONTRUSTEDDEVICES":
            return REMEMBERMULTIFACTORAUTHENTICATIONONTRUSTEDDEVICES_EXPIRATIONREQUIREMENT, nil
        case "TENANTTOKENLIFETIMEPOLICY":
            return TENANTTOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT, nil
        case "AUDIENCETOKENLIFETIMEPOLICY":
            return AUDIENCETOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT, nil
        case "SIGNINFREQUENCYPERIODICREAUTHENTICATION":
            return SIGNINFREQUENCYPERIODICREAUTHENTICATION_EXPIRATIONREQUIREMENT, nil
        case "NGCMFA":
            return NGCMFA_EXPIRATIONREQUIREMENT, nil
        case "SIGNINFREQUENCYEVERYTIME":
            return SIGNINFREQUENCYEVERYTIME_EXPIRATIONREQUIREMENT, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EXPIRATIONREQUIREMENT, nil
    }
    return 0, errors.New("Unknown ExpirationRequirement value: " + v)
}
func SerializeExpirationRequirement(values []ExpirationRequirement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
