package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal instantiates a new AccessPackageResourceRoleItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceRole/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder instantiates a new AccessPackageResourceRoleItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResourceRoles for identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
// Patch update the navigation property accessPackageResourceRoles in identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageResourceRoles for identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property accessPackageResourceRoles in identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
