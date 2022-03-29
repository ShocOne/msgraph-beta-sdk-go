package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecordOperationable 
type RecordOperationable interface {
    CommsOperationable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompletionReason()(*RecordCompletionReason)
    GetRecordingAccessToken()(*string)
    GetRecordingLocation()(*string)
    SetCompletionReason(value *RecordCompletionReason)()
    SetRecordingAccessToken(value *string)()
    SetRecordingLocation(value *string)()
}