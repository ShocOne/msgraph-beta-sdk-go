package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernancePolicy 
type GovernancePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The decisionMakerCriteria property
    decisionMakerCriteria []GovernanceCriteriaable
    // The notificationPolicy property
    notificationPolicy GovernanceNotificationPolicyable
    // The OdataType property
    odataType *string
}
// NewGovernancePolicy instantiates a new governancePolicy and sets the default values.
func NewGovernancePolicy()(*GovernancePolicy) {
    m := &GovernancePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.governancePolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGovernancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernancePolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernancePolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDecisionMakerCriteria gets the decisionMakerCriteria property value. The decisionMakerCriteria property
func (m *GovernancePolicy) GetDecisionMakerCriteria()([]GovernanceCriteriaable) {
    return m.decisionMakerCriteria
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decisionMakerCriteria"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceCriteriaFromDiscriminatorValue , m.SetDecisionMakerCriteria)
    res["notificationPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGovernanceNotificationPolicyFromDiscriminatorValue , m.SetNotificationPolicy)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetNotificationPolicy gets the notificationPolicy property value. The notificationPolicy property
func (m *GovernancePolicy) GetNotificationPolicy()(GovernanceNotificationPolicyable) {
    return m.notificationPolicy
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GovernancePolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GovernancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDecisionMakerCriteria() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDecisionMakerCriteria())
        err := writer.WriteCollectionOfObjectValues("decisionMakerCriteria", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notificationPolicy", m.GetNotificationPolicy())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernancePolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDecisionMakerCriteria sets the decisionMakerCriteria property value. The decisionMakerCriteria property
func (m *GovernancePolicy) SetDecisionMakerCriteria(value []GovernanceCriteriaable)() {
    m.decisionMakerCriteria = value
}
// SetNotificationPolicy sets the notificationPolicy property value. The notificationPolicy property
func (m *GovernancePolicy) SetNotificationPolicy(value GovernanceNotificationPolicyable)() {
    m.notificationPolicy = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GovernancePolicy) SetOdataType(value *string)() {
    m.odataType = value
}
