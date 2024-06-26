package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRoleassignmentsExportRequestBuilder provides operations to call the export method.
type ItemRoleassignmentsExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRoleassignmentsExportRequestBuilderGetQueryParameters invoke function export
type ItemRoleassignmentsExportRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemRoleassignmentsExportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleassignmentsExportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRoleassignmentsExportRequestBuilderGetQueryParameters
}
// NewItemRoleassignmentsExportRequestBuilderInternal instantiates a new ItemRoleassignmentsExportRequestBuilder and sets the default values.
func NewItemRoleassignmentsExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleassignmentsExportRequestBuilder) {
    m := &ItemRoleassignmentsExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/roleAssignments/export(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemRoleassignmentsExportRequestBuilder instantiates a new ItemRoleassignmentsExportRequestBuilder and sets the default values.
func NewItemRoleassignmentsExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleassignmentsExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRoleassignmentsExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function export
// Deprecated: This method is obsolete. Use GetAsExportGetResponse instead.
// returns a ItemRoleassignmentsExportResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleassignmentsExportRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRoleassignmentsExportRequestBuilderGetRequestConfiguration)(ItemRoleassignmentsExportResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRoleassignmentsExportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRoleassignmentsExportResponseable), nil
}
// GetAsExportGetResponse invoke function export
// returns a ItemRoleassignmentsExportGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleassignmentsExportRequestBuilder) GetAsExportGetResponse(ctx context.Context, requestConfiguration *ItemRoleassignmentsExportRequestBuilderGetRequestConfiguration)(ItemRoleassignmentsExportGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRoleassignmentsExportGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRoleassignmentsExportGetResponseable), nil
}
// ToGetRequestInformation invoke function export
// returns a *RequestInformation when successful
func (m *ItemRoleassignmentsExportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRoleassignmentsExportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRoleassignmentsExportRequestBuilder when successful
func (m *ItemRoleassignmentsExportRequestBuilder) WithUrl(rawUrl string)(*ItemRoleassignmentsExportRequestBuilder) {
    return NewItemRoleassignmentsExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
