package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationTeacher 
type EducationTeacher struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Id of the Teacher in external source system.
    externalId *string
    // The OdataType property
    odataType *string
    // Teacher number.
    teacherNumber *string
}
// NewEducationTeacher instantiates a new educationTeacher and sets the default values.
func NewEducationTeacher()(*EducationTeacher) {
    m := &EducationTeacher{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.educationTeacher";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationTeacherFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationTeacherFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationTeacher(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationTeacher) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExternalId gets the externalId property value. Id of the Teacher in external source system.
func (m *EducationTeacher) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationTeacher) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["teacherNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTeacherNumber)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationTeacher) GetOdataType()(*string) {
    return m.odataType
}
// GetTeacherNumber gets the teacherNumber property value. Teacher number.
func (m *EducationTeacher) GetTeacherNumber()(*string) {
    return m.teacherNumber
}
// Serialize serializes information the current object
func (m *EducationTeacher) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
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
        err := writer.WriteStringValue("teacherNumber", m.GetTeacherNumber())
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
func (m *EducationTeacher) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExternalId sets the externalId property value. Id of the Teacher in external source system.
func (m *EducationTeacher) SetExternalId(value *string)() {
    m.externalId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationTeacher) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTeacherNumber sets the teacherNumber property value. Teacher number.
func (m *EducationTeacher) SetTeacherNumber(value *string)() {
    m.teacherNumber = value
}
