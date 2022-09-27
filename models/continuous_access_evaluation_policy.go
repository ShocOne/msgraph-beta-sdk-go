package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContinuousAccessEvaluationPolicy 
type ContinuousAccessEvaluationPolicy struct {
    Entity
    // Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
    description *string
    // The value is always Continuous Access Evaluation. Read-only.
    displayName *string
    // The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty. Read-only.
    groups []string
    // true to indicate whether continuous access evaluation should be performed; otherwise false. Read-only.
    isEnabled *bool
    // true to indicate that the continuous access evaluation policy settings should be or has been migrated to the conditional access policy.
    migrate *bool
    // The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty. Read-only.
    users []string
}
// NewContinuousAccessEvaluationPolicy instantiates a new ContinuousAccessEvaluationPolicy and sets the default values.
func NewContinuousAccessEvaluationPolicy()(*ContinuousAccessEvaluationPolicy) {
    m := &ContinuousAccessEvaluationPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.continuousAccessEvaluationPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateContinuousAccessEvaluationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContinuousAccessEvaluationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContinuousAccessEvaluationPolicy(), nil
}
// GetDescription gets the description property value. Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The value is always Continuous Access Evaluation. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContinuousAccessEvaluationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["groups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetGroups)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["migrate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetMigrate)
    res["users"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetUsers)
    return res
}
// GetGroups gets the groups property value. The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetGroups()([]string) {
    return m.groups
}
// GetIsEnabled gets the isEnabled property value. true to indicate whether continuous access evaluation should be performed; otherwise false. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetMigrate gets the migrate property value. true to indicate that the continuous access evaluation policy settings should be or has been migrated to the conditional access policy.
func (m *ContinuousAccessEvaluationPolicy) GetMigrate()(*bool) {
    return m.migrate
}
// GetUsers gets the users property value. The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) GetUsers()([]string) {
    return m.users
}
// Serialize serializes information the current object
func (m *ContinuousAccessEvaluationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        err = writer.WriteCollectionOfStringValues("groups", m.GetGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("migrate", m.GetMigrate())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        err = writer.WriteCollectionOfStringValues("users", m.GetUsers())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Continuous access evaluation automatically blocks access to resources and applications in near real time when a user's access is removed or a client IP address changes. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The value is always Continuous Access Evaluation. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroups sets the groups property value. The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetGroups(value []string)() {
    m.groups = value
}
// SetIsEnabled sets the isEnabled property value. true to indicate whether continuous access evaluation should be performed; otherwise false. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetMigrate sets the migrate property value. true to indicate that the continuous access evaluation policy settings should be or has been migrated to the conditional access policy.
func (m *ContinuousAccessEvaluationPolicy) SetMigrate(value *bool)() {
    m.migrate = value
}
// SetUsers sets the users property value. The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty. Read-only.
func (m *ContinuousAccessEvaluationPolicy) SetUsers(value []string)() {
    m.users = value
}
