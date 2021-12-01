package reminderviewwithstartdatetimewithenddatetime

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ReminderViewWithStartDateTimeWithEndDateTime 
type ReminderViewWithStartDateTimeWithEndDateTime struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
    changeKey *string;
    // The date, time and time zone that the event ends.
    eventEndTime *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
    // The unique ID of the event. Read only.
    eventId *string;
    // The location of the event.
    eventLocation *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location;
    // The date, time, and time zone that the event starts.
    eventStartTime *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
    // The text of the event's subject line.
    eventSubject *string;
    // The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
    eventWebLink *string;
    // The date, time, and time zone that the reminder is set to occur.
    reminderFireTime *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
}
// NewReminderViewWithStartDateTimeWithEndDateTime instantiates a new reminderViewWithStartDateTimeWithEndDateTime and sets the default values.
func NewReminderViewWithStartDateTimeWithEndDateTime()(*ReminderViewWithStartDateTimeWithEndDateTime) {
    m := &ReminderViewWithStartDateTimeWithEndDateTime{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChangeKey gets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// GetEventEndTime gets the eventEndTime property value. The date, time and time zone that the event ends.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventEndTime()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.eventEndTime
    }
}
// GetEventId gets the eventId property value. The unique ID of the event. Read only.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventId
    }
}
// GetEventLocation gets the eventLocation property value. The location of the event.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventLocation()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location) {
    if m == nil {
        return nil
    } else {
        return m.eventLocation
    }
}
// GetEventStartTime gets the eventStartTime property value. The date, time, and time zone that the event starts.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventStartTime()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.eventStartTime
    }
}
// GetEventSubject gets the eventSubject property value. The text of the event's subject line.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventSubject
    }
}
// GetEventWebLink gets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventWebLink
    }
}
// GetReminderFireTime gets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetReminderFireTime()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.reminderFireTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["eventEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventEndTime(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
        }
        return nil
    }
    res["eventId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventId(val)
        }
        return nil
    }
    res["eventLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventLocation(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location))
        }
        return nil
    }
    res["eventStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventStartTime(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
        }
        return nil
    }
    res["eventSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventSubject(val)
        }
        return nil
    }
    res["eventWebLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventWebLink(val)
        }
        return nil
    }
    res["reminderFireTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderFireTime(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
        }
        return nil
    }
    return res
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ReminderViewWithStartDateTimeWithEndDateTime) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventEndTime", m.GetEventEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventId", m.GetEventId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventLocation", m.GetEventLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventStartTime", m.GetEventStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventSubject", m.GetEventSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventWebLink", m.GetEventWebLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reminderFireTime", m.GetReminderFireTime())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChangeKey sets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetChangeKey(value *string)() {
    if m != nil {
        m.changeKey = value
    }
}
// SetEventEndTime sets the eventEndTime property value. The date, time and time zone that the event ends.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventEndTime(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    if m != nil {
        m.eventEndTime = value
    }
}
// SetEventId sets the eventId property value. The unique ID of the event. Read only.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventId(value *string)() {
    if m != nil {
        m.eventId = value
    }
}
// SetEventLocation sets the eventLocation property value. The location of the event.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventLocation(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location)() {
    if m != nil {
        m.eventLocation = value
    }
}
// SetEventStartTime sets the eventStartTime property value. The date, time, and time zone that the event starts.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventStartTime(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    if m != nil {
        m.eventStartTime = value
    }
}
// SetEventSubject sets the eventSubject property value. The text of the event's subject line.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventSubject(value *string)() {
    if m != nil {
        m.eventSubject = value
    }
}
// SetEventWebLink sets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventWebLink(value *string)() {
    if m != nil {
        m.eventWebLink = value
    }
}
// SetReminderFireTime sets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetReminderFireTime(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    if m != nil {
        m.reminderFireTime = value
    }
}
