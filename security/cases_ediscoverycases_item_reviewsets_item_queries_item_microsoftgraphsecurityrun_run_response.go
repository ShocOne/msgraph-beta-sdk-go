package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponseable instead.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponse struct {
    CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponse
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponse instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponse and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponse()(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponse) {
    m := &CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponse{
        CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponse: *NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponse(),
    }
    return m
}
// CreateCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponse(), nil
}
// Deprecated: This class is obsolete. Use CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponseable instead.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponseable interface {
    CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
