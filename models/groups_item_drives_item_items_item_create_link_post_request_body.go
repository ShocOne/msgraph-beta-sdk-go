package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody provides operations to call the createLink method.
type GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expirationDateTime property
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The message property
    message *string
    // The password property
    password *string
    // The recipients property
    recipients []DriveRecipientable
    // The retainInheritedPermissions property
    retainInheritedPermissions *bool
    // The scope property
    scope *string
    // The type property
    type_escaped *string
}
// NewGroupsItemDrivesItemItemsItemCreateLinkPostRequestBody instantiates a new GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody and sets the default values.
func NewGroupsItemDrivesItemItemsItemCreateLinkPostRequestBody()(*GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) {
    m := &GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemDrivesItemItemsItemCreateLinkPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemDrivesItemItemsItemCreateLinkPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemDrivesItemItemsItemCreateLinkPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(DriveRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["retainInheritedPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetainInheritedPermissions(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetPassword gets the password property value. The password property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetRecipients gets the recipients property value. The recipients property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetRecipients()([]DriveRecipientable) {
    return m.recipients
}
// GetRetainInheritedPermissions gets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetRetainInheritedPermissions()(*bool) {
    return m.retainInheritedPermissions
}
// GetScope gets the scope property value. The scope property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetScope()(*string) {
    return m.scope
}
// GetType gets the type property value. The type property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("retainInheritedPermissions", m.GetRetainInheritedPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetMessage sets the message property value. The message property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetPassword sets the password property value. The password property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetRecipients(value []DriveRecipientable)() {
    m.recipients = value
}
// SetRetainInheritedPermissions sets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetRetainInheritedPermissions(value *bool)() {
    m.retainInheritedPermissions = value
}
// SetScope sets the scope property value. The scope property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetScope(value *string)() {
    m.scope = value
}
// SetType sets the type property value. The type property
func (m *GroupsItemDrivesItemItemsItemCreateLinkPostRequestBody) SetType(value *string)() {
    m.type_escaped = value
}
