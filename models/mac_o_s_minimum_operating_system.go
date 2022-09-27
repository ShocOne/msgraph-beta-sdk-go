package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSMinimumOperatingSystem the minimum operating system required for a MacOS app.
type MacOSMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Mac OS 10.10 or later.
    v10_10 *bool
    // Mac OS 10.11 or later.
    v10_11 *bool
    // Mac OS 10.12 or later.
    v10_12 *bool
    // Mac OS 10.13 or later.
    v10_13 *bool
    // Mac OS 10.14 or later.
    v10_14 *bool
    // Mac OS 10.15 or later.
    v10_15 *bool
    // Mac OS 10.7 or later.
    v10_7 *bool
    // Mac OS 10.8 or later.
    v10_8 *bool
    // Mac OS 10.9 or later.
    v10_9 *bool
    // Mac OS 11.0 or later.
    v11_0 *bool
    // Mac OS 12.0 or later.
    v12_0 *bool
}
// NewMacOSMinimumOperatingSystem instantiates a new macOSMinimumOperatingSystem and sets the default values.
func NewMacOSMinimumOperatingSystem()(*MacOSMinimumOperatingSystem) {
    m := &MacOSMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.macOSMinimumOperatingSystem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSMinimumOperatingSystem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["v10_10"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_10)
    res["v10_11"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_11)
    res["v10_12"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_12)
    res["v10_13"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_13)
    res["v10_14"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_14)
    res["v10_15"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_15)
    res["v10_7"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_7)
    res["v10_8"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_8)
    res["v10_9"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV10_9)
    res["v11_0"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV11_0)
    res["v12_0"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetV12_0)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV10_10 gets the v10_10 property value. Mac OS 10.10 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_10()(*bool) {
    return m.v10_10
}
// GetV10_11 gets the v10_11 property value. Mac OS 10.11 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_11()(*bool) {
    return m.v10_11
}
// GetV10_12 gets the v10_12 property value. Mac OS 10.12 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_12()(*bool) {
    return m.v10_12
}
// GetV10_13 gets the v10_13 property value. Mac OS 10.13 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_13()(*bool) {
    return m.v10_13
}
// GetV10_14 gets the v10_14 property value. Mac OS 10.14 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_14()(*bool) {
    return m.v10_14
}
// GetV10_15 gets the v10_15 property value. Mac OS 10.15 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_15()(*bool) {
    return m.v10_15
}
// GetV10_7 gets the v10_7 property value. Mac OS 10.7 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_7()(*bool) {
    return m.v10_7
}
// GetV10_8 gets the v10_8 property value. Mac OS 10.8 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_8()(*bool) {
    return m.v10_8
}
// GetV10_9 gets the v10_9 property value. Mac OS 10.9 or later.
func (m *MacOSMinimumOperatingSystem) GetV10_9()(*bool) {
    return m.v10_9
}
// GetV11_0 gets the v11_0 property value. Mac OS 11.0 or later.
func (m *MacOSMinimumOperatingSystem) GetV11_0()(*bool) {
    return m.v11_0
}
// GetV12_0 gets the v12_0 property value. Mac OS 12.0 or later.
func (m *MacOSMinimumOperatingSystem) GetV12_0()(*bool) {
    return m.v12_0
}
// Serialize serializes information the current object
func (m *MacOSMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_10", m.GetV10_10())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_11", m.GetV10_11())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_12", m.GetV10_12())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_13", m.GetV10_13())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_14", m.GetV10_14())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_15", m.GetV10_15())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_7", m.GetV10_7())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_8", m.GetV10_8())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_9", m.GetV10_9())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v11_0", m.GetV11_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v12_0", m.GetV12_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSMinimumOperatingSystem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV10_10 sets the v10_10 property value. Mac OS 10.10 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_10(value *bool)() {
    m.v10_10 = value
}
// SetV10_11 sets the v10_11 property value. Mac OS 10.11 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_11(value *bool)() {
    m.v10_11 = value
}
// SetV10_12 sets the v10_12 property value. Mac OS 10.12 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_12(value *bool)() {
    m.v10_12 = value
}
// SetV10_13 sets the v10_13 property value. Mac OS 10.13 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_13(value *bool)() {
    m.v10_13 = value
}
// SetV10_14 sets the v10_14 property value. Mac OS 10.14 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_14(value *bool)() {
    m.v10_14 = value
}
// SetV10_15 sets the v10_15 property value. Mac OS 10.15 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_15(value *bool)() {
    m.v10_15 = value
}
// SetV10_7 sets the v10_7 property value. Mac OS 10.7 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_7(value *bool)() {
    m.v10_7 = value
}
// SetV10_8 sets the v10_8 property value. Mac OS 10.8 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_8(value *bool)() {
    m.v10_8 = value
}
// SetV10_9 sets the v10_9 property value. Mac OS 10.9 or later.
func (m *MacOSMinimumOperatingSystem) SetV10_9(value *bool)() {
    m.v10_9 = value
}
// SetV11_0 sets the v11_0 property value. Mac OS 11.0 or later.
func (m *MacOSMinimumOperatingSystem) SetV11_0(value *bool)() {
    m.v11_0 = value
}
// SetV12_0 sets the v12_0 property value. Mac OS 12.0 or later.
func (m *MacOSMinimumOperatingSystem) SetV12_0(value *bool)() {
    m.v12_0 = value
}
