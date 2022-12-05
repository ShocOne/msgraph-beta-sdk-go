package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileGeneralDeviceConfiguration 
type AndroidWorkProfileGeneralDeviceConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to block face unlock.
    passwordBlockFaceUnlock *bool
    // Indicates whether or not to block fingerprint unlock.
    passwordBlockFingerprintUnlock *bool
    // Indicates whether or not to block iris unlock.
    passwordBlockIrisUnlock *bool
    // Indicates whether or not to block Smart Lock and other trust agents.
    passwordBlockTrustAgents *bool
    // Number of days before the password expires. Valid values 1 to 365
    passwordExpirationDays *int32
    // Minimum length of passwords. Valid values 4 to 16
    passwordMinimumLength *int32
    // Minutes of inactivity before the screen times out.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous passwords to block. Valid values 0 to 24
    passwordPreviousPasswordBlockCount *int32
    // Android Work Profile required password type.
    passwordRequiredType *AndroidWorkProfileRequiredPasswordType
    // Number of sign in failures allowed before factory reset. Valid values 1 to 16
    passwordSignInFailureCountBeforeFactoryReset *int32
    // The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
    requiredPasswordComplexity *AndroidRequiredPasswordComplexity
    // Require the Android Verify apps feature is turned on.
    securityRequireVerifyApps *bool
    // Enable lockdown mode for always-on VPN.
    vpnAlwaysOnPackageIdentifier *string
    // Enable lockdown mode for always-on VPN.
    vpnEnableAlwaysOnLockdownMode *bool
    // Indicates whether to allow installation of apps from unknown sources.
    workProfileAllowAppInstallsFromUnknownSources *bool
    // Allow widgets from work profile apps.
    workProfileAllowWidgets *bool
    // Block users from adding/removing accounts in work profile.
    workProfileBlockAddingAccounts *bool
    // Block work profile camera.
    workProfileBlockCamera *bool
    // Block display work profile caller ID in personal profile.
    workProfileBlockCrossProfileCallerId *bool
    // Block work profile contacts availability in personal profile.
    workProfileBlockCrossProfileContactsSearch *bool
    // Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
    workProfileBlockCrossProfileCopyPaste *bool
    // Indicates whether or not to block notifications while device locked.
    workProfileBlockNotificationsWhileDeviceLocked *bool
    // Prevent app installations from unknown sources in the personal profile.
    workProfileBlockPersonalAppInstallsFromUnknownSources *bool
    // Block screen capture in work profile.
    workProfileBlockScreenCapture *bool
    // Allow bluetooth devices to access enterprise contacts.
    workProfileBluetoothEnableContactSharing *bool
    // Android Work Profile cross profile data sharing type.
    workProfileDataSharingType *AndroidWorkProfileCrossProfileDataSharingType
    // Android Work Profile default app permission policy type.
    workProfileDefaultAppPermissionPolicy *AndroidWorkProfileDefaultAppPermissionPolicyType
    // Indicates whether or not to block face unlock for work profile.
    workProfilePasswordBlockFaceUnlock *bool
    // Indicates whether or not to block fingerprint unlock for work profile.
    workProfilePasswordBlockFingerprintUnlock *bool
    // Indicates whether or not to block iris unlock for work profile.
    workProfilePasswordBlockIrisUnlock *bool
    // Indicates whether or not to block Smart Lock and other trust agents for work profile.
    workProfilePasswordBlockTrustAgents *bool
    // Number of days before the work profile password expires. Valid values 1 to 365
    workProfilePasswordExpirationDays *int32
    // Minimum length of work profile password. Valid values 4 to 16
    workProfilePasswordMinimumLength *int32
    // Minimum # of letter characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinLetterCharacters *int32
    // Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinLowerCaseCharacters *int32
    // Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinNonLetterCharacters *int32
    // Minimum # of numeric characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinNumericCharacters *int32
    // Minimum # of symbols required in work profile password. Valid values 1 to 10
    workProfilePasswordMinSymbolCharacters *int32
    // Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
    workProfilePasswordMinUpperCaseCharacters *int32
    // Minutes of inactivity before the screen times out.
    workProfilePasswordMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous work profile passwords to block. Valid values 0 to 24
    workProfilePasswordPreviousPasswordBlockCount *int32
    // Android Work Profile required password type.
    workProfilePasswordRequiredType *AndroidWorkProfileRequiredPasswordType
    // Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
    workProfilePasswordSignInFailureCountBeforeFactoryReset *int32
    // The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
    workProfileRequiredPasswordComplexity *AndroidRequiredPasswordComplexity
    // Password is required or not for work profile
    workProfileRequirePassword *bool
}
// NewAndroidWorkProfileGeneralDeviceConfiguration instantiates a new AndroidWorkProfileGeneralDeviceConfiguration and sets the default values.
func NewAndroidWorkProfileGeneralDeviceConfiguration()(*AndroidWorkProfileGeneralDeviceConfiguration) {
    m := &AndroidWorkProfileGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileGeneralDeviceConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidWorkProfileGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileGeneralDeviceConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["passwordBlockFaceUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockFaceUnlock(val)
        }
        return nil
    }
    res["passwordBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["passwordBlockIrisUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockIrisUnlock(val)
        }
        return nil
    }
    res["passwordBlockTrustAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockTrustAgents(val)
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
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
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
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidWorkProfileRequiredPasswordType))
        }
        return nil
    }
    res["passwordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["requiredPasswordComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidRequiredPasswordComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPasswordComplexity(val.(*AndroidRequiredPasswordComplexity))
        }
        return nil
    }
    res["securityRequireVerifyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireVerifyApps(val)
        }
        return nil
    }
    res["vpnAlwaysOnPackageIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnAlwaysOnPackageIdentifier(val)
        }
        return nil
    }
    res["vpnEnableAlwaysOnLockdownMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnEnableAlwaysOnLockdownMode(val)
        }
        return nil
    }
    res["workProfileAllowAppInstallsFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileAllowAppInstallsFromUnknownSources(val)
        }
        return nil
    }
    res["workProfileAllowWidgets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileAllowWidgets(val)
        }
        return nil
    }
    res["workProfileBlockAddingAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockAddingAccounts(val)
        }
        return nil
    }
    res["workProfileBlockCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCamera(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileCallerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileCallerId(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileContactsSearch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileContactsSearch(val)
        }
        return nil
    }
    res["workProfileBlockCrossProfileCopyPaste"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockCrossProfileCopyPaste(val)
        }
        return nil
    }
    res["workProfileBlockNotificationsWhileDeviceLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockNotificationsWhileDeviceLocked(val)
        }
        return nil
    }
    res["workProfileBlockPersonalAppInstallsFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockPersonalAppInstallsFromUnknownSources(val)
        }
        return nil
    }
    res["workProfileBlockScreenCapture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBlockScreenCapture(val)
        }
        return nil
    }
    res["workProfileBluetoothEnableContactSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileBluetoothEnableContactSharing(val)
        }
        return nil
    }
    res["workProfileDataSharingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileCrossProfileDataSharingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileDataSharingType(val.(*AndroidWorkProfileCrossProfileDataSharingType))
        }
        return nil
    }
    res["workProfileDefaultAppPermissionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileDefaultAppPermissionPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileDefaultAppPermissionPolicy(val.(*AndroidWorkProfileDefaultAppPermissionPolicyType))
        }
        return nil
    }
    res["workProfilePasswordBlockFaceUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockFaceUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockIrisUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockIrisUnlock(val)
        }
        return nil
    }
    res["workProfilePasswordBlockTrustAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordBlockTrustAgents(val)
        }
        return nil
    }
    res["workProfilePasswordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordExpirationDays(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLength(val)
        }
        return nil
    }
    res["workProfilePasswordMinLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinLowerCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinNonLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinNumericCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinSymbolCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinUpperCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["workProfilePasswordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["workProfilePasswordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWorkProfileRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordRequiredType(val.(*AndroidWorkProfileRequiredPasswordType))
        }
        return nil
    }
    res["workProfilePasswordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["workProfileRequiredPasswordComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidRequiredPasswordComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileRequiredPasswordComplexity(val.(*AndroidRequiredPasswordComplexity))
        }
        return nil
    }
    res["workProfileRequirePassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfileRequirePassword(val)
        }
        return nil
    }
    return res
}
// GetPasswordBlockFaceUnlock gets the passwordBlockFaceUnlock property value. Indicates whether or not to block face unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFaceUnlock()(*bool) {
    return m.passwordBlockFaceUnlock
}
// GetPasswordBlockFingerprintUnlock gets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock()(*bool) {
    return m.passwordBlockFingerprintUnlock
}
// GetPasswordBlockIrisUnlock gets the passwordBlockIrisUnlock property value. Indicates whether or not to block iris unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockIrisUnlock()(*bool) {
    return m.passwordBlockIrisUnlock
}
// GetPasswordBlockTrustAgents gets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockTrustAgents()(*bool) {
    return m.passwordBlockTrustAgents
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passwordMinutesOfInactivityBeforeScreenTimeout
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordRequiredType()(*AndroidWorkProfileRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.passwordSignInFailureCountBeforeFactoryReset
}
// GetRequiredPasswordComplexity gets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity) {
    return m.requiredPasswordComplexity
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetSecurityRequireVerifyApps()(*bool) {
    return m.securityRequireVerifyApps
}
// GetVpnAlwaysOnPackageIdentifier gets the vpnAlwaysOnPackageIdentifier property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetVpnAlwaysOnPackageIdentifier()(*string) {
    return m.vpnAlwaysOnPackageIdentifier
}
// GetVpnEnableAlwaysOnLockdownMode gets the vpnEnableAlwaysOnLockdownMode property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetVpnEnableAlwaysOnLockdownMode()(*bool) {
    return m.vpnEnableAlwaysOnLockdownMode
}
// GetWorkProfileAllowAppInstallsFromUnknownSources gets the workProfileAllowAppInstallsFromUnknownSources property value. Indicates whether to allow installation of apps from unknown sources.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileAllowAppInstallsFromUnknownSources()(*bool) {
    return m.workProfileAllowAppInstallsFromUnknownSources
}
// GetWorkProfileAllowWidgets gets the workProfileAllowWidgets property value. Allow widgets from work profile apps.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileAllowWidgets()(*bool) {
    return m.workProfileAllowWidgets
}
// GetWorkProfileBlockAddingAccounts gets the workProfileBlockAddingAccounts property value. Block users from adding/removing accounts in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockAddingAccounts()(*bool) {
    return m.workProfileBlockAddingAccounts
}
// GetWorkProfileBlockCamera gets the workProfileBlockCamera property value. Block work profile camera.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCamera()(*bool) {
    return m.workProfileBlockCamera
}
// GetWorkProfileBlockCrossProfileCallerId gets the workProfileBlockCrossProfileCallerId property value. Block display work profile caller ID in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCallerId()(*bool) {
    return m.workProfileBlockCrossProfileCallerId
}
// GetWorkProfileBlockCrossProfileContactsSearch gets the workProfileBlockCrossProfileContactsSearch property value. Block work profile contacts availability in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileContactsSearch()(*bool) {
    return m.workProfileBlockCrossProfileContactsSearch
}
// GetWorkProfileBlockCrossProfileCopyPaste gets the workProfileBlockCrossProfileCopyPaste property value. Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCopyPaste()(*bool) {
    return m.workProfileBlockCrossProfileCopyPaste
}
// GetWorkProfileBlockNotificationsWhileDeviceLocked gets the workProfileBlockNotificationsWhileDeviceLocked property value. Indicates whether or not to block notifications while device locked.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockNotificationsWhileDeviceLocked()(*bool) {
    return m.workProfileBlockNotificationsWhileDeviceLocked
}
// GetWorkProfileBlockPersonalAppInstallsFromUnknownSources gets the workProfileBlockPersonalAppInstallsFromUnknownSources property value. Prevent app installations from unknown sources in the personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockPersonalAppInstallsFromUnknownSources()(*bool) {
    return m.workProfileBlockPersonalAppInstallsFromUnknownSources
}
// GetWorkProfileBlockScreenCapture gets the workProfileBlockScreenCapture property value. Block screen capture in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockScreenCapture()(*bool) {
    return m.workProfileBlockScreenCapture
}
// GetWorkProfileBluetoothEnableContactSharing gets the workProfileBluetoothEnableContactSharing property value. Allow bluetooth devices to access enterprise contacts.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBluetoothEnableContactSharing()(*bool) {
    return m.workProfileBluetoothEnableContactSharing
}
// GetWorkProfileDataSharingType gets the workProfileDataSharingType property value. Android Work Profile cross profile data sharing type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDataSharingType()(*AndroidWorkProfileCrossProfileDataSharingType) {
    return m.workProfileDataSharingType
}
// GetWorkProfileDefaultAppPermissionPolicy gets the workProfileDefaultAppPermissionPolicy property value. Android Work Profile default app permission policy type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDefaultAppPermissionPolicy()(*AndroidWorkProfileDefaultAppPermissionPolicyType) {
    return m.workProfileDefaultAppPermissionPolicy
}
// GetWorkProfilePasswordBlockFaceUnlock gets the workProfilePasswordBlockFaceUnlock property value. Indicates whether or not to block face unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFaceUnlock()(*bool) {
    return m.workProfilePasswordBlockFaceUnlock
}
// GetWorkProfilePasswordBlockFingerprintUnlock gets the workProfilePasswordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFingerprintUnlock()(*bool) {
    return m.workProfilePasswordBlockFingerprintUnlock
}
// GetWorkProfilePasswordBlockIrisUnlock gets the workProfilePasswordBlockIrisUnlock property value. Indicates whether or not to block iris unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockIrisUnlock()(*bool) {
    return m.workProfilePasswordBlockIrisUnlock
}
// GetWorkProfilePasswordBlockTrustAgents gets the workProfilePasswordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockTrustAgents()(*bool) {
    return m.workProfilePasswordBlockTrustAgents
}
// GetWorkProfilePasswordExpirationDays gets the workProfilePasswordExpirationDays property value. Number of days before the work profile password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDays()(*int32) {
    return m.workProfilePasswordExpirationDays
}
// GetWorkProfilePasswordMinimumLength gets the workProfilePasswordMinimumLength property value. Minimum length of work profile password. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLength()(*int32) {
    return m.workProfilePasswordMinimumLength
}
// GetWorkProfilePasswordMinLetterCharacters gets the workProfilePasswordMinLetterCharacters property value. Minimum # of letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLetterCharacters()(*int32) {
    return m.workProfilePasswordMinLetterCharacters
}
// GetWorkProfilePasswordMinLowerCaseCharacters gets the workProfilePasswordMinLowerCaseCharacters property value. Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLowerCaseCharacters()(*int32) {
    return m.workProfilePasswordMinLowerCaseCharacters
}
// GetWorkProfilePasswordMinNonLetterCharacters gets the workProfilePasswordMinNonLetterCharacters property value. Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNonLetterCharacters()(*int32) {
    return m.workProfilePasswordMinNonLetterCharacters
}
// GetWorkProfilePasswordMinNumericCharacters gets the workProfilePasswordMinNumericCharacters property value. Minimum # of numeric characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNumericCharacters()(*int32) {
    return m.workProfilePasswordMinNumericCharacters
}
// GetWorkProfilePasswordMinSymbolCharacters gets the workProfilePasswordMinSymbolCharacters property value. Minimum # of symbols required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinSymbolCharacters()(*int32) {
    return m.workProfilePasswordMinSymbolCharacters
}
// GetWorkProfilePasswordMinUpperCaseCharacters gets the workProfilePasswordMinUpperCaseCharacters property value. Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinUpperCaseCharacters()(*int32) {
    return m.workProfilePasswordMinUpperCaseCharacters
}
// GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout gets the workProfilePasswordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.workProfilePasswordMinutesOfInactivityBeforeScreenTimeout
}
// GetWorkProfilePasswordPreviousPasswordBlockCount gets the workProfilePasswordPreviousPasswordBlockCount property value. Number of previous work profile passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordBlockCount()(*int32) {
    return m.workProfilePasswordPreviousPasswordBlockCount
}
// GetWorkProfilePasswordRequiredType gets the workProfilePasswordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredType()(*AndroidWorkProfileRequiredPasswordType) {
    return m.workProfilePasswordRequiredType
}
// GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset gets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.workProfilePasswordSignInFailureCountBeforeFactoryReset
}
// GetWorkProfileRequiredPasswordComplexity gets the workProfileRequiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity) {
    return m.workProfileRequiredPasswordComplexity
}
// GetWorkProfileRequirePassword gets the workProfileRequirePassword property value. Password is required or not for work profile
func (m *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequirePassword()(*bool) {
    return m.workProfileRequirePassword
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("passwordBlockFaceUnlock", m.GetPasswordBlockFaceUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockFingerprintUnlock", m.GetPasswordBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockIrisUnlock", m.GetPasswordBlockIrisUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockTrustAgents", m.GetPasswordBlockTrustAgents())
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
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
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
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredPasswordComplexity() != nil {
        cast := (*m.GetRequiredPasswordComplexity()).String()
        err = writer.WriteStringValue("requiredPasswordComplexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireVerifyApps", m.GetSecurityRequireVerifyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vpnAlwaysOnPackageIdentifier", m.GetVpnAlwaysOnPackageIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("vpnEnableAlwaysOnLockdownMode", m.GetVpnEnableAlwaysOnLockdownMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileAllowAppInstallsFromUnknownSources", m.GetWorkProfileAllowAppInstallsFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileAllowWidgets", m.GetWorkProfileAllowWidgets())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockAddingAccounts", m.GetWorkProfileBlockAddingAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCamera", m.GetWorkProfileBlockCamera())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileCallerId", m.GetWorkProfileBlockCrossProfileCallerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileContactsSearch", m.GetWorkProfileBlockCrossProfileContactsSearch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockCrossProfileCopyPaste", m.GetWorkProfileBlockCrossProfileCopyPaste())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockNotificationsWhileDeviceLocked", m.GetWorkProfileBlockNotificationsWhileDeviceLocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockPersonalAppInstallsFromUnknownSources", m.GetWorkProfileBlockPersonalAppInstallsFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBlockScreenCapture", m.GetWorkProfileBlockScreenCapture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileBluetoothEnableContactSharing", m.GetWorkProfileBluetoothEnableContactSharing())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileDataSharingType() != nil {
        cast := (*m.GetWorkProfileDataSharingType()).String()
        err = writer.WriteStringValue("workProfileDataSharingType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileDefaultAppPermissionPolicy() != nil {
        cast := (*m.GetWorkProfileDefaultAppPermissionPolicy()).String()
        err = writer.WriteStringValue("workProfileDefaultAppPermissionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockFaceUnlock", m.GetWorkProfilePasswordBlockFaceUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockFingerprintUnlock", m.GetWorkProfilePasswordBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockIrisUnlock", m.GetWorkProfilePasswordBlockIrisUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfilePasswordBlockTrustAgents", m.GetWorkProfilePasswordBlockTrustAgents())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordExpirationDays", m.GetWorkProfilePasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLength", m.GetWorkProfilePasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinLetterCharacters", m.GetWorkProfilePasswordMinLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinLowerCaseCharacters", m.GetWorkProfilePasswordMinLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinNonLetterCharacters", m.GetWorkProfilePasswordMinNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinNumericCharacters", m.GetWorkProfilePasswordMinNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinSymbolCharacters", m.GetWorkProfilePasswordMinSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinUpperCaseCharacters", m.GetWorkProfilePasswordMinUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinutesOfInactivityBeforeScreenTimeout", m.GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordPreviousPasswordBlockCount", m.GetWorkProfilePasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfilePasswordRequiredType() != nil {
        cast := (*m.GetWorkProfilePasswordRequiredType()).String()
        err = writer.WriteStringValue("workProfilePasswordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordSignInFailureCountBeforeFactoryReset", m.GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfileRequiredPasswordComplexity() != nil {
        cast := (*m.GetWorkProfileRequiredPasswordComplexity()).String()
        err = writer.WriteStringValue("workProfileRequiredPasswordComplexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("workProfileRequirePassword", m.GetWorkProfileRequirePassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPasswordBlockFaceUnlock sets the passwordBlockFaceUnlock property value. Indicates whether or not to block face unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockFaceUnlock(value *bool)() {
    m.passwordBlockFaceUnlock = value
}
// SetPasswordBlockFingerprintUnlock sets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(value *bool)() {
    m.passwordBlockFingerprintUnlock = value
}
// SetPasswordBlockIrisUnlock sets the passwordBlockIrisUnlock property value. Indicates whether or not to block iris unlock.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockIrisUnlock(value *bool)() {
    m.passwordBlockIrisUnlock = value
}
// SetPasswordBlockTrustAgents sets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(value *bool)() {
    m.passwordBlockTrustAgents = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passwordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordRequiredType(value *AndroidWorkProfileRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.passwordSignInFailureCountBeforeFactoryReset = value
}
// SetRequiredPasswordComplexity sets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)() {
    m.requiredPasswordComplexity = value
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(value *bool)() {
    m.securityRequireVerifyApps = value
}
// SetVpnAlwaysOnPackageIdentifier sets the vpnAlwaysOnPackageIdentifier property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetVpnAlwaysOnPackageIdentifier(value *string)() {
    m.vpnAlwaysOnPackageIdentifier = value
}
// SetVpnEnableAlwaysOnLockdownMode sets the vpnEnableAlwaysOnLockdownMode property value. Enable lockdown mode for always-on VPN.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetVpnEnableAlwaysOnLockdownMode(value *bool)() {
    m.vpnEnableAlwaysOnLockdownMode = value
}
// SetWorkProfileAllowAppInstallsFromUnknownSources sets the workProfileAllowAppInstallsFromUnknownSources property value. Indicates whether to allow installation of apps from unknown sources.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileAllowAppInstallsFromUnknownSources(value *bool)() {
    m.workProfileAllowAppInstallsFromUnknownSources = value
}
// SetWorkProfileAllowWidgets sets the workProfileAllowWidgets property value. Allow widgets from work profile apps.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileAllowWidgets(value *bool)() {
    m.workProfileAllowWidgets = value
}
// SetWorkProfileBlockAddingAccounts sets the workProfileBlockAddingAccounts property value. Block users from adding/removing accounts in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockAddingAccounts(value *bool)() {
    m.workProfileBlockAddingAccounts = value
}
// SetWorkProfileBlockCamera sets the workProfileBlockCamera property value. Block work profile camera.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCamera(value *bool)() {
    m.workProfileBlockCamera = value
}
// SetWorkProfileBlockCrossProfileCallerId sets the workProfileBlockCrossProfileCallerId property value. Block display work profile caller ID in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCallerId(value *bool)() {
    m.workProfileBlockCrossProfileCallerId = value
}
// SetWorkProfileBlockCrossProfileContactsSearch sets the workProfileBlockCrossProfileContactsSearch property value. Block work profile contacts availability in personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileContactsSearch(value *bool)() {
    m.workProfileBlockCrossProfileContactsSearch = value
}
// SetWorkProfileBlockCrossProfileCopyPaste sets the workProfileBlockCrossProfileCopyPaste property value. Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCopyPaste(value *bool)() {
    m.workProfileBlockCrossProfileCopyPaste = value
}
// SetWorkProfileBlockNotificationsWhileDeviceLocked sets the workProfileBlockNotificationsWhileDeviceLocked property value. Indicates whether or not to block notifications while device locked.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockNotificationsWhileDeviceLocked(value *bool)() {
    m.workProfileBlockNotificationsWhileDeviceLocked = value
}
// SetWorkProfileBlockPersonalAppInstallsFromUnknownSources sets the workProfileBlockPersonalAppInstallsFromUnknownSources property value. Prevent app installations from unknown sources in the personal profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockPersonalAppInstallsFromUnknownSources(value *bool)() {
    m.workProfileBlockPersonalAppInstallsFromUnknownSources = value
}
// SetWorkProfileBlockScreenCapture sets the workProfileBlockScreenCapture property value. Block screen capture in work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockScreenCapture(value *bool)() {
    m.workProfileBlockScreenCapture = value
}
// SetWorkProfileBluetoothEnableContactSharing sets the workProfileBluetoothEnableContactSharing property value. Allow bluetooth devices to access enterprise contacts.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBluetoothEnableContactSharing(value *bool)() {
    m.workProfileBluetoothEnableContactSharing = value
}
// SetWorkProfileDataSharingType sets the workProfileDataSharingType property value. Android Work Profile cross profile data sharing type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDataSharingType(value *AndroidWorkProfileCrossProfileDataSharingType)() {
    m.workProfileDataSharingType = value
}
// SetWorkProfileDefaultAppPermissionPolicy sets the workProfileDefaultAppPermissionPolicy property value. Android Work Profile default app permission policy type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDefaultAppPermissionPolicy(value *AndroidWorkProfileDefaultAppPermissionPolicyType)() {
    m.workProfileDefaultAppPermissionPolicy = value
}
// SetWorkProfilePasswordBlockFaceUnlock sets the workProfilePasswordBlockFaceUnlock property value. Indicates whether or not to block face unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFaceUnlock(value *bool)() {
    m.workProfilePasswordBlockFaceUnlock = value
}
// SetWorkProfilePasswordBlockFingerprintUnlock sets the workProfilePasswordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFingerprintUnlock(value *bool)() {
    m.workProfilePasswordBlockFingerprintUnlock = value
}
// SetWorkProfilePasswordBlockIrisUnlock sets the workProfilePasswordBlockIrisUnlock property value. Indicates whether or not to block iris unlock for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockIrisUnlock(value *bool)() {
    m.workProfilePasswordBlockIrisUnlock = value
}
// SetWorkProfilePasswordBlockTrustAgents sets the workProfilePasswordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents for work profile.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockTrustAgents(value *bool)() {
    m.workProfilePasswordBlockTrustAgents = value
}
// SetWorkProfilePasswordExpirationDays sets the workProfilePasswordExpirationDays property value. Number of days before the work profile password expires. Valid values 1 to 365
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDays(value *int32)() {
    m.workProfilePasswordExpirationDays = value
}
// SetWorkProfilePasswordMinimumLength sets the workProfilePasswordMinimumLength property value. Minimum length of work profile password. Valid values 4 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLength(value *int32)() {
    m.workProfilePasswordMinimumLength = value
}
// SetWorkProfilePasswordMinLetterCharacters sets the workProfilePasswordMinLetterCharacters property value. Minimum # of letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLetterCharacters(value *int32)() {
    m.workProfilePasswordMinLetterCharacters = value
}
// SetWorkProfilePasswordMinLowerCaseCharacters sets the workProfilePasswordMinLowerCaseCharacters property value. Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLowerCaseCharacters(value *int32)() {
    m.workProfilePasswordMinLowerCaseCharacters = value
}
// SetWorkProfilePasswordMinNonLetterCharacters sets the workProfilePasswordMinNonLetterCharacters property value. Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNonLetterCharacters(value *int32)() {
    m.workProfilePasswordMinNonLetterCharacters = value
}
// SetWorkProfilePasswordMinNumericCharacters sets the workProfilePasswordMinNumericCharacters property value. Minimum # of numeric characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNumericCharacters(value *int32)() {
    m.workProfilePasswordMinNumericCharacters = value
}
// SetWorkProfilePasswordMinSymbolCharacters sets the workProfilePasswordMinSymbolCharacters property value. Minimum # of symbols required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinSymbolCharacters(value *int32)() {
    m.workProfilePasswordMinSymbolCharacters = value
}
// SetWorkProfilePasswordMinUpperCaseCharacters sets the workProfilePasswordMinUpperCaseCharacters property value. Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinUpperCaseCharacters(value *int32)() {
    m.workProfilePasswordMinUpperCaseCharacters = value
}
// SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout sets the workProfilePasswordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.workProfilePasswordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetWorkProfilePasswordPreviousPasswordBlockCount sets the workProfilePasswordPreviousPasswordBlockCount property value. Number of previous work profile passwords to block. Valid values 0 to 24
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordBlockCount(value *int32)() {
    m.workProfilePasswordPreviousPasswordBlockCount = value
}
// SetWorkProfilePasswordRequiredType sets the workProfilePasswordRequiredType property value. Android Work Profile required password type.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordRequiredType(value *AndroidWorkProfileRequiredPasswordType)() {
    m.workProfilePasswordRequiredType = value
}
// SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset sets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.workProfilePasswordSignInFailureCountBeforeFactoryReset = value
}
// SetWorkProfileRequiredPasswordComplexity sets the workProfileRequiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)() {
    m.workProfileRequiredPasswordComplexity = value
}
// SetWorkProfileRequirePassword sets the workProfileRequirePassword property value. Password is required or not for work profile
func (m *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileRequirePassword(value *bool)() {
    m.workProfileRequirePassword = value
}
