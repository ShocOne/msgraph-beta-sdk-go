package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupPolicyDefinitionFile struct {
    Entity
    // The group policy definitions associated with the file.
    definitions []GroupPolicyDefinition;
    // The localized description of the policy settings in the ADMX file. The default value is empty.
    description *string;
    // The localized friendly name of the ADMX file.
    displayName *string;
    // The supported language codes for the ADMX file.
    languageCodes []string;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
    policyType *GroupPolicyType;
    // The revision version associated with the file.
    revision *string;
    // Specifies the URI used to identify the namespace within the ADMX file.
    targetNamespace *string;
    // Specifies the logical name that refers to the namespace within the ADMX file.
    targetPrefix *string;
}
// Instantiates a new groupPolicyDefinitionFile and sets the default values.
func NewGroupPolicyDefinitionFile()(*GroupPolicyDefinitionFile) {
    m := &GroupPolicyDefinitionFile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the definitions property value. The group policy definitions associated with the file.
func (m *GroupPolicyDefinitionFile) GetDefinitions()([]GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definitions
    }
}
// Gets the description property value. The localized description of the policy settings in the ADMX file. The default value is empty.
func (m *GroupPolicyDefinitionFile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The localized friendly name of the ADMX file.
func (m *GroupPolicyDefinitionFile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the languageCodes property value. The supported language codes for the ADMX file.
func (m *GroupPolicyDefinitionFile) GetLanguageCodes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.languageCodes
    }
}
// Gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
func (m *GroupPolicyDefinitionFile) GetPolicyType()(*GroupPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
// Gets the revision property value. The revision version associated with the file.
func (m *GroupPolicyDefinitionFile) GetRevision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.revision
    }
}
// Gets the targetNamespace property value. Specifies the URI used to identify the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) GetTargetNamespace()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetNamespace
    }
}
// Gets the targetPrefix property value. Specifies the logical name that refers to the namespace within the ADMX file.
func (m *GroupPolicyDefinitionFile) GetTargetPrefix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetPrefix
    }
}
// The deserialization information for the current model
func (m *GroupPolicyDefinitionFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyDefinition))
        }
        m.SetDefinitions(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["languageCodes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetLanguageCodes(res)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["policyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyType)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicyType)
        m.SetPolicyType(&cast)
        return nil
    }
    res["revision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRevision(val)
        return nil
    }
    res["targetNamespace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetNamespace(val)
        return nil
    }
    res["targetPrefix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetPrefix(val)
        return nil
    }
    return res
}
func (m *GroupPolicyDefinitionFile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GroupPolicyDefinitionFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("languageCodes", m.GetLanguageCodes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := m.GetPolicyType().String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("revision", m.GetRevision())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetNamespace", m.GetTargetNamespace())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetPrefix", m.GetTargetPrefix())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the definitions property value. The group policy definitions associated with the file.
// Parameters:
//  - value : Value to set for the definitions property.
func (m *GroupPolicyDefinitionFile) SetDefinitions(value []GroupPolicyDefinition)() {
    m.definitions = value
}
// Sets the description property value. The localized description of the policy settings in the ADMX file. The default value is empty.
// Parameters:
//  - value : Value to set for the description property.
func (m *GroupPolicyDefinitionFile) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The localized friendly name of the ADMX file.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GroupPolicyDefinitionFile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the languageCodes property value. The supported language codes for the ADMX file.
// Parameters:
//  - value : Value to set for the languageCodes property.
func (m *GroupPolicyDefinitionFile) SetLanguageCodes(value []string)() {
    m.languageCodes = value
}
// Sets the lastModifiedDateTime property value. The date and time the entity was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *GroupPolicyDefinitionFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the policyType property value. Specifies the type of group policy. Possible values are: admxBacked, admxIngested.
// Parameters:
//  - value : Value to set for the policyType property.
func (m *GroupPolicyDefinitionFile) SetPolicyType(value *GroupPolicyType)() {
    m.policyType = value
}
// Sets the revision property value. The revision version associated with the file.
// Parameters:
//  - value : Value to set for the revision property.
func (m *GroupPolicyDefinitionFile) SetRevision(value *string)() {
    m.revision = value
}
// Sets the targetNamespace property value. Specifies the URI used to identify the namespace within the ADMX file.
// Parameters:
//  - value : Value to set for the targetNamespace property.
func (m *GroupPolicyDefinitionFile) SetTargetNamespace(value *string)() {
    m.targetNamespace = value
}
// Sets the targetPrefix property value. Specifies the logical name that refers to the namespace within the ADMX file.
// Parameters:
//  - value : Value to set for the targetPrefix property.
func (m *GroupPolicyDefinitionFile) SetTargetPrefix(value *string)() {
    m.targetPrefix = value
}
