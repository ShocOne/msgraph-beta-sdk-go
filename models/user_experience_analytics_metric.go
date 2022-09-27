package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsMetric the user experience analytics metric contains the score and units of a metric of a user experience anlaytics category.
type UserExperienceAnalyticsMetric struct {
    Entity
    // The unit of the user experience analytics metric.
    unit *string
    // The value of the user experience analytics metric.
    value *float64
}
// NewUserExperienceAnalyticsMetric instantiates a new userExperienceAnalyticsMetric and sets the default values.
func NewUserExperienceAnalyticsMetric()(*UserExperienceAnalyticsMetric) {
    m := &UserExperienceAnalyticsMetric{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsMetric";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsMetric(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["unit"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUnit)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetValue)
    return res
}
// GetUnit gets the unit property value. The unit of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) GetUnit()(*string) {
    return m.unit
}
// GetValue gets the value property value. The value of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) GetValue()(*float64) {
    return m.value
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("unit", m.GetUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUnit sets the unit property value. The unit of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) SetUnit(value *string)() {
    m.unit = value
}
// SetValue sets the value property value. The value of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) SetValue(value *float64)() {
    m.value = value
}
