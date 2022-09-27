package redirect

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// RedirectPostRequestBody provides operations to call the redirect method.
type RedirectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The callbackUri property
    callbackUri *string
    // The maskCallee property
    maskCallee *bool
    // The maskCaller property
    maskCaller *bool
    // The targetDisposition property
    targetDisposition *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition
    // The targets property
    targets []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable
    // The timeout property
    timeout *int32
}
// NewRedirectPostRequestBody instantiates a new redirectPostRequestBody and sets the default values.
func NewRedirectPostRequestBody()(*RedirectPostRequestBody) {
    m := &RedirectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRedirectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRedirectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRedirectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *RedirectPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedirectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCallbackUri)
    res["maskCallee"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetMaskCallee)
    res["maskCaller"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetMaskCaller)
    res["targetDisposition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCallDisposition , m.SetTargetDisposition)
    res["targets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInvitationParticipantInfoFromDiscriminatorValue , m.SetTargets)
    res["timeout"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTimeout)
    return res
}
// GetMaskCallee gets the maskCallee property value. The maskCallee property
func (m *RedirectPostRequestBody) GetMaskCallee()(*bool) {
    return m.maskCallee
}
// GetMaskCaller gets the maskCaller property value. The maskCaller property
func (m *RedirectPostRequestBody) GetMaskCaller()(*bool) {
    return m.maskCaller
}
// GetTargetDisposition gets the targetDisposition property value. The targetDisposition property
func (m *RedirectPostRequestBody) GetTargetDisposition()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition) {
    return m.targetDisposition
}
// GetTargets gets the targets property value. The targets property
func (m *RedirectPostRequestBody) GetTargets()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable) {
    return m.targets
}
// GetTimeout gets the timeout property value. The timeout property
func (m *RedirectPostRequestBody) GetTimeout()(*int32) {
    return m.timeout
}
// Serialize serializes information the current object
func (m *RedirectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCallee", m.GetMaskCallee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCaller", m.GetMaskCaller())
        if err != nil {
            return err
        }
    }
    if m.GetTargetDisposition() != nil {
        cast := (*m.GetTargetDisposition()).String()
        err := writer.WriteStringValue("targetDisposition", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTargets())
        err := writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeout", m.GetTimeout())
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
func (m *RedirectPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *RedirectPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetMaskCallee sets the maskCallee property value. The maskCallee property
func (m *RedirectPostRequestBody) SetMaskCallee(value *bool)() {
    m.maskCallee = value
}
// SetMaskCaller sets the maskCaller property value. The maskCaller property
func (m *RedirectPostRequestBody) SetMaskCaller(value *bool)() {
    m.maskCaller = value
}
// SetTargetDisposition sets the targetDisposition property value. The targetDisposition property
func (m *RedirectPostRequestBody) SetTargetDisposition(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition)() {
    m.targetDisposition = value
}
// SetTargets sets the targets property value. The targets property
func (m *RedirectPostRequestBody) SetTargets(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable)() {
    m.targets = value
}
// SetTimeout sets the timeout property value. The timeout property
func (m *RedirectPostRequestBody) SetTimeout(value *int32)() {
    m.timeout = value
}
