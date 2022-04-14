package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type WindowsDefenderProductStatus int

const (
    NOSTATUS_WINDOWSDEFENDERPRODUCTSTATUS WindowsDefenderProductStatus = iota
    SERVICENOTRUNNING_WINDOWSDEFENDERPRODUCTSTATUS
    SERVICESTARTEDWITHOUTMALWAREPROTECTION_WINDOWSDEFENDERPRODUCTSTATUS
    PENDINGFULLSCANDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    PENDINGREBOOTDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    PENDINGMANUALSTEPSDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    AVSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    ASSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
    NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
    SYSTEMINITIATEDSCANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    SYSTEMINITIATEDCLEANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    SAMPLESPENDINGSUBMISSION_WINDOWSDEFENDERPRODUCTSTATUS
    PRODUCTRUNNINGINEVALUATIONMODE_WINDOWSDEFENDERPRODUCTSTATUS
    PRODUCTRUNNINGINNONGENUINEMODE_WINDOWSDEFENDERPRODUCTSTATUS
    PRODUCTEXPIRED_WINDOWSDEFENDERPRODUCTSTATUS
    OFFLINESCANREQUIRED_WINDOWSDEFENDERPRODUCTSTATUS
    SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN_WINDOWSDEFENDERPRODUCTSTATUS
    THREATREMEDIATIONFAILEDCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
    THREATREMEDIATIONFAILEDNONCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
    NOSTATUSFLAGSSET_WINDOWSDEFENDERPRODUCTSTATUS
    PLATFORMOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    PLATFORMUPDATEINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    PLATFORMABOUTTOBEOUTDATED_WINDOWSDEFENDERPRODUCTSTATUS
    SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING_WINDOWSDEFENDERPRODUCTSTATUS
    WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL_WINDOWSDEFENDERPRODUCTSTATUS
)

func (i WindowsDefenderProductStatus) String() string {
    return []string{"NOSTATUS", "SERVICENOTRUNNING", "SERVICESTARTEDWITHOUTMALWAREPROTECTION", "PENDINGFULLSCANDUETOTHREATACTION", "PENDINGREBOOTDUETOTHREATACTION", "PENDINGMANUALSTEPSDUETOTHREATACTION", "AVSIGNATURESOUTOFDATE", "ASSIGNATURESOUTOFDATE", "NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD", "NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD", "SYSTEMINITIATEDSCANINPROGRESS", "SYSTEMINITIATEDCLEANINPROGRESS", "SAMPLESPENDINGSUBMISSION", "PRODUCTRUNNINGINEVALUATIONMODE", "PRODUCTRUNNINGINNONGENUINEMODE", "PRODUCTEXPIRED", "OFFLINESCANREQUIRED", "SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN", "THREATREMEDIATIONFAILEDCRITICALLY", "THREATREMEDIATIONFAILEDNONCRITICALLY", "NOSTATUSFLAGSSET", "PLATFORMOUTOFDATE", "PLATFORMUPDATEINPROGRESS", "PLATFORMABOUTTOBEOUTDATED", "SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING", "WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL"}[i]
}
func ParseWindowsDefenderProductStatus(v string) (interface{}, error) {
    result := NOSTATUS_WINDOWSDEFENDERPRODUCTSTATUS
    switch strings.ToUpper(v) {
        case "NOSTATUS":
            result = NOSTATUS_WINDOWSDEFENDERPRODUCTSTATUS
        case "SERVICENOTRUNNING":
            result = SERVICENOTRUNNING_WINDOWSDEFENDERPRODUCTSTATUS
        case "SERVICESTARTEDWITHOUTMALWAREPROTECTION":
            result = SERVICESTARTEDWITHOUTMALWAREPROTECTION_WINDOWSDEFENDERPRODUCTSTATUS
        case "PENDINGFULLSCANDUETOTHREATACTION":
            result = PENDINGFULLSCANDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
        case "PENDINGREBOOTDUETOTHREATACTION":
            result = PENDINGREBOOTDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
        case "PENDINGMANUALSTEPSDUETOTHREATACTION":
            result = PENDINGMANUALSTEPSDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
        case "AVSIGNATURESOUTOFDATE":
            result = AVSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
        case "ASSIGNATURESOUTOFDATE":
            result = ASSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
        case "NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD":
            result = NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
        case "NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD":
            result = NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
        case "SYSTEMINITIATEDSCANINPROGRESS":
            result = SYSTEMINITIATEDSCANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
        case "SYSTEMINITIATEDCLEANINPROGRESS":
            result = SYSTEMINITIATEDCLEANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
        case "SAMPLESPENDINGSUBMISSION":
            result = SAMPLESPENDINGSUBMISSION_WINDOWSDEFENDERPRODUCTSTATUS
        case "PRODUCTRUNNINGINEVALUATIONMODE":
            result = PRODUCTRUNNINGINEVALUATIONMODE_WINDOWSDEFENDERPRODUCTSTATUS
        case "PRODUCTRUNNINGINNONGENUINEMODE":
            result = PRODUCTRUNNINGINNONGENUINEMODE_WINDOWSDEFENDERPRODUCTSTATUS
        case "PRODUCTEXPIRED":
            result = PRODUCTEXPIRED_WINDOWSDEFENDERPRODUCTSTATUS
        case "OFFLINESCANREQUIRED":
            result = OFFLINESCANREQUIRED_WINDOWSDEFENDERPRODUCTSTATUS
        case "SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN":
            result = SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN_WINDOWSDEFENDERPRODUCTSTATUS
        case "THREATREMEDIATIONFAILEDCRITICALLY":
            result = THREATREMEDIATIONFAILEDCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
        case "THREATREMEDIATIONFAILEDNONCRITICALLY":
            result = THREATREMEDIATIONFAILEDNONCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
        case "NOSTATUSFLAGSSET":
            result = NOSTATUSFLAGSSET_WINDOWSDEFENDERPRODUCTSTATUS
        case "PLATFORMOUTOFDATE":
            result = PLATFORMOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
        case "PLATFORMUPDATEINPROGRESS":
            result = PLATFORMUPDATEINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
        case "PLATFORMABOUTTOBEOUTDATED":
            result = PLATFORMABOUTTOBEOUTDATED_WINDOWSDEFENDERPRODUCTSTATUS
        case "SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING":
            result = SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING_WINDOWSDEFENDERPRODUCTSTATUS
        case "WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL":
            result = WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL_WINDOWSDEFENDERPRODUCTSTATUS
        default:
            return 0, errors.New("Unknown WindowsDefenderProductStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDefenderProductStatus(values []WindowsDefenderProductStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
