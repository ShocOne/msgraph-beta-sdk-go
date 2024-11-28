package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CopilotAdminLimitedMode struct {
    Entity
}
// NewCopilotAdminLimitedMode instantiates a new CopilotAdminLimitedMode and sets the default values.
func NewCopilotAdminLimitedMode()(*CopilotAdminLimitedMode) {
    m := &CopilotAdminLimitedMode{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCopilotAdminLimitedModeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotAdminLimitedModeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotAdminLimitedMode(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotAdminLimitedMode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["isEnabledForGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForGroup(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The ID of a Microsoft Entra group to which the value of the isEnabledForGroup property is applied value. The default value is null. This parameter is optional. If isEnabledForGroup is set to true, the groupId value must be provided in order for Copilot limited mode in Teams Meetings to be enabled for the members of the group.
// returns a *string when successful
func (m *CopilotAdminLimitedMode) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsEnabledForGroup gets the isEnabledForGroup property value. Enables the user to be in limited mode for Copilot in Teams meetings. When copilotAdminLimitedMode=true, users in this mode can ask any questions, but Copilot doesn't respond to certain questions related to inferring emotions, behavior, or judgments. When copilotAdminLimitedMode=false, it responds to any types of questions grounded to the meeting conversation. The default value is false.
// returns a *bool when successful
func (m *CopilotAdminLimitedMode) GetIsEnabledForGroup()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabledForGroup")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CopilotAdminLimitedMode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledForGroup", m.GetIsEnabledForGroup())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupId sets the groupId property value. The ID of a Microsoft Entra group to which the value of the isEnabledForGroup property is applied value. The default value is null. This parameter is optional. If isEnabledForGroup is set to true, the groupId value must be provided in order for Copilot limited mode in Teams Meetings to be enabled for the members of the group.
func (m *CopilotAdminLimitedMode) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabledForGroup sets the isEnabledForGroup property value. Enables the user to be in limited mode for Copilot in Teams meetings. When copilotAdminLimitedMode=true, users in this mode can ask any questions, but Copilot doesn't respond to certain questions related to inferring emotions, behavior, or judgments. When copilotAdminLimitedMode=false, it responds to any types of questions grounded to the meeting conversation. The default value is false.
func (m *CopilotAdminLimitedMode) SetIsEnabledForGroup(value *bool)() {
    err := m.GetBackingStore().Set("isEnabledForGroup", value)
    if err != nil {
        panic(err)
    }
}
type CopilotAdminLimitedModeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupId()(*string)
    GetIsEnabledForGroup()(*bool)
    SetGroupId(value *string)()
    SetIsEnabledForGroup(value *bool)()
}
