package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaPrompt 
type MediaPrompt struct {
    Prompt
    // The loop property
    loop *int32
    // The mediaInfo property
    mediaInfo MediaInfoable
}
// NewMediaPrompt instantiates a new MediaPrompt and sets the default values.
func NewMediaPrompt()(*MediaPrompt) {
    m := &MediaPrompt{
        Prompt: *NewPrompt(),
    }
    odataTypeValue := "#microsoft.graph.mediaPrompt";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMediaPromptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaPromptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaPrompt(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaPrompt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Prompt.GetFieldDeserializers()
    res["loop"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetLoop)
    res["mediaInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMediaInfoFromDiscriminatorValue , m.SetMediaInfo)
    return res
}
// GetLoop gets the loop property value. The loop property
func (m *MediaPrompt) GetLoop()(*int32) {
    return m.loop
}
// GetMediaInfo gets the mediaInfo property value. The mediaInfo property
func (m *MediaPrompt) GetMediaInfo()(MediaInfoable) {
    return m.mediaInfo
}
// Serialize serializes information the current object
func (m *MediaPrompt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Prompt.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("loop", m.GetLoop())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaInfo", m.GetMediaInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLoop sets the loop property value. The loop property
func (m *MediaPrompt) SetLoop(value *int32)() {
    m.loop = value
}
// SetMediaInfo sets the mediaInfo property value. The mediaInfo property
func (m *MediaPrompt) SetMediaInfo(value MediaInfoable)() {
    m.mediaInfo = value
}
