package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWifiEnterpriseEAPConfigurationable 
type WindowsWifiEnterpriseEAPConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsWifiConfigurationable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetAuthenticationPeriodInSeconds()(*int32)
    GetAuthenticationRetryDelayPeriodInSeconds()(*int32)
    GetAuthenticationType()(*WifiAuthenticationType)
    GetCacheCredentials()(*bool)
    GetDisableUserPromptForServerValidation()(*bool)
    GetEapolStartPeriodInSeconds()(*int32)
    GetEapType()(*EapType)
    GetEnablePairwiseMasterKeyCaching()(*bool)
    GetEnablePreAuthentication()(*bool)
    GetIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEAPTTLS()(*NonEapAuthenticationMethodForEapTtlsType)
    GetMaximumAuthenticationFailures()(*int32)
    GetMaximumAuthenticationTimeoutInSeconds()(*int32)
    GetMaximumEAPOLStartMessages()(*int32)
    GetMaximumNumberOfPairwiseMasterKeysInCache()(*int32)
    GetMaximumPairwiseMasterKeyCacheTimeInMinutes()(*int32)
    GetMaximumPreAuthenticationAttempts()(*int32)
    GetNetworkSingleSignOn()(*NetworkSingleSignOnType)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetPerformServerValidation()(*bool)
    GetPromptForAdditionalAuthenticationCredentials()(*bool)
    GetRequireCryptographicBinding()(*bool)
    GetRootCertificateForClientValidation()(Windows81TrustedRootCertificateable)
    GetRootCertificatesForServerValidation()([]Windows81TrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    GetUserBasedVirtualLan()(*bool)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetAuthenticationPeriodInSeconds(value *int32)()
    SetAuthenticationRetryDelayPeriodInSeconds(value *int32)()
    SetAuthenticationType(value *WifiAuthenticationType)()
    SetCacheCredentials(value *bool)()
    SetDisableUserPromptForServerValidation(value *bool)()
    SetEapolStartPeriodInSeconds(value *int32)()
    SetEapType(value *EapType)()
    SetEnablePairwiseMasterKeyCaching(value *bool)()
    SetEnablePreAuthentication(value *bool)()
    SetIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEAPTTLS(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetMaximumAuthenticationFailures(value *int32)()
    SetMaximumAuthenticationTimeoutInSeconds(value *int32)()
    SetMaximumEAPOLStartMessages(value *int32)()
    SetMaximumNumberOfPairwiseMasterKeysInCache(value *int32)()
    SetMaximumPairwiseMasterKeyCacheTimeInMinutes(value *int32)()
    SetMaximumPreAuthenticationAttempts(value *int32)()
    SetNetworkSingleSignOn(value *NetworkSingleSignOnType)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetPerformServerValidation(value *bool)()
    SetPromptForAdditionalAuthenticationCredentials(value *bool)()
    SetRequireCryptographicBinding(value *bool)()
    SetRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)()
    SetRootCertificatesForServerValidation(value []Windows81TrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
    SetUserBasedVirtualLan(value *bool)()
}