package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecordOperation 
type RecordOperation struct {
    CommsOperation
    // Possible values are: operationCanceled, stopToneDetected, maxRecordDurationReached, initialSilenceTimeout, maxSilenceTimeout, playPromptFailed, playBeepFailed, mediaReceiveTimeout, unspecifiedError, none.
    completionReason *RecordCompletionReason
    // The access token required to retrieve the recording.
    recordingAccessToken *string
    // The location where the recording is located.
    recordingLocation *string
}
// NewRecordOperation instantiates a new RecordOperation and sets the default values.
func NewRecordOperation()(*RecordOperation) {
    m := &RecordOperation{
        CommsOperation: *NewCommsOperation(),
    }
    odataTypeValue := "#microsoft.graph.recordOperation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRecordOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecordOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecordOperation(), nil
}
// GetCompletionReason gets the completionReason property value. Possible values are: operationCanceled, stopToneDetected, maxRecordDurationReached, initialSilenceTimeout, maxSilenceTimeout, playPromptFailed, playBeepFailed, mediaReceiveTimeout, unspecifiedError, none.
func (m *RecordOperation) GetCompletionReason()(*RecordCompletionReason) {
    return m.completionReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecordOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["completionReason"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRecordCompletionReason , m.SetCompletionReason)
    res["recordingAccessToken"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRecordingAccessToken)
    res["recordingLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRecordingLocation)
    return res
}
// GetRecordingAccessToken gets the recordingAccessToken property value. The access token required to retrieve the recording.
func (m *RecordOperation) GetRecordingAccessToken()(*string) {
    return m.recordingAccessToken
}
// GetRecordingLocation gets the recordingLocation property value. The location where the recording is located.
func (m *RecordOperation) GetRecordingLocation()(*string) {
    return m.recordingLocation
}
// Serialize serializes information the current object
func (m *RecordOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCompletionReason() != nil {
        cast := (*m.GetCompletionReason()).String()
        err = writer.WriteStringValue("completionReason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordingAccessToken", m.GetRecordingAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordingLocation", m.GetRecordingLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletionReason sets the completionReason property value. Possible values are: operationCanceled, stopToneDetected, maxRecordDurationReached, initialSilenceTimeout, maxSilenceTimeout, playPromptFailed, playBeepFailed, mediaReceiveTimeout, unspecifiedError, none.
func (m *RecordOperation) SetCompletionReason(value *RecordCompletionReason)() {
    m.completionReason = value
}
// SetRecordingAccessToken sets the recordingAccessToken property value. The access token required to retrieve the recording.
func (m *RecordOperation) SetRecordingAccessToken(value *string)() {
    m.recordingAccessToken = value
}
// SetRecordingLocation sets the recordingLocation property value. The location where the recording is located.
func (m *RecordOperation) SetRecordingLocation(value *string)() {
    m.recordingLocation = value
}
