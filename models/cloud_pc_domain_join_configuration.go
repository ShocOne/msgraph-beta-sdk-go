package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcDomainJoinConfiguration 
type CloudPcDomainJoinConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Azure AD joined. If you enter an onPremisesConnectionId, leave regionName as empty.
    onPremisesConnectionId *string
    // The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. The underlying virtual network will be created and managed by the Windows 365 service. This can only be entered if the IT admin chooses Azure AD joined as the domain join type. If you enter a regionName, leave onPremisesConnectionId as empty.
    regionName *string
    // Specifies how the provisioned Cloud PC will be joined to Azure AD. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave regionName as empty. If you choose the azureADJoin type, provide a value for either onPremisesConnectionId or regionName. The possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
    type_escaped *CloudPcDomainJoinType
}
// NewCloudPcDomainJoinConfiguration instantiates a new cloudPcDomainJoinConfiguration and sets the default values.
func NewCloudPcDomainJoinConfiguration()(*CloudPcDomainJoinConfiguration) {
    m := &CloudPcDomainJoinConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.cloudPcDomainJoinConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDomainJoinConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcDomainJoinConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcDomainJoinConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["onPremisesConnectionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnPremisesConnectionId)
    res["regionName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegionName)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcDomainJoinType , m.SetType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcDomainJoinConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetOnPremisesConnectionId gets the onPremisesConnectionId property value. The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Azure AD joined. If you enter an onPremisesConnectionId, leave regionName as empty.
func (m *CloudPcDomainJoinConfiguration) GetOnPremisesConnectionId()(*string) {
    return m.onPremisesConnectionId
}
// GetRegionName gets the regionName property value. The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. The underlying virtual network will be created and managed by the Windows 365 service. This can only be entered if the IT admin chooses Azure AD joined as the domain join type. If you enter a regionName, leave onPremisesConnectionId as empty.
func (m *CloudPcDomainJoinConfiguration) GetRegionName()(*string) {
    return m.regionName
}
// GetType gets the type property value. Specifies how the provisioned Cloud PC will be joined to Azure AD. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave regionName as empty. If you choose the azureADJoin type, provide a value for either onPremisesConnectionId or regionName. The possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcDomainJoinConfiguration) GetType()(*CloudPcDomainJoinType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *CloudPcDomainJoinConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onPremisesConnectionId", m.GetOnPremisesConnectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("regionName", m.GetRegionName())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *CloudPcDomainJoinConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcDomainJoinConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnPremisesConnectionId sets the onPremisesConnectionId property value. The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Azure AD joined. If you enter an onPremisesConnectionId, leave regionName as empty.
func (m *CloudPcDomainJoinConfiguration) SetOnPremisesConnectionId(value *string)() {
    m.onPremisesConnectionId = value
}
// SetRegionName sets the regionName property value. The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. The underlying virtual network will be created and managed by the Windows 365 service. This can only be entered if the IT admin chooses Azure AD joined as the domain join type. If you enter a regionName, leave onPremisesConnectionId as empty.
func (m *CloudPcDomainJoinConfiguration) SetRegionName(value *string)() {
    m.regionName = value
}
// SetType sets the type property value. Specifies how the provisioned Cloud PC will be joined to Azure AD. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave regionName as empty. If you choose the azureADJoin type, provide a value for either onPremisesConnectionId or regionName. The possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcDomainJoinConfiguration) SetType(value *CloudPcDomainJoinType)() {
    m.type_escaped = value
}
