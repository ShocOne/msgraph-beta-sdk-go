package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserTrainingStatusInfo 
type UserTrainingStatusInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Date and time of assignment of the training to the user.
    assignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time of completion of the training by the user.
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Display name of the assigned training.
    displayName *string
    // The OdataType property
    odataType *string
    // Status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, unknownFutureValue.
    trainingStatus *TrainingStatus
}
// NewUserTrainingStatusInfo instantiates a new userTrainingStatusInfo and sets the default values.
func NewUserTrainingStatusInfo()(*UserTrainingStatusInfo) {
    m := &UserTrainingStatusInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userTrainingStatusInfo";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserTrainingStatusInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserTrainingStatusInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserTrainingStatusInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserTrainingStatusInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignedDateTime gets the assignedDateTime property value. Date and time of assignment of the training to the user.
func (m *UserTrainingStatusInfo) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.assignedDateTime
}
// GetCompletionDateTime gets the completionDateTime property value. Date and time of completion of the training by the user.
func (m *UserTrainingStatusInfo) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completionDateTime
}
// GetDisplayName gets the displayName property value. Display name of the assigned training.
func (m *UserTrainingStatusInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserTrainingStatusInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetAssignedDateTime)
    res["completionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompletionDateTime)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["trainingStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTrainingStatus , m.SetTrainingStatus)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserTrainingStatusInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetTrainingStatus gets the trainingStatus property value. Status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, unknownFutureValue.
func (m *UserTrainingStatusInfo) GetTrainingStatus()(*TrainingStatus) {
    return m.trainingStatus
}
// Serialize serializes information the current object
func (m *UserTrainingStatusInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("assignedDateTime", m.GetAssignedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetTrainingStatus() != nil {
        cast := (*m.GetTrainingStatus()).String()
        err := writer.WriteStringValue("trainingStatus", &cast)
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
func (m *UserTrainingStatusInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignedDateTime sets the assignedDateTime property value. Date and time of assignment of the training to the user.
func (m *UserTrainingStatusInfo) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignedDateTime = value
}
// SetCompletionDateTime sets the completionDateTime property value. Date and time of completion of the training by the user.
func (m *UserTrainingStatusInfo) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completionDateTime = value
}
// SetDisplayName sets the displayName property value. Display name of the assigned training.
func (m *UserTrainingStatusInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserTrainingStatusInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTrainingStatus sets the trainingStatus property value. Status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, unknownFutureValue.
func (m *UserTrainingStatusInfo) SetTrainingStatus(value *TrainingStatus)() {
    m.trainingStatus = value
}
