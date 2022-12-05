package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody provides operations to call the generateServerLogCollectionRequest method.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody instantiates a new DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) {
    m := &DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestPostRequestBody) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
