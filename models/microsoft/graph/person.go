package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Person 
type Person struct {
    Entity
    // The person's birthday.
    birthday *string;
    // The name of the person's company.
    companyName *string;
    // The person's department.
    department *string;
    // The person's display name.
    displayName *string;
    // The person's email addresses.
    emailAddresses []RankedEmailAddress;
    // The person's given name.
    givenName *string;
    // true if the user has flagged this person as a favorite.
    isFavorite *bool;
    // The type of mailbox that is represented by the person's email address.
    mailboxType *string;
    // The location of the person's office.
    officeLocation *string;
    // Free-form notes that the user has taken about this person.
    personNotes *string;
    // The type of person.
    personType *string;
    // The person's phone numbers.
    phones []Phone;
    // The person's addresses.
    postalAddresses []Location;
    // The person's profession.
    profession *string;
    // The sources the user data comes from, for example Directory or Outlook Contacts.
    sources []PersonDataSource;
    // The person's surname.
    surname *string;
    // The person's title.
    title *string;
    // The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
    userPrincipalName *string;
    // The person's websites.
    websites []Website;
    // The phonetic Japanese name of the person's company.
    yomiCompany *string;
}
// NewPerson instantiates a new person and sets the default values.
func NewPerson()(*Person) {
    m := &Person{
        Entity: *NewEntity(),
    }
    return m
}
// GetBirthday gets the birthday property value. The person's birthday.
func (m *Person) GetBirthday()(*string) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// GetCompanyName gets the companyName property value. The name of the person's company.
func (m *Person) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// GetDepartment gets the department property value. The person's department.
func (m *Person) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// GetDisplayName gets the displayName property value. The person's display name.
func (m *Person) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmailAddresses gets the emailAddresses property value. The person's email addresses.
func (m *Person) GetEmailAddresses()([]RankedEmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// GetGivenName gets the givenName property value. The person's given name.
func (m *Person) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// GetIsFavorite gets the isFavorite property value. true if the user has flagged this person as a favorite.
func (m *Person) GetIsFavorite()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFavorite
    }
}
// GetMailboxType gets the mailboxType property value. The type of mailbox that is represented by the person's email address.
func (m *Person) GetMailboxType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailboxType
    }
}
// GetOfficeLocation gets the officeLocation property value. The location of the person's office.
func (m *Person) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetPersonNotes gets the personNotes property value. Free-form notes that the user has taken about this person.
func (m *Person) GetPersonNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personNotes
    }
}
// GetPersonType gets the personType property value. The type of person.
func (m *Person) GetPersonType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personType
    }
}
// GetPhones gets the phones property value. The person's phone numbers.
func (m *Person) GetPhones()([]Phone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// GetPostalAddresses gets the postalAddresses property value. The person's addresses.
func (m *Person) GetPostalAddresses()([]Location) {
    if m == nil {
        return nil
    } else {
        return m.postalAddresses
    }
}
// GetProfession gets the profession property value. The person's profession.
func (m *Person) GetProfession()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profession
    }
}
// GetSources gets the sources property value. The sources the user data comes from, for example Directory or Outlook Contacts.
func (m *Person) GetSources()([]PersonDataSource) {
    if m == nil {
        return nil
    } else {
        return m.sources
    }
}
// GetSurname gets the surname property value. The person's surname.
func (m *Person) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// GetTitle gets the title property value. The person's title.
func (m *Person) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
func (m *Person) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetWebsites gets the websites property value. The person's websites.
func (m *Person) GetWebsites()([]Website) {
    if m == nil {
        return nil
    } else {
        return m.websites
    }
}
// GetYomiCompany gets the yomiCompany property value. The phonetic Japanese name of the person's company.
func (m *Person) GetYomiCompany()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiCompany
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Person) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["birthday"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthday(val)
        }
        return nil
    }
    res["companyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRankedEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RankedEmailAddress, len(val))
            for i, v := range val {
                res[i] = *(v.(*RankedEmailAddress))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["isFavorite"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavorite(val)
        }
        return nil
    }
    res["mailboxType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxType(val)
        }
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["personNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonNotes(val)
        }
        return nil
    }
    res["personType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonType(val)
        }
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhone() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phone, len(val))
            for i, v := range val {
                res[i] = *(v.(*Phone))
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["postalAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Location, len(val))
            for i, v := range val {
                res[i] = *(v.(*Location))
            }
            m.SetPostalAddresses(res)
        }
        return nil
    }
    res["profession"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfession(val)
        }
        return nil
    }
    res["sources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonDataSource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonDataSource, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonDataSource))
            }
            m.SetSources(res)
        }
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["websites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWebsite() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Website, len(val))
            for i, v := range val {
                res[i] = *(v.(*Website))
            }
            m.SetWebsites(res)
        }
        return nil
    }
    res["yomiCompany"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiCompany(val)
        }
        return nil
    }
    return res
}
func (m *Person) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Person) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmailAddresses()))
        for i, v := range m.GetEmailAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("emailAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFavorite", m.GetIsFavorite())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailboxType", m.GetMailboxType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personNotes", m.GetPersonNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personType", m.GetPersonType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPostalAddresses()))
        for i, v := range m.GetPostalAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("postalAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profession", m.GetProfession())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSources()))
        for i, v := range m.GetSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWebsites()))
        for i, v := range m.GetWebsites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("websites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiCompany", m.GetYomiCompany())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBirthday sets the birthday property value. The person's birthday.
func (m *Person) SetBirthday(value *string)() {
    if m != nil {
        m.birthday = value
    }
}
// SetCompanyName sets the companyName property value. The name of the person's company.
func (m *Person) SetCompanyName(value *string)() {
    if m != nil {
        m.companyName = value
    }
}
// SetDepartment sets the department property value. The person's department.
func (m *Person) SetDepartment(value *string)() {
    if m != nil {
        m.department = value
    }
}
// SetDisplayName sets the displayName property value. The person's display name.
func (m *Person) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmailAddresses sets the emailAddresses property value. The person's email addresses.
func (m *Person) SetEmailAddresses(value []RankedEmailAddress)() {
    if m != nil {
        m.emailAddresses = value
    }
}
// SetGivenName sets the givenName property value. The person's given name.
func (m *Person) SetGivenName(value *string)() {
    if m != nil {
        m.givenName = value
    }
}
// SetIsFavorite sets the isFavorite property value. true if the user has flagged this person as a favorite.
func (m *Person) SetIsFavorite(value *bool)() {
    if m != nil {
        m.isFavorite = value
    }
}
// SetMailboxType sets the mailboxType property value. The type of mailbox that is represented by the person's email address.
func (m *Person) SetMailboxType(value *string)() {
    if m != nil {
        m.mailboxType = value
    }
}
// SetOfficeLocation sets the officeLocation property value. The location of the person's office.
func (m *Person) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetPersonNotes sets the personNotes property value. Free-form notes that the user has taken about this person.
func (m *Person) SetPersonNotes(value *string)() {
    if m != nil {
        m.personNotes = value
    }
}
// SetPersonType sets the personType property value. The type of person.
func (m *Person) SetPersonType(value *string)() {
    if m != nil {
        m.personType = value
    }
}
// SetPhones sets the phones property value. The person's phone numbers.
func (m *Person) SetPhones(value []Phone)() {
    if m != nil {
        m.phones = value
    }
}
// SetPostalAddresses sets the postalAddresses property value. The person's addresses.
func (m *Person) SetPostalAddresses(value []Location)() {
    if m != nil {
        m.postalAddresses = value
    }
}
// SetProfession sets the profession property value. The person's profession.
func (m *Person) SetProfession(value *string)() {
    if m != nil {
        m.profession = value
    }
}
// SetSources sets the sources property value. The sources the user data comes from, for example Directory or Outlook Contacts.
func (m *Person) SetSources(value []PersonDataSource)() {
    if m != nil {
        m.sources = value
    }
}
// SetSurname sets the surname property value. The person's surname.
func (m *Person) SetSurname(value *string)() {
    if m != nil {
        m.surname = value
    }
}
// SetTitle sets the title property value. The person's title.
func (m *Person) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
func (m *Person) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetWebsites sets the websites property value. The person's websites.
func (m *Person) SetWebsites(value []Website)() {
    if m != nil {
        m.websites = value
    }
}
// SetYomiCompany sets the yomiCompany property value. The phonetic Japanese name of the person's company.
func (m *Person) SetYomiCompany(value *string)() {
    if m != nil {
        m.yomiCompany = value
    }
}
