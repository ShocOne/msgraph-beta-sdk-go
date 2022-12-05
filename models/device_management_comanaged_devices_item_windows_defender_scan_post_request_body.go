package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody provides operations to call the windowsDefenderScan method.
type DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The quickScan property
    quickScan *bool
}
// NewDeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody instantiates a new DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody and sets the default values.
func NewDeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody()(*DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody) {
    m := &DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quickScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickScan(val)
        }
        return nil
    }
    return res
}
// GetQuickScan gets the quickScan property value. The quickScan property
func (m *DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody) GetQuickScan()(*bool) {
    return m.quickScan
}
// Serialize serializes information the current object
func (m *DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("quickScan", m.GetQuickScan())
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
func (m *DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetQuickScan sets the quickScan property value. The quickScan property
func (m *DeviceManagementComanagedDevicesItemWindowsDefenderScanPostRequestBody) SetQuickScan(value *bool)() {
    m.quickScan = value
}
