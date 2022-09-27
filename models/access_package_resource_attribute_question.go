package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceAttributeQuestion 
type AccessPackageResourceAttributeQuestion struct {
    AccessPackageResourceAttributeSource
    // The question asked in order to get the value of the attribute
    question AccessPackageQuestionable
}
// NewAccessPackageResourceAttributeQuestion instantiates a new AccessPackageResourceAttributeQuestion and sets the default values.
func NewAccessPackageResourceAttributeQuestion()(*AccessPackageResourceAttributeQuestion) {
    m := &AccessPackageResourceAttributeQuestion{
        AccessPackageResourceAttributeSource: *NewAccessPackageResourceAttributeSource(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageResourceAttributeQuestion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageResourceAttributeQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageResourceAttributeQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageResourceAttributeQuestion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageResourceAttributeQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageResourceAttributeSource.GetFieldDeserializers()
    res["question"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageQuestionFromDiscriminatorValue , m.SetQuestion)
    return res
}
// GetQuestion gets the question property value. The question asked in order to get the value of the attribute
func (m *AccessPackageResourceAttributeQuestion) GetQuestion()(AccessPackageQuestionable) {
    return m.question
}
// Serialize serializes information the current object
func (m *AccessPackageResourceAttributeQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageResourceAttributeSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("question", m.GetQuestion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetQuestion sets the question property value. The question asked in order to get the value of the attribute
func (m *AccessPackageResourceAttributeQuestion) SetQuestion(value AccessPackageQuestionable)() {
    m.question = value
}
