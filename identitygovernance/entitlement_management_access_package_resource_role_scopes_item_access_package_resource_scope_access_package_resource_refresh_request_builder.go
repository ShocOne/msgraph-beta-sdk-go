package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder provides operations to call the refresh method.
type EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderInternal instantiates a new RefreshRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) {
    m := &EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceScope/accessPackageResource/refresh", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder instantiates a new RefreshRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action refresh
func (m *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action refresh
func (m *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
