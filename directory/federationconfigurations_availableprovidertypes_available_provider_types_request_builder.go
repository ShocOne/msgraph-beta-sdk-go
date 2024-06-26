package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder provides operations to call the availableProviderTypes method.
type FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters get all identity providers supported in a directory.
type FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters struct {
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
// FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetQueryParameters
}
// NewFederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal instantiates a new FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder and sets the default values.
func NewFederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    m := &FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/federationConfigurations/availableProviderTypes(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder instantiates a new FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder and sets the default values.
func NewFederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all identity providers supported in a directory.
// Deprecated: This method is obsolete. Use GetAsAvailableProviderTypesGetResponse instead.
// returns a FederationconfigurationsAvailableprovidertypesAvailableProviderTypesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityproviderbase-availableprovidertypes?view=graph-rest-beta
func (m *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(FederationconfigurationsAvailableprovidertypesAvailableProviderTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFederationconfigurationsAvailableprovidertypesAvailableProviderTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FederationconfigurationsAvailableprovidertypesAvailableProviderTypesResponseable), nil
}
// GetAsAvailableProviderTypesGetResponse get all identity providers supported in a directory.
// returns a FederationconfigurationsAvailableprovidertypesAvailableProviderTypesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityproviderbase-availableprovidertypes?view=graph-rest-beta
func (m *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) GetAsAvailableProviderTypesGetResponse(ctx context.Context, requestConfiguration *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(FederationconfigurationsAvailableprovidertypesAvailableProviderTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFederationconfigurationsAvailableprovidertypesAvailableProviderTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FederationconfigurationsAvailableprovidertypesAvailableProviderTypesGetResponseable), nil
}
// ToGetRequestInformation get all identity providers supported in a directory.
// returns a *RequestInformation when successful
func (m *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder when successful
func (m *FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) WithUrl(rawUrl string)(*FederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder) {
    return NewFederationconfigurationsAvailableprovidertypesAvailableProviderTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
