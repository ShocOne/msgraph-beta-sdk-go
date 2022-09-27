package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// User 
type User struct {
    DirectoryObject
    // A freeform text entry field for the user to describe themselves. Returned only on $select.
    aboutMe *string
    // true if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter (eq, ne, not, and in).
    accountEnabled *bool
    // The activities property
    activities []UserActivityable
    // Sets the age group of the user. Allowed values: null, Minor, NotAdult and Adult. Refer to the legal age group property definitions for further information. Supports $filter (eq, ne, not, and in).
    ageGroup *string
    // The user's terms of use acceptance statuses. Read-only. Nullable.
    agreementAcceptances []AgreementAcceptanceable
    // The analytics property
    analytics UserAnalyticsable
    // The appConsentRequestsForApproval property
    appConsentRequestsForApproval []AppConsentRequestable
    // The appRoleAssignedResources property
    appRoleAssignedResources []ServicePrincipalable
    // Represents the app roles a user has been granted for an application. Supports $expand.
    appRoleAssignments []AppRoleAssignmentable
    // The approvals property
    approvals []Approvalable
    // The licenses that are assigned to the user, including inherited (group-based) licenses. Not nullable. Supports $filter (eq, not, and counting empty collections).
    assignedLicenses []AssignedLicenseable
    // The plans that are assigned to the user. Read-only. Not nullable.Supports $filter (eq and not).
    assignedPlans []AssignedPlanable
    // The authentication property
    authentication Authenticationable
    // Identifiers that can be used to identify and authenticate a user in non-Azure AD environments. This property can be used to store identifiers for smartcard-based certificates that a user uses for access to on-premises Active Directory deployments or for federated access. It can also be used to store the Subject Alternate Name (SAN) that's associated with a Common Access Card (CAC). Nullable.Supports $filter (eq and startsWith).
    authorizationInfo AuthorizationInfoable
    // The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select.
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced from on-premises directory. Supports $filter (eq, not, ge, le, startsWith).
    businessPhones []string
    // The user's primary calendar. Read-only.
    calendar Calendarable
    // The user's calendar groups. Read-only. Nullable.
    calendarGroups []CalendarGroupable
    // The user's calendars. Read-only. Nullable.
    calendars []Calendarable
    // The calendar view for the calendar. Read-only. Nullable.
    calendarView []Eventable
    // The chats property
    chats []Chatable
    // The city in which the user is located. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    city *string
    // The cloudPCs property
    cloudPCs []CloudPCable
    // The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length is 64 characters.Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    companyName *string
    // Sets whether consent has been obtained for minors. Allowed values: null, Granted, Denied and NotRequired. Refer to the legal age group property definitions for further information. Supports $filter (eq, ne, not, and in).
    consentProvidedForMinor *string
    // The user's contacts folders. Read-only. Nullable.
    contactFolders []ContactFolderable
    // The user's contacts. Read-only. Nullable.
    contacts []Contactable
    // The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    country *string
    // The date and time the user was created. The value cannot be modified and is automatically populated when the entity is created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. Property is nullable. A null value indicates that an accurate creation time couldn't be determined for the user. Read-only. Supports $filter (eq, ne, not , ge, le, in).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Directory objects that were created by the user. Read-only. Nullable.
    createdObjects []DirectoryObjectable
    // Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp).  Read-only.Supports $filter (eq, ne, not, and in).
    creationType *string
    // An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
    customSecurityAttributes CustomSecurityAttributeValueable
    // The name for the department in which the user works. Maximum length is 64 characters.Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
    department *string
    // Get enrollment configurations targeted to the user
    deviceEnrollmentConfigurations []DeviceEnrollmentConfigurationable
    // The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
    deviceEnrollmentLimit *int32
    // The deviceKeys property
    deviceKeys []DeviceKeyable
    // The list of troubleshooting events for this user.
    deviceManagementTroubleshootingEvents []DeviceManagementTroubleshootingEventable
    // The devices property
    devices []Deviceable
    // The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
    directReports []DirectoryObjectable
    // The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderBy, and $search.
    displayName *string
    // The user's OneDrive. Read-only.
    drive Driveable
    // A collection of drives available for this user. Read-only.
    drives []Driveable
    // The date and time when the user was hired or will start work in case of a future hire. Supports $filter (eq, ne, not , ge, le, in).
    employeeHireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The employee identifier assigned to the user by the organization. The maximum length is 16 characters.Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    employeeId *string
    // The date and time when the user left or will leave the organization. Read: Requires User-LifeCycleInfo.Read.All. For delegated scenarios, the admin needs one of the following Azure AD roles: Lifecycle Workflows Administrator, Global Reader, or Global Admin.  Write: Requires User-LifeCycleInfo.ReadWrite.All. For delegated scenarios, the admin needs the Global Administrator Azure AD role. Supports $filter (eq, ne, not , ge, le, in).
    employeeLeaveDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Represents organization data (e.g. division and costCenter) associated with a user. Supports $filter (eq, ne, not , ge, le, in).
    employeeOrgData EmployeeOrgDataable
    // Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Supports $filter (eq, ne, not , ge, le, in, startsWith).
    employeeType *string
    // The user's events. Default is to show events under the Default Calendar. Read-only. Nullable.
    events []Eventable
    // The collection of open extensions defined for the user. Supports $expand. Nullable.
    extensions []Extensionable
    // For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Supports $filter (eq, ne, not , in).
    externalUserState *string
    // Shows the timestamp for the latest change to the externalUserState property. Supports $filter (eq, ne, not , in).
    externalUserStateChangeDateTime *string
    // The fax number of the user. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    faxNumber *string
    // The followedSites property
    followedSites []Siteable
    // The given name (first name) of the user. Maximum length is 64 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    givenName *string
    // The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
    hireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Supports $filter (eq) including on null values, only where the signInType is not userPrincipalName.
    identities []ObjectIdentityable
    // The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Supports $filter (eq, not, ge, le, startsWith).
    imAddresses []string
    // Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
    inferenceClassification InferenceClassificationable
    // Identifies the info segments assigned to the user.  Supports $filter (eq, not, ge, le, startsWith).
    infoCatalogs []string
    // The informationProtection property
    informationProtection InformationProtectionable
    // The insights property
    insights ItemInsightsable
    // A list for the user to describe their interests. Returned only on $select.
    interests []string
    // The isManagementRestricted property
    isManagementRestricted *bool
    // Do not use – reserved for future use.
    isResourceAccount *bool
    // The user's job title. Maximum length is 128 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    jobTitle *string
    // The joinedGroups property
    joinedGroups []Groupable
    // The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
    joinedTeams []Teamable
    // The time when this Azure AD user last changed their password or when their password was created, , whichever date the latest action was performed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select.
    lastPasswordChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, MinorWithOutParentalConsent, MinorWithParentalConsent, MinorNoParentalConsentRequired, NotAdult and Adult. Refer to the legal age group property definitions for further information. Returned only on $select.
    legalAgeGroupClassification *string
    // State of license assignments for this user. Read-only. Returned only on $select.
    licenseAssignmentStates []LicenseAssignmentStateable
    // The licenseDetails property
    licenseDetails []LicenseDetailsable
    // The SMTP address for the user, for example, admin@contoso.com. Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. This property cannot contain accent characters.  NOTE: We do not recommend updating this property for Azure AD B2C user profiles. Use the otherMails property instead.  Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
    mail *string
    // Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. For more information, see User preferences for languages and regional formats. Returned only on $select.
    mailboxSettings MailboxSettingsable
    // The user's mail folders. Read-only. Nullable.
    mailFolders []MailFolderable
    // The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    mailNickname *string
    // Zero or more managed app registrations that belong to the user.
    managedAppRegistrations []ManagedAppRegistrationable
    // The managed devices associated with the user.
    managedDevices []ManagedDeviceable
    // The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
    manager DirectoryObjectable
    // The groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // The messages in a mailbox or folder. Read-only. Nullable.
    messages []Messageable
    // The list of troubleshooting events for this user.
    mobileAppIntentAndStates []MobileAppIntentAndStateable
    // The list of mobile app troubleshooting events for this user.
    mobileAppTroubleshootingEvents []MobileAppTroubleshootingEventable
    // The primary cellular telephone number for the user. Read-only for users synced from on-premises directory.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    mobilePhone *string
    // The URL for the user's personal site. Returned only on $select.
    mySite *string
    // The notifications property
    notifications []Notificationable
    // The oauth2PermissionGrants property
    oauth2PermissionGrants []OAuth2PermissionGrantable
    // The office location in the user's place of business. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    officeLocation *string
    // The onenote property
    onenote Onenoteable
    // The onlineMeetings property
    onlineMeetings []OnlineMeetingable
    // Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
    onPremisesDistinguishedName *string
    // Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
    onPremisesDomainName *string
    // Contains extensionAttributes1-15 for the user. These extension attributes are also known as Exchange custom attributes 1-15. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. Supports $filter (eq, ne, not, in).
    onPremisesExtensionAttributes OnPremisesExtensionAttributesable
    // This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. Note: The $ and _ characters cannot be used when specifying this property. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesImmutableId *string
    // Indicates the last time at which the object was synced with the on-premises directory; for example: '2013-02-16T03:04:54Z'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Errors when using Microsoft synchronization product during provisioning.  Supports $filter (eq, not, ge, le).
    onPremisesProvisioningErrors []OnPremisesProvisioningErrorable
    // Contains the on-premises sAMAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith).
    onPremisesSamAccountName *string
    // Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Supports $filter (eq including on null values).
    onPremisesSecurityIdentifier *string
    // true if this user object is currently being synced from an on-premises Active Directory (AD); otherwise the user isn't being synced and can be managed in Azure Active Directory (Azure AD). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool
    // Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith).
    onPremisesUserPrincipalName *string
    // A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com'].NOTE: This property cannot contain accent characters.Supports $filter (eq, not, ge, le, in, startsWith, endsWith, and counting empty collections).
    otherMails []string
    // Selective Outlook services available to the user. Read-only. Nullable.
    outlook OutlookUserable
    // Devices that are owned by the user. Read-only. Nullable. Supports $expand.
    ownedDevices []DirectoryObjectable
    // Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
    ownedObjects []DirectoryObjectable
    // Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, not, and eq on null values).
    passwordPolicies *string
    // Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon. Supports $filter (eq, ne, not, in, and eq on null values).
    passwordProfile PasswordProfileable
    // A list for the user to enumerate their past projects. Returned only on $select.
    pastProjects []string
    // Navigation property to get list of access reviews pending approval by reviewer.
    pendingAccessReviewInstances []AccessReviewInstanceable
    // Read-only. The most relevant people to the user. The collection is ordered by their relevance to the user, which is determined by the user's communication, collaboration and business relationships. A person is an aggregation of information from across mail, contacts and social networks.
    people []Personable
    // The user's profile photo. Read-only.
    photo ProfilePhotoable
    // The photos property
    photos []ProfilePhotoable
    // Selective Planner services available to the user. Read-only. Nullable.
    planner PlannerUserable
    // The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    postalCode *string
    // The preferred data location for the user. For more information, see OneDrive Online Multi-Geo.
    preferredDataLocation *string
    // The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    preferredLanguage *string
    // The preferred name for the user. Not Supported. This attribute returns an empty string.Returned only on $select.
    preferredName *string
    // The presence property
    presence Presenceable
    // The print property
    print UserPrintable
    // Represents properties that are descriptive of a user in a tenant.
    profile Profileable
    // The plans that are provisioned for the user. Read-only. Not nullable. Supports $filter (eq, not, ge, le).
    provisionedPlans []ProvisionedPlanable
    // For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. Changes to the mail property will also update this collection to include the value as an SMTP address. For more information, see mail and proxyAddresses properties. The proxy address prefixed with SMTP (capitalized) is the primary proxy address while those prefixed with smtp are the secondary proxy addresses. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only in Microsoft Graph; you can update this property only through the Microsoft 365 admin center. Not nullable. Supports $filter (eq, not, ge, le, startsWith, endsWith, and counting empty collections).
    proxyAddresses []string
    // Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use invalidateAllRefreshTokens to reset.
    refreshTokensValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Devices that are registered for the user. Read-only. Nullable. Supports $expand.
    registeredDevices []DirectoryObjectable
    // A list for the user to enumerate their responsibilities. Returned only on $select.
    responsibilities []string
    // A list for the user to enumerate the schools they have attended. Returned only on $select.
    schools []string
    // The scoped-role administrative unit memberships for this user. Read-only. Nullable.
    scopedRoleMemberOf []ScopedRoleMembershipable
    // Security identifier (SID) of the user, used in Windows scenarios. Read-only. Returned by default. Supports $select and $filter (eq, not, ge, le, startsWith).
    securityIdentifier *string
    // The settings property
    settings UserSettingsable
    // Do not use in Microsoft Graph. Manage this property through the Microsoft 365 admin center instead. Represents whether the user should be included in the Outlook global address list. See Known issue.
    showInAddressList *bool
    // Get the last signed-in date and request ID of the sign-in for a given user. Read-only.Returned only on $select. Supports $filter (eq, ne, not, ge, le) but, not with any other filterable properties. Note: Details for this property require an Azure AD Premium P1/P2 license and the AuditLog.Read.All permission.Note: There's a known issue with retrieving this property.This property is not returned for a user who has never signed in or last signed in before April 2020.
    signInActivity SignInActivityable
    // Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset.
    signInSessionsValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A list for the user to enumerate their skills. Returned only on $select.
    skills []string
    // The state or province in the user's address. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    state *string
    // The street address of the user's place of business. Maximum length is 1024 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    streetAddress *string
    // The user's surname (family name or last name). Maximum length is 64 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    surname *string
    // The tasks property
    tasks Tasksable
    // A container for Microsoft Teams features available for the user. Read-only. Nullable.
    teamwork UserTeamworkable
    // Represents the To Do services available to a user.
    todo Todoable
    // The groups, including nested groups, and directory roles that a user is a member of. Nullable.
    transitiveMemberOf []DirectoryObjectable
    // The transitive reports for a user. Read-only.
    transitiveReports []DirectoryObjectable
    // A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    usageLocation *string
    // Represents the usage rights a user has been granted.
    usageRights []UsageRightable
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Only the following characters are allowed A - Z, a - z, 0 - 9, ' . - _ ! # ^ ~. For the complete list of allowed characters, see username policies. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderBy.
    userPrincipalName *string
    // A String value that can be used to classify user types in your directory, such as Member and Guest. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
    userType *string
    // Zero or more WIP device registrations that belong to the user.
    windowsInformationProtectionDeviceRegistrations []WindowsInformationProtectionDeviceRegistrationable
}
// NewUser instantiates a new User and sets the default values.
func NewUser()(*User) {
    m := &User{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.user";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser(), nil
}
// GetAboutMe gets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
func (m *User) GetAboutMe()(*string) {
    return m.aboutMe
}
// GetAccountEnabled gets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter (eq, ne, not, and in).
func (m *User) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetActivities gets the activities property value. The activities property
func (m *User) GetActivities()([]UserActivityable) {
    return m.activities
}
// GetAgeGroup gets the ageGroup property value. Sets the age group of the user. Allowed values: null, Minor, NotAdult and Adult. Refer to the legal age group property definitions for further information. Supports $filter (eq, ne, not, and in).
func (m *User) GetAgeGroup()(*string) {
    return m.ageGroup
}
// GetAgreementAcceptances gets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
func (m *User) GetAgreementAcceptances()([]AgreementAcceptanceable) {
    return m.agreementAcceptances
}
// GetAnalytics gets the analytics property value. The analytics property
func (m *User) GetAnalytics()(UserAnalyticsable) {
    return m.analytics
}
// GetAppConsentRequestsForApproval gets the appConsentRequestsForApproval property value. The appConsentRequestsForApproval property
func (m *User) GetAppConsentRequestsForApproval()([]AppConsentRequestable) {
    return m.appConsentRequestsForApproval
}
// GetAppRoleAssignedResources gets the appRoleAssignedResources property value. The appRoleAssignedResources property
func (m *User) GetAppRoleAssignedResources()([]ServicePrincipalable) {
    return m.appRoleAssignedResources
}
// GetAppRoleAssignments gets the appRoleAssignments property value. Represents the app roles a user has been granted for an application. Supports $expand.
func (m *User) GetAppRoleAssignments()([]AppRoleAssignmentable) {
    return m.appRoleAssignments
}
// GetApprovals gets the approvals property value. The approvals property
func (m *User) GetApprovals()([]Approvalable) {
    return m.approvals
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses. Not nullable. Supports $filter (eq, not, and counting empty collections).
func (m *User) GetAssignedLicenses()([]AssignedLicenseable) {
    return m.assignedLicenses
}
// GetAssignedPlans gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.Supports $filter (eq and not).
func (m *User) GetAssignedPlans()([]AssignedPlanable) {
    return m.assignedPlans
}
// GetAuthentication gets the authentication property value. The authentication property
func (m *User) GetAuthentication()(Authenticationable) {
    return m.authentication
}
// GetAuthorizationInfo gets the authorizationInfo property value. Identifiers that can be used to identify and authenticate a user in non-Azure AD environments. This property can be used to store identifiers for smartcard-based certificates that a user uses for access to on-premises Active Directory deployments or for federated access. It can also be used to store the Subject Alternate Name (SAN) that's associated with a Common Access Card (CAC). Nullable.Supports $filter (eq and startsWith).
func (m *User) GetAuthorizationInfo()(AuthorizationInfoable) {
    return m.authorizationInfo
}
// GetBirthday gets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select.
func (m *User) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.birthday
}
// GetBusinessPhones gets the businessPhones property value. The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced from on-premises directory. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) GetBusinessPhones()([]string) {
    return m.businessPhones
}
// GetCalendar gets the calendar property value. The user's primary calendar. Read-only.
func (m *User) GetCalendar()(Calendarable) {
    return m.calendar
}
// GetCalendarGroups gets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
func (m *User) GetCalendarGroups()([]CalendarGroupable) {
    return m.calendarGroups
}
// GetCalendars gets the calendars property value. The user's calendars. Read-only. Nullable.
func (m *User) GetCalendars()([]Calendarable) {
    return m.calendars
}
// GetCalendarView gets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
func (m *User) GetCalendarView()([]Eventable) {
    return m.calendarView
}
// GetChats gets the chats property value. The chats property
func (m *User) GetChats()([]Chatable) {
    return m.chats
}
// GetCity gets the city property value. The city in which the user is located. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetCity()(*string) {
    return m.city
}
// GetCloudPCs gets the cloudPCs property value. The cloudPCs property
func (m *User) GetCloudPCs()([]CloudPCable) {
    return m.cloudPCs
}
// GetCompanyName gets the companyName property value. The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length is 64 characters.Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetCompanyName()(*string) {
    return m.companyName
}
// GetConsentProvidedForMinor gets the consentProvidedForMinor property value. Sets whether consent has been obtained for minors. Allowed values: null, Granted, Denied and NotRequired. Refer to the legal age group property definitions for further information. Supports $filter (eq, ne, not, and in).
func (m *User) GetConsentProvidedForMinor()(*string) {
    return m.consentProvidedForMinor
}
// GetContactFolders gets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
func (m *User) GetContactFolders()([]ContactFolderable) {
    return m.contactFolders
}
// GetContacts gets the contacts property value. The user's contacts. Read-only. Nullable.
func (m *User) GetContacts()([]Contactable) {
    return m.contacts
}
// GetCountry gets the country property value. The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetCountry()(*string) {
    return m.country
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the user was created. The value cannot be modified and is automatically populated when the entity is created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. Property is nullable. A null value indicates that an accurate creation time couldn't be determined for the user. Read-only. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreatedObjects gets the createdObjects property value. Directory objects that were created by the user. Read-only. Nullable.
func (m *User) GetCreatedObjects()([]DirectoryObjectable) {
    return m.createdObjects
}
// GetCreationType gets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp).  Read-only.Supports $filter (eq, ne, not, and in).
func (m *User) GetCreationType()(*string) {
    return m.creationType
}
// GetCustomSecurityAttributes gets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
func (m *User) GetCustomSecurityAttributes()(CustomSecurityAttributeValueable) {
    return m.customSecurityAttributes
}
// GetDepartment gets the department property value. The name for the department in which the user works. Maximum length is 64 characters.Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
func (m *User) GetDepartment()(*string) {
    return m.department
}
// GetDeviceEnrollmentConfigurations gets the deviceEnrollmentConfigurations property value. Get enrollment configurations targeted to the user
func (m *User) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable) {
    return m.deviceEnrollmentConfigurations
}
// GetDeviceEnrollmentLimit gets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
func (m *User) GetDeviceEnrollmentLimit()(*int32) {
    return m.deviceEnrollmentLimit
}
// GetDeviceKeys gets the deviceKeys property value. The deviceKeys property
func (m *User) GetDeviceKeys()([]DeviceKeyable) {
    return m.deviceKeys
}
// GetDeviceManagementTroubleshootingEvents gets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
func (m *User) GetDeviceManagementTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable) {
    return m.deviceManagementTroubleshootingEvents
}
// GetDevices gets the devices property value. The devices property
func (m *User) GetDevices()([]Deviceable) {
    return m.devices
}
// GetDirectReports gets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
func (m *User) GetDirectReports()([]DirectoryObjectable) {
    return m.directReports
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderBy, and $search.
func (m *User) GetDisplayName()(*string) {
    return m.displayName
}
// GetDrive gets the drive property value. The user's OneDrive. Read-only.
func (m *User) GetDrive()(Driveable) {
    return m.drive
}
// GetDrives gets the drives property value. A collection of drives available for this user. Read-only.
func (m *User) GetDrives()([]Driveable) {
    return m.drives
}
// GetEmployeeHireDate gets the employeeHireDate property value. The date and time when the user was hired or will start work in case of a future hire. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) GetEmployeeHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.employeeHireDate
}
// GetEmployeeId gets the employeeId property value. The employee identifier assigned to the user by the organization. The maximum length is 16 characters.Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetEmployeeId()(*string) {
    return m.employeeId
}
// GetEmployeeLeaveDateTime gets the employeeLeaveDateTime property value. The date and time when the user left or will leave the organization. Read: Requires User-LifeCycleInfo.Read.All. For delegated scenarios, the admin needs one of the following Azure AD roles: Lifecycle Workflows Administrator, Global Reader, or Global Admin.  Write: Requires User-LifeCycleInfo.ReadWrite.All. For delegated scenarios, the admin needs the Global Administrator Azure AD role. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) GetEmployeeLeaveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.employeeLeaveDateTime
}
// GetEmployeeOrgData gets the employeeOrgData property value. Represents organization data (e.g. division and costCenter) associated with a user. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) GetEmployeeOrgData()(EmployeeOrgDataable) {
    return m.employeeOrgData
}
// GetEmployeeType gets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Supports $filter (eq, ne, not , ge, le, in, startsWith).
func (m *User) GetEmployeeType()(*string) {
    return m.employeeType
}
// GetEvents gets the events property value. The user's events. Default is to show events under the Default Calendar. Read-only. Nullable.
func (m *User) GetEvents()([]Eventable) {
    return m.events
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the user. Supports $expand. Nullable.
func (m *User) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetExternalUserState gets the externalUserState property value. For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Supports $filter (eq, ne, not , in).
func (m *User) GetExternalUserState()(*string) {
    return m.externalUserState
}
// GetExternalUserStateChangeDateTime gets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Supports $filter (eq, ne, not , in).
func (m *User) GetExternalUserStateChangeDateTime()(*string) {
    return m.externalUserStateChangeDateTime
}
// GetFaxNumber gets the faxNumber property value. The fax number of the user. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetFaxNumber()(*string) {
    return m.faxNumber
}
// GetFieldDeserializers the deserialization information for the current model
func (m *User) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["aboutMe"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAboutMe)
    res["accountEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAccountEnabled)
    res["activities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserActivityFromDiscriminatorValue , m.SetActivities)
    res["ageGroup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAgeGroup)
    res["agreementAcceptances"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAgreementAcceptanceFromDiscriminatorValue , m.SetAgreementAcceptances)
    res["analytics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserAnalyticsFromDiscriminatorValue , m.SetAnalytics)
    res["appConsentRequestsForApproval"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppConsentRequestFromDiscriminatorValue , m.SetAppConsentRequestsForApproval)
    res["appRoleAssignedResources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateServicePrincipalFromDiscriminatorValue , m.SetAppRoleAssignedResources)
    res["appRoleAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppRoleAssignmentFromDiscriminatorValue , m.SetAppRoleAssignments)
    res["approvals"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue , m.SetApprovals)
    res["assignedLicenses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAssignedLicenseFromDiscriminatorValue , m.SetAssignedLicenses)
    res["assignedPlans"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAssignedPlanFromDiscriminatorValue , m.SetAssignedPlans)
    res["authentication"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationFromDiscriminatorValue , m.SetAuthentication)
    res["authorizationInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthorizationInfoFromDiscriminatorValue , m.SetAuthorizationInfo)
    res["birthday"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetBirthday)
    res["businessPhones"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetBusinessPhones)
    res["calendar"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCalendarFromDiscriminatorValue , m.SetCalendar)
    res["calendarGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCalendarGroupFromDiscriminatorValue , m.SetCalendarGroups)
    res["calendars"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCalendarFromDiscriminatorValue , m.SetCalendars)
    res["calendarView"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEventFromDiscriminatorValue , m.SetCalendarView)
    res["chats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateChatFromDiscriminatorValue , m.SetChats)
    res["city"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCity)
    res["cloudPCs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPCFromDiscriminatorValue , m.SetCloudPCs)
    res["companyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompanyName)
    res["consentProvidedForMinor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConsentProvidedForMinor)
    res["contactFolders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateContactFolderFromDiscriminatorValue , m.SetContactFolders)
    res["contacts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateContactFromDiscriminatorValue , m.SetContacts)
    res["country"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCountry)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["createdObjects"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetCreatedObjects)
    res["creationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreationType)
    res["customSecurityAttributes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCustomSecurityAttributeValueFromDiscriminatorValue , m.SetCustomSecurityAttributes)
    res["department"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDepartment)
    res["deviceEnrollmentConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceEnrollmentConfigurationFromDiscriminatorValue , m.SetDeviceEnrollmentConfigurations)
    res["deviceEnrollmentLimit"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeviceEnrollmentLimit)
    res["deviceKeys"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceKeyFromDiscriminatorValue , m.SetDeviceKeys)
    res["deviceManagementTroubleshootingEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue , m.SetDeviceManagementTroubleshootingEvents)
    res["devices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceFromDiscriminatorValue , m.SetDevices)
    res["directReports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetDirectReports)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["drive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDriveFromDiscriminatorValue , m.SetDrive)
    res["drives"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDriveFromDiscriminatorValue , m.SetDrives)
    res["employeeHireDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEmployeeHireDate)
    res["employeeId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmployeeId)
    res["employeeLeaveDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEmployeeLeaveDateTime)
    res["employeeOrgData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEmployeeOrgDataFromDiscriminatorValue , m.SetEmployeeOrgData)
    res["employeeType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmployeeType)
    res["events"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEventFromDiscriminatorValue , m.SetEvents)
    res["extensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue , m.SetExtensions)
    res["externalUserState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalUserState)
    res["externalUserStateChangeDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalUserStateChangeDateTime)
    res["faxNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFaxNumber)
    res["followedSites"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSiteFromDiscriminatorValue , m.SetFollowedSites)
    res["givenName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGivenName)
    res["hireDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetHireDate)
    res["identities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateObjectIdentityFromDiscriminatorValue , m.SetIdentities)
    res["imAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetImAddresses)
    res["inferenceClassification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInferenceClassificationFromDiscriminatorValue , m.SetInferenceClassification)
    res["infoCatalogs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetInfoCatalogs)
    res["informationProtection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInformationProtectionFromDiscriminatorValue , m.SetInformationProtection)
    res["insights"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemInsightsFromDiscriminatorValue , m.SetInsights)
    res["interests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetInterests)
    res["isManagementRestricted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsManagementRestricted)
    res["isResourceAccount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsResourceAccount)
    res["jobTitle"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJobTitle)
    res["joinedGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue , m.SetJoinedGroups)
    res["joinedTeams"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTeamFromDiscriminatorValue , m.SetJoinedTeams)
    res["lastPasswordChangeDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastPasswordChangeDateTime)
    res["legalAgeGroupClassification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLegalAgeGroupClassification)
    res["licenseAssignmentStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLicenseAssignmentStateFromDiscriminatorValue , m.SetLicenseAssignmentStates)
    res["licenseDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLicenseDetailsFromDiscriminatorValue , m.SetLicenseDetails)
    res["mail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMail)
    res["mailboxSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMailboxSettingsFromDiscriminatorValue , m.SetMailboxSettings)
    res["mailFolders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMailFolderFromDiscriminatorValue , m.SetMailFolders)
    res["mailNickname"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMailNickname)
    res["managedAppRegistrations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedAppRegistrationFromDiscriminatorValue , m.SetManagedAppRegistrations)
    res["managedDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue , m.SetManagedDevices)
    res["manager"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDirectoryObjectFromDiscriminatorValue , m.SetManager)
    res["memberOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetMemberOf)
    res["messages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMessageFromDiscriminatorValue , m.SetMessages)
    res["mobileAppIntentAndStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppIntentAndStateFromDiscriminatorValue , m.SetMobileAppIntentAndStates)
    res["mobileAppTroubleshootingEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppTroubleshootingEventFromDiscriminatorValue , m.SetMobileAppTroubleshootingEvents)
    res["mobilePhone"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMobilePhone)
    res["mySite"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMySite)
    res["notifications"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateNotificationFromDiscriminatorValue , m.SetNotifications)
    res["oauth2PermissionGrants"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOAuth2PermissionGrantFromDiscriminatorValue , m.SetOauth2PermissionGrants)
    res["officeLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOfficeLocation)
    res["onenote"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOnenoteFromDiscriminatorValue , m.SetOnenote)
    res["onlineMeetings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOnlineMeetingFromDiscriminatorValue , m.SetOnlineMeetings)
    res["onPremisesDistinguishedName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnPremisesDistinguishedName)
    res["onPremisesDomainName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnPremisesDomainName)
    res["onPremisesExtensionAttributes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOnPremisesExtensionAttributesFromDiscriminatorValue , m.SetOnPremisesExtensionAttributes)
    res["onPremisesImmutableId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnPremisesImmutableId)
    res["onPremisesLastSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetOnPremisesLastSyncDateTime)
    res["onPremisesProvisioningErrors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOnPremisesProvisioningErrorFromDiscriminatorValue , m.SetOnPremisesProvisioningErrors)
    res["onPremisesSamAccountName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnPremisesSamAccountName)
    res["onPremisesSecurityIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnPremisesSecurityIdentifier)
    res["onPremisesSyncEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetOnPremisesSyncEnabled)
    res["onPremisesUserPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnPremisesUserPrincipalName)
    res["otherMails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetOtherMails)
    res["outlook"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOutlookUserFromDiscriminatorValue , m.SetOutlook)
    res["ownedDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetOwnedDevices)
    res["ownedObjects"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetOwnedObjects)
    res["passwordPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPasswordPolicies)
    res["passwordProfile"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePasswordProfileFromDiscriminatorValue , m.SetPasswordProfile)
    res["pastProjects"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetPastProjects)
    res["pendingAccessReviewInstances"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewInstanceFromDiscriminatorValue , m.SetPendingAccessReviewInstances)
    res["people"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonFromDiscriminatorValue , m.SetPeople)
    res["photo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateProfilePhotoFromDiscriminatorValue , m.SetPhoto)
    res["photos"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateProfilePhotoFromDiscriminatorValue , m.SetPhotos)
    res["planner"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePlannerUserFromDiscriminatorValue , m.SetPlanner)
    res["postalCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPostalCode)
    res["preferredDataLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreferredDataLocation)
    res["preferredLanguage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreferredLanguage)
    res["preferredName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPreferredName)
    res["presence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePresenceFromDiscriminatorValue , m.SetPresence)
    res["print"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserPrintFromDiscriminatorValue , m.SetPrint)
    res["profile"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateProfileFromDiscriminatorValue , m.SetProfile)
    res["provisionedPlans"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateProvisionedPlanFromDiscriminatorValue , m.SetProvisionedPlans)
    res["proxyAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetProxyAddresses)
    res["refreshTokensValidFromDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRefreshTokensValidFromDateTime)
    res["registeredDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetRegisteredDevices)
    res["responsibilities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetResponsibilities)
    res["schools"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSchools)
    res["scopedRoleMemberOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateScopedRoleMembershipFromDiscriminatorValue , m.SetScopedRoleMemberOf)
    res["securityIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSecurityIdentifier)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserSettingsFromDiscriminatorValue , m.SetSettings)
    res["showInAddressList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowInAddressList)
    res["signInActivity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSignInActivityFromDiscriminatorValue , m.SetSignInActivity)
    res["signInSessionsValidFromDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetSignInSessionsValidFromDateTime)
    res["skills"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSkills)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetState)
    res["streetAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStreetAddress)
    res["surname"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSurname)
    res["tasks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTasksFromDiscriminatorValue , m.SetTasks)
    res["teamwork"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserTeamworkFromDiscriminatorValue , m.SetTeamwork)
    res["todo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTodoFromDiscriminatorValue , m.SetTodo)
    res["transitiveMemberOf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetTransitiveMemberOf)
    res["transitiveReports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetTransitiveReports)
    res["usageLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUsageLocation)
    res["usageRights"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUsageRightFromDiscriminatorValue , m.SetUsageRights)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    res["userType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserType)
    res["windowsInformationProtectionDeviceRegistrations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue , m.SetWindowsInformationProtectionDeviceRegistrations)
    return res
}
// GetFollowedSites gets the followedSites property value. The followedSites property
func (m *User) GetFollowedSites()([]Siteable) {
    return m.followedSites
}
// GetGivenName gets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetGivenName()(*string) {
    return m.givenName
}
// GetHireDate gets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
func (m *User) GetHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.hireDate
}
// GetIdentities gets the identities property value. Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Supports $filter (eq) including on null values, only where the signInType is not userPrincipalName.
func (m *User) GetIdentities()([]ObjectIdentityable) {
    return m.identities
}
// GetImAddresses gets the imAddresses property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) GetImAddresses()([]string) {
    return m.imAddresses
}
// GetInferenceClassification gets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *User) GetInferenceClassification()(InferenceClassificationable) {
    return m.inferenceClassification
}
// GetInfoCatalogs gets the infoCatalogs property value. Identifies the info segments assigned to the user.  Supports $filter (eq, not, ge, le, startsWith).
func (m *User) GetInfoCatalogs()([]string) {
    return m.infoCatalogs
}
// GetInformationProtection gets the informationProtection property value. The informationProtection property
func (m *User) GetInformationProtection()(InformationProtectionable) {
    return m.informationProtection
}
// GetInsights gets the insights property value. The insights property
func (m *User) GetInsights()(ItemInsightsable) {
    return m.insights
}
// GetInterests gets the interests property value. A list for the user to describe their interests. Returned only on $select.
func (m *User) GetInterests()([]string) {
    return m.interests
}
// GetIsManagementRestricted gets the isManagementRestricted property value. The isManagementRestricted property
func (m *User) GetIsManagementRestricted()(*bool) {
    return m.isManagementRestricted
}
// GetIsResourceAccount gets the isResourceAccount property value. Do not use – reserved for future use.
func (m *User) GetIsResourceAccount()(*bool) {
    return m.isResourceAccount
}
// GetJobTitle gets the jobTitle property value. The user's job title. Maximum length is 128 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetJoinedGroups gets the joinedGroups property value. The joinedGroups property
func (m *User) GetJoinedGroups()([]Groupable) {
    return m.joinedGroups
}
// GetJoinedTeams gets the joinedTeams property value. The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *User) GetJoinedTeams()([]Teamable) {
    return m.joinedTeams
}
// GetLastPasswordChangeDateTime gets the lastPasswordChangeDateTime property value. The time when this Azure AD user last changed their password or when their password was created, , whichever date the latest action was performed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select.
func (m *User) GetLastPasswordChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastPasswordChangeDateTime
}
// GetLegalAgeGroupClassification gets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, MinorWithOutParentalConsent, MinorWithParentalConsent, MinorNoParentalConsentRequired, NotAdult and Adult. Refer to the legal age group property definitions for further information. Returned only on $select.
func (m *User) GetLegalAgeGroupClassification()(*string) {
    return m.legalAgeGroupClassification
}
// GetLicenseAssignmentStates gets the licenseAssignmentStates property value. State of license assignments for this user. Read-only. Returned only on $select.
func (m *User) GetLicenseAssignmentStates()([]LicenseAssignmentStateable) {
    return m.licenseAssignmentStates
}
// GetLicenseDetails gets the licenseDetails property value. The licenseDetails property
func (m *User) GetLicenseDetails()([]LicenseDetailsable) {
    return m.licenseDetails
}
// GetMail gets the mail property value. The SMTP address for the user, for example, admin@contoso.com. Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. This property cannot contain accent characters.  NOTE: We do not recommend updating this property for Azure AD B2C user profiles. Use the otherMails property instead.  Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
func (m *User) GetMail()(*string) {
    return m.mail
}
// GetMailboxSettings gets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. For more information, see User preferences for languages and regional formats. Returned only on $select.
func (m *User) GetMailboxSettings()(MailboxSettingsable) {
    return m.mailboxSettings
}
// GetMailFolders gets the mailFolders property value. The user's mail folders. Read-only. Nullable.
func (m *User) GetMailFolders()([]MailFolderable) {
    return m.mailFolders
}
// GetMailNickname gets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
func (m *User) GetManagedAppRegistrations()([]ManagedAppRegistrationable) {
    return m.managedAppRegistrations
}
// GetManagedDevices gets the managedDevices property value. The managed devices associated with the user.
func (m *User) GetManagedDevices()([]ManagedDeviceable) {
    return m.managedDevices
}
// GetManager gets the manager property value. The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
func (m *User) GetManager()(DirectoryObjectable) {
    return m.manager
}
// GetMemberOf gets the memberOf property value. The groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *User) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetMessages gets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
func (m *User) GetMessages()([]Messageable) {
    return m.messages
}
// GetMobileAppIntentAndStates gets the mobileAppIntentAndStates property value. The list of troubleshooting events for this user.
func (m *User) GetMobileAppIntentAndStates()([]MobileAppIntentAndStateable) {
    return m.mobileAppIntentAndStates
}
// GetMobileAppTroubleshootingEvents gets the mobileAppTroubleshootingEvents property value. The list of mobile app troubleshooting events for this user.
func (m *User) GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable) {
    return m.mobileAppTroubleshootingEvents
}
// GetMobilePhone gets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from on-premises directory.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetMobilePhone()(*string) {
    return m.mobilePhone
}
// GetMySite gets the mySite property value. The URL for the user's personal site. Returned only on $select.
func (m *User) GetMySite()(*string) {
    return m.mySite
}
// GetNotifications gets the notifications property value. The notifications property
func (m *User) GetNotifications()([]Notificationable) {
    return m.notifications
}
// GetOauth2PermissionGrants gets the oauth2PermissionGrants property value. The oauth2PermissionGrants property
func (m *User) GetOauth2PermissionGrants()([]OAuth2PermissionGrantable) {
    return m.oauth2PermissionGrants
}
// GetOfficeLocation gets the officeLocation property value. The office location in the user's place of business. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetOnenote gets the onenote property value. The onenote property
func (m *User) GetOnenote()(Onenoteable) {
    return m.onenote
}
// GetOnlineMeetings gets the onlineMeetings property value. The onlineMeetings property
func (m *User) GetOnlineMeetings()([]OnlineMeetingable) {
    return m.onlineMeetings
}
// GetOnPremisesDistinguishedName gets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
func (m *User) GetOnPremisesDistinguishedName()(*string) {
    return m.onPremisesDistinguishedName
}
// GetOnPremisesDomainName gets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
func (m *User) GetOnPremisesDomainName()(*string) {
    return m.onPremisesDomainName
}
// GetOnPremisesExtensionAttributes gets the onPremisesExtensionAttributes property value. Contains extensionAttributes1-15 for the user. These extension attributes are also known as Exchange custom attributes 1-15. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. Supports $filter (eq, ne, not, in).
func (m *User) GetOnPremisesExtensionAttributes()(OnPremisesExtensionAttributesable) {
    return m.onPremisesExtensionAttributes
}
// GetOnPremisesImmutableId gets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. Note: The $ and _ characters cannot be used when specifying this property. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) GetOnPremisesImmutableId()(*string) {
    return m.onPremisesImmutableId
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: '2013-02-16T03:04:54Z'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesProvisioningErrors gets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning.  Supports $filter (eq, not, ge, le).
func (m *User) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable) {
    return m.onPremisesProvisioningErrors
}
// GetOnPremisesSamAccountName gets the onPremisesSamAccountName property value. Contains the on-premises sAMAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) GetOnPremisesSamAccountName()(*string) {
    return m.onPremisesSamAccountName
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Supports $filter (eq including on null values).
func (m *User) GetOnPremisesSecurityIdentifier()(*string) {
    return m.onPremisesSecurityIdentifier
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this user object is currently being synced from an on-premises Active Directory (AD); otherwise the user isn't being synced and can be managed in Azure Active Directory (Azure AD). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetOnPremisesUserPrincipalName gets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) GetOnPremisesUserPrincipalName()(*string) {
    return m.onPremisesUserPrincipalName
}
// GetOtherMails gets the otherMails property value. A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com'].NOTE: This property cannot contain accent characters.Supports $filter (eq, not, ge, le, in, startsWith, endsWith, and counting empty collections).
func (m *User) GetOtherMails()([]string) {
    return m.otherMails
}
// GetOutlook gets the outlook property value. Selective Outlook services available to the user. Read-only. Nullable.
func (m *User) GetOutlook()(OutlookUserable) {
    return m.outlook
}
// GetOwnedDevices gets the ownedDevices property value. Devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) GetOwnedDevices()([]DirectoryObjectable) {
    return m.ownedDevices
}
// GetOwnedObjects gets the ownedObjects property value. Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) GetOwnedObjects()([]DirectoryObjectable) {
    return m.ownedObjects
}
// GetPasswordPolicies gets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, not, and eq on null values).
func (m *User) GetPasswordPolicies()(*string) {
    return m.passwordPolicies
}
// GetPasswordProfile gets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) GetPasswordProfile()(PasswordProfileable) {
    return m.passwordProfile
}
// GetPastProjects gets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
func (m *User) GetPastProjects()([]string) {
    return m.pastProjects
}
// GetPendingAccessReviewInstances gets the pendingAccessReviewInstances property value. Navigation property to get list of access reviews pending approval by reviewer.
func (m *User) GetPendingAccessReviewInstances()([]AccessReviewInstanceable) {
    return m.pendingAccessReviewInstances
}
// GetPeople gets the people property value. Read-only. The most relevant people to the user. The collection is ordered by their relevance to the user, which is determined by the user's communication, collaboration and business relationships. A person is an aggregation of information from across mail, contacts and social networks.
func (m *User) GetPeople()([]Personable) {
    return m.people
}
// GetPhoto gets the photo property value. The user's profile photo. Read-only.
func (m *User) GetPhoto()(ProfilePhotoable) {
    return m.photo
}
// GetPhotos gets the photos property value. The photos property
func (m *User) GetPhotos()([]ProfilePhotoable) {
    return m.photos
}
// GetPlanner gets the planner property value. Selective Planner services available to the user. Read-only. Nullable.
func (m *User) GetPlanner()(PlannerUserable) {
    return m.planner
}
// GetPostalCode gets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetPostalCode()(*string) {
    return m.postalCode
}
// GetPreferredDataLocation gets the preferredDataLocation property value. The preferred data location for the user. For more information, see OneDrive Online Multi-Geo.
func (m *User) GetPreferredDataLocation()(*string) {
    return m.preferredDataLocation
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetPreferredLanguage()(*string) {
    return m.preferredLanguage
}
// GetPreferredName gets the preferredName property value. The preferred name for the user. Not Supported. This attribute returns an empty string.Returned only on $select.
func (m *User) GetPreferredName()(*string) {
    return m.preferredName
}
// GetPresence gets the presence property value. The presence property
func (m *User) GetPresence()(Presenceable) {
    return m.presence
}
// GetPrint gets the print property value. The print property
func (m *User) GetPrint()(UserPrintable) {
    return m.print
}
// GetProfile gets the profile property value. Represents properties that are descriptive of a user in a tenant.
func (m *User) GetProfile()(Profileable) {
    return m.profile
}
// GetProvisionedPlans gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *User) GetProvisionedPlans()([]ProvisionedPlanable) {
    return m.provisionedPlans
}
// GetProxyAddresses gets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. Changes to the mail property will also update this collection to include the value as an SMTP address. For more information, see mail and proxyAddresses properties. The proxy address prefixed with SMTP (capitalized) is the primary proxy address while those prefixed with smtp are the secondary proxy addresses. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only in Microsoft Graph; you can update this property only through the Microsoft 365 admin center. Not nullable. Supports $filter (eq, not, ge, le, startsWith, endsWith, and counting empty collections).
func (m *User) GetProxyAddresses()([]string) {
    return m.proxyAddresses
}
// GetRefreshTokensValidFromDateTime gets the refreshTokensValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use invalidateAllRefreshTokens to reset.
func (m *User) GetRefreshTokensValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.refreshTokensValidFromDateTime
}
// GetRegisteredDevices gets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *User) GetRegisteredDevices()([]DirectoryObjectable) {
    return m.registeredDevices
}
// GetResponsibilities gets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
func (m *User) GetResponsibilities()([]string) {
    return m.responsibilities
}
// GetSchools gets the schools property value. A list for the user to enumerate the schools they have attended. Returned only on $select.
func (m *User) GetSchools()([]string) {
    return m.schools
}
// GetScopedRoleMemberOf gets the scopedRoleMemberOf property value. The scoped-role administrative unit memberships for this user. Read-only. Nullable.
func (m *User) GetScopedRoleMemberOf()([]ScopedRoleMembershipable) {
    return m.scopedRoleMemberOf
}
// GetSecurityIdentifier gets the securityIdentifier property value. Security identifier (SID) of the user, used in Windows scenarios. Read-only. Returned by default. Supports $select and $filter (eq, not, ge, le, startsWith).
func (m *User) GetSecurityIdentifier()(*string) {
    return m.securityIdentifier
}
// GetSettings gets the settings property value. The settings property
func (m *User) GetSettings()(UserSettingsable) {
    return m.settings
}
// GetShowInAddressList gets the showInAddressList property value. Do not use in Microsoft Graph. Manage this property through the Microsoft 365 admin center instead. Represents whether the user should be included in the Outlook global address list. See Known issue.
func (m *User) GetShowInAddressList()(*bool) {
    return m.showInAddressList
}
// GetSignInActivity gets the signInActivity property value. Get the last signed-in date and request ID of the sign-in for a given user. Read-only.Returned only on $select. Supports $filter (eq, ne, not, ge, le) but, not with any other filterable properties. Note: Details for this property require an Azure AD Premium P1/P2 license and the AuditLog.Read.All permission.Note: There's a known issue with retrieving this property.This property is not returned for a user who has never signed in or last signed in before April 2020.
func (m *User) GetSignInActivity()(SignInActivityable) {
    return m.signInActivity
}
// GetSignInSessionsValidFromDateTime gets the signInSessionsValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset.
func (m *User) GetSignInSessionsValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.signInSessionsValidFromDateTime
}
// GetSkills gets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
func (m *User) GetSkills()([]string) {
    return m.skills
}
// GetState gets the state property value. The state or province in the user's address. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetState()(*string) {
    return m.state
}
// GetStreetAddress gets the streetAddress property value. The street address of the user's place of business. Maximum length is 1024 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetStreetAddress()(*string) {
    return m.streetAddress
}
// GetSurname gets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetSurname()(*string) {
    return m.surname
}
// GetTasks gets the tasks property value. The tasks property
func (m *User) GetTasks()(Tasksable) {
    return m.tasks
}
// GetTeamwork gets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *User) GetTeamwork()(UserTeamworkable) {
    return m.teamwork
}
// GetTodo gets the todo property value. Represents the To Do services available to a user.
func (m *User) GetTodo()(Todoable) {
    return m.todo
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. The groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *User) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetTransitiveReports gets the transitiveReports property value. The transitive reports for a user. Read-only.
func (m *User) GetTransitiveReports()([]DirectoryObjectable) {
    return m.transitiveReports
}
// GetUsageLocation gets the usageLocation property value. A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetUsageLocation()(*string) {
    return m.usageLocation
}
// GetUsageRights gets the usageRights property value. Represents the usage rights a user has been granted.
func (m *User) GetUsageRights()([]UsageRightable) {
    return m.usageRights
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Only the following characters are allowed A - Z, a - z, 0 - 9, ' . - _ ! # ^ ~. For the complete list of allowed characters, see username policies. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderBy.
func (m *User) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserType gets the userType property value. A String value that can be used to classify user types in your directory, such as Member and Guest. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
func (m *User) GetUserType()(*string) {
    return m.userType
}
// GetWindowsInformationProtectionDeviceRegistrations gets the windowsInformationProtectionDeviceRegistrations property value. Zero or more WIP device registrations that belong to the user.
func (m *User) GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistrationable) {
    return m.windowsInformationProtectionDeviceRegistrations
}
// Serialize serializes information the current object
func (m *User) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aboutMe", m.GetAboutMe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetActivities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActivities())
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ageGroup", m.GetAgeGroup())
        if err != nil {
            return err
        }
    }
    if m.GetAgreementAcceptances() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAgreementAcceptances())
        err = writer.WriteCollectionOfObjectValues("agreementAcceptances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("analytics", m.GetAnalytics())
        if err != nil {
            return err
        }
    }
    if m.GetAppConsentRequestsForApproval() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppConsentRequestsForApproval())
        err = writer.WriteCollectionOfObjectValues("appConsentRequestsForApproval", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppRoleAssignedResources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppRoleAssignedResources())
        err = writer.WriteCollectionOfObjectValues("appRoleAssignedResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppRoleAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppRoleAssignments())
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetApprovals() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApprovals())
        err = writer.WriteCollectionOfObjectValues("approvals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignedLicenses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignedLicenses())
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignedPlans() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignedPlans())
        err = writer.WriteCollectionOfObjectValues("assignedPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authentication", m.GetAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizationInfo", m.GetAuthorizationInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    if m.GetBusinessPhones() != nil {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    if m.GetCalendarGroups() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCalendarGroups())
        err = writer.WriteCollectionOfObjectValues("calendarGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCalendars() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCalendars())
        err = writer.WriteCollectionOfObjectValues("calendars", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCalendarView() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCalendarView())
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChats() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChats())
        err = writer.WriteCollectionOfObjectValues("chats", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    if m.GetCloudPCs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudPCs())
        err = writer.WriteCollectionOfObjectValues("cloudPCs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("consentProvidedForMinor", m.GetConsentProvidedForMinor())
        if err != nil {
            return err
        }
    }
    if m.GetContactFolders() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContactFolders())
        err = writer.WriteCollectionOfObjectValues("contactFolders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContacts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContacts())
        err = writer.WriteCollectionOfObjectValues("contacts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
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
    if m.GetCreatedObjects() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCreatedObjects())
        err = writer.WriteCollectionOfObjectValues("createdObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creationType", m.GetCreationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customSecurityAttributes", m.GetCustomSecurityAttributes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceEnrollmentConfigurations())
        err = writer.WriteCollectionOfObjectValues("deviceEnrollmentConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceEnrollmentLimit", m.GetDeviceEnrollmentLimit())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceKeys() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceKeys())
        err = writer.WriteCollectionOfObjectValues("deviceKeys", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementTroubleshootingEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceManagementTroubleshootingEvents())
        err = writer.WriteCollectionOfObjectValues("deviceManagementTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDevices())
        err = writer.WriteCollectionOfObjectValues("devices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDirectReports() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDirectReports())
        err = writer.WriteCollectionOfObjectValues("directReports", cast)
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
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    if m.GetDrives() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDrives())
        err = writer.WriteCollectionOfObjectValues("drives", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("employeeHireDate", m.GetEmployeeHireDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("employeeId", m.GetEmployeeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("employeeLeaveDateTime", m.GetEmployeeLeaveDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("employeeOrgData", m.GetEmployeeOrgData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("employeeType", m.GetEmployeeType())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEvents())
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensions())
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalUserState", m.GetExternalUserState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalUserStateChangeDateTime", m.GetExternalUserStateChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faxNumber", m.GetFaxNumber())
        if err != nil {
            return err
        }
    }
    if m.GetFollowedSites() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFollowedSites())
        err = writer.WriteCollectionOfObjectValues("followedSites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("hireDate", m.GetHireDate())
        if err != nil {
            return err
        }
    }
    if m.GetIdentities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIdentities())
        err = writer.WriteCollectionOfObjectValues("identities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("imAddresses", m.GetImAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inferenceClassification", m.GetInferenceClassification())
        if err != nil {
            return err
        }
    }
    if m.GetInfoCatalogs() != nil {
        err = writer.WriteCollectionOfStringValues("infoCatalogs", m.GetInfoCatalogs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("informationProtection", m.GetInformationProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    if m.GetInterests() != nil {
        err = writer.WriteCollectionOfStringValues("interests", m.GetInterests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManagementRestricted", m.GetIsManagementRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isResourceAccount", m.GetIsResourceAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    if m.GetJoinedGroups() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetJoinedGroups())
        err = writer.WriteCollectionOfObjectValues("joinedGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJoinedTeams() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetJoinedTeams())
        err = writer.WriteCollectionOfObjectValues("joinedTeams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastPasswordChangeDateTime", m.GetLastPasswordChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("legalAgeGroupClassification", m.GetLegalAgeGroupClassification())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseAssignmentStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLicenseAssignmentStates())
        err = writer.WriteCollectionOfObjectValues("licenseAssignmentStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLicenseDetails())
        err = writer.WriteCollectionOfObjectValues("licenseDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mailboxSettings", m.GetMailboxSettings())
        if err != nil {
            return err
        }
    }
    if m.GetMailFolders() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMailFolders())
        err = writer.WriteCollectionOfObjectValues("mailFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppRegistrations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedAppRegistrations())
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDevices())
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    if m.GetMemberOf() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMemberOf())
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMessages())
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppIntentAndStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileAppIntentAndStates())
        err = writer.WriteCollectionOfObjectValues("mobileAppIntentAndStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppTroubleshootingEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileAppTroubleshootingEvents())
        err = writer.WriteCollectionOfObjectValues("mobileAppTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mySite", m.GetMySite())
        if err != nil {
            return err
        }
    }
    if m.GetNotifications() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNotifications())
        err = writer.WriteCollectionOfObjectValues("notifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOauth2PermissionGrants() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOauth2PermissionGrants())
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionGrants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onenote", m.GetOnenote())
        if err != nil {
            return err
        }
    }
    if m.GetOnlineMeetings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOnlineMeetings())
        err = writer.WriteCollectionOfObjectValues("onlineMeetings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDistinguishedName", m.GetOnPremisesDistinguishedName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDomainName", m.GetOnPremisesDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremisesExtensionAttributes", m.GetOnPremisesExtensionAttributes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesImmutableId", m.GetOnPremisesImmutableId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesProvisioningErrors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOnPremisesProvisioningErrors())
        err = writer.WriteCollectionOfObjectValues("onPremisesProvisioningErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSamAccountName", m.GetOnPremisesSamAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesUserPrincipalName", m.GetOnPremisesUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetOtherMails() != nil {
        err = writer.WriteCollectionOfStringValues("otherMails", m.GetOtherMails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("outlook", m.GetOutlook())
        if err != nil {
            return err
        }
    }
    if m.GetOwnedDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOwnedDevices())
        err = writer.WriteCollectionOfObjectValues("ownedDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnedObjects() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOwnedObjects())
        err = writer.WriteCollectionOfObjectValues("ownedObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordPolicies", m.GetPasswordPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("passwordProfile", m.GetPasswordProfile())
        if err != nil {
            return err
        }
    }
    if m.GetPastProjects() != nil {
        err = writer.WriteCollectionOfStringValues("pastProjects", m.GetPastProjects())
        if err != nil {
            return err
        }
    }
    if m.GetPendingAccessReviewInstances() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPendingAccessReviewInstances())
        err = writer.WriteCollectionOfObjectValues("pendingAccessReviewInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPeople() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPeople())
        err = writer.WriteCollectionOfObjectValues("people", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    if m.GetPhotos() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPhotos())
        err = writer.WriteCollectionOfObjectValues("photos", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("planner", m.GetPlanner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredDataLocation", m.GetPreferredDataLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredName", m.GetPreferredName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("presence", m.GetPresence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("print", m.GetPrint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("profile", m.GetProfile())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionedPlans() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProvisionedPlans())
        err = writer.WriteCollectionOfObjectValues("provisionedPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProxyAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("refreshTokensValidFromDateTime", m.GetRefreshTokensValidFromDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRegisteredDevices())
        err = writer.WriteCollectionOfObjectValues("registeredDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResponsibilities() != nil {
        err = writer.WriteCollectionOfStringValues("responsibilities", m.GetResponsibilities())
        if err != nil {
            return err
        }
    }
    if m.GetSchools() != nil {
        err = writer.WriteCollectionOfStringValues("schools", m.GetSchools())
        if err != nil {
            return err
        }
    }
    if m.GetScopedRoleMemberOf() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetScopedRoleMemberOf())
        err = writer.WriteCollectionOfObjectValues("scopedRoleMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("securityIdentifier", m.GetSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showInAddressList", m.GetShowInAddressList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signInActivity", m.GetSignInActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("signInSessionsValidFromDateTime", m.GetSignInSessionsValidFromDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSkills() != nil {
        err = writer.WriteCollectionOfStringValues("skills", m.GetSkills())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("streetAddress", m.GetStreetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tasks", m.GetTasks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamwork", m.GetTeamwork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("todo", m.GetTodo())
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveMemberOf() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTransitiveMemberOf())
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveReports() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTransitiveReports())
        err = writer.WriteCollectionOfObjectValues("transitiveReports", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usageLocation", m.GetUsageLocation())
        if err != nil {
            return err
        }
    }
    if m.GetUsageRights() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUsageRights())
        err = writer.WriteCollectionOfObjectValues("usageRights", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userType", m.GetUserType())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionDeviceRegistrations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsInformationProtectionDeviceRegistrations())
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionDeviceRegistrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAboutMe sets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
func (m *User) SetAboutMe(value *string)() {
    m.aboutMe = value
}
// SetAccountEnabled sets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter (eq, ne, not, and in).
func (m *User) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetActivities sets the activities property value. The activities property
func (m *User) SetActivities(value []UserActivityable)() {
    m.activities = value
}
// SetAgeGroup sets the ageGroup property value. Sets the age group of the user. Allowed values: null, Minor, NotAdult and Adult. Refer to the legal age group property definitions for further information. Supports $filter (eq, ne, not, and in).
func (m *User) SetAgeGroup(value *string)() {
    m.ageGroup = value
}
// SetAgreementAcceptances sets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
func (m *User) SetAgreementAcceptances(value []AgreementAcceptanceable)() {
    m.agreementAcceptances = value
}
// SetAnalytics sets the analytics property value. The analytics property
func (m *User) SetAnalytics(value UserAnalyticsable)() {
    m.analytics = value
}
// SetAppConsentRequestsForApproval sets the appConsentRequestsForApproval property value. The appConsentRequestsForApproval property
func (m *User) SetAppConsentRequestsForApproval(value []AppConsentRequestable)() {
    m.appConsentRequestsForApproval = value
}
// SetAppRoleAssignedResources sets the appRoleAssignedResources property value. The appRoleAssignedResources property
func (m *User) SetAppRoleAssignedResources(value []ServicePrincipalable)() {
    m.appRoleAssignedResources = value
}
// SetAppRoleAssignments sets the appRoleAssignments property value. Represents the app roles a user has been granted for an application. Supports $expand.
func (m *User) SetAppRoleAssignments(value []AppRoleAssignmentable)() {
    m.appRoleAssignments = value
}
// SetApprovals sets the approvals property value. The approvals property
func (m *User) SetApprovals(value []Approvalable)() {
    m.approvals = value
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses. Not nullable. Supports $filter (eq, not, and counting empty collections).
func (m *User) SetAssignedLicenses(value []AssignedLicenseable)() {
    m.assignedLicenses = value
}
// SetAssignedPlans sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.Supports $filter (eq and not).
func (m *User) SetAssignedPlans(value []AssignedPlanable)() {
    m.assignedPlans = value
}
// SetAuthentication sets the authentication property value. The authentication property
func (m *User) SetAuthentication(value Authenticationable)() {
    m.authentication = value
}
// SetAuthorizationInfo sets the authorizationInfo property value. Identifiers that can be used to identify and authenticate a user in non-Azure AD environments. This property can be used to store identifiers for smartcard-based certificates that a user uses for access to on-premises Active Directory deployments or for federated access. It can also be used to store the Subject Alternate Name (SAN) that's associated with a Common Access Card (CAC). Nullable.Supports $filter (eq and startsWith).
func (m *User) SetAuthorizationInfo(value AuthorizationInfoable)() {
    m.authorizationInfo = value
}
// SetBirthday sets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select.
func (m *User) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.birthday = value
}
// SetBusinessPhones sets the businessPhones property value. The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced from on-premises directory. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// SetCalendar sets the calendar property value. The user's primary calendar. Read-only.
func (m *User) SetCalendar(value Calendarable)() {
    m.calendar = value
}
// SetCalendarGroups sets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
func (m *User) SetCalendarGroups(value []CalendarGroupable)() {
    m.calendarGroups = value
}
// SetCalendars sets the calendars property value. The user's calendars. Read-only. Nullable.
func (m *User) SetCalendars(value []Calendarable)() {
    m.calendars = value
}
// SetCalendarView sets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
func (m *User) SetCalendarView(value []Eventable)() {
    m.calendarView = value
}
// SetChats sets the chats property value. The chats property
func (m *User) SetChats(value []Chatable)() {
    m.chats = value
}
// SetCity sets the city property value. The city in which the user is located. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCity(value *string)() {
    m.city = value
}
// SetCloudPCs sets the cloudPCs property value. The cloudPCs property
func (m *User) SetCloudPCs(value []CloudPCable)() {
    m.cloudPCs = value
}
// SetCompanyName sets the companyName property value. The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length is 64 characters.Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetConsentProvidedForMinor sets the consentProvidedForMinor property value. Sets whether consent has been obtained for minors. Allowed values: null, Granted, Denied and NotRequired. Refer to the legal age group property definitions for further information. Supports $filter (eq, ne, not, and in).
func (m *User) SetConsentProvidedForMinor(value *string)() {
    m.consentProvidedForMinor = value
}
// SetContactFolders sets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
func (m *User) SetContactFolders(value []ContactFolderable)() {
    m.contactFolders = value
}
// SetContacts sets the contacts property value. The user's contacts. Read-only. Nullable.
func (m *User) SetContacts(value []Contactable)() {
    m.contacts = value
}
// SetCountry sets the country property value. The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCountry(value *string)() {
    m.country = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the user was created. The value cannot be modified and is automatically populated when the entity is created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. Property is nullable. A null value indicates that an accurate creation time couldn't be determined for the user. Read-only. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedObjects sets the createdObjects property value. Directory objects that were created by the user. Read-only. Nullable.
func (m *User) SetCreatedObjects(value []DirectoryObjectable)() {
    m.createdObjects = value
}
// SetCreationType sets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp).  Read-only.Supports $filter (eq, ne, not, and in).
func (m *User) SetCreationType(value *string)() {
    m.creationType = value
}
// SetCustomSecurityAttributes sets the customSecurityAttributes property value. An open complex type that holds the value of a custom security attribute that is assigned to a directory object. Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith).
func (m *User) SetCustomSecurityAttributes(value CustomSecurityAttributeValueable)() {
    m.customSecurityAttributes = value
}
// SetDepartment sets the department property value. The name for the department in which the user works. Maximum length is 64 characters.Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
func (m *User) SetDepartment(value *string)() {
    m.department = value
}
// SetDeviceEnrollmentConfigurations sets the deviceEnrollmentConfigurations property value. Get enrollment configurations targeted to the user
func (m *User) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)() {
    m.deviceEnrollmentConfigurations = value
}
// SetDeviceEnrollmentLimit sets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
func (m *User) SetDeviceEnrollmentLimit(value *int32)() {
    m.deviceEnrollmentLimit = value
}
// SetDeviceKeys sets the deviceKeys property value. The deviceKeys property
func (m *User) SetDeviceKeys(value []DeviceKeyable)() {
    m.deviceKeys = value
}
// SetDeviceManagementTroubleshootingEvents sets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
func (m *User) SetDeviceManagementTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)() {
    m.deviceManagementTroubleshootingEvents = value
}
// SetDevices sets the devices property value. The devices property
func (m *User) SetDevices(value []Deviceable)() {
    m.devices = value
}
// SetDirectReports sets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
func (m *User) SetDirectReports(value []DirectoryObjectable)() {
    m.directReports = value
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderBy, and $search.
func (m *User) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDrive sets the drive property value. The user's OneDrive. Read-only.
func (m *User) SetDrive(value Driveable)() {
    m.drive = value
}
// SetDrives sets the drives property value. A collection of drives available for this user. Read-only.
func (m *User) SetDrives(value []Driveable)() {
    m.drives = value
}
// SetEmployeeHireDate sets the employeeHireDate property value. The date and time when the user was hired or will start work in case of a future hire. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetEmployeeHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.employeeHireDate = value
}
// SetEmployeeId sets the employeeId property value. The employee identifier assigned to the user by the organization. The maximum length is 16 characters.Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetEmployeeId(value *string)() {
    m.employeeId = value
}
// SetEmployeeLeaveDateTime sets the employeeLeaveDateTime property value. The date and time when the user left or will leave the organization. Read: Requires User-LifeCycleInfo.Read.All. For delegated scenarios, the admin needs one of the following Azure AD roles: Lifecycle Workflows Administrator, Global Reader, or Global Admin.  Write: Requires User-LifeCycleInfo.ReadWrite.All. For delegated scenarios, the admin needs the Global Administrator Azure AD role. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetEmployeeLeaveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.employeeLeaveDateTime = value
}
// SetEmployeeOrgData sets the employeeOrgData property value. Represents organization data (e.g. division and costCenter) associated with a user. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetEmployeeOrgData(value EmployeeOrgDataable)() {
    m.employeeOrgData = value
}
// SetEmployeeType sets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Supports $filter (eq, ne, not , ge, le, in, startsWith).
func (m *User) SetEmployeeType(value *string)() {
    m.employeeType = value
}
// SetEvents sets the events property value. The user's events. Default is to show events under the Default Calendar. Read-only. Nullable.
func (m *User) SetEvents(value []Eventable)() {
    m.events = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the user. Supports $expand. Nullable.
func (m *User) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetExternalUserState sets the externalUserState property value. For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Supports $filter (eq, ne, not , in).
func (m *User) SetExternalUserState(value *string)() {
    m.externalUserState = value
}
// SetExternalUserStateChangeDateTime sets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Supports $filter (eq, ne, not , in).
func (m *User) SetExternalUserStateChangeDateTime(value *string)() {
    m.externalUserStateChangeDateTime = value
}
// SetFaxNumber sets the faxNumber property value. The fax number of the user. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetFaxNumber(value *string)() {
    m.faxNumber = value
}
// SetFollowedSites sets the followedSites property value. The followedSites property
func (m *User) SetFollowedSites(value []Siteable)() {
    m.followedSites = value
}
// SetGivenName sets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetGivenName(value *string)() {
    m.givenName = value
}
// SetHireDate sets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
func (m *User) SetHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.hireDate = value
}
// SetIdentities sets the identities property value. Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Supports $filter (eq) including on null values, only where the signInType is not userPrincipalName.
func (m *User) SetIdentities(value []ObjectIdentityable)() {
    m.identities = value
}
// SetImAddresses sets the imAddresses property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetImAddresses(value []string)() {
    m.imAddresses = value
}
// SetInferenceClassification sets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *User) SetInferenceClassification(value InferenceClassificationable)() {
    m.inferenceClassification = value
}
// SetInfoCatalogs sets the infoCatalogs property value. Identifies the info segments assigned to the user.  Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetInfoCatalogs(value []string)() {
    m.infoCatalogs = value
}
// SetInformationProtection sets the informationProtection property value. The informationProtection property
func (m *User) SetInformationProtection(value InformationProtectionable)() {
    m.informationProtection = value
}
// SetInsights sets the insights property value. The insights property
func (m *User) SetInsights(value ItemInsightsable)() {
    m.insights = value
}
// SetInterests sets the interests property value. A list for the user to describe their interests. Returned only on $select.
func (m *User) SetInterests(value []string)() {
    m.interests = value
}
// SetIsManagementRestricted sets the isManagementRestricted property value. The isManagementRestricted property
func (m *User) SetIsManagementRestricted(value *bool)() {
    m.isManagementRestricted = value
}
// SetIsResourceAccount sets the isResourceAccount property value. Do not use – reserved for future use.
func (m *User) SetIsResourceAccount(value *bool)() {
    m.isResourceAccount = value
}
// SetJobTitle sets the jobTitle property value. The user's job title. Maximum length is 128 characters. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetJoinedGroups sets the joinedGroups property value. The joinedGroups property
func (m *User) SetJoinedGroups(value []Groupable)() {
    m.joinedGroups = value
}
// SetJoinedTeams sets the joinedTeams property value. The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *User) SetJoinedTeams(value []Teamable)() {
    m.joinedTeams = value
}
// SetLastPasswordChangeDateTime sets the lastPasswordChangeDateTime property value. The time when this Azure AD user last changed their password or when their password was created, , whichever date the latest action was performed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select.
func (m *User) SetLastPasswordChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastPasswordChangeDateTime = value
}
// SetLegalAgeGroupClassification sets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, MinorWithOutParentalConsent, MinorWithParentalConsent, MinorNoParentalConsentRequired, NotAdult and Adult. Refer to the legal age group property definitions for further information. Returned only on $select.
func (m *User) SetLegalAgeGroupClassification(value *string)() {
    m.legalAgeGroupClassification = value
}
// SetLicenseAssignmentStates sets the licenseAssignmentStates property value. State of license assignments for this user. Read-only. Returned only on $select.
func (m *User) SetLicenseAssignmentStates(value []LicenseAssignmentStateable)() {
    m.licenseAssignmentStates = value
}
// SetLicenseDetails sets the licenseDetails property value. The licenseDetails property
func (m *User) SetLicenseDetails(value []LicenseDetailsable)() {
    m.licenseDetails = value
}
// SetMail sets the mail property value. The SMTP address for the user, for example, admin@contoso.com. Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. This property cannot contain accent characters.  NOTE: We do not recommend updating this property for Azure AD B2C user profiles. Use the otherMails property instead.  Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
func (m *User) SetMail(value *string)() {
    m.mail = value
}
// SetMailboxSettings sets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale, and time zone. For more information, see User preferences for languages and regional formats. Returned only on $select.
func (m *User) SetMailboxSettings(value MailboxSettingsable)() {
    m.mailboxSettings = value
}
// SetMailFolders sets the mailFolders property value. The user's mail folders. Read-only. Nullable.
func (m *User) SetMailFolders(value []MailFolderable)() {
    m.mailFolders = value
}
// SetMailNickname sets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
func (m *User) SetManagedAppRegistrations(value []ManagedAppRegistrationable)() {
    m.managedAppRegistrations = value
}
// SetManagedDevices sets the managedDevices property value. The managed devices associated with the user.
func (m *User) SetManagedDevices(value []ManagedDeviceable)() {
    m.managedDevices = value
}
// SetManager sets the manager property value. The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
func (m *User) SetManager(value DirectoryObjectable)() {
    m.manager = value
}
// SetMemberOf sets the memberOf property value. The groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *User) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetMessages sets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
func (m *User) SetMessages(value []Messageable)() {
    m.messages = value
}
// SetMobileAppIntentAndStates sets the mobileAppIntentAndStates property value. The list of troubleshooting events for this user.
func (m *User) SetMobileAppIntentAndStates(value []MobileAppIntentAndStateable)() {
    m.mobileAppIntentAndStates = value
}
// SetMobileAppTroubleshootingEvents sets the mobileAppTroubleshootingEvents property value. The list of mobile app troubleshooting events for this user.
func (m *User) SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)() {
    m.mobileAppTroubleshootingEvents = value
}
// SetMobilePhone sets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from on-premises directory.  Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// SetMySite sets the mySite property value. The URL for the user's personal site. Returned only on $select.
func (m *User) SetMySite(value *string)() {
    m.mySite = value
}
// SetNotifications sets the notifications property value. The notifications property
func (m *User) SetNotifications(value []Notificationable)() {
    m.notifications = value
}
// SetOauth2PermissionGrants sets the oauth2PermissionGrants property value. The oauth2PermissionGrants property
func (m *User) SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)() {
    m.oauth2PermissionGrants = value
}
// SetOfficeLocation sets the officeLocation property value. The office location in the user's place of business. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetOnenote sets the onenote property value. The onenote property
func (m *User) SetOnenote(value Onenoteable)() {
    m.onenote = value
}
// SetOnlineMeetings sets the onlineMeetings property value. The onlineMeetings property
func (m *User) SetOnlineMeetings(value []OnlineMeetingable)() {
    m.onlineMeetings = value
}
// SetOnPremisesDistinguishedName sets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
func (m *User) SetOnPremisesDistinguishedName(value *string)() {
    m.onPremisesDistinguishedName = value
}
// SetOnPremisesDomainName sets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only.
func (m *User) SetOnPremisesDomainName(value *string)() {
    m.onPremisesDomainName = value
}
// SetOnPremisesExtensionAttributes sets the onPremisesExtensionAttributes property value. Contains extensionAttributes1-15 for the user. These extension attributes are also known as Exchange custom attributes 1-15. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. Supports $filter (eq, ne, not, in).
func (m *User) SetOnPremisesExtensionAttributes(value OnPremisesExtensionAttributesable)() {
    m.onPremisesExtensionAttributes = value
}
// SetOnPremisesImmutableId sets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. Note: The $ and _ characters cannot be used when specifying this property. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) SetOnPremisesImmutableId(value *string)() {
    m.onPremisesImmutableId = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: '2013-02-16T03:04:54Z'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesProvisioningErrors sets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning.  Supports $filter (eq, not, ge, le).
func (m *User) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)() {
    m.onPremisesProvisioningErrors = value
}
// SetOnPremisesSamAccountName sets the onPremisesSamAccountName property value. Contains the on-premises sAMAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) SetOnPremisesSamAccountName(value *string)() {
    m.onPremisesSamAccountName = value
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Supports $filter (eq including on null values).
func (m *User) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this user object is currently being synced from an on-premises Active Directory (AD); otherwise the user isn't being synced and can be managed in Azure Active Directory (Azure AD). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetOnPremisesUserPrincipalName sets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) SetOnPremisesUserPrincipalName(value *string)() {
    m.onPremisesUserPrincipalName = value
}
// SetOtherMails sets the otherMails property value. A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com'].NOTE: This property cannot contain accent characters.Supports $filter (eq, not, ge, le, in, startsWith, endsWith, and counting empty collections).
func (m *User) SetOtherMails(value []string)() {
    m.otherMails = value
}
// SetOutlook sets the outlook property value. Selective Outlook services available to the user. Read-only. Nullable.
func (m *User) SetOutlook(value OutlookUserable)() {
    m.outlook = value
}
// SetOwnedDevices sets the ownedDevices property value. Devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) SetOwnedDevices(value []DirectoryObjectable)() {
    m.ownedDevices = value
}
// SetOwnedObjects sets the ownedObjects property value. Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) SetOwnedObjects(value []DirectoryObjectable)() {
    m.ownedObjects = value
}
// SetPasswordPolicies sets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, not, and eq on null values).
func (m *User) SetPasswordPolicies(value *string)() {
    m.passwordPolicies = value
}
// SetPasswordProfile sets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) SetPasswordProfile(value PasswordProfileable)() {
    m.passwordProfile = value
}
// SetPastProjects sets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
func (m *User) SetPastProjects(value []string)() {
    m.pastProjects = value
}
// SetPendingAccessReviewInstances sets the pendingAccessReviewInstances property value. Navigation property to get list of access reviews pending approval by reviewer.
func (m *User) SetPendingAccessReviewInstances(value []AccessReviewInstanceable)() {
    m.pendingAccessReviewInstances = value
}
// SetPeople sets the people property value. Read-only. The most relevant people to the user. The collection is ordered by their relevance to the user, which is determined by the user's communication, collaboration and business relationships. A person is an aggregation of information from across mail, contacts and social networks.
func (m *User) SetPeople(value []Personable)() {
    m.people = value
}
// SetPhoto sets the photo property value. The user's profile photo. Read-only.
func (m *User) SetPhoto(value ProfilePhotoable)() {
    m.photo = value
}
// SetPhotos sets the photos property value. The photos property
func (m *User) SetPhotos(value []ProfilePhotoable)() {
    m.photos = value
}
// SetPlanner sets the planner property value. Selective Planner services available to the user. Read-only. Nullable.
func (m *User) SetPlanner(value PlannerUserable)() {
    m.planner = value
}
// SetPostalCode sets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetPostalCode(value *string)() {
    m.postalCode = value
}
// SetPreferredDataLocation sets the preferredDataLocation property value. The preferred data location for the user. For more information, see OneDrive Online Multi-Geo.
func (m *User) SetPreferredDataLocation(value *string)() {
    m.preferredDataLocation = value
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// SetPreferredName sets the preferredName property value. The preferred name for the user. Not Supported. This attribute returns an empty string.Returned only on $select.
func (m *User) SetPreferredName(value *string)() {
    m.preferredName = value
}
// SetPresence sets the presence property value. The presence property
func (m *User) SetPresence(value Presenceable)() {
    m.presence = value
}
// SetPrint sets the print property value. The print property
func (m *User) SetPrint(value UserPrintable)() {
    m.print = value
}
// SetProfile sets the profile property value. Represents properties that are descriptive of a user in a tenant.
func (m *User) SetProfile(value Profileable)() {
    m.profile = value
}
// SetProvisionedPlans sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *User) SetProvisionedPlans(value []ProvisionedPlanable)() {
    m.provisionedPlans = value
}
// SetProxyAddresses sets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. Changes to the mail property will also update this collection to include the value as an SMTP address. For more information, see mail and proxyAddresses properties. The proxy address prefixed with SMTP (capitalized) is the primary proxy address while those prefixed with smtp are the secondary proxy addresses. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only in Microsoft Graph; you can update this property only through the Microsoft 365 admin center. Not nullable. Supports $filter (eq, not, ge, le, startsWith, endsWith, and counting empty collections).
func (m *User) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// SetRefreshTokensValidFromDateTime sets the refreshTokensValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use invalidateAllRefreshTokens to reset.
func (m *User) SetRefreshTokensValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.refreshTokensValidFromDateTime = value
}
// SetRegisteredDevices sets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *User) SetRegisteredDevices(value []DirectoryObjectable)() {
    m.registeredDevices = value
}
// SetResponsibilities sets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
func (m *User) SetResponsibilities(value []string)() {
    m.responsibilities = value
}
// SetSchools sets the schools property value. A list for the user to enumerate the schools they have attended. Returned only on $select.
func (m *User) SetSchools(value []string)() {
    m.schools = value
}
// SetScopedRoleMemberOf sets the scopedRoleMemberOf property value. The scoped-role administrative unit memberships for this user. Read-only. Nullable.
func (m *User) SetScopedRoleMemberOf(value []ScopedRoleMembershipable)() {
    m.scopedRoleMemberOf = value
}
// SetSecurityIdentifier sets the securityIdentifier property value. Security identifier (SID) of the user, used in Windows scenarios. Read-only. Returned by default. Supports $select and $filter (eq, not, ge, le, startsWith).
func (m *User) SetSecurityIdentifier(value *string)() {
    m.securityIdentifier = value
}
// SetSettings sets the settings property value. The settings property
func (m *User) SetSettings(value UserSettingsable)() {
    m.settings = value
}
// SetShowInAddressList sets the showInAddressList property value. Do not use in Microsoft Graph. Manage this property through the Microsoft 365 admin center instead. Represents whether the user should be included in the Outlook global address list. See Known issue.
func (m *User) SetShowInAddressList(value *bool)() {
    m.showInAddressList = value
}
// SetSignInActivity sets the signInActivity property value. Get the last signed-in date and request ID of the sign-in for a given user. Read-only.Returned only on $select. Supports $filter (eq, ne, not, ge, le) but, not with any other filterable properties. Note: Details for this property require an Azure AD Premium P1/P2 license and the AuditLog.Read.All permission.Note: There's a known issue with retrieving this property.This property is not returned for a user who has never signed in or last signed in before April 2020.
func (m *User) SetSignInActivity(value SignInActivityable)() {
    m.signInActivity = value
}
// SetSignInSessionsValidFromDateTime sets the signInSessionsValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset.
func (m *User) SetSignInSessionsValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.signInSessionsValidFromDateTime = value
}
// SetSkills sets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
func (m *User) SetSkills(value []string)() {
    m.skills = value
}
// SetState sets the state property value. The state or province in the user's address. Maximum length is 128 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetState(value *string)() {
    m.state = value
}
// SetStreetAddress sets the streetAddress property value. The street address of the user's place of business. Maximum length is 1024 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetStreetAddress(value *string)() {
    m.streetAddress = value
}
// SetSurname sets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetSurname(value *string)() {
    m.surname = value
}
// SetTasks sets the tasks property value. The tasks property
func (m *User) SetTasks(value Tasksable)() {
    m.tasks = value
}
// SetTeamwork sets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *User) SetTeamwork(value UserTeamworkable)() {
    m.teamwork = value
}
// SetTodo sets the todo property value. Represents the To Do services available to a user.
func (m *User) SetTodo(value Todoable)() {
    m.todo = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. The groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *User) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetTransitiveReports sets the transitiveReports property value. The transitive reports for a user. Read-only.
func (m *User) SetTransitiveReports(value []DirectoryObjectable)() {
    m.transitiveReports = value
}
// SetUsageLocation sets the usageLocation property value. A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetUsageLocation(value *string)() {
    m.usageLocation = value
}
// SetUsageRights sets the usageRights property value. Represents the usage rights a user has been granted.
func (m *User) SetUsageRights(value []UsageRightable)() {
    m.usageRights = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Only the following characters are allowed A - Z, a - z, 0 - 9, ' . - _ ! # ^ ~. For the complete list of allowed characters, see username policies. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderBy.
func (m *User) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserType sets the userType property value. A String value that can be used to classify user types in your directory, such as Member and Guest. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
func (m *User) SetUserType(value *string)() {
    m.userType = value
}
// SetWindowsInformationProtectionDeviceRegistrations sets the windowsInformationProtectionDeviceRegistrations property value. Zero or more WIP device registrations that belong to the user.
func (m *User) SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistrationable)() {
    m.windowsInformationProtectionDeviceRegistrations = value
}
