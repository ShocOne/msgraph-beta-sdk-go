package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CasesRoot 
type CasesRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The ediscoveryCases property
    ediscoveryCases []EdiscoveryCaseable
}
// NewCasesRoot instantiates a new casesRoot and sets the default values.
func NewCasesRoot()(*CasesRoot) {
    m := &CasesRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security.casesRoot";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCasesRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesRoot(), nil
}
// GetEdiscoveryCases gets the ediscoveryCases property value. The ediscoveryCases property
func (m *CasesRoot) GetEdiscoveryCases()([]EdiscoveryCaseable) {
    return m.ediscoveryCases
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ediscoveryCases"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEdiscoveryCaseFromDiscriminatorValue , m.SetEdiscoveryCases)
    return res
}
// Serialize serializes information the current object
func (m *CasesRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEdiscoveryCases() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEdiscoveryCases())
        err = writer.WriteCollectionOfObjectValues("ediscoveryCases", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEdiscoveryCases sets the ediscoveryCases property value. The ediscoveryCases property
func (m *CasesRoot) SetEdiscoveryCases(value []EdiscoveryCaseable)() {
    m.ediscoveryCases = value
}
