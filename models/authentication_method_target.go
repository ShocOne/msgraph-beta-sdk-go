package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodTarget provides operations to manage the collection of accessReviewDecision entities.
type AuthenticationMethodTarget struct {
    Entity
    // Determines if the user is enforced to register the authentication method.
    isRegistrationRequired *bool
    // The targetType property
    targetType *AuthenticationMethodTargetType
}
// NewAuthenticationMethodTarget instantiates a new authenticationMethodTarget and sets the default values.
func NewAuthenticationMethodTarget()(*AuthenticationMethodTarget) {
    m := &AuthenticationMethodTarget{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationMethodTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationMethodTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodTarget":
                        return NewMicrosoftAuthenticatorAuthenticationMethodTarget(), nil
                    case "#microsoft.graph.smsAuthenticationMethodTarget":
                        return NewSmsAuthenticationMethodTarget(), nil
                }
            }
        }
    }
    return NewAuthenticationMethodTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isRegistrationRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRegistrationRequired)
    res["targetType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAuthenticationMethodTargetType , m.SetTargetType)
    return res
}
// GetIsRegistrationRequired gets the isRegistrationRequired property value. Determines if the user is enforced to register the authentication method.
func (m *AuthenticationMethodTarget) GetIsRegistrationRequired()(*bool) {
    return m.isRegistrationRequired
}
// GetTargetType gets the targetType property value. The targetType property
func (m *AuthenticationMethodTarget) GetTargetType()(*AuthenticationMethodTargetType) {
    return m.targetType
}
// Serialize serializes information the current object
func (m *AuthenticationMethodTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isRegistrationRequired", m.GetIsRegistrationRequired())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err = writer.WriteStringValue("targetType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsRegistrationRequired sets the isRegistrationRequired property value. Determines if the user is enforced to register the authentication method.
func (m *AuthenticationMethodTarget) SetIsRegistrationRequired(value *bool)() {
    m.isRegistrationRequired = value
}
// SetTargetType sets the targetType property value. The targetType property
func (m *AuthenticationMethodTarget) SetTargetType(value *AuthenticationMethodTargetType)() {
    m.targetType = value
}
