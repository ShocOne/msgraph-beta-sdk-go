package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CopilotAdmin struct {
    Entity
}
// NewCopilotAdmin instantiates a new CopilotAdmin and sets the default values.
func NewCopilotAdmin()(*CopilotAdmin) {
    m := &CopilotAdmin{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCopilotAdminFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotAdminFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotAdmin(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotAdmin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotAdminSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(CopilotAdminSettingable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. Set of Microsoft 365 Copilot settings that can be added or modified. Read-only. Nullable.
// returns a CopilotAdminSettingable when successful
func (m *CopilotAdmin) GetSettings()(CopilotAdminSettingable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CopilotAdminSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CopilotAdmin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettings sets the settings property value. Set of Microsoft 365 Copilot settings that can be added or modified. Read-only. Nullable.
func (m *CopilotAdmin) SetSettings(value CopilotAdminSettingable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
type CopilotAdminable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettings()(CopilotAdminSettingable)
    SetSettings(value CopilotAdminSettingable)()
}
