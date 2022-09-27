package callrecords

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TraceRouteHop 
type TraceRouteHop struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The network path count of this hop that was used to compute the round-trip time.
    hopCount *int32
    // IP address used for this hop in the network trace.
    ipAddress *string
    // The OdataType property
    odataType *string
    // The time from when the trace route packet was sent from the client to this hop and back to the client, denoted in [ISO 8601][] format. For example, 1 second is denoted as PT1S, where P is the duration designator, T is the time designator, and S is the second designator.
    roundTripTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewTraceRouteHop instantiates a new traceRouteHop and sets the default values.
func NewTraceRouteHop()(*TraceRouteHop) {
    m := &TraceRouteHop{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.callRecords.traceRouteHop";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTraceRouteHopFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTraceRouteHopFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTraceRouteHop(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TraceRouteHop) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TraceRouteHop) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hopCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetHopCount)
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["roundTripTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetRoundTripTime)
    return res
}
// GetHopCount gets the hopCount property value. The network path count of this hop that was used to compute the round-trip time.
func (m *TraceRouteHop) GetHopCount()(*int32) {
    return m.hopCount
}
// GetIpAddress gets the ipAddress property value. IP address used for this hop in the network trace.
func (m *TraceRouteHop) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TraceRouteHop) GetOdataType()(*string) {
    return m.odataType
}
// GetRoundTripTime gets the roundTripTime property value. The time from when the trace route packet was sent from the client to this hop and back to the client, denoted in [ISO 8601][] format. For example, 1 second is denoted as PT1S, where P is the duration designator, T is the time designator, and S is the second designator.
func (m *TraceRouteHop) GetRoundTripTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.roundTripTime
}
// Serialize serializes information the current object
func (m *TraceRouteHop) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("hopCount", m.GetHopCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
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
        err := writer.WriteISODurationValue("roundTripTime", m.GetRoundTripTime())
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
func (m *TraceRouteHop) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHopCount sets the hopCount property value. The network path count of this hop that was used to compute the round-trip time.
func (m *TraceRouteHop) SetHopCount(value *int32)() {
    m.hopCount = value
}
// SetIpAddress sets the ipAddress property value. IP address used for this hop in the network trace.
func (m *TraceRouteHop) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TraceRouteHop) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoundTripTime sets the roundTripTime property value. The time from when the trace route packet was sent from the client to this hop and back to the client, denoted in [ISO 8601][] format. For example, 1 second is denoted as PT1S, where P is the duration designator, T is the time designator, and S is the second designator.
func (m *TraceRouteHop) SetRoundTripTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.roundTripTime = value
}
