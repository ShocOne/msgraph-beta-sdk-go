package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementAssignmentTarget base type for assignment targets.
type DeviceAndAppManagementAssignmentTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Id of the filter for the target assignment.
    deviceAndAppManagementAssignmentFilterId *string
    // Represents type of the assignment filter.
    deviceAndAppManagementAssignmentFilterType *DeviceAndAppManagementAssignmentFilterType
    // The OdataType property
    odataType *string
}
// NewDeviceAndAppManagementAssignmentTarget instantiates a new deviceAndAppManagementAssignmentTarget and sets the default values.
func NewDeviceAndAppManagementAssignmentTarget()(*DeviceAndAppManagementAssignmentTarget) {
    m := &DeviceAndAppManagementAssignmentTarget{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceAndAppManagementAssignmentTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.allDevicesAssignmentTarget":
                        return NewAllDevicesAssignmentTarget(), nil
                    case "#microsoft.graph.allLicensedUsersAssignmentTarget":
                        return NewAllLicensedUsersAssignmentTarget(), nil
                    case "#microsoft.graph.configurationManagerCollectionAssignmentTarget":
                        return NewConfigurationManagerCollectionAssignmentTarget(), nil
                    case "#microsoft.graph.exclusionGroupAssignmentTarget":
                        return NewExclusionGroupAssignmentTarget(), nil
                    case "#microsoft.graph.groupAssignmentTarget":
                        return NewGroupAssignmentTarget(), nil
                }
            }
        }
    }
    return NewDeviceAndAppManagementAssignmentTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementAssignmentTarget) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceAndAppManagementAssignmentFilterId gets the deviceAndAppManagementAssignmentFilterId property value. The Id of the filter for the target assignment.
func (m *DeviceAndAppManagementAssignmentTarget) GetDeviceAndAppManagementAssignmentFilterId()(*string) {
    return m.deviceAndAppManagementAssignmentFilterId
}
// GetDeviceAndAppManagementAssignmentFilterType gets the deviceAndAppManagementAssignmentFilterType property value. Represents type of the assignment filter.
func (m *DeviceAndAppManagementAssignmentTarget) GetDeviceAndAppManagementAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    return m.deviceAndAppManagementAssignmentFilterType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceAndAppManagementAssignmentFilterId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceAndAppManagementAssignmentFilterId)
    res["deviceAndAppManagementAssignmentFilterType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType , m.SetDeviceAndAppManagementAssignmentFilterType)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceAndAppManagementAssignmentTarget) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceAndAppManagementAssignmentFilterId", m.GetDeviceAndAppManagementAssignmentFilterId())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceAndAppManagementAssignmentFilterType() != nil {
        cast := (*m.GetDeviceAndAppManagementAssignmentFilterType()).String()
        err := writer.WriteStringValue("deviceAndAppManagementAssignmentFilterType", &cast)
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
func (m *DeviceAndAppManagementAssignmentTarget) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceAndAppManagementAssignmentFilterId sets the deviceAndAppManagementAssignmentFilterId property value. The Id of the filter for the target assignment.
func (m *DeviceAndAppManagementAssignmentTarget) SetDeviceAndAppManagementAssignmentFilterId(value *string)() {
    m.deviceAndAppManagementAssignmentFilterId = value
}
// SetDeviceAndAppManagementAssignmentFilterType sets the deviceAndAppManagementAssignmentFilterType property value. Represents type of the assignment filter.
func (m *DeviceAndAppManagementAssignmentTarget) SetDeviceAndAppManagementAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.deviceAndAppManagementAssignmentFilterType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceAndAppManagementAssignmentTarget) SetOdataType(value *string)() {
    m.odataType = value
}
