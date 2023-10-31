package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder provides operations to call the refresh method.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilderInternal instantiates a new RefreshRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageResourceRole/accessPackageResource/refresh", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder instantiates a new RefreshRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action refresh
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration)(error) {
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
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceRoleAccessPackageResourceRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
