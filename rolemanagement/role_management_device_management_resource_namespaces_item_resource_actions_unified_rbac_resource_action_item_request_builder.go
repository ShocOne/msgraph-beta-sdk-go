package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
type RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderGetQueryParameters operations that an authorized principal are allowed to perform.
type RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderGetQueryParameters
}
// RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderInternal instantiates a new UnifiedRbacResourceActionItemRequestBuilder and sets the default values.
func NewRoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) {
    m := &RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/resourceActions/{unifiedRbacResourceAction%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder instantiates a new UnifiedRbacResourceActionItemRequestBuilder and sets the default values.
func NewRoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resourceActions for roleManagement
func (m *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation operations that an authorized principal are allowed to perform.
func (m *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property resourceActions in roleManagement
func (m *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, requestConfiguration *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property resourceActions for roleManagement
func (m *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get operations that an authorized principal are allowed to perform.
func (m *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceActionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable), nil
}
// Patch update the navigation property resourceActions in roleManagement
func (m *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, requestConfiguration *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceActionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable), nil
}
// ResourceScope provides operations to manage the resourceScope property of the microsoft.graph.unifiedRbacResourceAction entity.
func (m *RoleManagementDeviceManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) ResourceScope()(*RoleManagementDeviceManagementResourceNamespacesItemResourceActionsItemResourceScopeRequestBuilder) {
    return NewRoleManagementDeviceManagementResourceNamespacesItemResourceActionsItemResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
