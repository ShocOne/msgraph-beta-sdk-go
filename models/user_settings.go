package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserSettings struct {
    Entity
}
// NewUserSettings instantiates a new UserSettings and sets the default values.
func NewUserSettings()(*UserSettings) {
    m := &UserSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSettings(), nil
}
// GetContactMergeSuggestions gets the contactMergeSuggestions property value. The user's settings for the visibility of merge suggestion for the duplicate contacts in the user's contact list.
// returns a ContactMergeSuggestionsable when successful
func (m *UserSettings) GetContactMergeSuggestions()(ContactMergeSuggestionsable) {
    val, err := m.GetBackingStore().Get("contactMergeSuggestions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ContactMergeSuggestionsable)
    }
    return nil
}
// GetContributionToContentDiscoveryAsOrganizationDisabled gets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
// returns a *bool when successful
func (m *UserSettings) GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("contributionToContentDiscoveryAsOrganizationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContributionToContentDiscoveryDisabled gets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
// returns a *bool when successful
func (m *UserSettings) GetContributionToContentDiscoveryDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("contributionToContentDiscoveryDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExchange gets the exchange property value. The Exchange settings for mailbox discovery.
// returns a ExchangeSettingsable when successful
func (m *UserSettings) GetExchange()(ExchangeSettingsable) {
    val, err := m.GetBackingStore().Get("exchange")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ExchangeSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contactMergeSuggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContactMergeSuggestionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactMergeSuggestions(val.(ContactMergeSuggestionsable))
        }
        return nil
    }
    res["contributionToContentDiscoveryAsOrganizationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryAsOrganizationDisabled(val)
        }
        return nil
    }
    res["contributionToContentDiscoveryDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributionToContentDiscoveryDisabled(val)
        }
        return nil
    }
    res["exchange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExchangeSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchange(val.(ExchangeSettingsable))
        }
        return nil
    }
    res["itemInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(UserInsightsSettingsable))
        }
        return nil
    }
    res["regionalAndLanguageSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRegionalAndLanguageSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionalAndLanguageSettings(val.(RegionalAndLanguageSettingsable))
        }
        return nil
    }
    res["shiftPreferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShiftPreferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShiftPreferences(val.(ShiftPreferencesable))
        }
        return nil
    }
    res["storage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserStorageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorage(val.(UserStorageable))
        }
        return nil
    }
    res["windows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsSettingable)
                }
            }
            m.SetWindows(res)
        }
        return nil
    }
    return res
}
// GetItemInsights gets the itemInsights property value. The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
// returns a UserInsightsSettingsable when successful
func (m *UserSettings) GetItemInsights()(UserInsightsSettingsable) {
    val, err := m.GetBackingStore().Get("itemInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserInsightsSettingsable)
    }
    return nil
}
// GetRegionalAndLanguageSettings gets the regionalAndLanguageSettings property value. The user's preferences for languages, regional locale and date/time formatting.
// returns a RegionalAndLanguageSettingsable when successful
func (m *UserSettings) GetRegionalAndLanguageSettings()(RegionalAndLanguageSettingsable) {
    val, err := m.GetBackingStore().Get("regionalAndLanguageSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RegionalAndLanguageSettingsable)
    }
    return nil
}
// GetShiftPreferences gets the shiftPreferences property value. The shift preferences for the user.
// returns a ShiftPreferencesable when successful
func (m *UserSettings) GetShiftPreferences()(ShiftPreferencesable) {
    val, err := m.GetBackingStore().Get("shiftPreferences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ShiftPreferencesable)
    }
    return nil
}
// GetStorage gets the storage property value. The storage property
// returns a UserStorageable when successful
func (m *UserSettings) GetStorage()(UserStorageable) {
    val, err := m.GetBackingStore().Get("storage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserStorageable)
    }
    return nil
}
// GetWindows gets the windows property value. The Windows settings of the user stored in the cloud.
// returns a []WindowsSettingable when successful
func (m *UserSettings) GetWindows()([]WindowsSettingable) {
    val, err := m.GetBackingStore().Get("windows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contactMergeSuggestions", m.GetContactMergeSuggestions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryAsOrganizationDisabled", m.GetContributionToContentDiscoveryAsOrganizationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contributionToContentDiscoveryDisabled", m.GetContributionToContentDiscoveryDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("exchange", m.GetExchange())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("itemInsights", m.GetItemInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("regionalAndLanguageSettings", m.GetRegionalAndLanguageSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shiftPreferences", m.GetShiftPreferences())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("storage", m.GetStorage())
        if err != nil {
            return err
        }
    }
    if m.GetWindows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindows()))
        for i, v := range m.GetWindows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windows", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContactMergeSuggestions sets the contactMergeSuggestions property value. The user's settings for the visibility of merge suggestion for the duplicate contacts in the user's contact list.
func (m *UserSettings) SetContactMergeSuggestions(value ContactMergeSuggestionsable)() {
    err := m.GetBackingStore().Set("contactMergeSuggestions", value)
    if err != nil {
        panic(err)
    }
}
// SetContributionToContentDiscoveryAsOrganizationDisabled sets the contributionToContentDiscoveryAsOrganizationDisabled property value. Reflects the Office Delve organization level setting. When set to true, the organization doesn't have access to Office Delve. This setting is read-only and can only be changed by administrators in the SharePoint admin center.
func (m *UserSettings) SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("contributionToContentDiscoveryAsOrganizationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetContributionToContentDiscoveryDisabled sets the contributionToContentDiscoveryDisabled property value. When set to true, documents in the user's Office Delve are disabled. Users can control this setting in Office Delve.
func (m *UserSettings) SetContributionToContentDiscoveryDisabled(value *bool)() {
    err := m.GetBackingStore().Set("contributionToContentDiscoveryDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetExchange sets the exchange property value. The Exchange settings for mailbox discovery.
func (m *UserSettings) SetExchange(value ExchangeSettingsable)() {
    err := m.GetBackingStore().Set("exchange", value)
    if err != nil {
        panic(err)
    }
}
// SetItemInsights sets the itemInsights property value. The user's settings for the visibility of meeting hour insights, and insights derived between a user and other items in Microsoft 365, such as documents or sites. Get userInsightsSettings through this navigation property.
func (m *UserSettings) SetItemInsights(value UserInsightsSettingsable)() {
    err := m.GetBackingStore().Set("itemInsights", value)
    if err != nil {
        panic(err)
    }
}
// SetRegionalAndLanguageSettings sets the regionalAndLanguageSettings property value. The user's preferences for languages, regional locale and date/time formatting.
func (m *UserSettings) SetRegionalAndLanguageSettings(value RegionalAndLanguageSettingsable)() {
    err := m.GetBackingStore().Set("regionalAndLanguageSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetShiftPreferences sets the shiftPreferences property value. The shift preferences for the user.
func (m *UserSettings) SetShiftPreferences(value ShiftPreferencesable)() {
    err := m.GetBackingStore().Set("shiftPreferences", value)
    if err != nil {
        panic(err)
    }
}
// SetStorage sets the storage property value. The storage property
func (m *UserSettings) SetStorage(value UserStorageable)() {
    err := m.GetBackingStore().Set("storage", value)
    if err != nil {
        panic(err)
    }
}
// SetWindows sets the windows property value. The Windows settings of the user stored in the cloud.
func (m *UserSettings) SetWindows(value []WindowsSettingable)() {
    err := m.GetBackingStore().Set("windows", value)
    if err != nil {
        panic(err)
    }
}
type UserSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContactMergeSuggestions()(ContactMergeSuggestionsable)
    GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool)
    GetContributionToContentDiscoveryDisabled()(*bool)
    GetExchange()(ExchangeSettingsable)
    GetItemInsights()(UserInsightsSettingsable)
    GetRegionalAndLanguageSettings()(RegionalAndLanguageSettingsable)
    GetShiftPreferences()(ShiftPreferencesable)
    GetStorage()(UserStorageable)
    GetWindows()([]WindowsSettingable)
    SetContactMergeSuggestions(value ContactMergeSuggestionsable)()
    SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)()
    SetContributionToContentDiscoveryDisabled(value *bool)()
    SetExchange(value ExchangeSettingsable)()
    SetItemInsights(value UserInsightsSettingsable)()
    SetRegionalAndLanguageSettings(value RegionalAndLanguageSettingsable)()
    SetShiftPreferences(value ShiftPreferencesable)()
    SetStorage(value UserStorageable)()
    SetWindows(value []WindowsSettingable)()
}
