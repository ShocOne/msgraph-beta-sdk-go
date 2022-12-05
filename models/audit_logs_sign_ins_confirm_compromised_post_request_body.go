package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditLogsSignInsConfirmCompromisedPostRequestBody provides operations to call the confirmCompromised method.
type AuditLogsSignInsConfirmCompromisedPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The requestIds property
    requestIds []string
}
// NewAuditLogsSignInsConfirmCompromisedPostRequestBody instantiates a new AuditLogsSignInsConfirmCompromisedPostRequestBody and sets the default values.
func NewAuditLogsSignInsConfirmCompromisedPostRequestBody()(*AuditLogsSignInsConfirmCompromisedPostRequestBody) {
    m := &AuditLogsSignInsConfirmCompromisedPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuditLogsSignInsConfirmCompromisedPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditLogsSignInsConfirmCompromisedPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogsSignInsConfirmCompromisedPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditLogsSignInsConfirmCompromisedPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditLogsSignInsConfirmCompromisedPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["requestIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRequestIds(res)
        }
        return nil
    }
    return res
}
// GetRequestIds gets the requestIds property value. The requestIds property
func (m *AuditLogsSignInsConfirmCompromisedPostRequestBody) GetRequestIds()([]string) {
    return m.requestIds
}
// Serialize serializes information the current object
func (m *AuditLogsSignInsConfirmCompromisedPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRequestIds() != nil {
        err := writer.WriteCollectionOfStringValues("requestIds", m.GetRequestIds())
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
func (m *AuditLogsSignInsConfirmCompromisedPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRequestIds sets the requestIds property value. The requestIds property
func (m *AuditLogsSignInsConfirmCompromisedPostRequestBody) SetRequestIds(value []string)() {
    m.requestIds = value
}
