package connections

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponseable instead.
type ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponse struct {
    ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponse
}
// NewItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponse instantiates a new ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponse and sets the default values.
func NewItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponse()(*ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponse) {
    m := &ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponse{
        ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponse: *NewItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponse(),
    }
    return m
}
// CreateItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponseable instead.
type ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponseable interface {
    ItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
