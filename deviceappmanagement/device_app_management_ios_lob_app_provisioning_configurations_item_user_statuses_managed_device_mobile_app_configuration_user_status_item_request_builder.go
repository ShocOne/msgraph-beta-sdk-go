package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderGetQueryParameters the list of user installation states for this mobile app configuration.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal instantiates a new ManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    m := &DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}/userStatuses/{managedDeviceMobileAppConfigurationUserStatus%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder instantiates a new ManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userStatuses for deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of user installation states for this mobile app configuration.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property userStatuses in deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property userStatuses for deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of user installation states for this mobile app configuration.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable), nil
}
// Patch update the navigation property userStatuses in deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable), nil
}
