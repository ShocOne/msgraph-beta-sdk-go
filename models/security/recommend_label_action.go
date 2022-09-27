package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendLabelAction 
type RecommendLabelAction struct {
    InformationProtectionAction
    // Actions to take if the label is accepted by the user.
    actions []InformationProtectionActionable
    // The actionSource property
    actionSource *ActionSource
    // The sensitive information type GUIDs that caused the recommendation to be given.
    responsibleSensitiveTypeIds []string
    // The sensitivityLabelId property
    sensitivityLabelId *string
}
// NewRecommendLabelAction instantiates a new RecommendLabelAction and sets the default values.
func NewRecommendLabelAction()(*RecommendLabelAction) {
    m := &RecommendLabelAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.recommendLabelAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRecommendLabelActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecommendLabelActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecommendLabelAction(), nil
}
// GetActions gets the actions property value. Actions to take if the label is accepted by the user.
func (m *RecommendLabelAction) GetActions()([]InformationProtectionActionable) {
    return m.actions
}
// GetActionSource gets the actionSource property value. The actionSource property
func (m *RecommendLabelAction) GetActionSource()(*ActionSource) {
    return m.actionSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecommendLabelAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["actions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateInformationProtectionActionFromDiscriminatorValue , m.SetActions)
    res["actionSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseActionSource , m.SetActionSource)
    res["responsibleSensitiveTypeIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetResponsibleSensitiveTypeIds)
    res["sensitivityLabelId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSensitivityLabelId)
    return res
}
// GetResponsibleSensitiveTypeIds gets the responsibleSensitiveTypeIds property value. The sensitive information type GUIDs that caused the recommendation to be given.
func (m *RecommendLabelAction) GetResponsibleSensitiveTypeIds()([]string) {
    return m.responsibleSensitiveTypeIds
}
// GetSensitivityLabelId gets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *RecommendLabelAction) GetSensitivityLabelId()(*string) {
    return m.sensitivityLabelId
}
// Serialize serializes information the current object
func (m *RecommendLabelAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActions())
        err = writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetActionSource() != nil {
        cast := (*m.GetActionSource()).String()
        err = writer.WriteStringValue("actionSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetResponsibleSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("responsibleSensitiveTypeIds", m.GetResponsibleSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sensitivityLabelId", m.GetSensitivityLabelId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. Actions to take if the label is accepted by the user.
func (m *RecommendLabelAction) SetActions(value []InformationProtectionActionable)() {
    m.actions = value
}
// SetActionSource sets the actionSource property value. The actionSource property
func (m *RecommendLabelAction) SetActionSource(value *ActionSource)() {
    m.actionSource = value
}
// SetResponsibleSensitiveTypeIds sets the responsibleSensitiveTypeIds property value. The sensitive information type GUIDs that caused the recommendation to be given.
func (m *RecommendLabelAction) SetResponsibleSensitiveTypeIds(value []string)() {
    m.responsibleSensitiveTypeIds = value
}
// SetSensitivityLabelId sets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *RecommendLabelAction) SetSensitivityLabelId(value *string)() {
    m.sensitivityLabelId = value
}
