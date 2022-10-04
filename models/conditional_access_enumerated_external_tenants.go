package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessEnumeratedExternalTenants 
type ConditionalAccessEnumeratedExternalTenants struct {
    ConditionalAccessExternalTenants
    // Represents a collection of tenant ids in the scope of Conditional Access for guests and external users policy targeting.
    members []string
}
// NewConditionalAccessEnumeratedExternalTenants instantiates a new ConditionalAccessEnumeratedExternalTenants and sets the default values.
func NewConditionalAccessEnumeratedExternalTenants()(*ConditionalAccessEnumeratedExternalTenants) {
    m := &ConditionalAccessEnumeratedExternalTenants{
        ConditionalAccessExternalTenants: *NewConditionalAccessExternalTenants(),
    }
    odataTypeValue := "#microsoft.graph.conditionalAccessEnumeratedExternalTenants";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConditionalAccessEnumeratedExternalTenantsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessEnumeratedExternalTenantsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessEnumeratedExternalTenants(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessEnumeratedExternalTenants) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConditionalAccessExternalTenants.GetFieldDeserializers()
    res["members"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetMembers)
    return res
}
// GetMembers gets the members property value. Represents a collection of tenant ids in the scope of Conditional Access for guests and external users policy targeting.
func (m *ConditionalAccessEnumeratedExternalTenants) GetMembers()([]string) {
    return m.members
}
// Serialize serializes information the current object
func (m *ConditionalAccessEnumeratedExternalTenants) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConditionalAccessExternalTenants.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        err = writer.WriteCollectionOfStringValues("members", m.GetMembers())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMembers sets the members property value. Represents a collection of tenant ids in the scope of Conditional Access for guests and external users policy targeting.
func (m *ConditionalAccessEnumeratedExternalTenants) SetMembers(value []string)() {
    m.members = value
}
