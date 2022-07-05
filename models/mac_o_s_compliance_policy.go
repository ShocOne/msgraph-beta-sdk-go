package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSCompliancePolicy 
type MacOSCompliancePolicy struct {
    DeviceCompliancePolicy
    // MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
    advancedThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Require that devices have enabled device threat protection.
    deviceThreatProtectionEnabled *bool
    // Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
    deviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Corresponds to the 'Block all incoming connections' option.
    firewallBlockAllIncoming *bool
    // Whether the firewall should be enabled or not.
    firewallEnabled *bool
    // Corresponds to 'Enable stealth mode.'
    firewallEnableStealthMode *bool
    // System and Privacy setting that determines which download locations apps can be run from on a macOS device. Possible values are: notConfigured, macAppStore, macAppStoreAndIdentifiedDevelopers, anywhere.
    gatekeeperAllowedAppSource *MacOSGatekeeperAppSources
    // Maximum MacOS build version.
    osMaximumBuildVersion *string
    // Maximum MacOS version.
    osMaximumVersion *string
    // Minimum MacOS build version.
    osMinimumBuildVersion *string
    // Minimum MacOS version.
    osMinimumVersion *string
    // Indicates whether or not to block simple passwords.
    passwordBlockSimple *bool
    // Number of days before the password expires. Valid values 1 to 65535
    passwordExpirationDays *int32
    // The number of character sets required in the password.
    passwordMinimumCharacterSetCount *int32
    // Minimum length of password. Valid values 4 to 14
    passwordMinimumLength *int32
    // Minutes of inactivity before a password is required.
    passwordMinutesOfInactivityBeforeLock *int32
    // Number of previous passwords to block. Valid values 1 to 24
    passwordPreviousPasswordBlockCount *int32
    // Whether or not to require a password.
    passwordRequired *bool
    // The required password type. Possible values are: deviceDefault, alphanumeric, numeric.
    passwordRequiredType *RequiredPasswordType
    // Require encryption on Mac OS devices.
    storageRequireEncryption *bool
    // Require that devices have enabled system integrity protection.
    systemIntegrityProtectionEnabled *bool
}
// NewMacOSCompliancePolicy instantiates a new MacOSCompliancePolicy and sets the default values.
func NewMacOSCompliancePolicy()(*MacOSCompliancePolicy) {
    m := &MacOSCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    return m
}
// CreateMacOSCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSCompliancePolicy(), nil
}
// GetAdvancedThreatProtectionRequiredSecurityLevel gets the advancedThreatProtectionRequiredSecurityLevel property value. MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *MacOSCompliancePolicy) GetAdvancedThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    if m == nil {
        return nil
    } else {
        return m.advancedThreatProtectionRequiredSecurityLevel
    }
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *MacOSCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceThreatProtectionEnabled
    }
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *MacOSCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    if m == nil {
        return nil
    } else {
        return m.deviceThreatProtectionRequiredSecurityLevel
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["advancedThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["deviceThreatProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionEnabled(val)
        }
        return nil
    }
    res["deviceThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["firewallBlockAllIncoming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallBlockAllIncoming(val)
        }
        return nil
    }
    res["firewallEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallEnabled(val)
        }
        return nil
    }
    res["firewallEnableStealthMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallEnableStealthMode(val)
        }
        return nil
    }
    res["gatekeeperAllowedAppSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSGatekeeperAppSources)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGatekeeperAllowedAppSource(val.(*MacOSGatekeeperAppSources))
        }
        return nil
    }
    res["osMaximumBuildVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumBuildVersion(val)
        }
        return nil
    }
    res["osMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumVersion(val)
        }
        return nil
    }
    res["osMinimumBuildVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumBuildVersion(val)
        }
        return nil
    }
    res["osMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumVersion(val)
        }
        return nil
    }
    res["passwordBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockSimple(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumCharacterSetCount(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passwordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["passwordRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequired(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*RequiredPasswordType))
        }
        return nil
    }
    res["storageRequireEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRequireEncryption(val)
        }
        return nil
    }
    res["systemIntegrityProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemIntegrityProtectionEnabled(val)
        }
        return nil
    }
    return res
}
// GetFirewallBlockAllIncoming gets the firewallBlockAllIncoming property value. Corresponds to the 'Block all incoming connections' option.
func (m *MacOSCompliancePolicy) GetFirewallBlockAllIncoming()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallBlockAllIncoming
    }
}
// GetFirewallEnabled gets the firewallEnabled property value. Whether the firewall should be enabled or not.
func (m *MacOSCompliancePolicy) GetFirewallEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallEnabled
    }
}
// GetFirewallEnableStealthMode gets the firewallEnableStealthMode property value. Corresponds to 'Enable stealth mode.'
func (m *MacOSCompliancePolicy) GetFirewallEnableStealthMode()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallEnableStealthMode
    }
}
// GetGatekeeperAllowedAppSource gets the gatekeeperAllowedAppSource property value. System and Privacy setting that determines which download locations apps can be run from on a macOS device. Possible values are: notConfigured, macAppStore, macAppStoreAndIdentifiedDevelopers, anywhere.
func (m *MacOSCompliancePolicy) GetGatekeeperAllowedAppSource()(*MacOSGatekeeperAppSources) {
    if m == nil {
        return nil
    } else {
        return m.gatekeeperAllowedAppSource
    }
}
// GetOsMaximumBuildVersion gets the osMaximumBuildVersion property value. Maximum MacOS build version.
func (m *MacOSCompliancePolicy) GetOsMaximumBuildVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMaximumBuildVersion
    }
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum MacOS version.
func (m *MacOSCompliancePolicy) GetOsMaximumVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMaximumVersion
    }
}
// GetOsMinimumBuildVersion gets the osMinimumBuildVersion property value. Minimum MacOS build version.
func (m *MacOSCompliancePolicy) GetOsMinimumBuildVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMinimumBuildVersion
    }
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum MacOS version.
func (m *MacOSCompliancePolicy) GetOsMinimumVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osMinimumVersion
    }
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Indicates whether or not to block simple passwords.
func (m *MacOSCompliancePolicy) GetPasswordBlockSimple()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockSimple
    }
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 65535
func (m *MacOSCompliancePolicy) GetPasswordExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordExpirationDays
    }
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *MacOSCompliancePolicy) GetPasswordMinimumCharacterSetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinimumCharacterSetCount
    }
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of password. Valid values 4 to 14
func (m *MacOSCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinimumLength
    }
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *MacOSCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinutesOfInactivityBeforeLock
    }
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *MacOSCompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordPreviousPasswordBlockCount
    }
}
// GetPasswordRequired gets the passwordRequired property value. Whether or not to require a password.
func (m *MacOSCompliancePolicy) GetPasswordRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordRequired
    }
}
// GetPasswordRequiredType gets the passwordRequiredType property value. The required password type. Possible values are: deviceDefault, alphanumeric, numeric.
func (m *MacOSCompliancePolicy) GetPasswordRequiredType()(*RequiredPasswordType) {
    if m == nil {
        return nil
    } else {
        return m.passwordRequiredType
    }
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Mac OS devices.
func (m *MacOSCompliancePolicy) GetStorageRequireEncryption()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.storageRequireEncryption
    }
}
// GetSystemIntegrityProtectionEnabled gets the systemIntegrityProtectionEnabled property value. Require that devices have enabled system integrity protection.
func (m *MacOSCompliancePolicy) GetSystemIntegrityProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.systemIntegrityProtectionEnabled
    }
}
// Serialize serializes information the current object
func (m *MacOSCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetAdvancedThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("advancedThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceThreatProtectionEnabled", m.GetDeviceThreatProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetDeviceThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("deviceThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallBlockAllIncoming", m.GetFirewallBlockAllIncoming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallEnabled", m.GetFirewallEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallEnableStealthMode", m.GetFirewallEnableStealthMode())
        if err != nil {
            return err
        }
    }
    if m.GetGatekeeperAllowedAppSource() != nil {
        cast := (*m.GetGatekeeperAllowedAppSource()).String()
        err = writer.WriteStringValue("gatekeeperAllowedAppSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMaximumBuildVersion", m.GetOsMaximumBuildVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMaximumVersion", m.GetOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumBuildVersion", m.GetOsMinimumBuildVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumVersion", m.GetOsMinimumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockSimple", m.GetPasswordBlockSimple())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumCharacterSetCount", m.GetPasswordMinimumCharacterSetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeLock", m.GetPasswordMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordRequired", m.GetPasswordRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRequireEncryption", m.GetStorageRequireEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("systemIntegrityProtectionEnabled", m.GetSystemIntegrityProtectionEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedThreatProtectionRequiredSecurityLevel sets the advancedThreatProtectionRequiredSecurityLevel property value. MDATP Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *MacOSCompliancePolicy) SetAdvancedThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    if m != nil {
        m.advancedThreatProtectionRequiredSecurityLevel = value
    }
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *MacOSCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    if m != nil {
        m.deviceThreatProtectionEnabled = value
    }
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Require Mobile Threat Protection minimum risk level to report noncompliance. Possible values are: unavailable, secured, low, medium, high, notSet.
func (m *MacOSCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    if m != nil {
        m.deviceThreatProtectionRequiredSecurityLevel = value
    }
}
// SetFirewallBlockAllIncoming sets the firewallBlockAllIncoming property value. Corresponds to the 'Block all incoming connections' option.
func (m *MacOSCompliancePolicy) SetFirewallBlockAllIncoming(value *bool)() {
    if m != nil {
        m.firewallBlockAllIncoming = value
    }
}
// SetFirewallEnabled sets the firewallEnabled property value. Whether the firewall should be enabled or not.
func (m *MacOSCompliancePolicy) SetFirewallEnabled(value *bool)() {
    if m != nil {
        m.firewallEnabled = value
    }
}
// SetFirewallEnableStealthMode sets the firewallEnableStealthMode property value. Corresponds to 'Enable stealth mode.'
func (m *MacOSCompliancePolicy) SetFirewallEnableStealthMode(value *bool)() {
    if m != nil {
        m.firewallEnableStealthMode = value
    }
}
// SetGatekeeperAllowedAppSource sets the gatekeeperAllowedAppSource property value. System and Privacy setting that determines which download locations apps can be run from on a macOS device. Possible values are: notConfigured, macAppStore, macAppStoreAndIdentifiedDevelopers, anywhere.
func (m *MacOSCompliancePolicy) SetGatekeeperAllowedAppSource(value *MacOSGatekeeperAppSources)() {
    if m != nil {
        m.gatekeeperAllowedAppSource = value
    }
}
// SetOsMaximumBuildVersion sets the osMaximumBuildVersion property value. Maximum MacOS build version.
func (m *MacOSCompliancePolicy) SetOsMaximumBuildVersion(value *string)() {
    if m != nil {
        m.osMaximumBuildVersion = value
    }
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum MacOS version.
func (m *MacOSCompliancePolicy) SetOsMaximumVersion(value *string)() {
    if m != nil {
        m.osMaximumVersion = value
    }
}
// SetOsMinimumBuildVersion sets the osMinimumBuildVersion property value. Minimum MacOS build version.
func (m *MacOSCompliancePolicy) SetOsMinimumBuildVersion(value *string)() {
    if m != nil {
        m.osMinimumBuildVersion = value
    }
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum MacOS version.
func (m *MacOSCompliancePolicy) SetOsMinimumVersion(value *string)() {
    if m != nil {
        m.osMinimumVersion = value
    }
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Indicates whether or not to block simple passwords.
func (m *MacOSCompliancePolicy) SetPasswordBlockSimple(value *bool)() {
    if m != nil {
        m.passwordBlockSimple = value
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 65535
func (m *MacOSCompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    if m != nil {
        m.passwordExpirationDays = value
    }
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *MacOSCompliancePolicy) SetPasswordMinimumCharacterSetCount(value *int32)() {
    if m != nil {
        m.passwordMinimumCharacterSetCount = value
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of password. Valid values 4 to 14
func (m *MacOSCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    if m != nil {
        m.passwordMinimumLength = value
    }
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *MacOSCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    if m != nil {
        m.passwordMinutesOfInactivityBeforeLock = value
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *MacOSCompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    if m != nil {
        m.passwordPreviousPasswordBlockCount = value
    }
}
// SetPasswordRequired sets the passwordRequired property value. Whether or not to require a password.
func (m *MacOSCompliancePolicy) SetPasswordRequired(value *bool)() {
    if m != nil {
        m.passwordRequired = value
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. The required password type. Possible values are: deviceDefault, alphanumeric, numeric.
func (m *MacOSCompliancePolicy) SetPasswordRequiredType(value *RequiredPasswordType)() {
    if m != nil {
        m.passwordRequiredType = value
    }
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Mac OS devices.
func (m *MacOSCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    if m != nil {
        m.storageRequireEncryption = value
    }
}
// SetSystemIntegrityProtectionEnabled sets the systemIntegrityProtectionEnabled property value. Require that devices have enabled system integrity protection.
func (m *MacOSCompliancePolicy) SetSystemIntegrityProtectionEnabled(value *bool)() {
    if m != nil {
        m.systemIntegrityProtectionEnabled = value
    }
}