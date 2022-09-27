package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanDetails 
type PlannerPlanDetails struct {
    PlannerDelta
    // An object that specifies the descriptions of the 25 categories that can be associated with tasks in the plan
    categoryDescriptions PlannerCategoryDescriptionsable
    // Read-only. A collection of additional information associated with plannerPlanContext entries that are defined for the plannerPlan container.
    contextDetails PlannerPlanContextDetailsCollectionable
    // The set of user IDs that this plan is shared with. If you are using Microsoft 365 groups, use the groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection, although it is not required in order for them to access the plan owned by the group.
    sharedWith PlannerUserIdsable
}
// NewPlannerPlanDetails instantiates a new plannerPlanDetails and sets the default values.
func NewPlannerPlanDetails()(*PlannerPlanDetails) {
    m := &PlannerPlanDetails{
        PlannerDelta: *NewPlannerDelta(),
    }
    odataTypeValue := "#microsoft.graph.plannerPlanDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerPlanDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanDetails(), nil
}
// GetCategoryDescriptions gets the categoryDescriptions property value. An object that specifies the descriptions of the 25 categories that can be associated with tasks in the plan
func (m *PlannerPlanDetails) GetCategoryDescriptions()(PlannerCategoryDescriptionsable) {
    return m.categoryDescriptions
}
// GetContextDetails gets the contextDetails property value. Read-only. A collection of additional information associated with plannerPlanContext entries that are defined for the plannerPlan container.
func (m *PlannerPlanDetails) GetContextDetails()(PlannerPlanContextDetailsCollectionable) {
    return m.contextDetails
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["categoryDescriptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerCategoryDescriptionsFromDiscriminatorValue , m.SetCategoryDescriptions)
    res["contextDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerPlanContextDetailsCollectionFromDiscriminatorValue , m.SetContextDetails)
    res["sharedWith"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerUserIdsFromDiscriminatorValue , m.SetSharedWith)
    return res
}
// GetSharedWith gets the sharedWith property value. The set of user IDs that this plan is shared with. If you are using Microsoft 365 groups, use the groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection, although it is not required in order for them to access the plan owned by the group.
func (m *PlannerPlanDetails) GetSharedWith()(PlannerUserIdsable) {
    return m.sharedWith
}
// Serialize serializes information the current object
func (m *PlannerPlanDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("categoryDescriptions", m.GetCategoryDescriptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contextDetails", m.GetContextDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedWith", m.GetSharedWith())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategoryDescriptions sets the categoryDescriptions property value. An object that specifies the descriptions of the 25 categories that can be associated with tasks in the plan
func (m *PlannerPlanDetails) SetCategoryDescriptions(value PlannerCategoryDescriptionsable)() {
    m.categoryDescriptions = value
}
// SetContextDetails sets the contextDetails property value. Read-only. A collection of additional information associated with plannerPlanContext entries that are defined for the plannerPlan container.
func (m *PlannerPlanDetails) SetContextDetails(value PlannerPlanContextDetailsCollectionable)() {
    m.contextDetails = value
}
// SetSharedWith sets the sharedWith property value. The set of user IDs that this plan is shared with. If you are using Microsoft 365 groups, use the groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection, although it is not required in order for them to access the plan owned by the group.
func (m *PlannerPlanDetails) SetSharedWith(value PlannerUserIdsable)() {
    m.sharedWith = value
}
