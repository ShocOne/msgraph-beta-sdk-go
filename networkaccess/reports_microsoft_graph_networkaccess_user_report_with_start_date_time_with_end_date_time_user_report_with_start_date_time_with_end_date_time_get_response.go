package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse 
type ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse instantiates a new ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse()(*ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse) {
    m := &ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Userable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse) GetValue()([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Userable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Userable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponse) SetValue(value []i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Userable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponseable 
type ReportsMicrosoftGraphNetworkaccessUserReportWithStartDateTimeWithEndDateTimeUserReportWithStartDateTimeWithEndDateTimeGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Userable)
    SetValue(value []i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Userable)()
}
