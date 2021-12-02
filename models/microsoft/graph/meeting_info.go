package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingInfo 
type MeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    allowConversationWithoutHost *bool;
}
// NewMeetingInfo instantiates a new meetingInfo and sets the default values.
func NewMeetingInfo()(*MeetingInfo) {
    m := &MeetingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowConversationWithoutHost gets the allowConversationWithoutHost property value. 
func (m *MeetingInfo) GetAllowConversationWithoutHost()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowConversationWithoutHost
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowConversationWithoutHost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowConversationWithoutHost(val)
        }
        return nil
    }
    return res
}
func (m *MeetingInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowConversationWithoutHost", m.GetAllowConversationWithoutHost())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowConversationWithoutHost sets the allowConversationWithoutHost property value. 
func (m *MeetingInfo) SetAllowConversationWithoutHost(value *bool)() {
    if m != nil {
        m.allowConversationWithoutHost = value
    }
}
