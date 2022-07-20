package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerEnterpriseWiFiConfiguration 
type AndroidDeviceOwnerEnterpriseWiFiConfiguration struct {
    AndroidDeviceOwnerWiFiConfiguration
    // Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
    authenticationMethod *WiFiAuthenticationMethod
    // Tenant level settings for the Derived Credentials to be used for authentication.
    derivedCredentialSettings DeviceManagementDerivedCredentialSettingsable
    // Extensible Authentication Protocol (EAP) Configuration Types.
    eapType *AndroidEapType
    // Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
    identityCertificateForClientAuthentication AndroidDeviceOwnerCertificateProfileBaseable
    // Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
    innerAuthenticationProtocolForEapTtls *NonEapAuthenticationMethodForEapTtlsType
    // Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
    innerAuthenticationProtocolForPeap *NonEapAuthenticationMethodForPeap
    // Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
    outerIdentityPrivacyTemporaryValue *string
    // Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
    rootCertificateForServerValidation AndroidDeviceOwnerTrustedRootCertificateable
    // Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
    trustedServerCertificateNames []string
}
// NewAndroidDeviceOwnerEnterpriseWiFiConfiguration instantiates a new AndroidDeviceOwnerEnterpriseWiFiConfiguration and sets the default values.
func NewAndroidDeviceOwnerEnterpriseWiFiConfiguration()(*AndroidDeviceOwnerEnterpriseWiFiConfiguration) {
    m := &AndroidDeviceOwnerEnterpriseWiFiConfiguration{
        AndroidDeviceOwnerWiFiConfiguration: *NewAndroidDeviceOwnerWiFiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerEnterpriseWiFiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerEnterpriseWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerEnterpriseWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerEnterpriseWiFiConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentialSettings
    }
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetEapType()(*AndroidEapType) {
    if m == nil {
        return nil
    } else {
        return m.eapType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerWiFiConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*WiFiAuthenticationMethod))
        }
        return nil
    }
    res["derivedCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialSettings(val.(DeviceManagementDerivedCredentialSettingsable))
        }
        return nil
    }
    res["eapType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidEapType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapType(val.(*AndroidEapType))
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(AndroidDeviceOwnerCertificateProfileBaseable))
        }
        return nil
    }
    res["innerAuthenticationProtocolForEapTtls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForEapTtlsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForEapTtls(val.(*NonEapAuthenticationMethodForEapTtlsType))
        }
        return nil
    }
    res["innerAuthenticationProtocolForPeap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForPeap)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForPeap(val.(*NonEapAuthenticationMethodForPeap))
        }
        return nil
    }
    res["outerIdentityPrivacyTemporaryValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuterIdentityPrivacyTemporaryValue(val)
        }
        return nil
    }
    res["rootCertificateForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForServerValidation(val.(AndroidDeviceOwnerTrustedRootCertificateable))
        }
        return nil
    }
    res["trustedServerCertificateNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTrustedServerCertificateNames(res)
        }
        return nil
    }
    return res
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetIdentityCertificateForClientAuthentication()(AndroidDeviceOwnerCertificateProfileBaseable) {
    if m == nil {
        return nil
    } else {
        return m.identityCertificateForClientAuthentication
    }
}
// GetInnerAuthenticationProtocolForEapTtls gets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType) {
    if m == nil {
        return nil
    } else {
        return m.innerAuthenticationProtocolForEapTtls
    }
}
// GetInnerAuthenticationProtocolForPeap gets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap) {
    if m == nil {
        return nil
    } else {
        return m.innerAuthenticationProtocolForPeap
    }
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outerIdentityPrivacyTemporaryValue
    }
}
// GetRootCertificateForServerValidation gets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetRootCertificateForServerValidation()(AndroidDeviceOwnerTrustedRootCertificateable) {
    if m == nil {
        return nil
    } else {
        return m.rootCertificateForServerValidation
    }
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) GetTrustedServerCertificateNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.trustedServerCertificateNames
    }
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerWiFiConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("derivedCredentialSettings", m.GetDerivedCredentialSettings())
        if err != nil {
            return err
        }
    }
    if m.GetEapType() != nil {
        cast := (*m.GetEapType()).String()
        err = writer.WriteStringValue("eapType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificateForClientAuthentication", m.GetIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    if m.GetInnerAuthenticationProtocolForEapTtls() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForEapTtls()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForEapTtls", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInnerAuthenticationProtocolForPeap() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForPeap()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForPeap", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outerIdentityPrivacyTemporaryValue", m.GetOuterIdentityPrivacyTemporaryValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificateForServerValidation", m.GetRootCertificateForServerValidation())
        if err != nil {
            return err
        }
    }
    if m.GetTrustedServerCertificateNames() != nil {
        err = writer.WriteCollectionOfStringValues("trustedServerCertificateNames", m.GetTrustedServerCertificateNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    if m != nil {
        m.authenticationMethod = value
    }
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    if m != nil {
        m.derivedCredentialSettings = value
    }
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetEapType(value *AndroidEapType)() {
    if m != nil {
        m.eapType = value
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetIdentityCertificateForClientAuthentication(value AndroidDeviceOwnerCertificateProfileBaseable)() {
    if m != nil {
        m.identityCertificateForClientAuthentication = value
    }
}
// SetInnerAuthenticationProtocolForEapTtls sets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)() {
    if m != nil {
        m.innerAuthenticationProtocolForEapTtls = value
    }
}
// SetInnerAuthenticationProtocolForPeap sets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)() {
    if m != nil {
        m.innerAuthenticationProtocolForPeap = value
    }
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    if m != nil {
        m.outerIdentityPrivacyTemporaryValue = value
    }
}
// SetRootCertificateForServerValidation sets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetRootCertificateForServerValidation(value AndroidDeviceOwnerTrustedRootCertificateable)() {
    if m != nil {
        m.rootCertificateForServerValidation = value
    }
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *AndroidDeviceOwnerEnterpriseWiFiConfiguration) SetTrustedServerCertificateNames(value []string)() {
    if m != nil {
        m.trustedServerCertificateNames = value
    }
}
