package preview

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreviewPostRequestBody provides operations to call the preview method.
type PreviewPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allowEdit property
    allowEdit *bool
    // The chromeless property
    chromeless *bool
    // The page property
    page *string
    // The viewer property
    viewer *string
    // The zoom property
    zoom *float64
}
// NewPreviewPostRequestBody instantiates a new previewPostRequestBody and sets the default values.
func NewPreviewPostRequestBody()(*PreviewPostRequestBody) {
    m := &PreviewPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePreviewPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePreviewPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreviewPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreviewPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowEdit gets the allowEdit property value. The allowEdit property
func (m *PreviewPostRequestBody) GetAllowEdit()(*bool) {
    return m.allowEdit
}
// GetChromeless gets the chromeless property value. The chromeless property
func (m *PreviewPostRequestBody) GetChromeless()(*bool) {
    return m.chromeless
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreviewPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowEdit"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowEdit)
    res["chromeless"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetChromeless)
    res["page"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPage)
    res["viewer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetViewer)
    res["zoom"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetZoom)
    return res
}
// GetPage gets the page property value. The page property
func (m *PreviewPostRequestBody) GetPage()(*string) {
    return m.page
}
// GetViewer gets the viewer property value. The viewer property
func (m *PreviewPostRequestBody) GetViewer()(*string) {
    return m.viewer
}
// GetZoom gets the zoom property value. The zoom property
func (m *PreviewPostRequestBody) GetZoom()(*float64) {
    return m.zoom
}
// Serialize serializes information the current object
func (m *PreviewPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowEdit", m.GetAllowEdit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("chromeless", m.GetChromeless())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("page", m.GetPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("viewer", m.GetViewer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("zoom", m.GetZoom())
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
func (m *PreviewPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowEdit sets the allowEdit property value. The allowEdit property
func (m *PreviewPostRequestBody) SetAllowEdit(value *bool)() {
    m.allowEdit = value
}
// SetChromeless sets the chromeless property value. The chromeless property
func (m *PreviewPostRequestBody) SetChromeless(value *bool)() {
    m.chromeless = value
}
// SetPage sets the page property value. The page property
func (m *PreviewPostRequestBody) SetPage(value *string)() {
    m.page = value
}
// SetViewer sets the viewer property value. The viewer property
func (m *PreviewPostRequestBody) SetViewer(value *string)() {
    m.viewer = value
}
// SetZoom sets the zoom property value. The zoom property
func (m *PreviewPostRequestBody) SetZoom(value *float64)() {
    m.zoom = value
}
