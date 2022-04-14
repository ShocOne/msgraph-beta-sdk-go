package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceManagementExchangeAccessStateReason int

const (
    NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON DeviceManagementExchangeAccessStateReason = iota
    UNKNOWN_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEGLOBALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEINDIVIDUALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEDEVICERULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEUPGRADE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    EXCHANGEMAILBOXPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    OTHER_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    COMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    NOTCOMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    NOTENROLLED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    UNKNOWNLOCATION_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    MFAREQUIRED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    AZUREADBLOCKDUETOACCESSPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    COMPROMISEDPASSWORD_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    DEVICENOTKNOWNWITHMANAGEDAPP_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
)

func (i DeviceManagementExchangeAccessStateReason) String() string {
    return []string{"NONE", "UNKNOWN", "EXCHANGEGLOBALRULE", "EXCHANGEINDIVIDUALRULE", "EXCHANGEDEVICERULE", "EXCHANGEUPGRADE", "EXCHANGEMAILBOXPOLICY", "OTHER", "COMPLIANT", "NOTCOMPLIANT", "NOTENROLLED", "UNKNOWNLOCATION", "MFAREQUIRED", "AZUREADBLOCKDUETOACCESSPOLICY", "COMPROMISEDPASSWORD", "DEVICENOTKNOWNWITHMANAGEDAPP"}[i]
}
func ParseDeviceManagementExchangeAccessStateReason(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "EXCHANGEGLOBALRULE":
            result = EXCHANGEGLOBALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "EXCHANGEINDIVIDUALRULE":
            result = EXCHANGEINDIVIDUALRULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "EXCHANGEDEVICERULE":
            result = EXCHANGEDEVICERULE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "EXCHANGEUPGRADE":
            result = EXCHANGEUPGRADE_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "EXCHANGEMAILBOXPOLICY":
            result = EXCHANGEMAILBOXPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "OTHER":
            result = OTHER_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "COMPLIANT":
            result = COMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "NOTCOMPLIANT":
            result = NOTCOMPLIANT_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "NOTENROLLED":
            result = NOTENROLLED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "UNKNOWNLOCATION":
            result = UNKNOWNLOCATION_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "MFAREQUIRED":
            result = MFAREQUIRED_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "AZUREADBLOCKDUETOACCESSPOLICY":
            result = AZUREADBLOCKDUETOACCESSPOLICY_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "COMPROMISEDPASSWORD":
            result = COMPROMISEDPASSWORD_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        case "DEVICENOTKNOWNWITHMANAGEDAPP":
            result = DEVICENOTKNOWNWITHMANAGEDAPP_DEVICEMANAGEMENTEXCHANGEACCESSSTATEREASON
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeAccessStateReason value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeAccessStateReason(values []DeviceManagementExchangeAccessStateReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
