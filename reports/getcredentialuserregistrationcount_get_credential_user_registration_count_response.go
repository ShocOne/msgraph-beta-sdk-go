package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponseable instead.
type GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponse struct {
    GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponse
}
// NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponse instantiates a new GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponse and sets the default values.
func NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponse()(*GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponse) {
    m := &GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponse{
        GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponse: *NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponse(),
    }
    return m
}
// CreateGetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponse(), nil
}
// Deprecated: This class is obsolete. Use GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponseable instead.
type GetcredentialuserregistrationcountGetCredentialUserRegistrationCountResponseable interface {
    GetcredentialuserregistrationcountGetCredentialUserRegistrationCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
