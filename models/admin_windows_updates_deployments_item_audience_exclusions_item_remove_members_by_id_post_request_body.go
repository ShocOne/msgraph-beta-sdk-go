package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody provides operations to call the removeMembersById method.
type AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ids property
    ids []string
    // The memberEntityType property
    memberEntityType *string
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody instantiates a new AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody()(*AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) {
    m := &AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    res["memberEntityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEntityType(val)
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) GetMemberEntityType()(*string) {
    return m.memberEntityType
}
// Serialize serializes information the current object
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *AdminWindowsUpdatesDeploymentsItemAudienceExclusionsItemRemoveMembersByIdPostRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
