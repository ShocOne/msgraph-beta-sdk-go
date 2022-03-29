package replyall

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ReplyAllRequestBody provides operations to call the replyAll method.
type ReplyAllRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    comment *string;
    // 
    message i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Messageable;
}
// NewReplyAllRequestBody instantiates a new replyAllRequestBody and sets the default values.
func NewReplyAllRequestBody()(*ReplyAllRequestBody) {
    m := &ReplyAllRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReplyAllRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReplyAllRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewReplyAllRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReplyAllRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComment gets the comment property value. 
func (m *ReplyAllRequestBody) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReplyAllRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Messageable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. 
func (m *ReplyAllRequestBody) GetMessage()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Messageable) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Serialize serializes information the current object
func (m *ReplyAllRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
func (m *ReplyAllRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComment sets the comment property value. 
func (m *ReplyAllRequestBody) SetComment(value *string)() {
    if m != nil {
        m.comment = value
    }
}
// SetMessage sets the message property value. 
func (m *ReplyAllRequestBody) SetMessage(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Messageable)() {
    if m != nil {
        m.message = value
    }
}