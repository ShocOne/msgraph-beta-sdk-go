package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder provides operations to manage the itemActivityStats property of the microsoft.graph.itemAnalytics entity.
type ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetQueryParameters get itemActivityStats from groups
type ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetQueryParameters
}
// ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByItemActivityStatId provides operations to manage the itemActivityStats property of the microsoft.graph.itemAnalytics entity.
// returns a *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatItemRequestBuilder when successful
func (m *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) ByItemActivityStatId(itemActivityStatId string)(*ItemSitesItemAnalyticsItemactivitystatsItemActivityStatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if itemActivityStatId != "" {
        urlTplParams["itemActivityStat%2Did"] = itemActivityStatId
    }
    return NewItemSitesItemAnalyticsItemactivitystatsItemActivityStatItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderInternal instantiates a new ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder and sets the default values.
func NewItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) {
    m := &ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/analytics/itemActivityStats{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder instantiates a new ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder and sets the default values.
func NewItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemAnalyticsItemactivitystatsCountRequestBuilder when successful
func (m *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) Count()(*ItemSitesItemAnalyticsItemactivitystatsCountRequestBuilder) {
    return NewItemSitesItemAnalyticsItemactivitystatsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get itemActivityStats from groups
// returns a ItemActivityStatCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemActivityStatCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatCollectionResponseable), nil
}
// Post create new navigation property to itemActivityStats for groups
// returns a ItemActivityStatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable, requestConfiguration *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemActivityStatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable), nil
}
// ToGetRequestInformation get itemActivityStats from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to itemActivityStats for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemActivityStatable, requestConfiguration *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder when successful
func (m *ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) {
    return NewItemSitesItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
