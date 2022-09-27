package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AnalyzedMessageEvidence 
type AnalyzedMessageEvidence struct {
    AlertEvidence
    // Direction of the email relative to your network. The possible values are: Inbound, Outbound or Intraorg.
    antiSpamDirection *string
    // Number of attachments in the email.
    attachmentsCount *int64
    // Delivery action of the email. The possible values are: Delivered, DeliveredAsSpam, Junked, Blocked, or Replaced.
    deliveryAction *string
    // Location where the email was delivered. The possible values are: Inbox, External, JunkFolder, Quarantine, Failed, Dropped, DeletedFolder or Forwarded.
    deliveryLocation *string
    // Public-facing identifier for the email that is set by the sending email system.
    internetMessageId *string
    // Detected language of the email content.
    language *string
    // Unique identifier for the email, generated by Microsoft 365.
    networkMessageId *string
    // The P1 sender.
    p1Sender EmailSenderable
    // The P2 sender.
    p2Sender EmailSenderable
    // Date and time when the email was received.
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Email address of the recipient, or email address of the recipient after distribution list expansion.
    recipientEmailAddress *string
    // IP address of the last detected mail server that relayed the message.
    senderIp *string
    // Subject of the email.
    subject *string
    // Collection of methods used to detect malware, phishing, or other threats found in the email.
    threatDetectionMethods []string
    // Collection of detection names for malware or other threats found.
    threats []string
    // Number of embedded URLs in the email.
    urlCount *int64
    // Collection of the URLs contained in this email.
    urls []string
    // Uniform resource name (URN) of the automated investigation where the cluster was identified.
    urn *string
}
// NewAnalyzedMessageEvidence instantiates a new AnalyzedMessageEvidence and sets the default values.
func NewAnalyzedMessageEvidence()(*AnalyzedMessageEvidence) {
    m := &AnalyzedMessageEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.analyzedMessageEvidence";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAnalyzedMessageEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAnalyzedMessageEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnalyzedMessageEvidence(), nil
}
// GetAntiSpamDirection gets the antiSpamDirection property value. Direction of the email relative to your network. The possible values are: Inbound, Outbound or Intraorg.
func (m *AnalyzedMessageEvidence) GetAntiSpamDirection()(*string) {
    return m.antiSpamDirection
}
// GetAttachmentsCount gets the attachmentsCount property value. Number of attachments in the email.
func (m *AnalyzedMessageEvidence) GetAttachmentsCount()(*int64) {
    return m.attachmentsCount
}
// GetDeliveryAction gets the deliveryAction property value. Delivery action of the email. The possible values are: Delivered, DeliveredAsSpam, Junked, Blocked, or Replaced.
func (m *AnalyzedMessageEvidence) GetDeliveryAction()(*string) {
    return m.deliveryAction
}
// GetDeliveryLocation gets the deliveryLocation property value. Location where the email was delivered. The possible values are: Inbox, External, JunkFolder, Quarantine, Failed, Dropped, DeletedFolder or Forwarded.
func (m *AnalyzedMessageEvidence) GetDeliveryLocation()(*string) {
    return m.deliveryLocation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AnalyzedMessageEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["antiSpamDirection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAntiSpamDirection)
    res["attachmentsCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetAttachmentsCount)
    res["deliveryAction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeliveryAction)
    res["deliveryLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeliveryLocation)
    res["internetMessageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInternetMessageId)
    res["language"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLanguage)
    res["networkMessageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetworkMessageId)
    res["p1Sender"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEmailSenderFromDiscriminatorValue , m.SetP1Sender)
    res["p2Sender"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEmailSenderFromDiscriminatorValue , m.SetP2Sender)
    res["receivedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetReceivedDateTime)
    res["recipientEmailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRecipientEmailAddress)
    res["senderIp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSenderIp)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubject)
    res["threatDetectionMethods"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetThreatDetectionMethods)
    res["threats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetThreats)
    res["urlCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetUrlCount)
    res["urls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetUrls)
    res["urn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUrn)
    return res
}
// GetInternetMessageId gets the internetMessageId property value. Public-facing identifier for the email that is set by the sending email system.
func (m *AnalyzedMessageEvidence) GetInternetMessageId()(*string) {
    return m.internetMessageId
}
// GetLanguage gets the language property value. Detected language of the email content.
func (m *AnalyzedMessageEvidence) GetLanguage()(*string) {
    return m.language
}
// GetNetworkMessageId gets the networkMessageId property value. Unique identifier for the email, generated by Microsoft 365.
func (m *AnalyzedMessageEvidence) GetNetworkMessageId()(*string) {
    return m.networkMessageId
}
// GetP1Sender gets the p1Sender property value. The P1 sender.
func (m *AnalyzedMessageEvidence) GetP1Sender()(EmailSenderable) {
    return m.p1Sender
}
// GetP2Sender gets the p2Sender property value. The P2 sender.
func (m *AnalyzedMessageEvidence) GetP2Sender()(EmailSenderable) {
    return m.p2Sender
}
// GetReceivedDateTime gets the receivedDateTime property value. Date and time when the email was received.
func (m *AnalyzedMessageEvidence) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
// GetRecipientEmailAddress gets the recipientEmailAddress property value. Email address of the recipient, or email address of the recipient after distribution list expansion.
func (m *AnalyzedMessageEvidence) GetRecipientEmailAddress()(*string) {
    return m.recipientEmailAddress
}
// GetSenderIp gets the senderIp property value. IP address of the last detected mail server that relayed the message.
func (m *AnalyzedMessageEvidence) GetSenderIp()(*string) {
    return m.senderIp
}
// GetSubject gets the subject property value. Subject of the email.
func (m *AnalyzedMessageEvidence) GetSubject()(*string) {
    return m.subject
}
// GetThreatDetectionMethods gets the threatDetectionMethods property value. Collection of methods used to detect malware, phishing, or other threats found in the email.
func (m *AnalyzedMessageEvidence) GetThreatDetectionMethods()([]string) {
    return m.threatDetectionMethods
}
// GetThreats gets the threats property value. Collection of detection names for malware or other threats found.
func (m *AnalyzedMessageEvidence) GetThreats()([]string) {
    return m.threats
}
// GetUrlCount gets the urlCount property value. Number of embedded URLs in the email.
func (m *AnalyzedMessageEvidence) GetUrlCount()(*int64) {
    return m.urlCount
}
// GetUrls gets the urls property value. Collection of the URLs contained in this email.
func (m *AnalyzedMessageEvidence) GetUrls()([]string) {
    return m.urls
}
// GetUrn gets the urn property value. Uniform resource name (URN) of the automated investigation where the cluster was identified.
func (m *AnalyzedMessageEvidence) GetUrn()(*string) {
    return m.urn
}
// Serialize serializes information the current object
func (m *AnalyzedMessageEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("antiSpamDirection", m.GetAntiSpamDirection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("attachmentsCount", m.GetAttachmentsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deliveryAction", m.GetDeliveryAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deliveryLocation", m.GetDeliveryLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkMessageId", m.GetNetworkMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("p1Sender", m.GetP1Sender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("p2Sender", m.GetP2Sender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientEmailAddress", m.GetRecipientEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderIp", m.GetSenderIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    if m.GetThreatDetectionMethods() != nil {
        err = writer.WriteCollectionOfStringValues("threatDetectionMethods", m.GetThreatDetectionMethods())
        if err != nil {
            return err
        }
    }
    if m.GetThreats() != nil {
        err = writer.WriteCollectionOfStringValues("threats", m.GetThreats())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("urlCount", m.GetUrlCount())
        if err != nil {
            return err
        }
    }
    if m.GetUrls() != nil {
        err = writer.WriteCollectionOfStringValues("urls", m.GetUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("urn", m.GetUrn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAntiSpamDirection sets the antiSpamDirection property value. Direction of the email relative to your network. The possible values are: Inbound, Outbound or Intraorg.
func (m *AnalyzedMessageEvidence) SetAntiSpamDirection(value *string)() {
    m.antiSpamDirection = value
}
// SetAttachmentsCount sets the attachmentsCount property value. Number of attachments in the email.
func (m *AnalyzedMessageEvidence) SetAttachmentsCount(value *int64)() {
    m.attachmentsCount = value
}
// SetDeliveryAction sets the deliveryAction property value. Delivery action of the email. The possible values are: Delivered, DeliveredAsSpam, Junked, Blocked, or Replaced.
func (m *AnalyzedMessageEvidence) SetDeliveryAction(value *string)() {
    m.deliveryAction = value
}
// SetDeliveryLocation sets the deliveryLocation property value. Location where the email was delivered. The possible values are: Inbox, External, JunkFolder, Quarantine, Failed, Dropped, DeletedFolder or Forwarded.
func (m *AnalyzedMessageEvidence) SetDeliveryLocation(value *string)() {
    m.deliveryLocation = value
}
// SetInternetMessageId sets the internetMessageId property value. Public-facing identifier for the email that is set by the sending email system.
func (m *AnalyzedMessageEvidence) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetLanguage sets the language property value. Detected language of the email content.
func (m *AnalyzedMessageEvidence) SetLanguage(value *string)() {
    m.language = value
}
// SetNetworkMessageId sets the networkMessageId property value. Unique identifier for the email, generated by Microsoft 365.
func (m *AnalyzedMessageEvidence) SetNetworkMessageId(value *string)() {
    m.networkMessageId = value
}
// SetP1Sender sets the p1Sender property value. The P1 sender.
func (m *AnalyzedMessageEvidence) SetP1Sender(value EmailSenderable)() {
    m.p1Sender = value
}
// SetP2Sender sets the p2Sender property value. The P2 sender.
func (m *AnalyzedMessageEvidence) SetP2Sender(value EmailSenderable)() {
    m.p2Sender = value
}
// SetReceivedDateTime sets the receivedDateTime property value. Date and time when the email was received.
func (m *AnalyzedMessageEvidence) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// SetRecipientEmailAddress sets the recipientEmailAddress property value. Email address of the recipient, or email address of the recipient after distribution list expansion.
func (m *AnalyzedMessageEvidence) SetRecipientEmailAddress(value *string)() {
    m.recipientEmailAddress = value
}
// SetSenderIp sets the senderIp property value. IP address of the last detected mail server that relayed the message.
func (m *AnalyzedMessageEvidence) SetSenderIp(value *string)() {
    m.senderIp = value
}
// SetSubject sets the subject property value. Subject of the email.
func (m *AnalyzedMessageEvidence) SetSubject(value *string)() {
    m.subject = value
}
// SetThreatDetectionMethods sets the threatDetectionMethods property value. Collection of methods used to detect malware, phishing, or other threats found in the email.
func (m *AnalyzedMessageEvidence) SetThreatDetectionMethods(value []string)() {
    m.threatDetectionMethods = value
}
// SetThreats sets the threats property value. Collection of detection names for malware or other threats found.
func (m *AnalyzedMessageEvidence) SetThreats(value []string)() {
    m.threats = value
}
// SetUrlCount sets the urlCount property value. Number of embedded URLs in the email.
func (m *AnalyzedMessageEvidence) SetUrlCount(value *int64)() {
    m.urlCount = value
}
// SetUrls sets the urls property value. Collection of the URLs contained in this email.
func (m *AnalyzedMessageEvidence) SetUrls(value []string)() {
    m.urls = value
}
// SetUrn sets the urn property value. Uniform resource name (URN) of the automated investigation where the cluster was identified.
func (m *AnalyzedMessageEvidence) SetUrn(value *string)() {
    m.urn = value
}
