package getexpiringvpptokencountwithexpiringbeforedatetime

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse provides operations to call the getExpiringVppTokenCount method.
type GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value *int32
}
// NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse instantiates a new getExpiringVppTokenCountWithExpiringBeforeDateTimeResponse and sets the default values.
func NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse()(*GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse) {
    m := &GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("value", m.GetValue())
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
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeResponse) SetValue(value *int32)() {
    m.value = value
}
