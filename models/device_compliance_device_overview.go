package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceDeviceOverview 
type DeviceComplianceDeviceOverview struct {
    Entity
    // Version of the policy for that overview
    configurationVersion *int32
    // Number of devices in conflict
    conflictCount *int32
    // Number of error devices
    errorCount *int32
    // Number of failed devices
    failedCount *int32
    // Last update time
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of not applicable devices
    notApplicableCount *int32
    // Number of not applicable devices due to mismatch platform and policy
    notApplicablePlatformCount *int32
    // Number of pending devices
    pendingCount *int32
    // Number of succeeded devices
    successCount *int32
}
// NewDeviceComplianceDeviceOverview instantiates a new deviceComplianceDeviceOverview and sets the default values.
func NewDeviceComplianceDeviceOverview()(*DeviceComplianceDeviceOverview) {
    m := &DeviceComplianceDeviceOverview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceComplianceDeviceOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceDeviceOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceDeviceOverview(), nil
}
// GetConfigurationVersion gets the configurationVersion property value. Version of the policy for that overview
func (m *DeviceComplianceDeviceOverview) GetConfigurationVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationVersion
    }
}
// GetConflictCount gets the conflictCount property value. Number of devices in conflict
func (m *DeviceComplianceDeviceOverview) GetConflictCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictCount
    }
}
// GetErrorCount gets the errorCount property value. Number of error devices
func (m *DeviceComplianceDeviceOverview) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
// GetFailedCount gets the failedCount property value. Number of failed devices
func (m *DeviceComplianceDeviceOverview) GetFailedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceDeviceOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationVersion(val)
        }
        return nil
    }
    res["conflictCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictCount(val)
        }
        return nil
    }
    res["errorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["failedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["notApplicablePlatformCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicablePlatformCount(val)
        }
        return nil
    }
    res["pendingCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingCount(val)
        }
        return nil
    }
    res["successCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. Last update time
func (m *DeviceComplianceDeviceOverview) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceComplianceDeviceOverview) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
// GetNotApplicablePlatformCount gets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
func (m *DeviceComplianceDeviceOverview) GetNotApplicablePlatformCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicablePlatformCount
    }
}
// GetPendingCount gets the pendingCount property value. Number of pending devices
func (m *DeviceComplianceDeviceOverview) GetPendingCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingCount
    }
}
// GetSuccessCount gets the successCount property value. Number of succeeded devices
func (m *DeviceComplianceDeviceOverview) GetSuccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successCount
    }
}
// Serialize serializes information the current object
func (m *DeviceComplianceDeviceOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("configurationVersion", m.GetConfigurationVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictCount", m.GetConflictCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedCount", m.GetFailedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableCount", m.GetNotApplicableCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicablePlatformCount", m.GetNotApplicablePlatformCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingCount", m.GetPendingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationVersion sets the configurationVersion property value. Version of the policy for that overview
func (m *DeviceComplianceDeviceOverview) SetConfigurationVersion(value *int32)() {
    if m != nil {
        m.configurationVersion = value
    }
}
// SetConflictCount sets the conflictCount property value. Number of devices in conflict
func (m *DeviceComplianceDeviceOverview) SetConflictCount(value *int32)() {
    if m != nil {
        m.conflictCount = value
    }
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *DeviceComplianceDeviceOverview) SetErrorCount(value *int32)() {
    if m != nil {
        m.errorCount = value
    }
}
// SetFailedCount sets the failedCount property value. Number of failed devices
func (m *DeviceComplianceDeviceOverview) SetFailedCount(value *int32)() {
    if m != nil {
        m.failedCount = value
    }
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. Last update time
func (m *DeviceComplianceDeviceOverview) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdateDateTime = value
    }
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceComplianceDeviceOverview) SetNotApplicableCount(value *int32)() {
    if m != nil {
        m.notApplicableCount = value
    }
}
// SetNotApplicablePlatformCount sets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
func (m *DeviceComplianceDeviceOverview) SetNotApplicablePlatformCount(value *int32)() {
    if m != nil {
        m.notApplicablePlatformCount = value
    }
}
// SetPendingCount sets the pendingCount property value. Number of pending devices
func (m *DeviceComplianceDeviceOverview) SetPendingCount(value *int32)() {
    if m != nil {
        m.pendingCount = value
    }
}
// SetSuccessCount sets the successCount property value. Number of succeeded devices
func (m *DeviceComplianceDeviceOverview) SetSuccessCount(value *int32)() {
    if m != nil {
        m.successCount = value
    }
}