package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponseable instead.
type PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponse struct {
    PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponse
}
// NewPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponse instantiates a new PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponse and sets the default values.
func NewPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponse()(*PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponse) {
    m := &PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponse{
        PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponse: *NewPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponse(),
    }
    return m
}
// CreatePolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponse(), nil
}
// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponseable instead.
type PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponseable
}
