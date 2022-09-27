package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebApplication 
type WebApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Home page or landing page of the application.
    homePageUrl *string
    // Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
    implicitGrantSettings ImplicitGrantSettingsable
    // Specifies the URL that will be used by Microsoft's authorization service to logout an user using front-channel, back-channel or SAML logout protocols.
    logoutUrl *string
    // The oauth2AllowImplicitFlow property
    oauth2AllowImplicitFlow *bool
    // The OdataType property
    odataType *string
    // Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
    redirectUris []string
    // Specifies the index of the URLs where user tokens are sent for sign-in. This is only valid for applications using SAML.
    redirectUriSettings []RedirectUriSettingsable
}
// NewWebApplication instantiates a new webApplication and sets the default values.
func NewWebApplication()(*WebApplication) {
    m := &WebApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.webApplication";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWebApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebApplication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebApplication) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["homePageUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHomePageUrl)
    res["implicitGrantSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateImplicitGrantSettingsFromDiscriminatorValue , m.SetImplicitGrantSettings)
    res["logoutUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLogoutUrl)
    res["oauth2AllowImplicitFlow"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetOauth2AllowImplicitFlow)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["redirectUris"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRedirectUris)
    res["redirectUriSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRedirectUriSettingsFromDiscriminatorValue , m.SetRedirectUriSettings)
    return res
}
// GetHomePageUrl gets the homePageUrl property value. Home page or landing page of the application.
func (m *WebApplication) GetHomePageUrl()(*string) {
    return m.homePageUrl
}
// GetImplicitGrantSettings gets the implicitGrantSettings property value. Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
func (m *WebApplication) GetImplicitGrantSettings()(ImplicitGrantSettingsable) {
    return m.implicitGrantSettings
}
// GetLogoutUrl gets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using front-channel, back-channel or SAML logout protocols.
func (m *WebApplication) GetLogoutUrl()(*string) {
    return m.logoutUrl
}
// GetOauth2AllowImplicitFlow gets the oauth2AllowImplicitFlow property value. The oauth2AllowImplicitFlow property
func (m *WebApplication) GetOauth2AllowImplicitFlow()(*bool) {
    return m.oauth2AllowImplicitFlow
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WebApplication) GetOdataType()(*string) {
    return m.odataType
}
// GetRedirectUris gets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *WebApplication) GetRedirectUris()([]string) {
    return m.redirectUris
}
// GetRedirectUriSettings gets the redirectUriSettings property value. Specifies the index of the URLs where user tokens are sent for sign-in. This is only valid for applications using SAML.
func (m *WebApplication) GetRedirectUriSettings()([]RedirectUriSettingsable) {
    return m.redirectUriSettings
}
// Serialize serializes information the current object
func (m *WebApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("homePageUrl", m.GetHomePageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("implicitGrantSettings", m.GetImplicitGrantSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("oauth2AllowImplicitFlow", m.GetOauth2AllowImplicitFlow())
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
    if m.GetRedirectUris() != nil {
        err := writer.WriteCollectionOfStringValues("redirectUris", m.GetRedirectUris())
        if err != nil {
            return err
        }
    }
    if m.GetRedirectUriSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRedirectUriSettings())
        err := writer.WriteCollectionOfObjectValues("redirectUriSettings", cast)
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
func (m *WebApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHomePageUrl sets the homePageUrl property value. Home page or landing page of the application.
func (m *WebApplication) SetHomePageUrl(value *string)() {
    m.homePageUrl = value
}
// SetImplicitGrantSettings sets the implicitGrantSettings property value. Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
func (m *WebApplication) SetImplicitGrantSettings(value ImplicitGrantSettingsable)() {
    m.implicitGrantSettings = value
}
// SetLogoutUrl sets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using front-channel, back-channel or SAML logout protocols.
func (m *WebApplication) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
// SetOauth2AllowImplicitFlow sets the oauth2AllowImplicitFlow property value. The oauth2AllowImplicitFlow property
func (m *WebApplication) SetOauth2AllowImplicitFlow(value *bool)() {
    m.oauth2AllowImplicitFlow = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebApplication) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRedirectUris sets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *WebApplication) SetRedirectUris(value []string)() {
    m.redirectUris = value
}
// SetRedirectUriSettings sets the redirectUriSettings property value. Specifies the index of the URLs where user tokens are sent for sign-in. This is only valid for applications using SAML.
func (m *WebApplication) SetRedirectUriSettings(value []RedirectUriSettingsable)() {
    m.redirectUriSettings = value
}
