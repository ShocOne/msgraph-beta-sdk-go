package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody provides operations to call the changeDeploymentStatus method.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managementTemplateStepId property
    managementTemplateStepId *string
    // The status property
    status *string
    // The tenantId property
    tenantId *string
}
// NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody instantiates a new TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody()(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) {
    m := &TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementTemplateStepId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetManagementTemplateStepId gets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) GetManagementTemplateStepId()(*string) {
    return m.managementTemplateStepId
}
// GetStatus gets the status property value. The status property
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) GetStatus()(*string) {
    return m.status
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementTemplateStepId", m.GetManagementTemplateStepId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagementTemplateStepId sets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) SetManagementTemplateStepId(value *string)() {
    m.managementTemplateStepId = value
}
// SetStatus sets the status property value. The status property
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) SetStatus(value *string)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
