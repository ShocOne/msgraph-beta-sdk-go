package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse struct {
    ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponse
}
// NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse) {
    m := &ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse{
        ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponse: *NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponseable instead.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaResponseable interface {
    ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
