package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
type MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetQueryParameters managed device mobile app configuration states for this device.
type MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetQueryParameters
}
// MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal instantiates a new AssignmentFilterEvaluationStatusDetailsItemRequestBuilder and sets the default values.
func NewMeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    m := &MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/assignmentFilterEvaluationStatusDetails/{assignmentFilterEvaluationStatusDetails%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder instantiates a new AssignmentFilterEvaluationStatusDetailsItemRequestBuilder and sets the default values.
func NewMeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignmentFilterEvaluationStatusDetails for me
func (m *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation managed device mobile app configuration states for this device.
func (m *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property assignmentFilterEvaluationStatusDetails in me
func (m *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, requestConfiguration *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property assignmentFilterEvaluationStatusDetails for me
func (m *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get managed device mobile app configuration states for this device.
func (m *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable), nil
}
// Patch update the navigation property assignmentFilterEvaluationStatusDetails in me
func (m *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, requestConfiguration *MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterEvaluationStatusDetailsable), nil
}
