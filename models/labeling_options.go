package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelingOptions 
type LabelingOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentMethod property
    assignmentMethod *AssignmentMethod
    // The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
    downgradeJustification DowngradeJustificationable
    // Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
    extendedProperties []KeyValuePairable
    // The GUID of the label that should be applied to the information.
    labelId *string
    // The OdataType property
    odataType *string
}
// NewLabelingOptions instantiates a new labelingOptions and sets the default values.
func NewLabelingOptions()(*LabelingOptions) {
    m := &LabelingOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.labelingOptions";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLabelingOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelingOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelingOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LabelingOptions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentMethod gets the assignmentMethod property value. The assignmentMethod property
func (m *LabelingOptions) GetAssignmentMethod()(*AssignmentMethod) {
    return m.assignmentMethod
}
// GetDowngradeJustification gets the downgradeJustification property value. The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
func (m *LabelingOptions) GetDowngradeJustification()(DowngradeJustificationable) {
    return m.downgradeJustification
}
// GetExtendedProperties gets the extendedProperties property value. Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
func (m *LabelingOptions) GetExtendedProperties()([]KeyValuePairable) {
    return m.extendedProperties
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelingOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAssignmentMethod , m.SetAssignmentMethod)
    res["downgradeJustification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDowngradeJustificationFromDiscriminatorValue , m.SetDowngradeJustification)
    res["extendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetExtendedProperties)
    res["labelId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLabelId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetLabelId gets the labelId property value. The GUID of the label that should be applied to the information.
func (m *LabelingOptions) GetLabelId()(*string) {
    return m.labelId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LabelingOptions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LabelingOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("downgradeJustification", m.GetDowngradeJustification())
        if err != nil {
            return err
        }
    }
    if m.GetExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtendedProperties())
        err := writer.WriteCollectionOfObjectValues("extendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *LabelingOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentMethod sets the assignmentMethod property value. The assignmentMethod property
func (m *LabelingOptions) SetAssignmentMethod(value *AssignmentMethod)() {
    m.assignmentMethod = value
}
// SetDowngradeJustification sets the downgradeJustification property value. The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
func (m *LabelingOptions) SetDowngradeJustification(value DowngradeJustificationable)() {
    m.downgradeJustification = value
}
// SetExtendedProperties sets the extendedProperties property value. Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
func (m *LabelingOptions) SetExtendedProperties(value []KeyValuePairable)() {
    m.extendedProperties = value
}
// SetLabelId sets the labelId property value. The GUID of the label that should be applied to the information.
func (m *LabelingOptions) SetLabelId(value *string)() {
    m.labelId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LabelingOptions) SetOdataType(value *string)() {
    m.odataType = value
}
