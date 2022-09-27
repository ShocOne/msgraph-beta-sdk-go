package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteResource 
type OnenoteResource struct {
    OnenoteEntityBaseModel
    // The content property
    content []byte
    // The contentUrl property
    contentUrl *string
}
// NewOnenoteResource instantiates a new OnenoteResource and sets the default values.
func NewOnenoteResource()(*OnenoteResource) {
    m := &OnenoteResource{
        OnenoteEntityBaseModel: *NewOnenoteEntityBaseModel(),
    }
    odataTypeValue := "#microsoft.graph.onenoteResource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOnenoteResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenoteResource(), nil
}
// GetContent gets the content property value. The content property
func (m *OnenoteResource) GetContent()([]byte) {
    return m.content
}
// GetContentUrl gets the contentUrl property value. The contentUrl property
func (m *OnenoteResource) GetContentUrl()(*string) {
    return m.contentUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnenoteEntityBaseModel.GetFieldDeserializers()
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetContent)
    res["contentUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentUrl)
    return res
}
// Serialize serializes information the current object
func (m *OnenoteResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnenoteEntityBaseModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *OnenoteResource) SetContent(value []byte)() {
    m.content = value
}
// SetContentUrl sets the contentUrl property value. The contentUrl property
func (m *OnenoteResource) SetContentUrl(value *string)() {
    m.contentUrl = value
}
