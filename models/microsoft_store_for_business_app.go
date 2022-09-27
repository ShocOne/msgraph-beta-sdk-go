package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftStoreForBusinessApp 
type MicrosoftStoreForBusinessApp struct {
    MobileApp
    // The collection of contained apps in a mobileApp acting as a package.
    containedApps []MobileContainedAppable
    // The licenseType property
    licenseType *MicrosoftStoreForBusinessLicenseType
    // The supported License Type.
    licensingType VppLicensingTypeable
    // The app package identifier
    packageIdentityName *string
    // The app product key
    productKey *string
    // The total number of Microsoft Store for Business licenses.
    totalLicenseCount *int32
    // The number of Microsoft Store for Business licenses in use.
    usedLicenseCount *int32
}
// NewMicrosoftStoreForBusinessApp instantiates a new MicrosoftStoreForBusinessApp and sets the default values.
func NewMicrosoftStoreForBusinessApp()(*MicrosoftStoreForBusinessApp) {
    m := &MicrosoftStoreForBusinessApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.microsoftStoreForBusinessApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftStoreForBusinessApp(), nil
}
// GetContainedApps gets the containedApps property value. The collection of contained apps in a mobileApp acting as a package.
func (m *MicrosoftStoreForBusinessApp) GetContainedApps()([]MobileContainedAppable) {
    return m.containedApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftStoreForBusinessApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["containedApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileContainedAppFromDiscriminatorValue , m.SetContainedApps)
    res["licenseType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMicrosoftStoreForBusinessLicenseType , m.SetLicenseType)
    res["licensingType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateVppLicensingTypeFromDiscriminatorValue , m.SetLicensingType)
    res["packageIdentityName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPackageIdentityName)
    res["productKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProductKey)
    res["totalLicenseCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalLicenseCount)
    res["usedLicenseCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUsedLicenseCount)
    return res
}
// GetLicenseType gets the licenseType property value. The licenseType property
func (m *MicrosoftStoreForBusinessApp) GetLicenseType()(*MicrosoftStoreForBusinessLicenseType) {
    return m.licenseType
}
// GetLicensingType gets the licensingType property value. The supported License Type.
func (m *MicrosoftStoreForBusinessApp) GetLicensingType()(VppLicensingTypeable) {
    return m.licensingType
}
// GetPackageIdentityName gets the packageIdentityName property value. The app package identifier
func (m *MicrosoftStoreForBusinessApp) GetPackageIdentityName()(*string) {
    return m.packageIdentityName
}
// GetProductKey gets the productKey property value. The app product key
func (m *MicrosoftStoreForBusinessApp) GetProductKey()(*string) {
    return m.productKey
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of Microsoft Store for Business licenses.
func (m *MicrosoftStoreForBusinessApp) GetTotalLicenseCount()(*int32) {
    return m.totalLicenseCount
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of Microsoft Store for Business licenses in use.
func (m *MicrosoftStoreForBusinessApp) GetUsedLicenseCount()(*int32) {
    return m.usedLicenseCount
}
// Serialize serializes information the current object
func (m *MicrosoftStoreForBusinessApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContainedApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContainedApps())
        err = writer.WriteCollectionOfObjectValues("containedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseType() != nil {
        cast := (*m.GetLicenseType()).String()
        err = writer.WriteStringValue("licenseType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licensingType", m.GetLicensingType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageIdentityName", m.GetPackageIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicenseCount", m.GetTotalLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usedLicenseCount", m.GetUsedLicenseCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainedApps sets the containedApps property value. The collection of contained apps in a mobileApp acting as a package.
func (m *MicrosoftStoreForBusinessApp) SetContainedApps(value []MobileContainedAppable)() {
    m.containedApps = value
}
// SetLicenseType sets the licenseType property value. The licenseType property
func (m *MicrosoftStoreForBusinessApp) SetLicenseType(value *MicrosoftStoreForBusinessLicenseType)() {
    m.licenseType = value
}
// SetLicensingType sets the licensingType property value. The supported License Type.
func (m *MicrosoftStoreForBusinessApp) SetLicensingType(value VppLicensingTypeable)() {
    m.licensingType = value
}
// SetPackageIdentityName sets the packageIdentityName property value. The app package identifier
func (m *MicrosoftStoreForBusinessApp) SetPackageIdentityName(value *string)() {
    m.packageIdentityName = value
}
// SetProductKey sets the productKey property value. The app product key
func (m *MicrosoftStoreForBusinessApp) SetProductKey(value *string)() {
    m.productKey = value
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of Microsoft Store for Business licenses.
func (m *MicrosoftStoreForBusinessApp) SetTotalLicenseCount(value *int32)() {
    m.totalLicenseCount = value
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of Microsoft Store for Business licenses in use.
func (m *MicrosoftStoreForBusinessApp) SetUsedLicenseCount(value *int32)() {
    m.usedLicenseCount = value
}
