package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable instead.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponse struct {
    TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponse
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponse instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponse and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponse()(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponse) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponse{
        TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponse: *NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponse(),
    }
    return m
}
// CreateTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable instead.
type TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable
}
