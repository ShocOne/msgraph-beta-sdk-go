package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse struct {
    ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse
}
// NewItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse instantiates a new ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse and sets the default values.
func NewItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse()(*ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse) {
    m := &ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse{
        ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse: *NewItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable interface {
    ItemApprovalsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
