package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable instead.
type ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse struct {
    ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse
}
// NewServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse instantiates a new ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse and sets the default values.
func NewServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse()(*ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse) {
    m := &ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse{
        ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse: *NewServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse(),
    }
    return m
}
// CreateServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable instead.
type ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceActivityGetActiveUserMetricsForiOSOrAndroidMailByReadEmailWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable
}
