package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThumbnailSet provides operations to manage the collection of accessReview entities.
type ThumbnailSet struct {
    Entity
    // A 1920x1920 scaled thumbnail.
    large Thumbnailable
    // A 176x176 scaled thumbnail.
    medium Thumbnailable
    // A 48x48 cropped thumbnail.
    small Thumbnailable
    // A custom thumbnail image or the original image used to generate other thumbnails.
    source Thumbnailable
}
// NewThumbnailSet instantiates a new thumbnailSet and sets the default values.
func NewThumbnailSet()(*ThumbnailSet) {
    m := &ThumbnailSet{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.thumbnailSet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateThumbnailSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThumbnailSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThumbnailSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThumbnailSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["large"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateThumbnailFromDiscriminatorValue , m.SetLarge)
    res["medium"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateThumbnailFromDiscriminatorValue , m.SetMedium)
    res["small"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateThumbnailFromDiscriminatorValue , m.SetSmall)
    res["source"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateThumbnailFromDiscriminatorValue , m.SetSource)
    return res
}
// GetLarge gets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) GetLarge()(Thumbnailable) {
    return m.large
}
// GetMedium gets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) GetMedium()(Thumbnailable) {
    return m.medium
}
// GetSmall gets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) GetSmall()(Thumbnailable) {
    return m.small
}
// GetSource gets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) GetSource()(Thumbnailable) {
    return m.source
}
// Serialize serializes information the current object
func (m *ThumbnailSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("large", m.GetLarge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("medium", m.GetMedium())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("small", m.GetSmall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLarge sets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) SetLarge(value Thumbnailable)() {
    m.large = value
}
// SetMedium sets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) SetMedium(value Thumbnailable)() {
    m.medium = value
}
// SetSmall sets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) SetSmall(value Thumbnailable)() {
    m.small = value
}
// SetSource sets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) SetSource(value Thumbnailable)() {
    m.source = value
}
