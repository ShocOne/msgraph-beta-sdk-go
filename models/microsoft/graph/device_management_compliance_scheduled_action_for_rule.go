package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementComplianceScheduledActionForRule 
type DeviceManagementComplianceScheduledActionForRule struct {
    Entity
    // Name of the rule which this scheduled action applies to.
    ruleName *string;
    // The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
    scheduledActionConfigurations []DeviceManagementComplianceActionItem;
}
// NewDeviceManagementComplianceScheduledActionForRule instantiates a new deviceManagementComplianceScheduledActionForRule and sets the default values.
func NewDeviceManagementComplianceScheduledActionForRule()(*DeviceManagementComplianceScheduledActionForRule) {
    m := &DeviceManagementComplianceScheduledActionForRule{
        Entity: *NewEntity(),
    }
    return m
}
// GetRuleName gets the ruleName property value. Name of the rule which this scheduled action applies to.
func (m *DeviceManagementComplianceScheduledActionForRule) GetRuleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleName
    }
}
// GetScheduledActionConfigurations gets the scheduledActionConfigurations property value. The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceScheduledActionForRule) GetScheduledActionConfigurations()([]DeviceManagementComplianceActionItem) {
    if m == nil {
        return nil
    } else {
        return m.scheduledActionConfigurations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplianceScheduledActionForRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ruleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    res["scheduledActionConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementComplianceActionItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementComplianceActionItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementComplianceActionItem))
            }
            m.SetScheduledActionConfigurations(res)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementComplianceScheduledActionForRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementComplianceScheduledActionForRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScheduledActionConfigurations()))
        for i, v := range m.GetScheduledActionConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scheduledActionConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRuleName sets the ruleName property value. Name of the rule which this scheduled action applies to.
func (m *DeviceManagementComplianceScheduledActionForRule) SetRuleName(value *string)() {
    if m != nil {
        m.ruleName = value
    }
}
// SetScheduledActionConfigurations sets the scheduledActionConfigurations property value. The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceScheduledActionForRule) SetScheduledActionConfigurations(value []DeviceManagementComplianceActionItem)() {
    if m != nil {
        m.scheduledActionConfigurations = value
    }
}