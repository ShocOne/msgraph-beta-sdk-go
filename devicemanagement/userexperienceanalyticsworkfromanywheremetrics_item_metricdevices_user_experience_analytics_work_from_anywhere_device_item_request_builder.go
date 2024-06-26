package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder provides operations to manage the metricDevices property of the microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric entity.
type UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters the work from anywhere metric devices. Read-only.
type UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    m := &UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/{userExperienceAnalyticsWorkFromAnywhereMetric%2Did}/metricDevices/{userExperienceAnalyticsWorkFromAnywhereDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder instantiates a new UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property metricDevices for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the work from anywhere metric devices. Read-only.
// returns a UserExperienceAnalyticsWorkFromAnywhereDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable), nil
}
// Patch update the navigation property metricDevices in deviceManagement
// returns a UserExperienceAnalyticsWorkFromAnywhereDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable), nil
}
// ToDeleteRequestInformation delete navigation property metricDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the work from anywhere metric devices. Read-only.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property metricDevices in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
