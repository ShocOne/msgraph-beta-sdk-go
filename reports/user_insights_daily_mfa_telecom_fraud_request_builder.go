package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsDailyMfaTelecomFraudRequestBuilder provides operations to manage the mfaTelecomFraud property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
type UserInsightsDailyMfaTelecomFraudRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsDailyMfaTelecomFraudRequestBuilderGetQueryParameters get mfaTelecomFraud from reports
type UserInsightsDailyMfaTelecomFraudRequestBuilderGetQueryParameters struct {
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
// UserInsightsDailyMfaTelecomFraudRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyMfaTelecomFraudRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsDailyMfaTelecomFraudRequestBuilderGetQueryParameters
}
// UserInsightsDailyMfaTelecomFraudRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyMfaTelecomFraudRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMfaTelecomFraudMetricId provides operations to manage the mfaTelecomFraud property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder when successful
func (m *UserInsightsDailyMfaTelecomFraudRequestBuilder) ByMfaTelecomFraudMetricId(mfaTelecomFraudMetricId string)(*UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mfaTelecomFraudMetricId != "" {
        urlTplParams["mfaTelecomFraudMetric%2Did"] = mfaTelecomFraudMetricId
    }
    return NewUserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserInsightsDailyMfaTelecomFraudRequestBuilderInternal instantiates a new UserInsightsDailyMfaTelecomFraudRequestBuilder and sets the default values.
func NewUserInsightsDailyMfaTelecomFraudRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyMfaTelecomFraudRequestBuilder) {
    m := &UserInsightsDailyMfaTelecomFraudRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/daily/mfaTelecomFraud{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserInsightsDailyMfaTelecomFraudRequestBuilder instantiates a new UserInsightsDailyMfaTelecomFraudRequestBuilder and sets the default values.
func NewUserInsightsDailyMfaTelecomFraudRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyMfaTelecomFraudRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsDailyMfaTelecomFraudRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserInsightsDailyMfaTelecomFraudCountRequestBuilder when successful
func (m *UserInsightsDailyMfaTelecomFraudRequestBuilder) Count()(*UserInsightsDailyMfaTelecomFraudCountRequestBuilder) {
    return NewUserInsightsDailyMfaTelecomFraudCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get mfaTelecomFraud from reports
// returns a MfaTelecomFraudMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyMfaTelecomFraudRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsDailyMfaTelecomFraudRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMfaTelecomFraudMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricCollectionResponseable), nil
}
// Post create new navigation property to mfaTelecomFraud for reports
// returns a MfaTelecomFraudMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyMfaTelecomFraudRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable, requestConfiguration *UserInsightsDailyMfaTelecomFraudRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMfaTelecomFraudMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable), nil
}
// ToGetRequestInformation get mfaTelecomFraud from reports
// returns a *RequestInformation when successful
func (m *UserInsightsDailyMfaTelecomFraudRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsDailyMfaTelecomFraudRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mfaTelecomFraud for reports
// returns a *RequestInformation when successful
func (m *UserInsightsDailyMfaTelecomFraudRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable, requestConfiguration *UserInsightsDailyMfaTelecomFraudRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserInsightsDailyMfaTelecomFraudRequestBuilder when successful
func (m *UserInsightsDailyMfaTelecomFraudRequestBuilder) WithUrl(rawUrl string)(*UserInsightsDailyMfaTelecomFraudRequestBuilder) {
    return NewUserInsightsDailyMfaTelecomFraudRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
