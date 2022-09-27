package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationRepeatOffender 
type AttackSimulationRepeatOffender struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // User in an attack simulation and training campaign.
    attackSimulationUser AttackSimulationUserable
    // The OdataType property
    odataType *string
    // Number of repeat offences of the user in attack simulation and training campaigns.
    repeatOffenceCount *int32
}
// NewAttackSimulationRepeatOffender instantiates a new attackSimulationRepeatOffender and sets the default values.
func NewAttackSimulationRepeatOffender()(*AttackSimulationRepeatOffender) {
    m := &AttackSimulationRepeatOffender{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.attackSimulationRepeatOffender";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAttackSimulationRepeatOffenderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationRepeatOffenderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationRepeatOffender(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationRepeatOffender) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttackSimulationUser gets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationRepeatOffender) GetAttackSimulationUser()(AttackSimulationUserable) {
    return m.attackSimulationUser
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationRepeatOffender) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attackSimulationUser"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAttackSimulationUserFromDiscriminatorValue , m.SetAttackSimulationUser)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["repeatOffenceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRepeatOffenceCount)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttackSimulationRepeatOffender) GetOdataType()(*string) {
    return m.odataType
}
// GetRepeatOffenceCount gets the repeatOffenceCount property value. Number of repeat offences of the user in attack simulation and training campaigns.
func (m *AttackSimulationRepeatOffender) GetRepeatOffenceCount()(*int32) {
    return m.repeatOffenceCount
}
// Serialize serializes information the current object
func (m *AttackSimulationRepeatOffender) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attackSimulationUser", m.GetAttackSimulationUser())
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
        err := writer.WriteInt32Value("repeatOffenceCount", m.GetRepeatOffenceCount())
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
func (m *AttackSimulationRepeatOffender) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttackSimulationUser sets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationRepeatOffender) SetAttackSimulationUser(value AttackSimulationUserable)() {
    m.attackSimulationUser = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttackSimulationRepeatOffender) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRepeatOffenceCount sets the repeatOffenceCount property value. Number of repeat offences of the user in attack simulation and training campaigns.
func (m *AttackSimulationRepeatOffender) SetRepeatOffenceCount(value *int32)() {
    m.repeatOffenceCount = value
}
