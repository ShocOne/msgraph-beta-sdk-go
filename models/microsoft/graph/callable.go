package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Callable 
type Callable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActiveModalities()([]Modality)
    GetAnsweredBy()(ParticipantInfoable)
    GetAudioRoutingGroups()([]AudioRoutingGroupable)
    GetCallbackUri()(*string)
    GetCallChainId()(*string)
    GetCallOptions()(CallOptionsable)
    GetCallRoutes()([]CallRouteable)
    GetChatInfo()(ChatInfoable)
    GetContentSharingSessions()([]ContentSharingSessionable)
    GetDirection()(*CallDirection)
    GetIncomingContext()(IncomingContextable)
    GetMediaConfig()(MediaConfigable)
    GetMediaState()(CallMediaStateable)
    GetMeetingCapability()(MeetingCapabilityable)
    GetMeetingInfo()(MeetingInfoable)
    GetMyParticipantId()(*string)
    GetOperations()([]CommsOperationable)
    GetParticipants()([]Participantable)
    GetRequestedModalities()([]Modality)
    GetResultInfo()(ResultInfoable)
    GetRingingTimeoutInSeconds()(*int32)
    GetRoutingPolicies()([]RoutingPolicy)
    GetSource()(ParticipantInfoable)
    GetState()(*CallState)
    GetSubject()(*string)
    GetTargets()([]InvitationParticipantInfoable)
    GetTenantId()(*string)
    GetTerminationReason()(*string)
    GetToneInfo()(ToneInfoable)
    GetTranscription()(CallTranscriptionInfoable)
    SetActiveModalities(value []Modality)()
    SetAnsweredBy(value ParticipantInfoable)()
    SetAudioRoutingGroups(value []AudioRoutingGroupable)()
    SetCallbackUri(value *string)()
    SetCallChainId(value *string)()
    SetCallOptions(value CallOptionsable)()
    SetCallRoutes(value []CallRouteable)()
    SetChatInfo(value ChatInfoable)()
    SetContentSharingSessions(value []ContentSharingSessionable)()
    SetDirection(value *CallDirection)()
    SetIncomingContext(value IncomingContextable)()
    SetMediaConfig(value MediaConfigable)()
    SetMediaState(value CallMediaStateable)()
    SetMeetingCapability(value MeetingCapabilityable)()
    SetMeetingInfo(value MeetingInfoable)()
    SetMyParticipantId(value *string)()
    SetOperations(value []CommsOperationable)()
    SetParticipants(value []Participantable)()
    SetRequestedModalities(value []Modality)()
    SetResultInfo(value ResultInfoable)()
    SetRingingTimeoutInSeconds(value *int32)()
    SetRoutingPolicies(value []RoutingPolicy)()
    SetSource(value ParticipantInfoable)()
    SetState(value *CallState)()
    SetSubject(value *string)()
    SetTargets(value []InvitationParticipantInfoable)()
    SetTenantId(value *string)()
    SetTerminationReason(value *string)()
    SetToneInfo(value ToneInfoable)()
    SetTranscription(value CallTranscriptionInfoable)()
}
