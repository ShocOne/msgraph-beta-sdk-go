package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
type MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderGetQueryParameters the set of direct relationships for this app.
type MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderInternal instantiates a new MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) {
    m := &MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.iosVppApp/relationships/{mobileAppRelationship%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder instantiates a new MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property relationships for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the set of direct relationships for this app.
// returns a MobileAppRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable), nil
}
// Patch update the navigation property relationships in deviceAppManagement
// returns a MobileAppRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, requestConfiguration *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable), nil
}
// ToDeleteRequestInformation delete navigation property relationships for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the set of direct relationships for this app.
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property relationships in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, requestConfiguration *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder when successful
func (m *MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder) {
    return NewMobileAppsItemGraphIosVppAppRelationshipsMobileAppRelationshipItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
