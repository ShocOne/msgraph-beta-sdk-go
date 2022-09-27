package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcForensicStorageAccount 
type CloudPcForensicStorageAccount struct {
    Entity
    // The ID of the storage account.
    storageAccountId *string
    // The name of the storage account.
    storageAccountName *string
}
// NewCloudPcForensicStorageAccount instantiates a new cloudPcForensicStorageAccount and sets the default values.
func NewCloudPcForensicStorageAccount()(*CloudPcForensicStorageAccount) {
    m := &CloudPcForensicStorageAccount{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcForensicStorageAccount";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcForensicStorageAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcForensicStorageAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcForensicStorageAccount(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcForensicStorageAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["storageAccountId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStorageAccountId)
    res["storageAccountName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStorageAccountName)
    return res
}
// GetStorageAccountId gets the storageAccountId property value. The ID of the storage account.
func (m *CloudPcForensicStorageAccount) GetStorageAccountId()(*string) {
    return m.storageAccountId
}
// GetStorageAccountName gets the storageAccountName property value. The name of the storage account.
func (m *CloudPcForensicStorageAccount) GetStorageAccountName()(*string) {
    return m.storageAccountName
}
// Serialize serializes information the current object
func (m *CloudPcForensicStorageAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("storageAccountId", m.GetStorageAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("storageAccountName", m.GetStorageAccountName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStorageAccountId sets the storageAccountId property value. The ID of the storage account.
func (m *CloudPcForensicStorageAccount) SetStorageAccountId(value *string)() {
    m.storageAccountId = value
}
// SetStorageAccountName sets the storageAccountName property value. The name of the storage account.
func (m *CloudPcForensicStorageAccount) SetStorageAccountName(value *string)() {
    m.storageAccountName = value
}
