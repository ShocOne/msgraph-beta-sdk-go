package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidGeneralDeviceConfiguration 
type AndroidGeneralDeviceConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to block clipboard sharing to copy and paste between applications.
    appsBlockClipboardSharing *bool
    // Indicates whether or not to block copy and paste within applications.
    appsBlockCopyPaste *bool
    // Indicates whether or not to block the YouTube app.
    appsBlockYouTube *bool
    // List of apps to be hidden on the KNOX device. This collection can contain a maximum of 500 elements.
    appsHideList []AppListItemable
    // List of apps which can be installed on the KNOX device. This collection can contain a maximum of 500 elements.
    appsInstallAllowList []AppListItemable
    // List of apps which are blocked from being launched on the KNOX device. This collection can contain a maximum of 500 elements.
    appsLaunchBlockList []AppListItemable
    // Indicates whether or not to block Bluetooth.
    bluetoothBlocked *bool
    // Indicates whether or not to block the use of the camera.
    cameraBlocked *bool
    // Indicates whether or not to block data roaming.
    cellularBlockDataRoaming *bool
    // Indicates whether or not to block SMS/MMS messaging.
    cellularBlockMessaging *bool
    // Indicates whether or not to block voice roaming.
    cellularBlockVoiceRoaming *bool
    // Indicates whether or not to block syncing Wi-Fi tethering.
    cellularBlockWiFiTethering *bool
    // Possible values of the compliance app list.
    compliantAppListType *AppListType
    // List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
    compliantAppsList []AppListItemable
    // Indicates whether or not to block changing date and time while in KNOX Mode.
    dateAndTimeBlockChanges *bool
    // Indicates whether or not to allow device sharing mode.
    deviceSharingAllowed *bool
    // Indicates whether or not to block diagnostic data submission.
    diagnosticDataBlockSubmission *bool
    // Indicates whether or not to block user performing a factory reset.
    factoryResetBlocked *bool
    // Indicates whether or not to block Google account auto sync.
    googleAccountBlockAutoSync *bool
    // Indicates whether or not to block the Google Play store.
    googlePlayStoreBlocked *bool
    // A list of apps that will be allowed to run when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
    kioskModeApps []AppListItemable
    // Indicates whether or not to block the screen sleep button while in Kiosk Mode.
    kioskModeBlockSleepButton *bool
    // Indicates whether or not to block the volume buttons while in Kiosk Mode.
    kioskModeBlockVolumeButtons *bool
    // Indicates whether or not to block location services.
    locationServicesBlocked *bool
    // Indicates whether or not to block Near-Field Communication.
    nfcBlocked *bool
    // Indicates whether or not to block fingerprint unlock.
    passwordBlockFingerprintUnlock *bool
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
    // Indicates whether or not to require a password.
    passwordRequired *bool
    // Android required password type.
    passwordRequiredType *AndroidRequiredPasswordType
    // Number of sign in failures allowed before factory reset. Valid values 1 to 16
    passwordSignInFailureCountBeforeFactoryReset *int32
    // Indicates whether or not to block powering off the device.
    powerOffBlocked *bool
    // The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
    requiredPasswordComplexity *AndroidRequiredPasswordComplexity
    // Indicates whether or not to block screenshots.
    screenCaptureBlocked *bool
    // Require the Android Verify apps feature is turned on.
    securityRequireVerifyApps *bool
    // Indicates whether or not to block Google Backup.
    storageBlockGoogleBackup *bool
    // Indicates whether or not to block removable storage usage.
    storageBlockRemovableStorage *bool
    // Indicates whether or not to require device encryption.
    storageRequireDeviceEncryption *bool
    // Indicates whether or not to require removable storage encryption.
    storageRequireRemovableStorageEncryption *bool
    // Indicates whether or not to block the use of the Voice Assistant.
    voiceAssistantBlocked *bool
    // Indicates whether or not to block voice dialing.
    voiceDialingBlocked *bool
    // Indicates whether or not to block the web browser's auto fill feature.
    webBrowserBlockAutofill *bool
    // Indicates whether or not to block the web browser.
    webBrowserBlocked *bool
    // Indicates whether or not to block JavaScript within the web browser.
    webBrowserBlockJavaScript *bool
    // Indicates whether or not to block popups within the web browser.
    webBrowserBlockPopups *bool
    // Web Browser Cookie Settings.
    webBrowserCookieSettings *WebBrowserCookieSettings
    // Indicates whether or not to block syncing Wi-Fi.
    wiFiBlocked *bool
}
// NewAndroidGeneralDeviceConfiguration instantiates a new AndroidGeneralDeviceConfiguration and sets the default values.
func NewAndroidGeneralDeviceConfiguration()(*AndroidGeneralDeviceConfiguration) {
    m := &AndroidGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidGeneralDeviceConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidGeneralDeviceConfiguration(), nil
}
// GetAppsBlockClipboardSharing gets the appsBlockClipboardSharing property value. Indicates whether or not to block clipboard sharing to copy and paste between applications.
func (m *AndroidGeneralDeviceConfiguration) GetAppsBlockClipboardSharing()(*bool) {
    return m.appsBlockClipboardSharing
}
// GetAppsBlockCopyPaste gets the appsBlockCopyPaste property value. Indicates whether or not to block copy and paste within applications.
func (m *AndroidGeneralDeviceConfiguration) GetAppsBlockCopyPaste()(*bool) {
    return m.appsBlockCopyPaste
}
// GetAppsBlockYouTube gets the appsBlockYouTube property value. Indicates whether or not to block the YouTube app.
func (m *AndroidGeneralDeviceConfiguration) GetAppsBlockYouTube()(*bool) {
    return m.appsBlockYouTube
}
// GetAppsHideList gets the appsHideList property value. List of apps to be hidden on the KNOX device. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) GetAppsHideList()([]AppListItemable) {
    return m.appsHideList
}
// GetAppsInstallAllowList gets the appsInstallAllowList property value. List of apps which can be installed on the KNOX device. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) GetAppsInstallAllowList()([]AppListItemable) {
    return m.appsInstallAllowList
}
// GetAppsLaunchBlockList gets the appsLaunchBlockList property value. List of apps which are blocked from being launched on the KNOX device. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) GetAppsLaunchBlockList()([]AppListItemable) {
    return m.appsLaunchBlockList
}
// GetBluetoothBlocked gets the bluetoothBlocked property value. Indicates whether or not to block Bluetooth.
func (m *AndroidGeneralDeviceConfiguration) GetBluetoothBlocked()(*bool) {
    return m.bluetoothBlocked
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to block the use of the camera.
func (m *AndroidGeneralDeviceConfiguration) GetCameraBlocked()(*bool) {
    return m.cameraBlocked
}
// GetCellularBlockDataRoaming gets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *AndroidGeneralDeviceConfiguration) GetCellularBlockDataRoaming()(*bool) {
    return m.cellularBlockDataRoaming
}
// GetCellularBlockMessaging gets the cellularBlockMessaging property value. Indicates whether or not to block SMS/MMS messaging.
func (m *AndroidGeneralDeviceConfiguration) GetCellularBlockMessaging()(*bool) {
    return m.cellularBlockMessaging
}
// GetCellularBlockVoiceRoaming gets the cellularBlockVoiceRoaming property value. Indicates whether or not to block voice roaming.
func (m *AndroidGeneralDeviceConfiguration) GetCellularBlockVoiceRoaming()(*bool) {
    return m.cellularBlockVoiceRoaming
}
// GetCellularBlockWiFiTethering gets the cellularBlockWiFiTethering property value. Indicates whether or not to block syncing Wi-Fi tethering.
func (m *AndroidGeneralDeviceConfiguration) GetCellularBlockWiFiTethering()(*bool) {
    return m.cellularBlockWiFiTethering
}
// GetCompliantAppListType gets the compliantAppListType property value. Possible values of the compliance app list.
func (m *AndroidGeneralDeviceConfiguration) GetCompliantAppListType()(*AppListType) {
    return m.compliantAppListType
}
// GetCompliantAppsList gets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *AndroidGeneralDeviceConfiguration) GetCompliantAppsList()([]AppListItemable) {
    return m.compliantAppsList
}
// GetDateAndTimeBlockChanges gets the dateAndTimeBlockChanges property value. Indicates whether or not to block changing date and time while in KNOX Mode.
func (m *AndroidGeneralDeviceConfiguration) GetDateAndTimeBlockChanges()(*bool) {
    return m.dateAndTimeBlockChanges
}
// GetDeviceSharingAllowed gets the deviceSharingAllowed property value. Indicates whether or not to allow device sharing mode.
func (m *AndroidGeneralDeviceConfiguration) GetDeviceSharingAllowed()(*bool) {
    return m.deviceSharingAllowed
}
// GetDiagnosticDataBlockSubmission gets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *AndroidGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmission()(*bool) {
    return m.diagnosticDataBlockSubmission
}
// GetFactoryResetBlocked gets the factoryResetBlocked property value. Indicates whether or not to block user performing a factory reset.
func (m *AndroidGeneralDeviceConfiguration) GetFactoryResetBlocked()(*bool) {
    return m.factoryResetBlocked
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["appsBlockClipboardSharing"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAppsBlockClipboardSharing)
    res["appsBlockCopyPaste"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAppsBlockCopyPaste)
    res["appsBlockYouTube"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAppsBlockYouTube)
    res["appsHideList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue , m.SetAppsHideList)
    res["appsInstallAllowList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue , m.SetAppsInstallAllowList)
    res["appsLaunchBlockList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue , m.SetAppsLaunchBlockList)
    res["bluetoothBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBluetoothBlocked)
    res["cameraBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCameraBlocked)
    res["cellularBlockDataRoaming"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCellularBlockDataRoaming)
    res["cellularBlockMessaging"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCellularBlockMessaging)
    res["cellularBlockVoiceRoaming"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCellularBlockVoiceRoaming)
    res["cellularBlockWiFiTethering"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCellularBlockWiFiTethering)
    res["compliantAppListType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAppListType , m.SetCompliantAppListType)
    res["compliantAppsList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue , m.SetCompliantAppsList)
    res["dateAndTimeBlockChanges"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDateAndTimeBlockChanges)
    res["deviceSharingAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDeviceSharingAllowed)
    res["diagnosticDataBlockSubmission"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDiagnosticDataBlockSubmission)
    res["factoryResetBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFactoryResetBlocked)
    res["googleAccountBlockAutoSync"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetGoogleAccountBlockAutoSync)
    res["googlePlayStoreBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetGooglePlayStoreBlocked)
    res["kioskModeApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue , m.SetKioskModeApps)
    res["kioskModeBlockSleepButton"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetKioskModeBlockSleepButton)
    res["kioskModeBlockVolumeButtons"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetKioskModeBlockVolumeButtons)
    res["locationServicesBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetLocationServicesBlocked)
    res["nfcBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetNfcBlocked)
    res["passwordBlockFingerprintUnlock"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPasswordBlockFingerprintUnlock)
    res["passwordBlockTrustAgents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPasswordBlockTrustAgents)
    res["passwordExpirationDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordExpirationDays)
    res["passwordMinimumLength"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordMinimumLength)
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordMinutesOfInactivityBeforeScreenTimeout)
    res["passwordPreviousPasswordBlockCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordPreviousPasswordBlockCount)
    res["passwordRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPasswordRequired)
    res["passwordRequiredType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidRequiredPasswordType , m.SetPasswordRequiredType)
    res["passwordSignInFailureCountBeforeFactoryReset"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordSignInFailureCountBeforeFactoryReset)
    res["powerOffBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPowerOffBlocked)
    res["requiredPasswordComplexity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidRequiredPasswordComplexity , m.SetRequiredPasswordComplexity)
    res["screenCaptureBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetScreenCaptureBlocked)
    res["securityRequireVerifyApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSecurityRequireVerifyApps)
    res["storageBlockGoogleBackup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStorageBlockGoogleBackup)
    res["storageBlockRemovableStorage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStorageBlockRemovableStorage)
    res["storageRequireDeviceEncryption"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStorageRequireDeviceEncryption)
    res["storageRequireRemovableStorageEncryption"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStorageRequireRemovableStorageEncryption)
    res["voiceAssistantBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetVoiceAssistantBlocked)
    res["voiceDialingBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetVoiceDialingBlocked)
    res["webBrowserBlockAutofill"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetWebBrowserBlockAutofill)
    res["webBrowserBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetWebBrowserBlocked)
    res["webBrowserBlockJavaScript"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetWebBrowserBlockJavaScript)
    res["webBrowserBlockPopups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetWebBrowserBlockPopups)
    res["webBrowserCookieSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWebBrowserCookieSettings , m.SetWebBrowserCookieSettings)
    res["wiFiBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetWiFiBlocked)
    return res
}
// GetGoogleAccountBlockAutoSync gets the googleAccountBlockAutoSync property value. Indicates whether or not to block Google account auto sync.
func (m *AndroidGeneralDeviceConfiguration) GetGoogleAccountBlockAutoSync()(*bool) {
    return m.googleAccountBlockAutoSync
}
// GetGooglePlayStoreBlocked gets the googlePlayStoreBlocked property value. Indicates whether or not to block the Google Play store.
func (m *AndroidGeneralDeviceConfiguration) GetGooglePlayStoreBlocked()(*bool) {
    return m.googlePlayStoreBlocked
}
// GetKioskModeApps gets the kioskModeApps property value. A list of apps that will be allowed to run when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) GetKioskModeApps()([]AppListItemable) {
    return m.kioskModeApps
}
// GetKioskModeBlockSleepButton gets the kioskModeBlockSleepButton property value. Indicates whether or not to block the screen sleep button while in Kiosk Mode.
func (m *AndroidGeneralDeviceConfiguration) GetKioskModeBlockSleepButton()(*bool) {
    return m.kioskModeBlockSleepButton
}
// GetKioskModeBlockVolumeButtons gets the kioskModeBlockVolumeButtons property value. Indicates whether or not to block the volume buttons while in Kiosk Mode.
func (m *AndroidGeneralDeviceConfiguration) GetKioskModeBlockVolumeButtons()(*bool) {
    return m.kioskModeBlockVolumeButtons
}
// GetLocationServicesBlocked gets the locationServicesBlocked property value. Indicates whether or not to block location services.
func (m *AndroidGeneralDeviceConfiguration) GetLocationServicesBlocked()(*bool) {
    return m.locationServicesBlocked
}
// GetNfcBlocked gets the nfcBlocked property value. Indicates whether or not to block Near-Field Communication.
func (m *AndroidGeneralDeviceConfiguration) GetNfcBlocked()(*bool) {
    return m.nfcBlocked
}
// GetPasswordBlockFingerprintUnlock gets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock()(*bool) {
    return m.passwordBlockFingerprintUnlock
}
// GetPasswordBlockTrustAgents gets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidGeneralDeviceConfiguration) GetPasswordBlockTrustAgents()(*bool) {
    return m.passwordBlockTrustAgents
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passwordMinutesOfInactivityBeforeScreenTimeout
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequired gets the passwordRequired property value. Indicates whether or not to require a password.
func (m *AndroidGeneralDeviceConfiguration) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Android required password type.
func (m *AndroidGeneralDeviceConfiguration) GetPasswordRequiredType()(*AndroidRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.passwordSignInFailureCountBeforeFactoryReset
}
// GetPowerOffBlocked gets the powerOffBlocked property value. Indicates whether or not to block powering off the device.
func (m *AndroidGeneralDeviceConfiguration) GetPowerOffBlocked()(*bool) {
    return m.powerOffBlocked
}
// GetRequiredPasswordComplexity gets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidGeneralDeviceConfiguration) GetRequiredPasswordComplexity()(*AndroidRequiredPasswordComplexity) {
    return m.requiredPasswordComplexity
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to block screenshots.
func (m *AndroidGeneralDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidGeneralDeviceConfiguration) GetSecurityRequireVerifyApps()(*bool) {
    return m.securityRequireVerifyApps
}
// GetStorageBlockGoogleBackup gets the storageBlockGoogleBackup property value. Indicates whether or not to block Google Backup.
func (m *AndroidGeneralDeviceConfiguration) GetStorageBlockGoogleBackup()(*bool) {
    return m.storageBlockGoogleBackup
}
// GetStorageBlockRemovableStorage gets the storageBlockRemovableStorage property value. Indicates whether or not to block removable storage usage.
func (m *AndroidGeneralDeviceConfiguration) GetStorageBlockRemovableStorage()(*bool) {
    return m.storageBlockRemovableStorage
}
// GetStorageRequireDeviceEncryption gets the storageRequireDeviceEncryption property value. Indicates whether or not to require device encryption.
func (m *AndroidGeneralDeviceConfiguration) GetStorageRequireDeviceEncryption()(*bool) {
    return m.storageRequireDeviceEncryption
}
// GetStorageRequireRemovableStorageEncryption gets the storageRequireRemovableStorageEncryption property value. Indicates whether or not to require removable storage encryption.
func (m *AndroidGeneralDeviceConfiguration) GetStorageRequireRemovableStorageEncryption()(*bool) {
    return m.storageRequireRemovableStorageEncryption
}
// GetVoiceAssistantBlocked gets the voiceAssistantBlocked property value. Indicates whether or not to block the use of the Voice Assistant.
func (m *AndroidGeneralDeviceConfiguration) GetVoiceAssistantBlocked()(*bool) {
    return m.voiceAssistantBlocked
}
// GetVoiceDialingBlocked gets the voiceDialingBlocked property value. Indicates whether or not to block voice dialing.
func (m *AndroidGeneralDeviceConfiguration) GetVoiceDialingBlocked()(*bool) {
    return m.voiceDialingBlocked
}
// GetWebBrowserBlockAutofill gets the webBrowserBlockAutofill property value. Indicates whether or not to block the web browser's auto fill feature.
func (m *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockAutofill()(*bool) {
    return m.webBrowserBlockAutofill
}
// GetWebBrowserBlocked gets the webBrowserBlocked property value. Indicates whether or not to block the web browser.
func (m *AndroidGeneralDeviceConfiguration) GetWebBrowserBlocked()(*bool) {
    return m.webBrowserBlocked
}
// GetWebBrowserBlockJavaScript gets the webBrowserBlockJavaScript property value. Indicates whether or not to block JavaScript within the web browser.
func (m *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockJavaScript()(*bool) {
    return m.webBrowserBlockJavaScript
}
// GetWebBrowserBlockPopups gets the webBrowserBlockPopups property value. Indicates whether or not to block popups within the web browser.
func (m *AndroidGeneralDeviceConfiguration) GetWebBrowserBlockPopups()(*bool) {
    return m.webBrowserBlockPopups
}
// GetWebBrowserCookieSettings gets the webBrowserCookieSettings property value. Web Browser Cookie Settings.
func (m *AndroidGeneralDeviceConfiguration) GetWebBrowserCookieSettings()(*WebBrowserCookieSettings) {
    return m.webBrowserCookieSettings
}
// GetWiFiBlocked gets the wiFiBlocked property value. Indicates whether or not to block syncing Wi-Fi.
func (m *AndroidGeneralDeviceConfiguration) GetWiFiBlocked()(*bool) {
    return m.wiFiBlocked
}
// Serialize serializes information the current object
func (m *AndroidGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appsBlockClipboardSharing", m.GetAppsBlockClipboardSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appsBlockCopyPaste", m.GetAppsBlockCopyPaste())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appsBlockYouTube", m.GetAppsBlockYouTube())
        if err != nil {
            return err
        }
    }
    if m.GetAppsHideList() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppsHideList())
        err = writer.WriteCollectionOfObjectValues("appsHideList", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppsInstallAllowList() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppsInstallAllowList())
        err = writer.WriteCollectionOfObjectValues("appsInstallAllowList", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppsLaunchBlockList() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppsLaunchBlockList())
        err = writer.WriteCollectionOfObjectValues("appsLaunchBlockList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlocked", m.GetBluetoothBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cameraBlocked", m.GetCameraBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockDataRoaming", m.GetCellularBlockDataRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockMessaging", m.GetCellularBlockMessaging())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockVoiceRoaming", m.GetCellularBlockVoiceRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockWiFiTethering", m.GetCellularBlockWiFiTethering())
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppListType() != nil {
        cast := (*m.GetCompliantAppListType()).String()
        err = writer.WriteStringValue("compliantAppListType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppsList() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCompliantAppsList())
        err = writer.WriteCollectionOfObjectValues("compliantAppsList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dateAndTimeBlockChanges", m.GetDateAndTimeBlockChanges())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceSharingAllowed", m.GetDeviceSharingAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticDataBlockSubmission", m.GetDiagnosticDataBlockSubmission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("factoryResetBlocked", m.GetFactoryResetBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("googleAccountBlockAutoSync", m.GetGoogleAccountBlockAutoSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("googlePlayStoreBlocked", m.GetGooglePlayStoreBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetKioskModeApps())
        err = writer.WriteCollectionOfObjectValues("kioskModeApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockSleepButton", m.GetKioskModeBlockSleepButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockVolumeButtons", m.GetKioskModeBlockVolumeButtons())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locationServicesBlocked", m.GetLocationServicesBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("nfcBlocked", m.GetNfcBlocked())
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
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("powerOffBlocked", m.GetPowerOffBlocked())
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
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
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
        err = writer.WriteBoolValue("storageBlockGoogleBackup", m.GetStorageBlockGoogleBackup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockRemovableStorage", m.GetStorageBlockRemovableStorage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRequireDeviceEncryption", m.GetStorageRequireDeviceEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRequireRemovableStorageEncryption", m.GetStorageRequireRemovableStorageEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("voiceAssistantBlocked", m.GetVoiceAssistantBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("voiceDialingBlocked", m.GetVoiceDialingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("webBrowserBlockAutofill", m.GetWebBrowserBlockAutofill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("webBrowserBlocked", m.GetWebBrowserBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("webBrowserBlockJavaScript", m.GetWebBrowserBlockJavaScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("webBrowserBlockPopups", m.GetWebBrowserBlockPopups())
        if err != nil {
            return err
        }
    }
    if m.GetWebBrowserCookieSettings() != nil {
        cast := (*m.GetWebBrowserCookieSettings()).String()
        err = writer.WriteStringValue("webBrowserCookieSettings", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wiFiBlocked", m.GetWiFiBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppsBlockClipboardSharing sets the appsBlockClipboardSharing property value. Indicates whether or not to block clipboard sharing to copy and paste between applications.
func (m *AndroidGeneralDeviceConfiguration) SetAppsBlockClipboardSharing(value *bool)() {
    m.appsBlockClipboardSharing = value
}
// SetAppsBlockCopyPaste sets the appsBlockCopyPaste property value. Indicates whether or not to block copy and paste within applications.
func (m *AndroidGeneralDeviceConfiguration) SetAppsBlockCopyPaste(value *bool)() {
    m.appsBlockCopyPaste = value
}
// SetAppsBlockYouTube sets the appsBlockYouTube property value. Indicates whether or not to block the YouTube app.
func (m *AndroidGeneralDeviceConfiguration) SetAppsBlockYouTube(value *bool)() {
    m.appsBlockYouTube = value
}
// SetAppsHideList sets the appsHideList property value. List of apps to be hidden on the KNOX device. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) SetAppsHideList(value []AppListItemable)() {
    m.appsHideList = value
}
// SetAppsInstallAllowList sets the appsInstallAllowList property value. List of apps which can be installed on the KNOX device. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) SetAppsInstallAllowList(value []AppListItemable)() {
    m.appsInstallAllowList = value
}
// SetAppsLaunchBlockList sets the appsLaunchBlockList property value. List of apps which are blocked from being launched on the KNOX device. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) SetAppsLaunchBlockList(value []AppListItemable)() {
    m.appsLaunchBlockList = value
}
// SetBluetoothBlocked sets the bluetoothBlocked property value. Indicates whether or not to block Bluetooth.
func (m *AndroidGeneralDeviceConfiguration) SetBluetoothBlocked(value *bool)() {
    m.bluetoothBlocked = value
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to block the use of the camera.
func (m *AndroidGeneralDeviceConfiguration) SetCameraBlocked(value *bool)() {
    m.cameraBlocked = value
}
// SetCellularBlockDataRoaming sets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *AndroidGeneralDeviceConfiguration) SetCellularBlockDataRoaming(value *bool)() {
    m.cellularBlockDataRoaming = value
}
// SetCellularBlockMessaging sets the cellularBlockMessaging property value. Indicates whether or not to block SMS/MMS messaging.
func (m *AndroidGeneralDeviceConfiguration) SetCellularBlockMessaging(value *bool)() {
    m.cellularBlockMessaging = value
}
// SetCellularBlockVoiceRoaming sets the cellularBlockVoiceRoaming property value. Indicates whether or not to block voice roaming.
func (m *AndroidGeneralDeviceConfiguration) SetCellularBlockVoiceRoaming(value *bool)() {
    m.cellularBlockVoiceRoaming = value
}
// SetCellularBlockWiFiTethering sets the cellularBlockWiFiTethering property value. Indicates whether or not to block syncing Wi-Fi tethering.
func (m *AndroidGeneralDeviceConfiguration) SetCellularBlockWiFiTethering(value *bool)() {
    m.cellularBlockWiFiTethering = value
}
// SetCompliantAppListType sets the compliantAppListType property value. Possible values of the compliance app list.
func (m *AndroidGeneralDeviceConfiguration) SetCompliantAppListType(value *AppListType)() {
    m.compliantAppListType = value
}
// SetCompliantAppsList sets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *AndroidGeneralDeviceConfiguration) SetCompliantAppsList(value []AppListItemable)() {
    m.compliantAppsList = value
}
// SetDateAndTimeBlockChanges sets the dateAndTimeBlockChanges property value. Indicates whether or not to block changing date and time while in KNOX Mode.
func (m *AndroidGeneralDeviceConfiguration) SetDateAndTimeBlockChanges(value *bool)() {
    m.dateAndTimeBlockChanges = value
}
// SetDeviceSharingAllowed sets the deviceSharingAllowed property value. Indicates whether or not to allow device sharing mode.
func (m *AndroidGeneralDeviceConfiguration) SetDeviceSharingAllowed(value *bool)() {
    m.deviceSharingAllowed = value
}
// SetDiagnosticDataBlockSubmission sets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *AndroidGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmission(value *bool)() {
    m.diagnosticDataBlockSubmission = value
}
// SetFactoryResetBlocked sets the factoryResetBlocked property value. Indicates whether or not to block user performing a factory reset.
func (m *AndroidGeneralDeviceConfiguration) SetFactoryResetBlocked(value *bool)() {
    m.factoryResetBlocked = value
}
// SetGoogleAccountBlockAutoSync sets the googleAccountBlockAutoSync property value. Indicates whether or not to block Google account auto sync.
func (m *AndroidGeneralDeviceConfiguration) SetGoogleAccountBlockAutoSync(value *bool)() {
    m.googleAccountBlockAutoSync = value
}
// SetGooglePlayStoreBlocked sets the googlePlayStoreBlocked property value. Indicates whether or not to block the Google Play store.
func (m *AndroidGeneralDeviceConfiguration) SetGooglePlayStoreBlocked(value *bool)() {
    m.googlePlayStoreBlocked = value
}
// SetKioskModeApps sets the kioskModeApps property value. A list of apps that will be allowed to run when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidGeneralDeviceConfiguration) SetKioskModeApps(value []AppListItemable)() {
    m.kioskModeApps = value
}
// SetKioskModeBlockSleepButton sets the kioskModeBlockSleepButton property value. Indicates whether or not to block the screen sleep button while in Kiosk Mode.
func (m *AndroidGeneralDeviceConfiguration) SetKioskModeBlockSleepButton(value *bool)() {
    m.kioskModeBlockSleepButton = value
}
// SetKioskModeBlockVolumeButtons sets the kioskModeBlockVolumeButtons property value. Indicates whether or not to block the volume buttons while in Kiosk Mode.
func (m *AndroidGeneralDeviceConfiguration) SetKioskModeBlockVolumeButtons(value *bool)() {
    m.kioskModeBlockVolumeButtons = value
}
// SetLocationServicesBlocked sets the locationServicesBlocked property value. Indicates whether or not to block location services.
func (m *AndroidGeneralDeviceConfiguration) SetLocationServicesBlocked(value *bool)() {
    m.locationServicesBlocked = value
}
// SetNfcBlocked sets the nfcBlocked property value. Indicates whether or not to block Near-Field Communication.
func (m *AndroidGeneralDeviceConfiguration) SetNfcBlocked(value *bool)() {
    m.nfcBlocked = value
}
// SetPasswordBlockFingerprintUnlock sets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *AndroidGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(value *bool)() {
    m.passwordBlockFingerprintUnlock = value
}
// SetPasswordBlockTrustAgents sets the passwordBlockTrustAgents property value. Indicates whether or not to block Smart Lock and other trust agents.
func (m *AndroidGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(value *bool)() {
    m.passwordBlockTrustAgents = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of passwords. Valid values 4 to 16
func (m *AndroidGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passwordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *AndroidGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequired sets the passwordRequired property value. Indicates whether or not to require a password.
func (m *AndroidGeneralDeviceConfiguration) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Android required password type.
func (m *AndroidGeneralDeviceConfiguration) SetPasswordRequiredType(value *AndroidRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset. Valid values 1 to 16
func (m *AndroidGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.passwordSignInFailureCountBeforeFactoryReset = value
}
// SetPowerOffBlocked sets the powerOffBlocked property value. Indicates whether or not to block powering off the device.
func (m *AndroidGeneralDeviceConfiguration) SetPowerOffBlocked(value *bool)() {
    m.powerOffBlocked = value
}
// SetRequiredPasswordComplexity sets the requiredPasswordComplexity property value. The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to Android 11+.
func (m *AndroidGeneralDeviceConfiguration) SetRequiredPasswordComplexity(value *AndroidRequiredPasswordComplexity)() {
    m.requiredPasswordComplexity = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to block screenshots.
func (m *AndroidGeneralDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(value *bool)() {
    m.securityRequireVerifyApps = value
}
// SetStorageBlockGoogleBackup sets the storageBlockGoogleBackup property value. Indicates whether or not to block Google Backup.
func (m *AndroidGeneralDeviceConfiguration) SetStorageBlockGoogleBackup(value *bool)() {
    m.storageBlockGoogleBackup = value
}
// SetStorageBlockRemovableStorage sets the storageBlockRemovableStorage property value. Indicates whether or not to block removable storage usage.
func (m *AndroidGeneralDeviceConfiguration) SetStorageBlockRemovableStorage(value *bool)() {
    m.storageBlockRemovableStorage = value
}
// SetStorageRequireDeviceEncryption sets the storageRequireDeviceEncryption property value. Indicates whether or not to require device encryption.
func (m *AndroidGeneralDeviceConfiguration) SetStorageRequireDeviceEncryption(value *bool)() {
    m.storageRequireDeviceEncryption = value
}
// SetStorageRequireRemovableStorageEncryption sets the storageRequireRemovableStorageEncryption property value. Indicates whether or not to require removable storage encryption.
func (m *AndroidGeneralDeviceConfiguration) SetStorageRequireRemovableStorageEncryption(value *bool)() {
    m.storageRequireRemovableStorageEncryption = value
}
// SetVoiceAssistantBlocked sets the voiceAssistantBlocked property value. Indicates whether or not to block the use of the Voice Assistant.
func (m *AndroidGeneralDeviceConfiguration) SetVoiceAssistantBlocked(value *bool)() {
    m.voiceAssistantBlocked = value
}
// SetVoiceDialingBlocked sets the voiceDialingBlocked property value. Indicates whether or not to block voice dialing.
func (m *AndroidGeneralDeviceConfiguration) SetVoiceDialingBlocked(value *bool)() {
    m.voiceDialingBlocked = value
}
// SetWebBrowserBlockAutofill sets the webBrowserBlockAutofill property value. Indicates whether or not to block the web browser's auto fill feature.
func (m *AndroidGeneralDeviceConfiguration) SetWebBrowserBlockAutofill(value *bool)() {
    m.webBrowserBlockAutofill = value
}
// SetWebBrowserBlocked sets the webBrowserBlocked property value. Indicates whether or not to block the web browser.
func (m *AndroidGeneralDeviceConfiguration) SetWebBrowserBlocked(value *bool)() {
    m.webBrowserBlocked = value
}
// SetWebBrowserBlockJavaScript sets the webBrowserBlockJavaScript property value. Indicates whether or not to block JavaScript within the web browser.
func (m *AndroidGeneralDeviceConfiguration) SetWebBrowserBlockJavaScript(value *bool)() {
    m.webBrowserBlockJavaScript = value
}
// SetWebBrowserBlockPopups sets the webBrowserBlockPopups property value. Indicates whether or not to block popups within the web browser.
func (m *AndroidGeneralDeviceConfiguration) SetWebBrowserBlockPopups(value *bool)() {
    m.webBrowserBlockPopups = value
}
// SetWebBrowserCookieSettings sets the webBrowserCookieSettings property value. Web Browser Cookie Settings.
func (m *AndroidGeneralDeviceConfiguration) SetWebBrowserCookieSettings(value *WebBrowserCookieSettings)() {
    m.webBrowserCookieSettings = value
}
// SetWiFiBlocked sets the wiFiBlocked property value. Indicates whether or not to block syncing Wi-Fi.
func (m *AndroidGeneralDeviceConfiguration) SetWiFiBlocked(value *bool)() {
    m.wiFiBlocked = value
}
