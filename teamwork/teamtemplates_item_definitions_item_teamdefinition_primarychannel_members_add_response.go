package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponseable instead.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponse struct {
    TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponse
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponse instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponse and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponse()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponse) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponse{
        TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponse: *NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponse(),
    }
    return m
}
// CreateTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponseable instead.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersAddPostResponseable
}
