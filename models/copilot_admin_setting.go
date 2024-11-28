package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CopilotAdminSetting struct {
    Entity
}
// NewCopilotAdminSetting instantiates a new CopilotAdminSetting and sets the default values.
func NewCopilotAdminSetting()(*CopilotAdminSetting) {
    m := &CopilotAdminSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCopilotAdminSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotAdminSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotAdminSetting(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotAdminSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["limitedMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotAdminLimitedModeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimitedMode(val.(CopilotAdminLimitedModeable))
        }
        return nil
    }
    return res
}
// GetLimitedMode gets the limitedMode property value. Represents a setting that controls whether Microsoft 365 Copilot in Teams Meetings users can receive responses to sentiment-related prompts. Read-only. Nullable.
// returns a CopilotAdminLimitedModeable when successful
func (m *CopilotAdminSetting) GetLimitedMode()(CopilotAdminLimitedModeable) {
    val, err := m.GetBackingStore().Get("limitedMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CopilotAdminLimitedModeable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CopilotAdminSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("limitedMode", m.GetLimitedMode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLimitedMode sets the limitedMode property value. Represents a setting that controls whether Microsoft 365 Copilot in Teams Meetings users can receive responses to sentiment-related prompts. Read-only. Nullable.
func (m *CopilotAdminSetting) SetLimitedMode(value CopilotAdminLimitedModeable)() {
    err := m.GetBackingStore().Set("limitedMode", value)
    if err != nil {
        panic(err)
    }
}
type CopilotAdminSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLimitedMode()(CopilotAdminLimitedModeable)
    SetLimitedMode(value CopilotAdminLimitedModeable)()
}
