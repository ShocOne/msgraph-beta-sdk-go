package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivityStatistics provides operations to manage the collection of accessReview entities.
type ActivityStatistics struct {
    Entity
    // The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
    activity *AnalyticsActivityType
    // Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
    duration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
    startDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
    timeZoneUsed *string
}
// NewActivityStatistics instantiates a new activityStatistics and sets the default values.
func NewActivityStatistics()(*ActivityStatistics) {
    m := &ActivityStatistics{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.activityStatistics";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateActivityStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivityStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.callActivityStatistics":
                        return NewCallActivityStatistics(), nil
                    case "#microsoft.graph.chatActivityStatistics":
                        return NewChatActivityStatistics(), nil
                    case "#microsoft.graph.emailActivityStatistics":
                        return NewEmailActivityStatistics(), nil
                    case "#microsoft.graph.focusActivityStatistics":
                        return NewFocusActivityStatistics(), nil
                    case "#microsoft.graph.meetingActivityStatistics":
                        return NewMeetingActivityStatistics(), nil
                }
            }
        }
    }
    return NewActivityStatistics(), nil
}
// GetActivity gets the activity property value. The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
func (m *ActivityStatistics) GetActivity()(*AnalyticsActivityType) {
    return m.activity
}
// GetDuration gets the duration property value. Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
func (m *ActivityStatistics) GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.duration
}
// GetEndDate gets the endDate property value. Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivityStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAnalyticsActivityType , m.SetActivity)
    res["duration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetDuration)
    res["endDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetEndDate)
    res["startDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetStartDate)
    res["timeZoneUsed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTimeZoneUsed)
    return res
}
// GetStartDate gets the startDate property value. Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.startDate
}
// GetTimeZoneUsed gets the timeZoneUsed property value. The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
func (m *ActivityStatistics) GetTimeZoneUsed()(*string) {
    return m.timeZoneUsed
}
// Serialize serializes information the current object
func (m *ActivityStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivity() != nil {
        cast := (*m.GetActivity()).String()
        err = writer.WriteStringValue("activity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZoneUsed", m.GetTimeZoneUsed())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
func (m *ActivityStatistics) SetActivity(value *AnalyticsActivityType)() {
    m.activity = value
}
// SetDuration sets the duration property value. Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
func (m *ActivityStatistics) SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.duration = value
}
// SetEndDate sets the endDate property value. Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetStartDate sets the startDate property value. Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.startDate = value
}
// SetTimeZoneUsed sets the timeZoneUsed property value. The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
func (m *ActivityStatistics) SetTimeZoneUsed(value *string)() {
    m.timeZoneUsed = value
}
