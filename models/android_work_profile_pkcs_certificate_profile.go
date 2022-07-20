package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfilePkcsCertificateProfile 
type AndroidWorkProfilePkcsCertificateProfile struct {
    AndroidWorkProfileCertificateProfileBase
    // Target store certificate. Possible values are: user, machine.
    certificateStore *CertificateStore
    // PKCS Certificate Template Name
    certificateTemplateName *string
    // PKCS Certification Authority
    certificationAuthority *string
    // PKCS Certification Authority Name
    certificationAuthorityName *string
    // Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
    customSubjectAlternativeNames []CustomSubjectAlternativeNameable
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Custom String that defines the AAD Attribute.
    subjectAlternativeNameFormatString *string
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewAndroidWorkProfilePkcsCertificateProfile instantiates a new AndroidWorkProfilePkcsCertificateProfile and sets the default values.
func NewAndroidWorkProfilePkcsCertificateProfile()(*AndroidWorkProfilePkcsCertificateProfile) {
    m := &AndroidWorkProfilePkcsCertificateProfile{
        AndroidWorkProfileCertificateProfileBase: *NewAndroidWorkProfileCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfilePkcsCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidWorkProfilePkcsCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfilePkcsCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfilePkcsCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *AndroidWorkProfilePkcsCertificateProfile) GetCertificateStore()(*CertificateStore) {
    if m == nil {
        return nil
    } else {
        return m.certificateStore
    }
}
// GetCertificateTemplateName gets the certificateTemplateName property value. PKCS Certificate Template Name
func (m *AndroidWorkProfilePkcsCertificateProfile) GetCertificateTemplateName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateTemplateName
    }
}
// GetCertificationAuthority gets the certificationAuthority property value. PKCS Certification Authority
func (m *AndroidWorkProfilePkcsCertificateProfile) GetCertificationAuthority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationAuthority
    }
}
// GetCertificationAuthorityName gets the certificationAuthorityName property value. PKCS Certification Authority Name
func (m *AndroidWorkProfilePkcsCertificateProfile) GetCertificationAuthorityName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationAuthorityName
    }
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfilePkcsCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    if m == nil {
        return nil
    } else {
        return m.customSubjectAlternativeNames
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfilePkcsCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidWorkProfileCertificateProfileBase.GetFieldDeserializers()
    res["certificateStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateStore(val.(*CertificateStore))
        }
        return nil
    }
    res["certificateTemplateName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateTemplateName(val)
        }
        return nil
    }
    res["certificationAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthority(val)
        }
        return nil
    }
    res["certificationAuthorityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityName(val)
        }
        return nil
    }
    res["customSubjectAlternativeNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomSubjectAlternativeNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomSubjectAlternativeNameable, len(val))
            for i, v := range val {
                res[i] = v.(CustomSubjectAlternativeNameable)
            }
            m.SetCustomSubjectAlternativeNames(res)
        }
        return nil
    }
    res["managedDeviceCertificateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceCertificateStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceCertificateStateable)
            }
            m.SetManagedDeviceCertificateStates(res)
        }
        return nil
    }
    res["subjectAlternativeNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectAlternativeNameFormatString(val)
        }
        return nil
    }
    res["subjectNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormatString(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidWorkProfilePkcsCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCertificateStates
    }
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AndroidWorkProfilePkcsCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectAlternativeNameFormatString
    }
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *AndroidWorkProfilePkcsCertificateProfile) GetSubjectNameFormatString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectNameFormatString
    }
}
// Serialize serializes information the current object
func (m *AndroidWorkProfilePkcsCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidWorkProfileCertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateStore() != nil {
        cast := (*m.GetCertificateStore()).String()
        err = writer.WriteStringValue("certificateStore", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateTemplateName", m.GetCertificateTemplateName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthority", m.GetCertificationAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthorityName", m.GetCertificationAuthorityName())
        if err != nil {
            return err
        }
    }
    if m.GetCustomSubjectAlternativeNames() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomSubjectAlternativeNames()))
        for i, v := range m.GetCustomSubjectAlternativeNames() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customSubjectAlternativeNames", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCertificateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceCertificateStates()))
        for i, v := range m.GetManagedDeviceCertificateStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectAlternativeNameFormatString", m.GetSubjectAlternativeNameFormatString())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectNameFormatString", m.GetSubjectNameFormatString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateStore sets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *AndroidWorkProfilePkcsCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    if m != nil {
        m.certificateStore = value
    }
}
// SetCertificateTemplateName sets the certificateTemplateName property value. PKCS Certificate Template Name
func (m *AndroidWorkProfilePkcsCertificateProfile) SetCertificateTemplateName(value *string)() {
    if m != nil {
        m.certificateTemplateName = value
    }
}
// SetCertificationAuthority sets the certificationAuthority property value. PKCS Certification Authority
func (m *AndroidWorkProfilePkcsCertificateProfile) SetCertificationAuthority(value *string)() {
    if m != nil {
        m.certificationAuthority = value
    }
}
// SetCertificationAuthorityName sets the certificationAuthorityName property value. PKCS Certification Authority Name
func (m *AndroidWorkProfilePkcsCertificateProfile) SetCertificationAuthorityName(value *string)() {
    if m != nil {
        m.certificationAuthorityName = value
    }
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AndroidWorkProfilePkcsCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    if m != nil {
        m.customSubjectAlternativeNames = value
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidWorkProfilePkcsCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    if m != nil {
        m.managedDeviceCertificateStates = value
    }
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AndroidWorkProfilePkcsCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    if m != nil {
        m.subjectAlternativeNameFormatString = value
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *AndroidWorkProfilePkcsCertificateProfile) SetSubjectNameFormatString(value *string)() {
    if m != nil {
        m.subjectNameFormatString = value
    }
}
