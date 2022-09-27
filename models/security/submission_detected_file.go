package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubmissionDetectedFile 
type SubmissionDetectedFile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The file hash.
    fileHash *string
    // The file name.
    fileName *string
    // The OdataType property
    odataType *string
}
// NewSubmissionDetectedFile instantiates a new submissionDetectedFile and sets the default values.
func NewSubmissionDetectedFile()(*SubmissionDetectedFile) {
    m := &SubmissionDetectedFile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.security.submissionDetectedFile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSubmissionDetectedFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubmissionDetectedFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubmissionDetectedFile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubmissionDetectedFile) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubmissionDetectedFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fileHash"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileHash)
    res["fileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetFileHash gets the fileHash property value. The file hash.
func (m *SubmissionDetectedFile) GetFileHash()(*string) {
    return m.fileHash
}
// GetFileName gets the fileName property value. The file name.
func (m *SubmissionDetectedFile) GetFileName()(*string) {
    return m.fileName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SubmissionDetectedFile) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SubmissionDetectedFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fileHash", m.GetFileHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
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
func (m *SubmissionDetectedFile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFileHash sets the fileHash property value. The file hash.
func (m *SubmissionDetectedFile) SetFileHash(value *string)() {
    m.fileHash = value
}
// SetFileName sets the fileName property value. The file name.
func (m *SubmissionDetectedFile) SetFileName(value *string)() {
    m.fileName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SubmissionDetectedFile) SetOdataType(value *string)() {
    m.odataType = value
}
