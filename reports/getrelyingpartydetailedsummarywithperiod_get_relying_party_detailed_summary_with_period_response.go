package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodGetResponseable instead.
type GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponse struct {
    GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodGetResponse
}
// NewGetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponse instantiates a new GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponse and sets the default values.
func NewGetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponse()(*GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponse) {
    m := &GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponse{
        GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodGetResponse: *NewGetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodGetResponse(),
    }
    return m
}
// CreateGetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponse(), nil
}
// Deprecated: This class is obsolete. Use GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodGetResponseable instead.
type GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodResponseable interface {
    GetrelyingpartydetailedsummarywithperiodGetRelyingPartyDetailedSummaryWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
