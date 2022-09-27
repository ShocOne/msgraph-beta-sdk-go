package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentSharingSession provides operations to manage the collection of activityStatistics entities.
type ContentSharingSession struct {
    Entity
}
// NewContentSharingSession instantiates a new contentSharingSession and sets the default values.
func NewContentSharingSession()(*ContentSharingSession) {
    m := &ContentSharingSession{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.contentSharingSession";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateContentSharingSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentSharingSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentSharingSession(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentSharingSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ContentSharingSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
