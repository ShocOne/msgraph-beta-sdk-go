package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleEligibilityScheduleRequest 
type UnifiedRoleEligibilityScheduleRequest struct {
    Request
    // Represents the type of operation on the role eligibility request. The possible values are: AdminAssign: For administrators to assign eligible roles to principals.AdminRemove: For administrators to remove eligible roles from principals. AdminUpdate: For administrators to change existing role eligibilities.AdminExtend: For administrators to extend expiring role eligibilities.AdminRenew: For administrators to renew expired eligibilities.UserAdd: For users to activate their eligible assignments.UserRemove: For users to deactivate their active eligible assignments.UserExtend: For users to request to extend their expiring eligible assignments.UserRenew: For users to request to renew their expired eligible assignments.
    action *string
    // Read-only property with details of the app-specific scope when the role eligibility is scoped to an app. Nullable. Supports $expand.
    appScope AppScopeable
    // Identifier of the app-specific scope when the role eligibility is scoped to an app. The scope of a role eligibility determines the set of resources for which the principal is eligible to access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
    appScopeId *string
    // The directory object that is the scope of the role eligibility. Read-only. Supports $expand.
    directoryScope DirectoryObjectable
    // Identifier of the directory object representing the scope of the role eligibility. The scope of a role eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, ne, and on null values).
    directoryScopeId *string
    // Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
    isValidationOnly *bool
    // A message provided by users and administrators when create they create the unifiedRoleEligibilityScheduleRequest object.
    justification *string
    // The principal that's getting a role eligibility through the request. Supports $expand.
    principal DirectoryObjectable
    // Identifier of the principal that has been granted the role eligibility. Can be a user or a role-assignable group. You can grant only active assignments service principals. Supports $filter (eq, ne).
    principalId *string
    // Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
    roleDefinition UnifiedRoleDefinitionable
    // Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
    roleDefinitionId *string
    // The period of the role eligibility. Recurring schedules are currently unsupported.
    scheduleInfo RequestScheduleable
    // The schedule for a role eligibility that is referenced through the targetScheduleId property. Supports $expand.
    targetSchedule UnifiedRoleEligibilityScheduleable
    // Identifier of the schedule object that's linked to the eligibility request. Supports $filter (eq, ne).
    targetScheduleId *string
    // Ticket details linked to the role eligibility request including details of the ticket number and ticket system. Optional.
    ticketInfo TicketInfoable
}
// NewUnifiedRoleEligibilityScheduleRequest instantiates a new UnifiedRoleEligibilityScheduleRequest and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequest()(*UnifiedRoleEligibilityScheduleRequest) {
    m := &UnifiedRoleEligibilityScheduleRequest{
        Request: *NewRequest(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRoleEligibilityScheduleRequest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleEligibilityScheduleRequest(), nil
}
// GetAction gets the action property value. Represents the type of operation on the role eligibility request. The possible values are: AdminAssign: For administrators to assign eligible roles to principals.AdminRemove: For administrators to remove eligible roles from principals. AdminUpdate: For administrators to change existing role eligibilities.AdminExtend: For administrators to extend expiring role eligibilities.AdminRenew: For administrators to renew expired eligibilities.UserAdd: For users to activate their eligible assignments.UserRemove: For users to deactivate their active eligible assignments.UserExtend: For users to request to extend their expiring eligible assignments.UserRenew: For users to request to renew their expired eligible assignments.
func (m *UnifiedRoleEligibilityScheduleRequest) GetAction()(*string) {
    return m.action
}
// GetAppScope gets the appScope property value. Read-only property with details of the app-specific scope when the role eligibility is scoped to an app. Nullable. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) GetAppScope()(AppScopeable) {
    return m.appScope
}
// GetAppScopeId gets the appScopeId property value. Identifier of the app-specific scope when the role eligibility is scoped to an app. The scope of a role eligibility determines the set of resources for which the principal is eligible to access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleEligibilityScheduleRequest) GetAppScopeId()(*string) {
    return m.appScopeId
}
// GetDirectoryScope gets the directoryScope property value. The directory object that is the scope of the role eligibility. Read-only. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) GetDirectoryScope()(DirectoryObjectable) {
    return m.directoryScope
}
// GetDirectoryScopeId gets the directoryScopeId property value. Identifier of the directory object representing the scope of the role eligibility. The scope of a role eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleEligibilityScheduleRequest) GetDirectoryScopeId()(*string) {
    return m.directoryScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleEligibilityScheduleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["action"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAction)
    res["appScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAppScopeFromDiscriminatorValue , m.SetAppScope)
    res["appScopeId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppScopeId)
    res["directoryScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDirectoryObjectFromDiscriminatorValue , m.SetDirectoryScope)
    res["directoryScopeId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDirectoryScopeId)
    res["isValidationOnly"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsValidationOnly)
    res["justification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJustification)
    res["principal"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDirectoryObjectFromDiscriminatorValue , m.SetPrincipal)
    res["principalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrincipalId)
    res["roleDefinition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUnifiedRoleDefinitionFromDiscriminatorValue , m.SetRoleDefinition)
    res["roleDefinitionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRoleDefinitionId)
    res["scheduleInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRequestScheduleFromDiscriminatorValue , m.SetScheduleInfo)
    res["targetSchedule"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue , m.SetTargetSchedule)
    res["targetScheduleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetScheduleId)
    res["ticketInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTicketInfoFromDiscriminatorValue , m.SetTicketInfo)
    return res
}
// GetIsValidationOnly gets the isValidationOnly property value. Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
func (m *UnifiedRoleEligibilityScheduleRequest) GetIsValidationOnly()(*bool) {
    return m.isValidationOnly
}
// GetJustification gets the justification property value. A message provided by users and administrators when create they create the unifiedRoleEligibilityScheduleRequest object.
func (m *UnifiedRoleEligibilityScheduleRequest) GetJustification()(*string) {
    return m.justification
}
// GetPrincipal gets the principal property value. The principal that's getting a role eligibility through the request. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. Identifier of the principal that has been granted the role eligibility. Can be a user or a role-assignable group. You can grant only active assignments service principals. Supports $filter (eq, ne).
func (m *UnifiedRoleEligibilityScheduleRequest) GetPrincipalId()(*string) {
    return m.principalId
}
// GetRoleDefinition gets the roleDefinition property value. Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) GetRoleDefinition()(UnifiedRoleDefinitionable) {
    return m.roleDefinition
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
func (m *UnifiedRoleEligibilityScheduleRequest) GetRoleDefinitionId()(*string) {
    return m.roleDefinitionId
}
// GetScheduleInfo gets the scheduleInfo property value. The period of the role eligibility. Recurring schedules are currently unsupported.
func (m *UnifiedRoleEligibilityScheduleRequest) GetScheduleInfo()(RequestScheduleable) {
    return m.scheduleInfo
}
// GetTargetSchedule gets the targetSchedule property value. The schedule for a role eligibility that is referenced through the targetScheduleId property. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) GetTargetSchedule()(UnifiedRoleEligibilityScheduleable) {
    return m.targetSchedule
}
// GetTargetScheduleId gets the targetScheduleId property value. Identifier of the schedule object that's linked to the eligibility request. Supports $filter (eq, ne).
func (m *UnifiedRoleEligibilityScheduleRequest) GetTargetScheduleId()(*string) {
    return m.targetScheduleId
}
// GetTicketInfo gets the ticketInfo property value. Ticket details linked to the role eligibility request including details of the ticket number and ticket system. Optional.
func (m *UnifiedRoleEligibilityScheduleRequest) GetTicketInfo()(TicketInfoable) {
    return m.ticketInfo
}
// Serialize serializes information the current object
func (m *UnifiedRoleEligibilityScheduleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Request.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appScope", m.GetAppScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appScopeId", m.GetAppScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directoryScope", m.GetDirectoryScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryScopeId", m.GetDirectoryScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValidationOnly", m.GetIsValidationOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetSchedule", m.GetTargetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetScheduleId", m.GetTargetScheduleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ticketInfo", m.GetTicketInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. Represents the type of operation on the role eligibility request. The possible values are: AdminAssign: For administrators to assign eligible roles to principals.AdminRemove: For administrators to remove eligible roles from principals. AdminUpdate: For administrators to change existing role eligibilities.AdminExtend: For administrators to extend expiring role eligibilities.AdminRenew: For administrators to renew expired eligibilities.UserAdd: For users to activate their eligible assignments.UserRemove: For users to deactivate their active eligible assignments.UserExtend: For users to request to extend their expiring eligible assignments.UserRenew: For users to request to renew their expired eligible assignments.
func (m *UnifiedRoleEligibilityScheduleRequest) SetAction(value *string)() {
    m.action = value
}
// SetAppScope sets the appScope property value. Read-only property with details of the app-specific scope when the role eligibility is scoped to an app. Nullable. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) SetAppScope(value AppScopeable)() {
    m.appScope = value
}
// SetAppScopeId sets the appScopeId property value. Identifier of the app-specific scope when the role eligibility is scoped to an app. The scope of a role eligibility determines the set of resources for which the principal is eligible to access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleEligibilityScheduleRequest) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
// SetDirectoryScope sets the directoryScope property value. The directory object that is the scope of the role eligibility. Read-only. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) SetDirectoryScope(value DirectoryObjectable)() {
    m.directoryScope = value
}
// SetDirectoryScopeId sets the directoryScopeId property value. Identifier of the directory object representing the scope of the role eligibility. The scope of a role eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleEligibilityScheduleRequest) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
// SetIsValidationOnly sets the isValidationOnly property value. Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
func (m *UnifiedRoleEligibilityScheduleRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// SetJustification sets the justification property value. A message provided by users and administrators when create they create the unifiedRoleEligibilityScheduleRequest object.
func (m *UnifiedRoleEligibilityScheduleRequest) SetJustification(value *string)() {
    m.justification = value
}
// SetPrincipal sets the principal property value. The principal that's getting a role eligibility through the request. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. Identifier of the principal that has been granted the role eligibility. Can be a user or a role-assignable group. You can grant only active assignments service principals. Supports $filter (eq, ne).
func (m *UnifiedRoleEligibilityScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetRoleDefinition sets the roleDefinition property value. Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) SetRoleDefinition(value UnifiedRoleDefinitionable)() {
    m.roleDefinition = value
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
func (m *UnifiedRoleEligibilityScheduleRequest) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// SetScheduleInfo sets the scheduleInfo property value. The period of the role eligibility. Recurring schedules are currently unsupported.
func (m *UnifiedRoleEligibilityScheduleRequest) SetScheduleInfo(value RequestScheduleable)() {
    m.scheduleInfo = value
}
// SetTargetSchedule sets the targetSchedule property value. The schedule for a role eligibility that is referenced through the targetScheduleId property. Supports $expand.
func (m *UnifiedRoleEligibilityScheduleRequest) SetTargetSchedule(value UnifiedRoleEligibilityScheduleable)() {
    m.targetSchedule = value
}
// SetTargetScheduleId sets the targetScheduleId property value. Identifier of the schedule object that's linked to the eligibility request. Supports $filter (eq, ne).
func (m *UnifiedRoleEligibilityScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
// SetTicketInfo sets the ticketInfo property value. Ticket details linked to the role eligibility request including details of the ticket number and ticket system. Optional.
func (m *UnifiedRoleEligibilityScheduleRequest) SetTicketInfo(value TicketInfoable)() {
    m.ticketInfo = value
}
