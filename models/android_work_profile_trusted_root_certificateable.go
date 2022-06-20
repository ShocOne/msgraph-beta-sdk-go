package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileTrustedRootCertificateable 
type AndroidWorkProfileTrustedRootCertificateable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertFileName()(*string)
    GetTrustedRootCertificate()([]byte)
    SetCertFileName(value *string)()
    SetTrustedRootCertificate(value []byte)()
}
