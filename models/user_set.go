package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSet 
type UserSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
    isBackup *bool
    // The OdataType property
    odataType *string
}
// NewUserSet instantiates a new userSet and sets the default values.
func NewUserSet()(*UserSet) {
    m := &UserSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userSet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.connectedOrganizationMembers":
                        return NewConnectedOrganizationMembers(), nil
                    case "#microsoft.graph.externalSponsors":
                        return NewExternalSponsors(), nil
                    case "#microsoft.graph.groupMembers":
                        return NewGroupMembers(), nil
                    case "#microsoft.graph.internalSponsors":
                        return NewInternalSponsors(), nil
                    case "#microsoft.graph.requestorManager":
                        return NewRequestorManager(), nil
                    case "#microsoft.graph.singleUser":
                        return NewSingleUser(), nil
                }
            }
        }
    }
    return NewUserSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSet) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isBackup"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBackup)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIsBackup gets the isBackup property value. For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
func (m *UserSet) GetIsBackup()(*bool) {
    return m.isBackup
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserSet) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *UserSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isBackup", m.GetIsBackup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *UserSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsBackup sets the isBackup property value. For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
func (m *UserSet) SetIsBackup(value *bool)() {
    m.isBackup = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserSet) SetOdataType(value *string)() {
    m.odataType = value
}
