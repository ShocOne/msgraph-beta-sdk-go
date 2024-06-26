package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder provides operations to manage the appHealthMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
type UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderGetQueryParameters the scores and insights for the application health metrics.
type UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderGetQueryParameters
}
// NewUserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderInternal instantiates a new UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder) {
    m := &UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline%2Did}/appHealthMetrics{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder instantiates a new UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder and sets the default values.
func NewUserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the scores and insights for the application health metrics.
// returns a UserExperienceAnalyticsCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsCategoryable), nil
}
// ToGetRequestInformation the scores and insights for the application health metrics.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesItemApphealthmetricsAppHealthMetricsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
