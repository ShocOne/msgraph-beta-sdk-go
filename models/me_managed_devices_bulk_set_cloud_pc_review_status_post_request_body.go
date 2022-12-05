package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody provides operations to call the bulkSetCloudPcReviewStatus method.
type MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managedDeviceIds property
    managedDeviceIds []string
    // The reviewStatus property
    reviewStatus CloudPcReviewStatusable
}
// NewMeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody instantiates a new MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody and sets the default values.
func NewMeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody()(*MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) {
    m := &MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedDeviceIds(res)
        }
        return nil
    }
    res["reviewStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcReviewStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewStatus(val.(CloudPcReviewStatusable))
        }
        return nil
    }
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// GetReviewStatus gets the reviewStatus property value. The reviewStatus property
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) GetReviewStatus()(CloudPcReviewStatusable) {
    return m.reviewStatus
}
// Serialize serializes information the current object
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reviewStatus", m.GetReviewStatus())
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
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
// SetReviewStatus sets the reviewStatus property value. The reviewStatus property
func (m *MeManagedDevicesBulkSetCloudPcReviewStatusPostRequestBody) SetReviewStatus(value CloudPcReviewStatusable)() {
    m.reviewStatus = value
}
