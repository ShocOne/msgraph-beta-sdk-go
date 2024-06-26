package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder provides operations to manage the retentionEventTypes property of the microsoft.graph.security.triggerTypesRoot entity.
type TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderGetQueryParameters get a list of the retentionEventType objects and their properties.
type TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderGetQueryParameters
}
// TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRetentionEventTypeId provides operations to manage the retentionEventTypes property of the microsoft.graph.security.triggerTypesRoot entity.
// returns a *TriggertypesRetentioneventtypesRetentionEventTypeItemRequestBuilder when successful
func (m *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) ByRetentionEventTypeId(retentionEventTypeId string)(*TriggertypesRetentioneventtypesRetentionEventTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if retentionEventTypeId != "" {
        urlTplParams["retentionEventType%2Did"] = retentionEventTypeId
    }
    return NewTriggertypesRetentioneventtypesRetentionEventTypeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderInternal instantiates a new TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder and sets the default values.
func NewTriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) {
    m := &TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/triggerTypes/retentionEventTypes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder instantiates a new TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder and sets the default values.
func NewTriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TriggertypesRetentioneventtypesCountRequestBuilder when successful
func (m *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) Count()(*TriggertypesRetentioneventtypesCountRequestBuilder) {
    return NewTriggertypesRetentioneventtypesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the retentionEventType objects and their properties.
// returns a RetentionEventTypeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentioneventtype-list?view=graph-rest-beta
func (m *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateRetentionEventTypeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeCollectionResponseable), nil
}
// Post create a new retentionEventType object.
// returns a RetentionEventTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentioneventtype-post?view=graph-rest-beta
func (m *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, requestConfiguration *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateRetentionEventTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable), nil
}
// ToGetRequestInformation get a list of the retentionEventType objects and their properties.
// returns a *RequestInformation when successful
func (m *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new retentionEventType object.
// returns a *RequestInformation when successful
func (m *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, requestConfiguration *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder when successful
func (m *TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) WithUrl(rawUrl string)(*TriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder) {
    return NewTriggertypesRetentioneventtypesRetentionEventTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
