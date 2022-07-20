package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserFlowApiConnectorConfigurationable 
type UserFlowApiConnectorConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPostAttributeCollection()(IdentityApiConnectorable)
    GetPostFederationSignup()(IdentityApiConnectorable)
    GetPreTokenIssuance()(IdentityApiConnectorable)
    SetOdataType(value *string)()
    SetPostAttributeCollection(value IdentityApiConnectorable)()
    SetPostFederationSignup(value IdentityApiConnectorable)()
    SetPreTokenIssuance(value IdentityApiConnectorable)()
}
