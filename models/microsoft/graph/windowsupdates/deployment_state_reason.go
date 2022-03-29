package windowsupdates

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeploymentStateReason 
type DeploymentStateReason struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies a reason for the deployment state. Possible values are: scheduledByOfferWindow, offeringByRequest, pausedByRequest, pausedByMonitoring. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: faultedByContentOutdated. Read-only.
    value *DeploymentStateReasonValue;
}
// NewDeploymentStateReason instantiates a new deploymentStateReason and sets the default values.
func NewDeploymentStateReason()(*DeploymentStateReason) {
    m := &DeploymentStateReason{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeploymentStateReasonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentStateReasonFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeploymentStateReason(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentStateReason) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentStateReason) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentStateReasonValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(*DeploymentStateReasonValue))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Specifies a reason for the deployment state. Possible values are: scheduledByOfferWindow, offeringByRequest, pausedByRequest, pausedByMonitoring. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: faultedByContentOutdated. Read-only.
func (m *DeploymentStateReason) GetValue()(*DeploymentStateReasonValue) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *DeploymentStateReason) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := (*m.GetValue()).String()
        err := writer.WriteStringValue("value", &cast)
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
func (m *DeploymentStateReason) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. Specifies a reason for the deployment state. Possible values are: scheduledByOfferWindow, offeringByRequest, pausedByRequest, pausedByMonitoring. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: faultedByContentOutdated. Read-only.
func (m *DeploymentStateReason) SetValue(value *DeploymentStateReasonValue)() {
    if m != nil {
        m.value = value
    }
}