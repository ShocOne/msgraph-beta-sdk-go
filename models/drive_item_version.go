package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveItemVersion 
type DriveItemVersion struct {
    BaseItemVersion
    // The content property
    content []byte
    // Indicates the size of the content stream for this version of the item.
    size *int64
}
// NewDriveItemVersion instantiates a new DriveItemVersion and sets the default values.
func NewDriveItemVersion()(*DriveItemVersion) {
    m := &DriveItemVersion{
        BaseItemVersion: *NewBaseItemVersion(),
    }
    odataTypeValue := "#microsoft.graph.driveItemVersion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDriveItemVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveItemVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveItemVersion(), nil
}
// GetContent gets the content property value. The content property
func (m *DriveItemVersion) GetContent()([]byte) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItemVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItemVersion.GetFieldDeserializers()
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetContent)
    res["size"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetSize)
    return res
}
// GetSize gets the size property value. Indicates the size of the content stream for this version of the item.
func (m *DriveItemVersion) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *DriveItemVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItemVersion.Serialize(writer)
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
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *DriveItemVersion) SetContent(value []byte)() {
    m.content = value
}
// SetSize sets the size property value. Indicates the size of the content stream for this version of the item.
func (m *DriveItemVersion) SetSize(value *int64)() {
    m.size = value
}
