package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable instead.
type InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponse struct {
    InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse
}
// NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponse instantiates a new InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponse and sets the default values.
func NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponse()(*InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponse) {
    m := &InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponse{
        InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse: *NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse(),
    }
    return m
}
// CreateInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponse(), nil
}
// Deprecated: This class is obsolete. Use InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable instead.
type InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseable interface {
    InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
