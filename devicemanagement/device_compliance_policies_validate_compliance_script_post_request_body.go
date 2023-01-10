package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody 
type DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceCompliancePolicyScript property
    deviceCompliancePolicyScript ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyScriptable
}
// NewDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody instantiates a new DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody and sets the default values.
func NewDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody()(*DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) {
    m := &DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceCompliancePoliciesValidateComplianceScriptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesValidateComplianceScriptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePoliciesValidateComplianceScriptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceCompliancePolicyScript gets the deviceCompliancePolicyScript property value. The deviceCompliancePolicyScript property
func (m *DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) GetDeviceCompliancePolicyScript()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyScriptable) {
    return m.deviceCompliancePolicyScript
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceCompliancePolicyScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyScript(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyScriptable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceCompliancePolicyScript", m.GetDeviceCompliancePolicyScript())
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
func (m *DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceCompliancePolicyScript sets the deviceCompliancePolicyScript property value. The deviceCompliancePolicyScript property
func (m *DeviceCompliancePoliciesValidateComplianceScriptPostRequestBody) SetDeviceCompliancePolicyScript(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyScriptable)() {
    m.deviceCompliancePolicyScript = value
}
