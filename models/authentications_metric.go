package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationsMetric 
type AuthenticationsMetric struct {
    Entity
}
// NewAuthenticationsMetric instantiates a new authenticationsMetric and sets the default values.
func NewAuthenticationsMetric()(*AuthenticationsMetric) {
    m := &AuthenticationsMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationsMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationsMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationsMetric(), nil
}
// GetAppid gets the appid property value. The appid property
func (m *AuthenticationsMetric) GetAppid()(*string) {
    val, err := m.GetBackingStore().Get("appid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAttemptsCount gets the attemptsCount property value. The attemptsCount property
func (m *AuthenticationsMetric) GetAttemptsCount()(*int64) {
    val, err := m.GetBackingStore().Get("attemptsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetCountry gets the country property value. The country property
func (m *AuthenticationsMetric) GetCountry()(*string) {
    val, err := m.GetBackingStore().Get("country")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFactDate gets the factDate property value. The factDate property
func (m *AuthenticationsMetric) GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("factDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationsMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppid(val)
        }
        return nil
    }
    res["attemptsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttemptsCount(val)
        }
        return nil
    }
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["factDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactDate(val)
        }
        return nil
    }
    res["os"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
        }
        return nil
    }
    res["successCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
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
// GetOs gets the os property value. The os property
func (m *AuthenticationsMetric) GetOs()(*string) {
    val, err := m.GetBackingStore().Get("os")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuccessCount gets the successCount property value. The successCount property
func (m *AuthenticationsMetric) GetSuccessCount()(*int64) {
    val, err := m.GetBackingStore().Get("successCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationsMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appid", m.GetAppid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("attemptsCount", m.GetAttemptsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("factDate", m.GetFactDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("os", m.GetOs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppid sets the appid property value. The appid property
func (m *AuthenticationsMetric) SetAppid(value *string)() {
    err := m.GetBackingStore().Set("appid", value)
    if err != nil {
        panic(err)
    }
}
// SetAttemptsCount sets the attemptsCount property value. The attemptsCount property
func (m *AuthenticationsMetric) SetAttemptsCount(value *int64)() {
    err := m.GetBackingStore().Set("attemptsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCountry sets the country property value. The country property
func (m *AuthenticationsMetric) SetCountry(value *string)() {
    err := m.GetBackingStore().Set("country", value)
    if err != nil {
        panic(err)
    }
}
// SetFactDate sets the factDate property value. The factDate property
func (m *AuthenticationsMetric) SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("factDate", value)
    if err != nil {
        panic(err)
    }
}
// SetOs sets the os property value. The os property
func (m *AuthenticationsMetric) SetOs(value *string)() {
    err := m.GetBackingStore().Set("os", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessCount sets the successCount property value. The successCount property
func (m *AuthenticationsMetric) SetSuccessCount(value *int64)() {
    err := m.GetBackingStore().Set("successCount", value)
    if err != nil {
        panic(err)
    }
}
// AuthenticationsMetricable 
type AuthenticationsMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppid()(*string)
    GetAttemptsCount()(*int64)
    GetCountry()(*string)
    GetFactDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetOs()(*string)
    GetSuccessCount()(*int64)
    SetAppid(value *string)()
    SetAttemptsCount(value *int64)()
    SetCountry(value *string)()
    SetFactDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetOs(value *string)()
    SetSuccessCount(value *int64)()
}
