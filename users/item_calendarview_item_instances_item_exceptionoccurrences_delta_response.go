package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponse struct {
    ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaGetResponse
}
// NewItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponse instantiates a new ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponse and sets the default values.
func NewItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponse()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponse) {
    m := &ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponse{
        ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaGetResponse: *NewItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaResponseable interface {
    ItemCalendarviewItemInstancesItemExceptionoccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
