package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderGetQueryParameters get a list of the privilegedAccessGroupEligibilityScheduleRequest objects and their properties.
type PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderGetQueryParameters struct {
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
// PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrivilegedAccessGroupEligibilityScheduleRequestId provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) ByPrivilegedAccessGroupEligibilityScheduleRequestId(privilegedAccessGroupEligibilityScheduleRequestId string)(*PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegedAccessGroupEligibilityScheduleRequestId != "" {
        urlTplParams["privilegedAccessGroupEligibilityScheduleRequest%2Did"] = privilegedAccessGroupEligibilityScheduleRequestId
    }
    return NewPrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderInternal instantiates a new PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) {
    m := &PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder instantiates a new PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsCountRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) Count()(*PrivilegedaccessGroupEligibilityschedulerequestsCountRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedaccessGroupEligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get a list of the privilegedAccessGroupEligibilityScheduleRequest objects and their properties.
// returns a PrivilegedAccessGroupEligibilityScheduleRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-list-eligibilityschedulerequests?view=graph-rest-beta
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupEligibilityScheduleRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestCollectionResponseable), nil
}
// Post create a new privilegedAccessGroupEligibilityScheduleRequest object.
// returns a PrivilegedAccessGroupEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-post-eligibilityschedulerequests?view=graph-rest-beta
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable), nil
}
// ToGetRequestInformation get a list of the privilegedAccessGroupEligibilityScheduleRequest objects and their properties.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new privilegedAccessGroupEligibilityScheduleRequest object.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
