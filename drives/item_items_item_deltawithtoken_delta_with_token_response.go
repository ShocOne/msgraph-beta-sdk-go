package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemDeltawithtokenDeltaWithTokenGetResponseable instead.
type ItemItemsItemDeltawithtokenDeltaWithTokenResponse struct {
    ItemItemsItemDeltawithtokenDeltaWithTokenGetResponse
}
// NewItemItemsItemDeltawithtokenDeltaWithTokenResponse instantiates a new ItemItemsItemDeltawithtokenDeltaWithTokenResponse and sets the default values.
func NewItemItemsItemDeltawithtokenDeltaWithTokenResponse()(*ItemItemsItemDeltawithtokenDeltaWithTokenResponse) {
    m := &ItemItemsItemDeltawithtokenDeltaWithTokenResponse{
        ItemItemsItemDeltawithtokenDeltaWithTokenGetResponse: *NewItemItemsItemDeltawithtokenDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateItemItemsItemDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemDeltawithtokenDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemDeltawithtokenDeltaWithTokenGetResponseable instead.
type ItemItemsItemDeltawithtokenDeltaWithTokenResponseable interface {
    ItemItemsItemDeltawithtokenDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
