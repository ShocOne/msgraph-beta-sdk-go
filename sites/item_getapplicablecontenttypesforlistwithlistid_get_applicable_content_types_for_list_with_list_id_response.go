package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable instead.
type ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponse struct {
    ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponse
}
// NewItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponse instantiates a new ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponse and sets the default values.
func NewItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponse()(*ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponse) {
    m := &ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponse{
        ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponse: *NewItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponse(),
    }
    return m
}
// CreateItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable instead.
type ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdResponseable interface {
    ItemGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
