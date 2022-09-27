package decryptbuffer

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DecryptBufferPostRequestBody provides operations to call the decryptBuffer method.
type DecryptBufferPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The encryptedBuffer property
    encryptedBuffer []byte
    // The publishingLicense property
    publishingLicense []byte
}
// NewDecryptBufferPostRequestBody instantiates a new decryptBufferPostRequestBody and sets the default values.
func NewDecryptBufferPostRequestBody()(*DecryptBufferPostRequestBody) {
    m := &DecryptBufferPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDecryptBufferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDecryptBufferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDecryptBufferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DecryptBufferPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEncryptedBuffer gets the encryptedBuffer property value. The encryptedBuffer property
func (m *DecryptBufferPostRequestBody) GetEncryptedBuffer()([]byte) {
    return m.encryptedBuffer
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DecryptBufferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encryptedBuffer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetEncryptedBuffer)
    res["publishingLicense"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetPublishingLicense)
    return res
}
// GetPublishingLicense gets the publishingLicense property value. The publishingLicense property
func (m *DecryptBufferPostRequestBody) GetPublishingLicense()([]byte) {
    return m.publishingLicense
}
// Serialize serializes information the current object
func (m *DecryptBufferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("encryptedBuffer", m.GetEncryptedBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("publishingLicense", m.GetPublishingLicense())
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
func (m *DecryptBufferPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEncryptedBuffer sets the encryptedBuffer property value. The encryptedBuffer property
func (m *DecryptBufferPostRequestBody) SetEncryptedBuffer(value []byte)() {
    m.encryptedBuffer = value
}
// SetPublishingLicense sets the publishingLicense property value. The publishingLicense property
func (m *DecryptBufferPostRequestBody) SetPublishingLicense(value []byte)() {
    m.publishingLicense = value
}
