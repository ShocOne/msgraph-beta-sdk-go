package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridGetResponseable instead.
type ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponse struct {
    ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridGetResponse
}
// NewResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponse instantiates a new ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponse and sets the default values.
func NewResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponse()(*ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponse) {
    m := &ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponse{
        ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridGetResponse: *NewResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridGetResponse(),
    }
    return m
}
// CreateResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponse(), nil
}
// Deprecated: This class is obsolete. Use ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridGetResponseable instead.
type ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceoperationsItemGetscopesforuserwithuseridGetScopesForUserWithUseridGetResponseable
}
