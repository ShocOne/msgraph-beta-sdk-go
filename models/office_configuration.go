package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeConfiguration 
type OfficeConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of office Client configuration.
    clientConfigurations []OfficeClientConfigurationable
    // The OdataType property
    odataType *string
    // List of office Client check-in status.
    tenantCheckinStatuses []OfficeClientCheckinStatusable
    // Entity that describes tenant check-in statues
    tenantUserCheckinSummary OfficeUserCheckinSummaryable
}
// NewOfficeConfiguration instantiates a new OfficeConfiguration and sets the default values.
func NewOfficeConfiguration()(*OfficeConfiguration) {
    m := &OfficeConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.officeConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOfficeConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClientConfigurations gets the clientConfigurations property value. List of office Client configuration.
func (m *OfficeConfiguration) GetClientConfigurations()([]OfficeClientConfigurationable) {
    return m.clientConfigurations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOfficeClientConfigurationFromDiscriminatorValue , m.SetClientConfigurations)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["tenantCheckinStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOfficeClientCheckinStatusFromDiscriminatorValue , m.SetTenantCheckinStatuses)
    res["tenantUserCheckinSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOfficeUserCheckinSummaryFromDiscriminatorValue , m.SetTenantUserCheckinSummary)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OfficeConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantCheckinStatuses gets the tenantCheckinStatuses property value. List of office Client check-in status.
func (m *OfficeConfiguration) GetTenantCheckinStatuses()([]OfficeClientCheckinStatusable) {
    return m.tenantCheckinStatuses
}
// GetTenantUserCheckinSummary gets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
func (m *OfficeConfiguration) GetTenantUserCheckinSummary()(OfficeUserCheckinSummaryable) {
    return m.tenantUserCheckinSummary
}
// Serialize serializes information the current object
func (m *OfficeConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClientConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetClientConfigurations())
        err := writer.WriteCollectionOfObjectValues("clientConfigurations", cast)
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
    if m.GetTenantCheckinStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTenantCheckinStatuses())
        err := writer.WriteCollectionOfObjectValues("tenantCheckinStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tenantUserCheckinSummary", m.GetTenantUserCheckinSummary())
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
func (m *OfficeConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClientConfigurations sets the clientConfigurations property value. List of office Client configuration.
func (m *OfficeConfiguration) SetClientConfigurations(value []OfficeClientConfigurationable)() {
    m.clientConfigurations = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OfficeConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantCheckinStatuses sets the tenantCheckinStatuses property value. List of office Client check-in status.
func (m *OfficeConfiguration) SetTenantCheckinStatuses(value []OfficeClientCheckinStatusable)() {
    m.tenantCheckinStatuses = value
}
// SetTenantUserCheckinSummary sets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
func (m *OfficeConfiguration) SetTenantUserCheckinSummary(value OfficeUserCheckinSummaryable)() {
    m.tenantUserCheckinSummary = value
}
