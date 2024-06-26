package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder provides operations to call the compare method.
type IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetQueryParameters invoke function compare
type IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetQueryParameters struct {
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
// IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetQueryParameters
}
// NewIntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderInternal instantiates a new IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder and sets the default values.
func NewIntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, templateId *string)(*IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) {
    m := &IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/compare(templateId='{templateId}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if templateId != nil {
        m.BaseRequestBuilder.PathParameters["templateId"] = *templateId
    }
    return m
}
// NewIntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder instantiates a new IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder and sets the default values.
func NewIntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function compare
// Deprecated: This method is obsolete. Use GetAsCompareWithTemplateIdGetResponse instead.
// returns a IntentsItemComparewithtemplateidCompareWithTemplateIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(IntentsItemComparewithtemplateidCompareWithTemplateIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIntentsItemComparewithtemplateidCompareWithTemplateIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IntentsItemComparewithtemplateidCompareWithTemplateIdResponseable), nil
}
// GetAsCompareWithTemplateIdGetResponse invoke function compare
// returns a IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) GetAsCompareWithTemplateIdGetResponse(ctx context.Context, requestConfiguration *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIntentsItemComparewithtemplateidCompareWithTemplateIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponseable), nil
}
// ToGetRequestInformation invoke function compare
// returns a *RequestInformation when successful
func (m *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder when successful
func (m *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) WithUrl(rawUrl string)(*IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) {
    return NewIntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
