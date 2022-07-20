package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ControlScoreable 
type ControlScoreable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetControlCategory()(*string)
    GetControlName()(*string)
    GetDescription()(*string)
    GetOdataType()(*string)
    GetScore()(*float64)
    SetControlCategory(value *string)()
    SetControlName(value *string)()
    SetDescription(value *string)()
    SetOdataType(value *string)()
    SetScore(value *float64)()
}
