package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeFont 
type WorkbookRangeFont struct {
    Entity
    // Represents the bold status of font.
    bold *bool
    // HTML color code representation of the text color. E.g. #FF0000 represents Red.
    color *string
    // Represents the italic status of the font.
    italic *bool
    // Font name (e.g. 'Calibri')
    name *string
    // Font size.
    size *float64
    // Type of underline applied to the font. Possible values are: None, Single, Double, SingleAccountant, DoubleAccountant.
    underline *string
}
// NewWorkbookRangeFont instantiates a new workbookRangeFont and sets the default values.
func NewWorkbookRangeFont()(*WorkbookRangeFont) {
    m := &WorkbookRangeFont{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.workbookRangeFont";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkbookRangeFontFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookRangeFontFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeFont(), nil
}
// GetBold gets the bold property value. Represents the bold status of font.
func (m *WorkbookRangeFont) GetBold()(*bool) {
    return m.bold
}
// GetColor gets the color property value. HTML color code representation of the text color. E.g. #FF0000 represents Red.
func (m *WorkbookRangeFont) GetColor()(*string) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeFont) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bold"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBold)
    res["color"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetColor)
    res["italic"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetItalic)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["size"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetSize)
    res["underline"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUnderline)
    return res
}
// GetItalic gets the italic property value. Represents the italic status of the font.
func (m *WorkbookRangeFont) GetItalic()(*bool) {
    return m.italic
}
// GetName gets the name property value. Font name (e.g. 'Calibri')
func (m *WorkbookRangeFont) GetName()(*string) {
    return m.name
}
// GetSize gets the size property value. Font size.
func (m *WorkbookRangeFont) GetSize()(*float64) {
    return m.size
}
// GetUnderline gets the underline property value. Type of underline applied to the font. Possible values are: None, Single, Double, SingleAccountant, DoubleAccountant.
func (m *WorkbookRangeFont) GetUnderline()(*string) {
    return m.underline
}
// Serialize serializes information the current object
func (m *WorkbookRangeFont) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("bold", m.GetBold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("italic", m.GetItalic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("underline", m.GetUnderline())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBold sets the bold property value. Represents the bold status of font.
func (m *WorkbookRangeFont) SetBold(value *bool)() {
    m.bold = value
}
// SetColor sets the color property value. HTML color code representation of the text color. E.g. #FF0000 represents Red.
func (m *WorkbookRangeFont) SetColor(value *string)() {
    m.color = value
}
// SetItalic sets the italic property value. Represents the italic status of the font.
func (m *WorkbookRangeFont) SetItalic(value *bool)() {
    m.italic = value
}
// SetName sets the name property value. Font name (e.g. 'Calibri')
func (m *WorkbookRangeFont) SetName(value *string)() {
    m.name = value
}
// SetSize sets the size property value. Font size.
func (m *WorkbookRangeFont) SetSize(value *float64)() {
    m.size = value
}
// SetUnderline sets the underline property value. Type of underline applied to the font. Possible values are: None, Single, Double, SingleAccountant, DoubleAccountant.
func (m *WorkbookRangeFont) SetUnderline(value *string)() {
    m.underline = value
}
