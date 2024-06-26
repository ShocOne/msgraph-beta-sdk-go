package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder provides operations to call the getRemediationHistory method.
type DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderInternal instantiates a new DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder and sets the default values.
func NewDevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder) {
    m := &DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/getRemediationHistory()", pathParameters),
    }
    return m
}
// NewDevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder instantiates a new DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder and sets the default values.
func NewDevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get function to get the number of remediations by a device health scripts
// returns a DeviceHealthScriptRemediationHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptRemediationHistoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptRemediationHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptRemediationHistoryable), nil
}
// ToGetRequestInformation function to get the number of remediations by a device health scripts
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder when successful
func (m *DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder) WithUrl(rawUrl string)(*DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder) {
    return NewDevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
