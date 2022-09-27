package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TelecomExpenseManagementPartner telecomExpenseManagementPartner resources represent the metadata and status of a given TEM service. Once your organization has onboarded with a partner, the partner can be enabled or disabled to switch TEM functionality on or off.
type TelecomExpenseManagementPartner struct {
    Entity
    // Whether the partner's AAD app has been authorized to access Intune.
    appAuthorized *bool
    // Display name of the TEM partner.
    displayName *string
    // Whether Intune's connection to the TEM service is currently enabled or disabled.
    enabled *bool
    // Timestamp of the last request sent to Intune by the TEM partner.
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
    url *string
}
// NewTelecomExpenseManagementPartner instantiates a new telecomExpenseManagementPartner and sets the default values.
func NewTelecomExpenseManagementPartner()(*TelecomExpenseManagementPartner) {
    m := &TelecomExpenseManagementPartner{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.telecomExpenseManagementPartner";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTelecomExpenseManagementPartnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTelecomExpenseManagementPartnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTelecomExpenseManagementPartner(), nil
}
// GetAppAuthorized gets the appAuthorized property value. Whether the partner's AAD app has been authorized to access Intune.
func (m *TelecomExpenseManagementPartner) GetAppAuthorized()(*bool) {
    return m.appAuthorized
}
// GetDisplayName gets the displayName property value. Display name of the TEM partner.
func (m *TelecomExpenseManagementPartner) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnabled gets the enabled property value. Whether Intune's connection to the TEM service is currently enabled or disabled.
func (m *TelecomExpenseManagementPartner) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TelecomExpenseManagementPartner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appAuthorized"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAppAuthorized)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["enabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnabled)
    res["lastConnectionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastConnectionDateTime)
    res["url"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUrl)
    return res
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *TelecomExpenseManagementPartner) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastConnectionDateTime
}
// GetUrl gets the url property value. URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
func (m *TelecomExpenseManagementPartner) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *TelecomExpenseManagementPartner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appAuthorized", m.GetAppAuthorized())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastConnectionDateTime", m.GetLastConnectionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppAuthorized sets the appAuthorized property value. Whether the partner's AAD app has been authorized to access Intune.
func (m *TelecomExpenseManagementPartner) SetAppAuthorized(value *bool)() {
    m.appAuthorized = value
}
// SetDisplayName sets the displayName property value. Display name of the TEM partner.
func (m *TelecomExpenseManagementPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnabled sets the enabled property value. Whether Intune's connection to the TEM service is currently enabled or disabled.
func (m *TelecomExpenseManagementPartner) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *TelecomExpenseManagementPartner) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectionDateTime = value
}
// SetUrl sets the url property value. URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
func (m *TelecomExpenseManagementPartner) SetUrl(value *string)() {
    m.url = value
}
