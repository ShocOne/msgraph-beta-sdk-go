package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable instead.
type ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponse struct {
    ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponse
}
// NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponse instantiates a new ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponse and sets the default values.
func NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponse()(*ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponse) {
    m := &ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponse{
        ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponse: *NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponse(),
    }
    return m
}
// CreateItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable instead.
type ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponseable interface {
    ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
