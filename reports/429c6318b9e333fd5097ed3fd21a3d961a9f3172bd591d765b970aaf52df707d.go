package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable instead.
type ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse struct {
    ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse
}
// NewServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse instantiates a new ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse and sets the default values.
func NewServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse()(*ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse) {
    m := &ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse{
        ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse: *NewServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse(),
    }
    return m
}
// CreateServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable instead.
type ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceActivityGetActiveUserMetricsForEmailByModernAuthenticationWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable
}
