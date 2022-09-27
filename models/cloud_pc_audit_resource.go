package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcAuditResource 
type CloudPcAuditResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The resource entity display name.
    displayName *string
    // A list of modified properties.
    modifiedProperties []CloudPcAuditPropertyable
    // The OdataType property
    odataType *string
    // The ID of the audit resource.
    resourceId *string
    // The type of the audit resource.
    type_escaped *string
}
// NewCloudPcAuditResource instantiates a new cloudPcAuditResource and sets the default values.
func NewCloudPcAuditResource()(*CloudPcAuditResource) {
    m := &CloudPcAuditResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.cloudPcAuditResource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcAuditResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcAuditResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcAuditResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcAuditResource) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The resource entity display name.
func (m *CloudPcAuditResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcAuditResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["modifiedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcAuditPropertyFromDiscriminatorValue , m.SetModifiedProperties)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["resourceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResourceId)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    return res
}
// GetModifiedProperties gets the modifiedProperties property value. A list of modified properties.
func (m *CloudPcAuditResource) GetModifiedProperties()([]CloudPcAuditPropertyable) {
    return m.modifiedProperties
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcAuditResource) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceId gets the resourceId property value. The ID of the audit resource.
func (m *CloudPcAuditResource) GetResourceId()(*string) {
    return m.resourceId
}
// GetType gets the type property value. The type of the audit resource.
func (m *CloudPcAuditResource) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *CloudPcAuditResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetModifiedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetModifiedProperties())
        err := writer.WriteCollectionOfObjectValues("modifiedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *CloudPcAuditResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The resource entity display name.
func (m *CloudPcAuditResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetModifiedProperties sets the modifiedProperties property value. A list of modified properties.
func (m *CloudPcAuditResource) SetModifiedProperties(value []CloudPcAuditPropertyable)() {
    m.modifiedProperties = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcAuditResource) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceId sets the resourceId property value. The ID of the audit resource.
func (m *CloudPcAuditResource) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetType sets the type property value. The type of the audit resource.
func (m *CloudPcAuditResource) SetType(value *string)() {
    m.type_escaped = value
}
