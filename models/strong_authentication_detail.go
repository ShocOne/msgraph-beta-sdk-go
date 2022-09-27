package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StrongAuthenticationDetail 
type StrongAuthenticationDetail struct {
    Entity
    // The encryptedPinHashHistory property
    encryptedPinHashHistory []byte
    // The proofupTime property
    proofupTime *int64
}
// NewStrongAuthenticationDetail instantiates a new StrongAuthenticationDetail and sets the default values.
func NewStrongAuthenticationDetail()(*StrongAuthenticationDetail) {
    m := &StrongAuthenticationDetail{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.strongAuthenticationDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateStrongAuthenticationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStrongAuthenticationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStrongAuthenticationDetail(), nil
}
// GetEncryptedPinHashHistory gets the encryptedPinHashHistory property value. The encryptedPinHashHistory property
func (m *StrongAuthenticationDetail) GetEncryptedPinHashHistory()([]byte) {
    return m.encryptedPinHashHistory
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StrongAuthenticationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["encryptedPinHashHistory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetEncryptedPinHashHistory)
    res["proofupTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetProofupTime)
    return res
}
// GetProofupTime gets the proofupTime property value. The proofupTime property
func (m *StrongAuthenticationDetail) GetProofupTime()(*int64) {
    return m.proofupTime
}
// Serialize serializes information the current object
func (m *StrongAuthenticationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("encryptedPinHashHistory", m.GetEncryptedPinHashHistory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("proofupTime", m.GetProofupTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEncryptedPinHashHistory sets the encryptedPinHashHistory property value. The encryptedPinHashHistory property
func (m *StrongAuthenticationDetail) SetEncryptedPinHashHistory(value []byte)() {
    m.encryptedPinHashHistory = value
}
// SetProofupTime sets the proofupTime property value. The proofupTime property
func (m *StrongAuthenticationDetail) SetProofupTime(value *int64)() {
    m.proofupTime = value
}
