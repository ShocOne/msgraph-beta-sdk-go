package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenShift 
type OpenShift struct {
    ChangeTrackedEntity
    // An unpublished open shift.
    draftOpenShift OpenShiftItemable
    // The isStagedForDeletion property
    isStagedForDeletion *bool
    // ID for the scheduling group that the open shift belongs to.
    schedulingGroupId *string
    // A published open shift.
    sharedOpenShift OpenShiftItemable
}
// NewOpenShift instantiates a new OpenShift and sets the default values.
func NewOpenShift()(*OpenShift) {
    m := &OpenShift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.openShift";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOpenShiftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenShiftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenShift(), nil
}
// GetDraftOpenShift gets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) GetDraftOpenShift()(OpenShiftItemable) {
    return m.draftOpenShift
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenShift) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftOpenShift"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOpenShiftItemFromDiscriminatorValue , m.SetDraftOpenShift)
    res["isStagedForDeletion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsStagedForDeletion)
    res["schedulingGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSchedulingGroupId)
    res["sharedOpenShift"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOpenShiftItemFromDiscriminatorValue , m.SetSharedOpenShift)
    return res
}
// GetIsStagedForDeletion gets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *OpenShift) GetIsStagedForDeletion()(*bool) {
    return m.isStagedForDeletion
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) GetSchedulingGroupId()(*string) {
    return m.schedulingGroupId
}
// GetSharedOpenShift gets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) GetSharedOpenShift()(OpenShiftItemable) {
    return m.sharedOpenShift
}
// Serialize serializes information the current object
func (m *OpenShift) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftOpenShift", m.GetDraftOpenShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isStagedForDeletion", m.GetIsStagedForDeletion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schedulingGroupId", m.GetSchedulingGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedOpenShift", m.GetSharedOpenShift())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftOpenShift sets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) SetDraftOpenShift(value OpenShiftItemable)() {
    m.draftOpenShift = value
}
// SetIsStagedForDeletion sets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *OpenShift) SetIsStagedForDeletion(value *bool)() {
    m.isStagedForDeletion = value
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// SetSharedOpenShift sets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) SetSharedOpenShift(value OpenShiftItemable)() {
    m.sharedOpenShift = value
}
