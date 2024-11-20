package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PublicKeyInfrastructureRoot struct {
    Entity
}
// NewPublicKeyInfrastructureRoot instantiates a new PublicKeyInfrastructureRoot and sets the default values.
func NewPublicKeyInfrastructureRoot()(*PublicKeyInfrastructureRoot) {
    m := &PublicKeyInfrastructureRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePublicKeyInfrastructureRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePublicKeyInfrastructureRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicKeyInfrastructureRoot(), nil
}
// GetCertificateBasedAuthConfigurations gets the certificateBasedAuthConfigurations property value. The collection of public key infrastructure instances for the certificate-based authentication feature for users.
// returns a []CertificateBasedAuthPkiable when successful
func (m *PublicKeyInfrastructureRoot) GetCertificateBasedAuthConfigurations()([]CertificateBasedAuthPkiable) {
    val, err := m.GetBackingStore().Get("certificateBasedAuthConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CertificateBasedAuthPkiable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PublicKeyInfrastructureRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateBasedAuthConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateBasedAuthPkiFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateBasedAuthPkiable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CertificateBasedAuthPkiable)
                }
            }
            m.SetCertificateBasedAuthConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PublicKeyInfrastructureRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateBasedAuthConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateBasedAuthConfigurations()))
        for i, v := range m.GetCertificateBasedAuthConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("certificateBasedAuthConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateBasedAuthConfigurations sets the certificateBasedAuthConfigurations property value. The collection of public key infrastructure instances for the certificate-based authentication feature for users.
func (m *PublicKeyInfrastructureRoot) SetCertificateBasedAuthConfigurations(value []CertificateBasedAuthPkiable)() {
    err := m.GetBackingStore().Set("certificateBasedAuthConfigurations", value)
    if err != nil {
        panic(err)
    }
}
type PublicKeyInfrastructureRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateBasedAuthConfigurations()([]CertificateBasedAuthPkiable)
    SetCertificateBasedAuthConfigurations(value []CertificateBasedAuthPkiable)()
}
