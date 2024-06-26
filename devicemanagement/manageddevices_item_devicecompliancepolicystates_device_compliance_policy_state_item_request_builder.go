package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderGetQueryParameters device compliance policy states for this device.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderGetQueryParameters
}
// ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderInternal instantiates a new ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder and sets the default values.
func NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) {
    m := &ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceCompliancePolicyStates/{deviceCompliancePolicyState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder instantiates a new ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder and sets the default values.
func NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceCompliancePolicyStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get device compliance policy states for this device.
// returns a DeviceCompliancePolicyStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable), nil
}
// Patch update the navigation property deviceCompliancePolicyStates in deviceManagement
// returns a DeviceCompliancePolicyStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable), nil
}
// ToDeleteRequestInformation delete navigation property deviceCompliancePolicyStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation device compliance policy states for this device.
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceCompliancePolicyStates in deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyStateable, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder) {
    return NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
