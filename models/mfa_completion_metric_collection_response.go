package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MfaCompletionMetricCollectionResponse 
type MfaCompletionMetricCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewMfaCompletionMetricCollectionResponse instantiates a new mfaCompletionMetricCollectionResponse and sets the default values.
func NewMfaCompletionMetricCollectionResponse()(*MfaCompletionMetricCollectionResponse) {
    m := &MfaCompletionMetricCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMfaCompletionMetricCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMfaCompletionMetricCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMfaCompletionMetricCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MfaCompletionMetricCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMfaCompletionMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MfaCompletionMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MfaCompletionMetricable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MfaCompletionMetricCollectionResponse) GetValue()([]MfaCompletionMetricable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MfaCompletionMetricable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MfaCompletionMetricCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MfaCompletionMetricCollectionResponse) SetValue(value []MfaCompletionMetricable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// MfaCompletionMetricCollectionResponseable 
type MfaCompletionMetricCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]MfaCompletionMetricable)
    SetValue(value []MfaCompletionMetricable)()
}
