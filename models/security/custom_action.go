package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomAction 
type CustomAction struct {
    InformationProtectionAction
    // Name of the custom action.
    name *string
    // Properties, in key-value pair format, of the action.
    properties []KeyValuePairable
}
// NewCustomAction instantiates a new CustomAction and sets the default values.
func NewCustomAction()(*CustomAction) {
    m := &CustomAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.customAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["properties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetProperties)
    return res
}
// GetName gets the name property value. Name of the custom action.
func (m *CustomAction) GetName()(*string) {
    return m.name
}
// GetProperties gets the properties property value. Properties, in key-value pair format, of the action.
func (m *CustomAction) GetProperties()([]KeyValuePairable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *CustomAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProperties())
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name of the custom action.
func (m *CustomAction) SetName(value *string)() {
    m.name = value
}
// SetProperties sets the properties property value. Properties, in key-value pair format, of the action.
func (m *CustomAction) SetProperties(value []KeyValuePairable)() {
    m.properties = value
}
