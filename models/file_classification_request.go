package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileClassificationRequest 
type FileClassificationRequest struct {
    Entity
    // The file property
    file []byte
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
}
// NewFileClassificationRequest instantiates a new FileClassificationRequest and sets the default values.
func NewFileClassificationRequest()(*FileClassificationRequest) {
    m := &FileClassificationRequest{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.fileClassificationRequest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFileClassificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileClassificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileClassificationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileClassificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["file"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetFile)
    res["sensitiveTypeIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSensitiveTypeIds)
    return res
}
// GetFile gets the file property value. The file property
func (m *FileClassificationRequest) GetFile()([]byte) {
    return m.file
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *FileClassificationRequest) GetSensitiveTypeIds()([]string) {
    return m.sensitiveTypeIds
}
// Serialize serializes information the current object
func (m *FileClassificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFile sets the file property value. The file property
func (m *FileClassificationRequest) SetFile(value []byte)() {
    m.file = value
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *FileClassificationRequest) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
