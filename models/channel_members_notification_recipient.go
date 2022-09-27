package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelMembersNotificationRecipient 
type ChannelMembersNotificationRecipient struct {
    TeamworkNotificationRecipient
    // The channel's identifier.
    channelId *string
    // The team's identifier under which the channel resides.
    teamId *string
}
// NewChannelMembersNotificationRecipient instantiates a new ChannelMembersNotificationRecipient and sets the default values.
func NewChannelMembersNotificationRecipient()(*ChannelMembersNotificationRecipient) {
    m := &ChannelMembersNotificationRecipient{
        TeamworkNotificationRecipient: *NewTeamworkNotificationRecipient(),
    }
    odataTypeValue := "#microsoft.graph.channelMembersNotificationRecipient";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateChannelMembersNotificationRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelMembersNotificationRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelMembersNotificationRecipient(), nil
}
// GetChannelId gets the channelId property value. The channel's identifier.
func (m *ChannelMembersNotificationRecipient) GetChannelId()(*string) {
    return m.channelId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelMembersNotificationRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamworkNotificationRecipient.GetFieldDeserializers()
    res["channelId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetChannelId)
    res["teamId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTeamId)
    return res
}
// GetTeamId gets the teamId property value. The team's identifier under which the channel resides.
func (m *ChannelMembersNotificationRecipient) GetTeamId()(*string) {
    return m.teamId
}
// Serialize serializes information the current object
func (m *ChannelMembersNotificationRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamworkNotificationRecipient.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("channelId", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChannelId sets the channelId property value. The channel's identifier.
func (m *ChannelMembersNotificationRecipient) SetChannelId(value *string)() {
    m.channelId = value
}
// SetTeamId sets the teamId property value. The team's identifier under which the channel resides.
func (m *ChannelMembersNotificationRecipient) SetTeamId(value *string)() {
    m.teamId = value
}
