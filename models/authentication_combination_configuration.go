package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationCombinationConfiguration provides operations to manage the collection of administrativeUnit entities.
type AuthenticationCombinationConfiguration struct {
    Entity
    // Which authentication method combinations this configuration applies to. Must be an allowedCombinations object that's defined for the authenticationStrengthPolicy. The only possible value for fido2combinationConfigurations is 'fido2'.
    appliesToCombinations []AuthenticationMethodModes
}
// NewAuthenticationCombinationConfiguration instantiates a new authenticationCombinationConfiguration and sets the default values.
func NewAuthenticationCombinationConfiguration()(*AuthenticationCombinationConfiguration) {
    m := &AuthenticationCombinationConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationCombinationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationCombinationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.fido2CombinationConfiguration":
                        return NewFido2CombinationConfiguration(), nil
                }
            }
        }
    }
    return NewAuthenticationCombinationConfiguration(), nil
}
// GetAppliesToCombinations gets the appliesToCombinations property value. Which authentication method combinations this configuration applies to. Must be an allowedCombinations object that's defined for the authenticationStrengthPolicy. The only possible value for fido2combinationConfigurations is 'fido2'.
func (m *AuthenticationCombinationConfiguration) GetAppliesToCombinations()([]AuthenticationMethodModes) {
    return m.appliesToCombinations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationCombinationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesToCombinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAuthenticationMethodModes)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodModes, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthenticationMethodModes))
            }
            m.SetAppliesToCombinations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuthenticationCombinationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesToCombinations() != nil {
        err = writer.WriteCollectionOfStringValues("appliesToCombinations", SerializeAuthenticationMethodModes(m.GetAppliesToCombinations()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesToCombinations sets the appliesToCombinations property value. Which authentication method combinations this configuration applies to. Must be an allowedCombinations object that's defined for the authenticationStrengthPolicy. The only possible value for fido2combinationConfigurations is 'fido2'.
func (m *AuthenticationCombinationConfiguration) SetAppliesToCombinations(value []AuthenticationMethodModes)() {
    m.appliesToCombinations = value
}
