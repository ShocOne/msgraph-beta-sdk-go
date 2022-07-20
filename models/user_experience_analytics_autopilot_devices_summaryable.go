package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAutopilotDevicesSummaryable 
type UserExperienceAnalyticsAutopilotDevicesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevicesNotAutopilotRegistered()(*int32)
    GetDevicesWithoutAutopilotProfileAssigned()(*int32)
    GetOdataType()(*string)
    GetTotalWindows10DevicesWithoutTenantAttached()(*int32)
    SetDevicesNotAutopilotRegistered(value *int32)()
    SetDevicesWithoutAutopilotProfileAssigned(value *int32)()
    SetOdataType(value *string)()
    SetTotalWindows10DevicesWithoutTenantAttached(value *int32)()
}
