package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody provides operations to call the removeMembers method.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody instantiates a new AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody) {
    m := &AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
