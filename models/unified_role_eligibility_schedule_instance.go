package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleEligibilityScheduleInstance 
type UnifiedRoleEligibilityScheduleInstance struct {
    UnifiedRoleScheduleInstanceBase
    // Time that the roleEligibilityScheduleInstance will expire.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Membership type of the assignment. It can either be Inherited, Direct, or Group.
    memberType *string
    // Identifier of the parent roleEligibilitySchedule for this instance.
    roleEligibilityScheduleId *string
    // Time that the roleEligibilityScheduleInstance will start.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUnifiedRoleEligibilityScheduleInstance instantiates a new unifiedRoleEligibilityScheduleInstance and sets the default values.
func NewUnifiedRoleEligibilityScheduleInstance()(*UnifiedRoleEligibilityScheduleInstance) {
    m := &UnifiedRoleEligibilityScheduleInstance{
        UnifiedRoleScheduleInstanceBase: *NewUnifiedRoleScheduleInstanceBase(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRoleEligibilityScheduleInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleEligibilityScheduleInstance(), nil
}
// GetEndDateTime gets the endDateTime property value. Time that the roleEligibilityScheduleInstance will expire.
func (m *UnifiedRoleEligibilityScheduleInstance) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleEligibilityScheduleInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleInstanceBase.GetFieldDeserializers()
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEndDateTime)
    res["memberType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMemberType)
    res["roleEligibilityScheduleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRoleEligibilityScheduleId)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    return res
}
// GetMemberType gets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleEligibilityScheduleInstance) GetMemberType()(*string) {
    return m.memberType
}
// GetRoleEligibilityScheduleId gets the roleEligibilityScheduleId property value. Identifier of the parent roleEligibilitySchedule for this instance.
func (m *UnifiedRoleEligibilityScheduleInstance) GetRoleEligibilityScheduleId()(*string) {
    return m.roleEligibilityScheduleId
}
// GetStartDateTime gets the startDateTime property value. Time that the roleEligibilityScheduleInstance will start.
func (m *UnifiedRoleEligibilityScheduleInstance) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *UnifiedRoleEligibilityScheduleInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleInstanceBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleEligibilityScheduleId", m.GetRoleEligibilityScheduleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDateTime sets the endDateTime property value. Time that the roleEligibilityScheduleInstance will expire.
func (m *UnifiedRoleEligibilityScheduleInstance) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetMemberType sets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleEligibilityScheduleInstance) SetMemberType(value *string)() {
    m.memberType = value
}
// SetRoleEligibilityScheduleId sets the roleEligibilityScheduleId property value. Identifier of the parent roleEligibilitySchedule for this instance.
func (m *UnifiedRoleEligibilityScheduleInstance) SetRoleEligibilityScheduleId(value *string)() {
    m.roleEligibilityScheduleId = value
}
// SetStartDateTime sets the startDateTime property value. Time that the roleEligibilityScheduleInstance will start.
func (m *UnifiedRoleEligibilityScheduleInstance) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
