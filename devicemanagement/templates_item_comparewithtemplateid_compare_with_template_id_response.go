package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TemplatesItemComparewithtemplateidCompareWithTemplateIdGetResponseable instead.
type TemplatesItemComparewithtemplateidCompareWithTemplateIdResponse struct {
    TemplatesItemComparewithtemplateidCompareWithTemplateIdGetResponse
}
// NewTemplatesItemComparewithtemplateidCompareWithTemplateIdResponse instantiates a new TemplatesItemComparewithtemplateidCompareWithTemplateIdResponse and sets the default values.
func NewTemplatesItemComparewithtemplateidCompareWithTemplateIdResponse()(*TemplatesItemComparewithtemplateidCompareWithTemplateIdResponse) {
    m := &TemplatesItemComparewithtemplateidCompareWithTemplateIdResponse{
        TemplatesItemComparewithtemplateidCompareWithTemplateIdGetResponse: *NewTemplatesItemComparewithtemplateidCompareWithTemplateIdGetResponse(),
    }
    return m
}
// CreateTemplatesItemComparewithtemplateidCompareWithTemplateIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTemplatesItemComparewithtemplateidCompareWithTemplateIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplatesItemComparewithtemplateidCompareWithTemplateIdResponse(), nil
}
// Deprecated: This class is obsolete. Use TemplatesItemComparewithtemplateidCompareWithTemplateIdGetResponseable instead.
type TemplatesItemComparewithtemplateidCompareWithTemplateIdResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TemplatesItemComparewithtemplateidCompareWithTemplateIdGetResponseable
}
