package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse struct {
    AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse
}
// NewAccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse instantiates a new AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse()(*AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse) {
    m := &AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse{
        AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse: *NewAccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable interface {
    AccessreviewsDefinitionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
