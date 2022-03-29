package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplicationTemplateable 
type ApplicationTemplateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategories()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHomePageUrl()(*string)
    GetInformationalUrls()(InformationalUrlsable)
    GetLogoUrl()(*string)
    GetPublisher()(*string)
    GetSupportedProvisioningTypes()([]string)
    GetSupportedSingleSignOnModes()([]string)
    SetCategories(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHomePageUrl(value *string)()
    SetInformationalUrls(value InformationalUrlsable)()
    SetLogoUrl(value *string)()
    SetPublisher(value *string)()
    SetSupportedProvisioningTypes(value []string)()
    SetSupportedSingleSignOnModes(value []string)()
}