package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintCertificateSigningRequest 
type PrintCertificateSigningRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A base64-encoded pkcs10 certificate request. Read-only.
    content *string;
    // The base64-encoded public portion of an asymmetric key that is generated by the client. Read-only.
    transportKey *string;
}
// NewPrintCertificateSigningRequest instantiates a new printCertificateSigningRequest and sets the default values.
func NewPrintCertificateSigningRequest()(*PrintCertificateSigningRequest) {
    m := &PrintCertificateSigningRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintCertificateSigningRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContent gets the content property value. A base64-encoded pkcs10 certificate request. Read-only.
func (m *PrintCertificateSigningRequest) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetTransportKey gets the transportKey property value. The base64-encoded public portion of an asymmetric key that is generated by the client. Read-only.
func (m *PrintCertificateSigningRequest) GetTransportKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.transportKey
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintCertificateSigningRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["transportKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransportKey(val)
        }
        return nil
    }
    return res
}
func (m *PrintCertificateSigningRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrintCertificateSigningRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("transportKey", m.GetTransportKey())
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
func (m *PrintCertificateSigningRequest) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContent sets the content property value. A base64-encoded pkcs10 certificate request. Read-only.
func (m *PrintCertificateSigningRequest) SetContent(value *string)() {
    if m != nil {
        m.content = value
    }
}
// SetTransportKey sets the transportKey property value. The base64-encoded public portion of an asymmetric key that is generated by the client. Read-only.
func (m *PrintCertificateSigningRequest) SetTransportKey(value *string)() {
    if m != nil {
        m.transportKey = value
    }
}
