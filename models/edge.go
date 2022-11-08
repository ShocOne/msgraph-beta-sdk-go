package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Edge 
type Edge struct {
    Entity
    // A container for Internet Explorer mode resources.
    internetExplorerMode InternetExplorerModeable
}
// NewEdge instantiates a new edge and sets the default values.
func NewEdge()(*Edge) {
    m := &Edge{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.edge";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEdgeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdge(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Edge) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["internetExplorerMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInternetExplorerModeFromDiscriminatorValue , m.SetInternetExplorerMode)
    return res
}
// GetInternetExplorerMode gets the internetExplorerMode property value. A container for Internet Explorer mode resources.
func (m *Edge) GetInternetExplorerMode()(InternetExplorerModeable) {
    return m.internetExplorerMode
}
// Serialize serializes information the current object
func (m *Edge) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("internetExplorerMode", m.GetInternetExplorerMode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInternetExplorerMode sets the internetExplorerMode property value. A container for Internet Explorer mode resources.
func (m *Edge) SetInternetExplorerMode(value InternetExplorerModeable)() {
    m.internetExplorerMode = value
}
