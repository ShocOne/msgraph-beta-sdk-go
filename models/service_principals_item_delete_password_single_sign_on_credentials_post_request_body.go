package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody provides operations to call the deletePasswordSingleSignOnCredentials method.
type ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The id property
    id *string
}
// NewServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody instantiates a new ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody and sets the default values.
func NewServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody()(*ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody) {
    m := &ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *ServicePrincipalsItemDeletePasswordSingleSignOnCredentialsPostRequestBody) SetId(value *string)() {
    m.id = value
}
