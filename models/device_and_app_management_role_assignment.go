package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementRoleAssignment 
type DeviceAndAppManagementRoleAssignment struct {
    RoleAssignment
    // The list of ids of role member security groups. These are IDs from Azure Active Directory.
    members []string
    // The set of Role Scope Tags defined on the Role Assignment.
    roleScopeTags []RoleScopeTagable
}
// NewDeviceAndAppManagementRoleAssignment instantiates a new DeviceAndAppManagementRoleAssignment and sets the default values.
func NewDeviceAndAppManagementRoleAssignment()(*DeviceAndAppManagementRoleAssignment) {
    m := &DeviceAndAppManagementRoleAssignment{
        RoleAssignment: *NewRoleAssignment(),
    }
    odataTypeValue := "#microsoft.graph.deviceAndAppManagementRoleAssignment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAndAppManagementRoleAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RoleAssignment.GetFieldDeserializers()
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["roleScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleScopeTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleScopeTagable, len(val))
            for i, v := range val {
                res[i] = v.(RoleScopeTagable)
            }
            m.SetRoleScopeTags(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. The list of ids of role member security groups. These are IDs from Azure Active Directory.
func (m *DeviceAndAppManagementRoleAssignment) GetMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetRoleScopeTags gets the roleScopeTags property value. The set of Role Scope Tags defined on the Role Assignment.
func (m *DeviceAndAppManagementRoleAssignment) GetRoleScopeTags()([]RoleScopeTagable) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTags
    }
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RoleAssignment.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        err = writer.WriteCollectionOfStringValues("members", m.GetMembers())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleScopeTags()))
        for i, v := range m.GetRoleScopeTags() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("roleScopeTags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMembers sets the members property value. The list of ids of role member security groups. These are IDs from Azure Active Directory.
func (m *DeviceAndAppManagementRoleAssignment) SetMembers(value []string)() {
    if m != nil {
        m.members = value
    }
}
// SetRoleScopeTags sets the roleScopeTags property value. The set of Role Scope Tags defined on the Role Assignment.
func (m *DeviceAndAppManagementRoleAssignment) SetRoleScopeTags(value []RoleScopeTagable)() {
    if m != nil {
        m.roleScopeTags = value
    }
}
