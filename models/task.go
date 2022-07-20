package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Task 
type Task struct {
    BaseTask
}
// NewTask instantiates a new Task and sets the default values.
func NewTask()(*Task) {
    m := &Task{
        BaseTask: *NewBaseTask(),
    }
    odataTypeValue := "#microsoft.graph.task";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTask(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Task) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseTask.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Task) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseTask.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
