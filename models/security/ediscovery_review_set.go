package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryReviewSet 
type EdiscoveryReviewSet struct {
    DataSet
    // Represents files within the review set.
    files []EdiscoveryFileable
    // Represents queries within the review set.
    queries []EdiscoveryReviewSetQueryable
}
// NewEdiscoveryReviewSet instantiates a new EdiscoveryReviewSet and sets the default values.
func NewEdiscoveryReviewSet()(*EdiscoveryReviewSet) {
    m := &EdiscoveryReviewSet{
        DataSet: *NewDataSet(),
    }
    return m
}
// CreateEdiscoveryReviewSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryReviewSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryReviewSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryReviewSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSet.GetFieldDeserializers()
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryFileable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryFileable)
            }
            m.SetFiles(res)
        }
        return nil
    }
    res["queries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewSetQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryReviewSetQueryable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryReviewSetQueryable)
            }
            m.SetQueries(res)
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. Represents files within the review set.
func (m *EdiscoveryReviewSet) GetFiles()([]EdiscoveryFileable) {
    if m == nil {
        return nil
    } else {
        return m.files
    }
}
// GetQueries gets the queries property value. Represents queries within the review set.
func (m *EdiscoveryReviewSet) GetQueries()([]EdiscoveryReviewSetQueryable) {
    if m == nil {
        return nil
    } else {
        return m.queries
    }
}
// Serialize serializes information the current object
func (m *EdiscoveryReviewSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("files", cast)
        if err != nil {
            return err
        }
    }
    if m.GetQueries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetQueries()))
        for i, v := range m.GetQueries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("queries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFiles sets the files property value. Represents files within the review set.
func (m *EdiscoveryReviewSet) SetFiles(value []EdiscoveryFileable)() {
    if m != nil {
        m.files = value
    }
}
// SetQueries sets the queries property value. Represents queries within the review set.
func (m *EdiscoveryReviewSet) SetQueries(value []EdiscoveryReviewSetQueryable)() {
    if m != nil {
        m.queries = value
    }
}
