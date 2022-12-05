package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse provides operations to call the getAttackSimulationTrainingUserCoverage method.
type PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []AttackSimulationTrainingUserCoverageable
}
// NewPrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse instantiates a new PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse and sets the default values.
func NewPrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse()(*PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse) {
    m := &PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttackSimulationTrainingUserCoverageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttackSimulationTrainingUserCoverageable, len(val))
            for i, v := range val {
                res[i] = v.(AttackSimulationTrainingUserCoverageable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse) GetValue()([]AttackSimulationTrainingUserCoverageable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *PrintReportsSecurityGetAttackSimulationTrainingUserCoverageResponse) SetValue(value []AttackSimulationTrainingUserCoverageable)() {
    m.value = value
}
