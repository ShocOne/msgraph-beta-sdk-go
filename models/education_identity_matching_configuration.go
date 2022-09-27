package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationIdentityMatchingConfiguration 
type EducationIdentityMatchingConfiguration struct {
    EducationIdentitySynchronizationConfiguration
    // Mapping between the user account and the options to use to uniquely identify the user to update.
    matchingOptions []EducationIdentityMatchingOptionsable
}
// NewEducationIdentityMatchingConfiguration instantiates a new EducationIdentityMatchingConfiguration and sets the default values.
func NewEducationIdentityMatchingConfiguration()(*EducationIdentityMatchingConfiguration) {
    m := &EducationIdentityMatchingConfiguration{
        EducationIdentitySynchronizationConfiguration: *NewEducationIdentitySynchronizationConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.educationIdentityMatchingConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationIdentityMatchingConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationIdentityMatchingConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationIdentityMatchingConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationIdentityMatchingConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationIdentitySynchronizationConfiguration.GetFieldDeserializers()
    res["matchingOptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEducationIdentityMatchingOptionsFromDiscriminatorValue , m.SetMatchingOptions)
    return res
}
// GetMatchingOptions gets the matchingOptions property value. Mapping between the user account and the options to use to uniquely identify the user to update.
func (m *EducationIdentityMatchingConfiguration) GetMatchingOptions()([]EducationIdentityMatchingOptionsable) {
    return m.matchingOptions
}
// Serialize serializes information the current object
func (m *EducationIdentityMatchingConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationIdentitySynchronizationConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMatchingOptions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMatchingOptions())
        err = writer.WriteCollectionOfObjectValues("matchingOptions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatchingOptions sets the matchingOptions property value. Mapping between the user account and the options to use to uniquely identify the user to update.
func (m *EducationIdentityMatchingConfiguration) SetMatchingOptions(value []EducationIdentityMatchingOptionsable)() {
    m.matchingOptions = value
}
