package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailboxEvidence 
type MailboxEvidence struct {
    AlertEvidence
    // The name associated with the mailbox.
    displayName *string
    // The primary email address of the mailbox.
    primaryAddress *string
    // The user account of the mailbox.
    userAccount UserAccountable
}
// NewMailboxEvidence instantiates a new MailboxEvidence and sets the default values.
func NewMailboxEvidence()(*MailboxEvidence) {
    m := &MailboxEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.mailboxEvidence";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMailboxEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailboxEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxEvidence(), nil
}
// GetDisplayName gets the displayName property value. The name associated with the mailbox.
func (m *MailboxEvidence) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailboxEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["primaryAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrimaryAddress)
    res["userAccount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserAccountFromDiscriminatorValue , m.SetUserAccount)
    return res
}
// GetPrimaryAddress gets the primaryAddress property value. The primary email address of the mailbox.
func (m *MailboxEvidence) GetPrimaryAddress()(*string) {
    return m.primaryAddress
}
// GetUserAccount gets the userAccount property value. The user account of the mailbox.
func (m *MailboxEvidence) GetUserAccount()(UserAccountable) {
    return m.userAccount
}
// Serialize serializes information the current object
func (m *MailboxEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryAddress", m.GetPrimaryAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userAccount", m.GetUserAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name associated with the mailbox.
func (m *MailboxEvidence) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPrimaryAddress sets the primaryAddress property value. The primary email address of the mailbox.
func (m *MailboxEvidence) SetPrimaryAddress(value *string)() {
    m.primaryAddress = value
}
// SetUserAccount sets the userAccount property value. The user account of the mailbox.
func (m *MailboxEvidence) SetUserAccount(value UserAccountable)() {
    m.userAccount = value
}
