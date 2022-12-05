package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OrganizationItemBrandingRequestBuilder provides operations to manage the branding property of the microsoft.graph.organization entity.
type OrganizationItemBrandingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationItemBrandingRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationItemBrandingRequestBuilderGetQueryParameters retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
type OrganizationItemBrandingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrganizationItemBrandingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrganizationItemBrandingRequestBuilderGetQueryParameters
}
// OrganizationItemBrandingRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemBrandingRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackgroundImage provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingRequestBuilder) BackgroundImage()(*OrganizationItemBrandingBackgroundImageRequestBuilder) {
    return NewOrganizationItemBrandingBackgroundImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BannerLogo provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingRequestBuilder) BannerLogo()(*OrganizationItemBrandingBannerLogoRequestBuilder) {
    return NewOrganizationItemBrandingBannerLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrganizationItemBrandingRequestBuilderInternal instantiates a new BrandingRequestBuilder and sets the default values.
func NewOrganizationItemBrandingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingRequestBuilder) {
    m := &OrganizationItemBrandingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/branding{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationItemBrandingRequestBuilder instantiates a new BrandingRequestBuilder and sets the default values.
func NewOrganizationItemBrandingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemBrandingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemBrandingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete the default organizational branding object. To delete the organizationalBranding object, all images (Stream types) must first be removed from the object.
func (m *OrganizationItemBrandingRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemBrandingRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
func (m *OrganizationItemBrandingRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemBrandingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the properties of the default branding object specified by the organizationalBranding resource.
func (m *OrganizationItemBrandingRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, requestConfiguration *OrganizationItemBrandingRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CustomCSS provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingRequestBuilder) CustomCSS()(*OrganizationItemBrandingCustomCSSRequestBuilder) {
    return NewOrganizationItemBrandingCustomCSSRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete the default organizational branding object. To delete the organizationalBranding object, all images (Stream types) must first be removed from the object.
func (m *OrganizationItemBrandingRequestBuilder) Delete(ctx context.Context, requestConfiguration *OrganizationItemBrandingRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Favicon provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingRequestBuilder) Favicon()(*OrganizationItemBrandingFaviconRequestBuilder) {
    return NewOrganizationItemBrandingFaviconRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the default organizational branding object, if the **Accept-Language** header is set to `0` or `default`. If no default organizational branding object exists, this method returns a `404 Not Found` error. If the **Accept-Language** header is set to an existing locale identified by the value of its **id**, this method retrieves the branding for the specified locale. This method retrieves only non-Stream properties, for example, **usernameHintText** and **signInPageText**. To retrieve Stream types of the default branding, for example, **bannerLogo** and **backgroundImage**, use the GET organizationalBrandingLocalization method.
func (m *OrganizationItemBrandingRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationItemBrandingRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationalBrandingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable), nil
}
// HeaderLogo provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingRequestBuilder) HeaderLogo()(*OrganizationItemBrandingHeaderLogoRequestBuilder) {
    return NewOrganizationItemBrandingHeaderLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Localizations provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
func (m *OrganizationItemBrandingRequestBuilder) Localizations()(*OrganizationItemBrandingLocalizationsRequestBuilder) {
    return NewOrganizationItemBrandingLocalizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizationsById provides operations to manage the localizations property of the microsoft.graph.organizationalBranding entity.
func (m *OrganizationItemBrandingRequestBuilder) LocalizationsById(id string)(*OrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organizationalBrandingLocalization%2Did"] = id
    }
    return NewOrganizationItemBrandingLocalizationsOrganizationalBrandingLocalizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of the default branding object specified by the organizationalBranding resource.
func (m *OrganizationItemBrandingRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, requestConfiguration *OrganizationItemBrandingRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationalBrandingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationalBrandingable), nil
}
// SquareLogo provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingRequestBuilder) SquareLogo()(*OrganizationItemBrandingSquareLogoRequestBuilder) {
    return NewOrganizationItemBrandingSquareLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SquareLogoDark provides operations to manage the media for the organization entity.
func (m *OrganizationItemBrandingRequestBuilder) SquareLogoDark()(*OrganizationItemBrandingSquareLogoDarkRequestBuilder) {
    return NewOrganizationItemBrandingSquareLogoDarkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
