package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchUploadAgent 
type ExactMatchUploadAgent struct {
    Entity
    // The creationDateTime property
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
}
// NewExactMatchUploadAgent instantiates a new exactMatchUploadAgent and sets the default values.
func NewExactMatchUploadAgent()(*ExactMatchUploadAgent) {
    m := &ExactMatchUploadAgent{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.exactMatchUploadAgent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExactMatchUploadAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchUploadAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchUploadAgent(), nil
}
// GetCreationDateTime gets the creationDateTime property value. The creationDateTime property
func (m *ExactMatchUploadAgent) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetDescription gets the description property value. The description property
func (m *ExactMatchUploadAgent) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchUploadAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["creationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreationDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    return res
}
// Serialize serializes information the current object
func (m *ExactMatchUploadAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreationDateTime sets the creationDateTime property value. The creationDateTime property
func (m *ExactMatchUploadAgent) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetDescription sets the description property value. The description property
func (m *ExactMatchUploadAgent) SetDescription(value *string)() {
    m.description = value
}
