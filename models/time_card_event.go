package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCardEvent 
type TimeCardEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the entry was recorded at the approved location.
    atApprovedLocation *bool
    // The time the entry is recorded.
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Notes about the timeCardEvent.
    notes ItemBodyable
    // The OdataType property
    odataType *string
}
// NewTimeCardEvent instantiates a new timeCardEvent and sets the default values.
func NewTimeCardEvent()(*TimeCardEvent) {
    m := &TimeCardEvent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.timeCardEvent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTimeCardEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeCardEvent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEvent) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAtApprovedLocation gets the atApprovedLocation property value. Indicates whether the entry was recorded at the approved location.
func (m *TimeCardEvent) GetAtApprovedLocation()(*bool) {
    return m.atApprovedLocation
}
// GetDateTime gets the dateTime property value. The time the entry is recorded.
func (m *TimeCardEvent) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCardEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["atApprovedLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAtApprovedLocation)
    res["dateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDateTime)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemBodyFromDiscriminatorValue , m.SetNotes)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetNotes gets the notes property value. Notes about the timeCardEvent.
func (m *TimeCardEvent) GetNotes()(ItemBodyable) {
    return m.notes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TimeCardEvent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TimeCardEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("atApprovedLocation", m.GetAtApprovedLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardEvent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAtApprovedLocation sets the atApprovedLocation property value. Indicates whether the entry was recorded at the approved location.
func (m *TimeCardEvent) SetAtApprovedLocation(value *bool)() {
    m.atApprovedLocation = value
}
// SetDateTime sets the dateTime property value. The time the entry is recorded.
func (m *TimeCardEvent) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// SetNotes sets the notes property value. Notes about the timeCardEvent.
func (m *TimeCardEvent) SetNotes(value ItemBodyable)() {
    m.notes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimeCardEvent) SetOdataType(value *string)() {
    m.odataType = value
}
