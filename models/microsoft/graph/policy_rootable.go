package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PolicyRootable 
type PolicyRootable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessReviewPolicy()(AccessReviewPolicyable)
    GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable)
    GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable)
    GetAppManagementPolicies()([]AppManagementPolicyable)
    GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable)
    GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable)
    GetAuthorizationPolicy()([]AuthorizationPolicyable)
    GetB2cAuthenticationMethodsPolicy()(B2cAuthenticationMethodsPolicyable)
    GetClaimsMappingPolicies()([]ClaimsMappingPolicyable)
    GetConditionalAccessPolicies()([]ConditionalAccessPolicyable)
    GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable)
    GetDefaultAppManagementPolicy()(TenantAppManagementPolicyable)
    GetDirectoryRoleAccessReviewPolicy()(DirectoryRoleAccessReviewPolicyable)
    GetExternalIdentitiesPolicy()(ExternalIdentitiesPolicyable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable)
    GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable)
    GetMobileAppManagementPolicies()([]MobilityManagementPolicyable)
    GetMobileDeviceManagementPolicies()([]MobilityManagementPolicyable)
    GetPermissionGrantPolicies()([]PermissionGrantPolicyable)
    GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable)
    GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable)
    GetServicePrincipalCreationPolicies()([]ServicePrincipalCreationPolicyable)
    GetTokenIssuancePolicies()([]TokenIssuancePolicyable)
    GetTokenLifetimePolicies()([]TokenLifetimePolicyable)
    SetAccessReviewPolicy(value AccessReviewPolicyable)()
    SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)()
    SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)()
    SetAppManagementPolicies(value []AppManagementPolicyable)()
    SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)()
    SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)()
    SetAuthorizationPolicy(value []AuthorizationPolicyable)()
    SetB2cAuthenticationMethodsPolicy(value B2cAuthenticationMethodsPolicyable)()
    SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)()
    SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)()
    SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)()
    SetDefaultAppManagementPolicy(value TenantAppManagementPolicyable)()
    SetDirectoryRoleAccessReviewPolicy(value DirectoryRoleAccessReviewPolicyable)()
    SetExternalIdentitiesPolicy(value ExternalIdentitiesPolicyable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)()
    SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)()
    SetMobileAppManagementPolicies(value []MobilityManagementPolicyable)()
    SetMobileDeviceManagementPolicies(value []MobilityManagementPolicyable)()
    SetPermissionGrantPolicies(value []PermissionGrantPolicyable)()
    SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)()
    SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)()
    SetServicePrincipalCreationPolicies(value []ServicePrincipalCreationPolicyable)()
    SetTokenIssuancePolicies(value []TokenIssuancePolicyable)()
    SetTokenLifetimePolicies(value []TokenLifetimePolicyable)()
}