package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintReportsGetOffice365ServicesUserCountsWithPeriodResponseable 
type PrintReportsGetOffice365ServicesUserCountsWithPeriodResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]Office365ServicesUserCountsable)
    SetValue(value []Office365ServicesUserCountsable)()
}
