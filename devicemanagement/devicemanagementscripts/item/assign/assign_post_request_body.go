package assign

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AssignPostRequestBody provides operations to call the assign method.
type AssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceManagementScriptAssignments property
    deviceManagementScriptAssignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable
    // The deviceManagementScriptGroupAssignments property
    deviceManagementScriptGroupAssignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable
}
// NewAssignPostRequestBody instantiates a new assignPostRequestBody and sets the default values.
func NewAssignPostRequestBody()(*AssignPostRequestBody) {
    m := &AssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceManagementScriptAssignments gets the deviceManagementScriptAssignments property value. The deviceManagementScriptAssignments property
func (m *AssignPostRequestBody) GetDeviceManagementScriptAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable) {
    return m.deviceManagementScriptAssignments
}
// GetDeviceManagementScriptGroupAssignments gets the deviceManagementScriptGroupAssignments property value. The deviceManagementScriptGroupAssignments property
func (m *AssignPostRequestBody) GetDeviceManagementScriptGroupAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable) {
    return m.deviceManagementScriptGroupAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceManagementScriptAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptAssignmentFromDiscriminatorValue , m.SetDeviceManagementScriptAssignments)
    res["deviceManagementScriptGroupAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue , m.SetDeviceManagementScriptGroupAssignments)
    return res
}
// Serialize serializes information the current object
func (m *AssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceManagementScriptAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceManagementScriptAssignments())
        err := writer.WriteCollectionOfObjectValues("deviceManagementScriptAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementScriptGroupAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceManagementScriptGroupAssignments())
        err := writer.WriteCollectionOfObjectValues("deviceManagementScriptGroupAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceManagementScriptAssignments sets the deviceManagementScriptAssignments property value. The deviceManagementScriptAssignments property
func (m *AssignPostRequestBody) SetDeviceManagementScriptAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable)() {
    m.deviceManagementScriptAssignments = value
}
// SetDeviceManagementScriptGroupAssignments sets the deviceManagementScriptGroupAssignments property value. The deviceManagementScriptGroupAssignments property
func (m *AssignPostRequestBody) SetDeviceManagementScriptGroupAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable)() {
    m.deviceManagementScriptGroupAssignments = value
}
