package externalconnectors

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ExternalItem provides operations to manage the collection of activityStatistics entities.
type ExternalItem struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // An array of access control entries. Each entry specifies the access granted to a user or group. Required.
    acl []Aclable
    // Write-only property. Returns results.
    activities []ExternalActivityable
    // A plain-text representation of the contents of the item. The text in this property is full-text indexed. Optional.
    content ExternalItemContentable
    // A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
    properties Propertiesable
}
// NewExternalItem instantiates a new externalItem and sets the default values.
func NewExternalItem()(*ExternalItem) {
    m := &ExternalItem{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.externalConnectors.externalItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExternalItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalItem(), nil
}
// GetAcl gets the acl property value. An array of access control entries. Each entry specifies the access granted to a user or group. Required.
func (m *ExternalItem) GetAcl()([]Aclable) {
    return m.acl
}
// GetActivities gets the activities property value. Write-only property. Returns results.
func (m *ExternalItem) GetActivities()([]ExternalActivityable) {
    return m.activities
}
// GetContent gets the content property value. A plain-text representation of the contents of the item. The text in this property is full-text indexed. Optional.
func (m *ExternalItem) GetContent()(ExternalItemContentable) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAclFromDiscriminatorValue , m.SetAcl)
    res["activities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExternalActivityFromDiscriminatorValue , m.SetActivities)
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateExternalItemContentFromDiscriminatorValue , m.SetContent)
    res["properties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePropertiesFromDiscriminatorValue , m.SetProperties)
    return res
}
// GetProperties gets the properties property value. A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
func (m *ExternalItem) GetProperties()(Propertiesable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *ExternalItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAcl() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAcl())
        err = writer.WriteCollectionOfObjectValues("acl", cast)
        if err != nil {
            return err
        }
    }
    if m.GetActivities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActivities())
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("properties", m.GetProperties())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcl sets the acl property value. An array of access control entries. Each entry specifies the access granted to a user or group. Required.
func (m *ExternalItem) SetAcl(value []Aclable)() {
    m.acl = value
}
// SetActivities sets the activities property value. Write-only property. Returns results.
func (m *ExternalItem) SetActivities(value []ExternalActivityable)() {
    m.activities = value
}
// SetContent sets the content property value. A plain-text representation of the contents of the item. The text in this property is full-text indexed. Optional.
func (m *ExternalItem) SetContent(value ExternalItemContentable)() {
    m.content = value
}
// SetProperties sets the properties property value. A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
func (m *ExternalItem) SetProperties(value Propertiesable)() {
    m.properties = value
}
