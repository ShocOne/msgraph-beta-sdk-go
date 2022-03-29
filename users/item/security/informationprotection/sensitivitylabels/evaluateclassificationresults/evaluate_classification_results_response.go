package evaluateclassificationresults

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
)

// EvaluateClassificationResultsResponse provides operations to call the evaluateClassificationResults method.
type EvaluateClassificationResultsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    value []i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.InformationProtectionActionable;
}
// NewEvaluateClassificationResultsResponse instantiates a new evaluateClassificationResultsResponse and sets the default values.
func NewEvaluateClassificationResultsResponse()(*EvaluateClassificationResultsResponse) {
    m := &EvaluateClassificationResultsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateClassificationResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateClassificationResultsResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEvaluateClassificationResultsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateClassificationResultsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateClassificationResultsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.CreateInformationProtectionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.InformationProtectionActionable, len(val))
            for i, v := range val {
                res[i] = v.(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.InformationProtectionActionable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. 
func (m *EvaluateClassificationResultsResponse) GetValue()([]i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.InformationProtectionActionable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *EvaluateClassificationResultsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *EvaluateClassificationResultsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. 
func (m *EvaluateClassificationResultsResponse) SetValue(value []i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.InformationProtectionActionable)() {
    if m != nil {
        m.value = value
    }
}