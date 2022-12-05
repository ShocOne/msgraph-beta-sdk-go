package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemDrivesItemListItemsDeltaResponse provides operations to call the delta method.
type UsersItemDrivesItemListItemsDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []ListItemable
}
// NewUsersItemDrivesItemListItemsDeltaResponse instantiates a new UsersItemDrivesItemListItemsDeltaResponse and sets the default values.
func NewUsersItemDrivesItemListItemsDeltaResponse()(*UsersItemDrivesItemListItemsDeltaResponse) {
    m := &UsersItemDrivesItemListItemsDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateUsersItemDrivesItemListItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemDrivesItemListItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemDrivesItemListItemsDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemDrivesItemListItemsDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ListItemable, len(val))
            for i, v := range val {
                res[i] = v.(ListItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemDrivesItemListItemsDeltaResponse) GetValue()([]ListItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemDrivesItemListItemsDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UsersItemDrivesItemListItemsDeltaResponse) SetValue(value []ListItemable)() {
    m.value = value
}
