package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Shift 
type Shift struct {
    ChangeTrackedEntity
    // The draft version of this shift that is viewable by managers. Required.
    draftShift ShiftItemable
    // The isStagedForDeletion property
    isStagedForDeletion *bool
    // ID of the scheduling group the shift is part of. Required.
    schedulingGroupId *string
    // The shared version of this shift that is viewable by both employees and managers. Required.
    sharedShift ShiftItemable
    // ID of the user assigned to the shift. Required.
    userId *string
}
// NewShift instantiates a new Shift and sets the default values.
func NewShift()(*Shift) {
    m := &Shift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.shift";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateShiftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateShiftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShift(), nil
}
// GetDraftShift gets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
func (m *Shift) GetDraftShift()(ShiftItemable) {
    return m.draftShift
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Shift) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftShift"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateShiftItemFromDiscriminatorValue , m.SetDraftShift)
    res["isStagedForDeletion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsStagedForDeletion)
    res["schedulingGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSchedulingGroupId)
    res["sharedShift"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateShiftItemFromDiscriminatorValue , m.SetSharedShift)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetIsStagedForDeletion gets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *Shift) GetIsStagedForDeletion()(*bool) {
    return m.isStagedForDeletion
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) GetSchedulingGroupId()(*string) {
    return m.schedulingGroupId
}
// GetSharedShift gets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
func (m *Shift) GetSharedShift()(ShiftItemable) {
    return m.sharedShift
}
// GetUserId gets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *Shift) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftShift", m.GetDraftShift())
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
        err = writer.WriteObjectValue("sharedShift", m.GetSharedShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftShift sets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
func (m *Shift) SetDraftShift(value ShiftItemable)() {
    m.draftShift = value
}
// SetIsStagedForDeletion sets the isStagedForDeletion property value. The isStagedForDeletion property
func (m *Shift) SetIsStagedForDeletion(value *bool)() {
    m.isStagedForDeletion = value
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// SetSharedShift sets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
func (m *Shift) SetSharedShift(value ShiftItemable)() {
    m.sharedShift = value
}
// SetUserId sets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) SetUserId(value *string)() {
    m.userId = value
}
