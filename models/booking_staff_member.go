package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingStaffMember 
type BookingStaffMember struct {
    BookingPerson
    // True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
    availabilityIsAffectedByPersonalCalendar *bool
    // Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
    colorIndex *int32
    // The isEmailNotificationEnabled property
    isEmailNotificationEnabled *bool
    // The role property
    role *BookingStaffRole
    // The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
    timeZone *string
    // True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
    useBusinessHours *bool
    // The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
    workingHours []BookingWorkHoursable
}
// NewBookingStaffMember instantiates a new BookingStaffMember and sets the default values.
func NewBookingStaffMember()(*BookingStaffMember) {
    m := &BookingStaffMember{
        BookingPerson: *NewBookingPerson(),
    }
    odataTypeValue := "#microsoft.graph.bookingStaffMember";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBookingStaffMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingStaffMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingStaffMember(), nil
}
// GetAvailabilityIsAffectedByPersonalCalendar gets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) GetAvailabilityIsAffectedByPersonalCalendar()(*bool) {
    return m.availabilityIsAffectedByPersonalCalendar
}
// GetColorIndex gets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
func (m *BookingStaffMember) GetColorIndex()(*int32) {
    return m.colorIndex
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingStaffMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["availabilityIsAffectedByPersonalCalendar"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAvailabilityIsAffectedByPersonalCalendar)
    res["colorIndex"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetColorIndex)
    res["isEmailNotificationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEmailNotificationEnabled)
    res["role"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseBookingStaffRole , m.SetRole)
    res["timeZone"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTimeZone)
    res["useBusinessHours"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUseBusinessHours)
    res["workingHours"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue , m.SetWorkingHours)
    return res
}
// GetIsEmailNotificationEnabled gets the isEmailNotificationEnabled property value. The isEmailNotificationEnabled property
func (m *BookingStaffMember) GetIsEmailNotificationEnabled()(*bool) {
    return m.isEmailNotificationEnabled
}
// GetRole gets the role property value. The role property
func (m *BookingStaffMember) GetRole()(*BookingStaffRole) {
    return m.role
}
// GetTimeZone gets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) GetTimeZone()(*string) {
    return m.timeZone
}
// GetUseBusinessHours gets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) GetUseBusinessHours()(*bool) {
    return m.useBusinessHours
}
// GetWorkingHours gets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) GetWorkingHours()([]BookingWorkHoursable) {
    return m.workingHours
}
// Serialize serializes information the current object
func (m *BookingStaffMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingPerson.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("availabilityIsAffectedByPersonalCalendar", m.GetAvailabilityIsAffectedByPersonalCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("colorIndex", m.GetColorIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEmailNotificationEnabled", m.GetIsEmailNotificationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useBusinessHours", m.GetUseBusinessHours())
        if err != nil {
            return err
        }
    }
    if m.GetWorkingHours() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWorkingHours())
        err = writer.WriteCollectionOfObjectValues("workingHours", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityIsAffectedByPersonalCalendar sets the availabilityIsAffectedByPersonalCalendar property value. True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's availability in their personal calendar in Microsoft 365, before making a booking.
func (m *BookingStaffMember) SetAvailabilityIsAffectedByPersonalCalendar(value *bool)() {
    m.availabilityIsAffectedByPersonalCalendar = value
}
// SetColorIndex sets the colorIndex property value. Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details page in the Bookings app.
func (m *BookingStaffMember) SetColorIndex(value *int32)() {
    m.colorIndex = value
}
// SetIsEmailNotificationEnabled sets the isEmailNotificationEnabled property value. The isEmailNotificationEnabled property
func (m *BookingStaffMember) SetIsEmailNotificationEnabled(value *bool)() {
    m.isEmailNotificationEnabled = value
}
// SetRole sets the role property value. The role property
func (m *BookingStaffMember) SetRole(value *BookingStaffRole)() {
    m.role = value
}
// SetTimeZone sets the timeZone property value. The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
func (m *BookingStaffMember) SetTimeZone(value *string)() {
    m.timeZone = value
}
// SetUseBusinessHours sets the useBusinessHours property value. True means the staff member's availability is as specified in the businessHours property of the business. False means the availability is determined by the staff member's workingHours property setting.
func (m *BookingStaffMember) SetUseBusinessHours(value *bool)() {
    m.useBusinessHours = value
}
// SetWorkingHours sets the workingHours property value. The range of hours each day of the week that the staff member is available for booking. By default, they are initialized to be the same as the businessHours property of the business.
func (m *BookingStaffMember) SetWorkingHours(value []BookingWorkHoursable)() {
    m.workingHours = value
}
