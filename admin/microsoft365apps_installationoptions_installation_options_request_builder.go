package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder provides operations to manage the installationOptions property of the microsoft.graph.adminMicrosoft365Apps entity.
type Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderGetQueryParameters read the properties and relationships of an m365AppsInstallationOptions object.
type Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderGetQueryParameters
}
// Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoft365appsInstallationoptionsInstallationOptionsRequestBuilderInternal instantiates a new Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder and sets the default values.
func NewMicrosoft365appsInstallationoptionsInstallationOptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) {
    m := &Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/microsoft365Apps/installationOptions{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMicrosoft365appsInstallationoptionsInstallationOptionsRequestBuilder instantiates a new Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder and sets the default values.
func NewMicrosoft365appsInstallationoptionsInstallationOptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoft365appsInstallationoptionsInstallationOptionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property installationOptions for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) Delete(ctx context.Context, requestConfiguration *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of an m365AppsInstallationOptions object.
// returns a M365AppsInstallationOptionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/m365appsinstallationoptions-get?view=graph-rest-beta
func (m *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.M365AppsInstallationOptionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateM365AppsInstallationOptionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.M365AppsInstallationOptionsable), nil
}
// Patch update the properties of an m365AppsInstallationOptions object.
// returns a M365AppsInstallationOptionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/m365appsinstallationoptions-update?view=graph-rest-beta
func (m *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.M365AppsInstallationOptionsable, requestConfiguration *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.M365AppsInstallationOptionsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateM365AppsInstallationOptionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.M365AppsInstallationOptionsable), nil
}
// ToDeleteRequestInformation delete navigation property installationOptions for admin
// returns a *RequestInformation when successful
func (m *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an m365AppsInstallationOptions object.
// returns a *RequestInformation when successful
func (m *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of an m365AppsInstallationOptions object.
// returns a *RequestInformation when successful
func (m *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.M365AppsInstallationOptionsable, requestConfiguration *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder when successful
func (m *Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) WithUrl(rawUrl string)(*Microsoft365appsInstallationoptionsInstallationOptionsRequestBuilder) {
    return NewMicrosoft365appsInstallationoptionsInstallationOptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
