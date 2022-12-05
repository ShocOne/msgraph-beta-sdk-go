package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody provides operations to call the createPasswordSingleSignOnCredentials method.
type ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The credentials property
    credentials []Credentialable
    // The id property
    id *string
}
// NewServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody instantiates a new ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody and sets the default values.
func NewServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody()(*ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) {
    m := &ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCredentials gets the credentials property value. The credentials property
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) GetCredentials()([]Credentialable) {
    return m.credentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Credentialable, len(val))
            for i, v := range val {
                res[i] = v.(Credentialable)
            }
            m.SetCredentials(res)
        }
        return nil
    }
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
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredentials()))
        for i, v := range m.GetCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("credentials", cast)
        if err != nil {
            return err
        }
    }
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
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCredentials sets the credentials property value. The credentials property
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) SetCredentials(value []Credentialable)() {
    m.credentials = value
}
// SetId sets the id property value. The id property
func (m *ServicePrincipalsItemCreatePasswordSingleSignOnCredentialsPostRequestBody) SetId(value *string)() {
    m.id = value
}
