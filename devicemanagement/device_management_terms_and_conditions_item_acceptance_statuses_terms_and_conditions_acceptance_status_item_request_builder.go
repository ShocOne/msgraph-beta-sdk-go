package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder provides operations to manage the acceptanceStatuses property of the microsoft.graph.termsAndConditions entity.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderGetQueryParameters the list of acceptance statuses for this T&C policy.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderGetQueryParameters
}
// DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal instantiates a new TermsAndConditionsAcceptanceStatusItemRequestBuilder and sets the default values.
func NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) {
    m := &DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder instantiates a new TermsAndConditionsAcceptanceStatusItemRequestBuilder and sets the default values.
func NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property acceptanceStatuses for deviceManagement
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of acceptance statuses for this T&C policy.
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property acceptanceStatuses in deviceManagement
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsAcceptanceStatusable, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property acceptanceStatuses for deviceManagement
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of acceptance statuses for this T&C policy.
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsAcceptanceStatusable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsAcceptanceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsAcceptanceStatusable), nil
}
// Patch update the navigation property acceptanceStatuses in deviceManagement
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsAcceptanceStatusable, requestConfiguration *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsAcceptanceStatusable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsAcceptanceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsAcceptanceStatusable), nil
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.termsAndConditionsAcceptanceStatus entity.
func (m *DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) TermsAndConditions()(*DeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilder) {
    return NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesItemTermsAndConditionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
