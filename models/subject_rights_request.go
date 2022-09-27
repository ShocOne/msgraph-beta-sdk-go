package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequest provides operations to manage the collection of accessReviewDecision entities.
type SubjectRightsRequest struct {
    Entity
    // Identity that the request is assigned to.
    assignedTo Identityable
    // The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // KQL based content query that should be used for search. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    contentQuery *string
    // Identity information for the entity that created the request.
    createdBy IdentitySetable
    // The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Information about the data subject.
    dataSubject DataSubjectable
    // The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
    dataSubjectType *DataSubjectType
    // Description for the request.
    description *string
    // The name of the request.
    displayName *string
    // The external ID for the request that is immutable after creation and is used for tracking the request for the external system. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    externalId *string
    // Collection of history change events.
    history []SubjectRightsRequestHistoryable
    // Include all versions of the documents. By default, the current copies of the documents will be returned. If SharePoint sites have versioning enabled, including all versions will include the historical copies of the documents. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    includeAllVersions *bool
    // Include content authored by the data subject. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    includeAuthoredContent *bool
    // Insight about the request.
    insight SubjectRightsRequestDetailable
    // The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    internalDueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identity information for the entity that last modified the request.
    lastModifiedBy IdentitySetable
    // The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The mailboxlocations property
    mailboxlocations SubjectRightsRequestMailboxLocationable
    // List of notes associated with the request.
    notes []AuthoredNoteable
    // Pause the request after estimate has finished. By default, the data estimate will run and then pause, allowing you to preview results and then select the option to retrieve data in the UI. You can set this property to false if you want it to perform the estimate and then automatically begin with the retrieval of the content. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
    pauseAfterEstimate *bool
    // List of regulations that this request will fulfill.
    regulations []string
    // The sitelocations property
    sitelocations SubjectRightsRequestSiteLocationable
    // Information about the different stages for the request.
    stages []SubjectRightsRequestStageDetailable
    // The status of the request. Possible values are: active, closed, unknownFutureValue.
    status *SubjectRightsRequestStatus
    // Information about the Microsoft Teams team that was created for the request.
    team Teamable
    // The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
    type_escaped *SubjectRightsRequestType
}
// NewSubjectRightsRequest instantiates a new subjectRightsRequest and sets the default values.
func NewSubjectRightsRequest()(*SubjectRightsRequest) {
    m := &SubjectRightsRequest{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.subjectRightsRequest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSubjectRightsRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequest(), nil
}
// GetAssignedTo gets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) GetAssignedTo()(Identityable) {
    return m.assignedTo
}
// GetClosedDateTime gets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closedDateTime
}
// GetContentQuery gets the contentQuery property value. KQL based content query that should be used for search. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) GetContentQuery()(*string) {
    return m.contentQuery
}
// GetCreatedBy gets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDataSubject gets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) GetDataSubject()(DataSubjectable) {
    return m.dataSubject
}
// GetDataSubjectType gets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) GetDataSubjectType()(*DataSubjectType) {
    return m.dataSubjectType
}
// GetDescription gets the description property value. Description for the request.
func (m *SubjectRightsRequest) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. The external ID for the request that is immutable after creation and is used for tracking the request for the external system. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentityFromDiscriminatorValue , m.SetAssignedTo)
    res["closedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetClosedDateTime)
    res["contentQuery"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentQuery)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["dataSubject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDataSubjectFromDiscriminatorValue , m.SetDataSubject)
    res["dataSubjectType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDataSubjectType , m.SetDataSubjectType)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["externalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalId)
    res["history"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSubjectRightsRequestHistoryFromDiscriminatorValue , m.SetHistory)
    res["includeAllVersions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIncludeAllVersions)
    res["includeAuthoredContent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIncludeAuthoredContent)
    res["insight"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSubjectRightsRequestDetailFromDiscriminatorValue , m.SetInsight)
    res["internalDueDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetInternalDueDateTime)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["mailboxlocations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSubjectRightsRequestMailboxLocationFromDiscriminatorValue , m.SetMailboxlocations)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthoredNoteFromDiscriminatorValue , m.SetNotes)
    res["pauseAfterEstimate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPauseAfterEstimate)
    res["regulations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRegulations)
    res["sitelocations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSubjectRightsRequestSiteLocationFromDiscriminatorValue , m.SetSitelocations)
    res["stages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSubjectRightsRequestStageDetailFromDiscriminatorValue , m.SetStages)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubjectRightsRequestStatus , m.SetStatus)
    res["team"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamFromDiscriminatorValue , m.SetTeam)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubjectRightsRequestType , m.SetType)
    return res
}
// GetHistory gets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) GetHistory()([]SubjectRightsRequestHistoryable) {
    return m.history
}
// GetIncludeAllVersions gets the includeAllVersions property value. Include all versions of the documents. By default, the current copies of the documents will be returned. If SharePoint sites have versioning enabled, including all versions will include the historical copies of the documents. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) GetIncludeAllVersions()(*bool) {
    return m.includeAllVersions
}
// GetIncludeAuthoredContent gets the includeAuthoredContent property value. Include content authored by the data subject. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) GetIncludeAuthoredContent()(*bool) {
    return m.includeAuthoredContent
}
// GetInsight gets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) GetInsight()(SubjectRightsRequestDetailable) {
    return m.insight
}
// GetInternalDueDateTime gets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.internalDueDateTime
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMailboxlocations gets the mailboxlocations property value. The mailboxlocations property
func (m *SubjectRightsRequest) GetMailboxlocations()(SubjectRightsRequestMailboxLocationable) {
    return m.mailboxlocations
}
// GetNotes gets the notes property value. List of notes associated with the request.
func (m *SubjectRightsRequest) GetNotes()([]AuthoredNoteable) {
    return m.notes
}
// GetPauseAfterEstimate gets the pauseAfterEstimate property value. Pause the request after estimate has finished. By default, the data estimate will run and then pause, allowing you to preview results and then select the option to retrieve data in the UI. You can set this property to false if you want it to perform the estimate and then automatically begin with the retrieval of the content. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) GetPauseAfterEstimate()(*bool) {
    return m.pauseAfterEstimate
}
// GetRegulations gets the regulations property value. List of regulations that this request will fulfill.
func (m *SubjectRightsRequest) GetRegulations()([]string) {
    return m.regulations
}
// GetSitelocations gets the sitelocations property value. The sitelocations property
func (m *SubjectRightsRequest) GetSitelocations()(SubjectRightsRequestSiteLocationable) {
    return m.sitelocations
}
// GetStages gets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) GetStages()([]SubjectRightsRequestStageDetailable) {
    return m.stages
}
// GetStatus gets the status property value. The status of the request. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) GetStatus()(*SubjectRightsRequestStatus) {
    return m.status
}
// GetTeam gets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) GetTeam()(Teamable) {
    return m.team
}
// GetType gets the type property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) GetType()(*SubjectRightsRequestType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *SubjectRightsRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closedDateTime", m.GetClosedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataSubject", m.GetDataSubject())
        if err != nil {
            return err
        }
    }
    if m.GetDataSubjectType() != nil {
        cast := (*m.GetDataSubjectType()).String()
        err = writer.WriteStringValue("dataSubjectType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHistory())
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeAllVersions", m.GetIncludeAllVersions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeAuthoredContent", m.GetIncludeAuthoredContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("insight", m.GetInsight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("internalDueDateTime", m.GetInternalDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mailboxlocations", m.GetMailboxlocations())
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNotes())
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pauseAfterEstimate", m.GetPauseAfterEstimate())
        if err != nil {
            return err
        }
    }
    if m.GetRegulations() != nil {
        err = writer.WriteCollectionOfStringValues("regulations", m.GetRegulations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sitelocations", m.GetSitelocations())
        if err != nil {
            return err
        }
    }
    if m.GetStages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetStages())
        err = writer.WriteCollectionOfObjectValues("stages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. Identity that the request is assigned to.
func (m *SubjectRightsRequest) SetAssignedTo(value Identityable)() {
    m.assignedTo = value
}
// SetClosedDateTime sets the closedDateTime property value. The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
// SetContentQuery sets the contentQuery property value. KQL based content query that should be used for search. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetContentQuery(value *string)() {
    m.contentQuery = value
}
// SetCreatedBy sets the createdBy property value. Identity information for the entity that created the request.
func (m *SubjectRightsRequest) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDataSubject sets the dataSubject property value. Information about the data subject.
func (m *SubjectRightsRequest) SetDataSubject(value DataSubjectable)() {
    m.dataSubject = value
}
// SetDataSubjectType sets the dataSubjectType property value. The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue.
func (m *SubjectRightsRequest) SetDataSubjectType(value *DataSubjectType)() {
    m.dataSubjectType = value
}
// SetDescription sets the description property value. Description for the request.
func (m *SubjectRightsRequest) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the request.
func (m *SubjectRightsRequest) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. The external ID for the request that is immutable after creation and is used for tracking the request for the external system. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetExternalId(value *string)() {
    m.externalId = value
}
// SetHistory sets the history property value. Collection of history change events.
func (m *SubjectRightsRequest) SetHistory(value []SubjectRightsRequestHistoryable)() {
    m.history = value
}
// SetIncludeAllVersions sets the includeAllVersions property value. Include all versions of the documents. By default, the current copies of the documents will be returned. If SharePoint sites have versioning enabled, including all versions will include the historical copies of the documents. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetIncludeAllVersions(value *bool)() {
    m.includeAllVersions = value
}
// SetIncludeAuthoredContent sets the includeAuthoredContent property value. Include content authored by the data subject. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetIncludeAuthoredContent(value *bool)() {
    m.includeAuthoredContent = value
}
// SetInsight sets the insight property value. Insight about the request.
func (m *SubjectRightsRequest) SetInsight(value SubjectRightsRequestDetailable)() {
    m.insight = value
}
// SetInternalDueDateTime sets the internalDueDateTime property value. The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.internalDueDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity information for the entity that last modified the request.
func (m *SubjectRightsRequest) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SubjectRightsRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMailboxlocations sets the mailboxlocations property value. The mailboxlocations property
func (m *SubjectRightsRequest) SetMailboxlocations(value SubjectRightsRequestMailboxLocationable)() {
    m.mailboxlocations = value
}
// SetNotes sets the notes property value. List of notes associated with the request.
func (m *SubjectRightsRequest) SetNotes(value []AuthoredNoteable)() {
    m.notes = value
}
// SetPauseAfterEstimate sets the pauseAfterEstimate property value. Pause the request after estimate has finished. By default, the data estimate will run and then pause, allowing you to preview results and then select the option to retrieve data in the UI. You can set this property to false if you want it to perform the estimate and then automatically begin with the retrieval of the content. This property is defined only for APIs accessed using the /security query path and not the /privacy query path.
func (m *SubjectRightsRequest) SetPauseAfterEstimate(value *bool)() {
    m.pauseAfterEstimate = value
}
// SetRegulations sets the regulations property value. List of regulations that this request will fulfill.
func (m *SubjectRightsRequest) SetRegulations(value []string)() {
    m.regulations = value
}
// SetSitelocations sets the sitelocations property value. The sitelocations property
func (m *SubjectRightsRequest) SetSitelocations(value SubjectRightsRequestSiteLocationable)() {
    m.sitelocations = value
}
// SetStages sets the stages property value. Information about the different stages for the request.
func (m *SubjectRightsRequest) SetStages(value []SubjectRightsRequestStageDetailable)() {
    m.stages = value
}
// SetStatus sets the status property value. The status of the request. Possible values are: active, closed, unknownFutureValue.
func (m *SubjectRightsRequest) SetStatus(value *SubjectRightsRequestStatus)() {
    m.status = value
}
// SetTeam sets the team property value. Information about the Microsoft Teams team that was created for the request.
func (m *SubjectRightsRequest) SetTeam(value Teamable)() {
    m.team = value
}
// SetType sets the type property value. The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
func (m *SubjectRightsRequest) SetType(value *SubjectRightsRequestType)() {
    m.type_escaped = value
}
