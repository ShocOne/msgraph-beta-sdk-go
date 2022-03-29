package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptable 
type DeviceHealthScriptable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]DeviceHealthScriptAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetectionScriptContent()([]byte)
    GetDetectionScriptParameters()([]DeviceHealthScriptParameterable)
    GetDeviceRunStates()([]DeviceHealthScriptDeviceStateable)
    GetDisplayName()(*string)
    GetEnforceSignatureCheck()(*bool)
    GetHighestAvailableVersion()(*string)
    GetIsGlobalScript()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublisher()(*string)
    GetRemediationScriptContent()([]byte)
    GetRemediationScriptParameters()([]DeviceHealthScriptParameterable)
    GetRoleScopeTagIds()([]string)
    GetRunAs32Bit()(*bool)
    GetRunAsAccount()(*RunAsAccountType)
    GetRunSummary()(DeviceHealthScriptRunSummaryable)
    GetVersion()(*string)
    SetAssignments(value []DeviceHealthScriptAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetectionScriptContent(value []byte)()
    SetDetectionScriptParameters(value []DeviceHealthScriptParameterable)()
    SetDeviceRunStates(value []DeviceHealthScriptDeviceStateable)()
    SetDisplayName(value *string)()
    SetEnforceSignatureCheck(value *bool)()
    SetHighestAvailableVersion(value *string)()
    SetIsGlobalScript(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublisher(value *string)()
    SetRemediationScriptContent(value []byte)()
    SetRemediationScriptParameters(value []DeviceHealthScriptParameterable)()
    SetRoleScopeTagIds(value []string)()
    SetRunAs32Bit(value *bool)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetRunSummary(value DeviceHealthScriptRunSummaryable)()
    SetVersion(value *string)()
}