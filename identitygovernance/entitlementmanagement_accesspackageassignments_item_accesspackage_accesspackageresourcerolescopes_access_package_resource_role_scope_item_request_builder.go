package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderGetQueryParameters get accessPackageResourceRoleScopes from identityGovernance
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceRole provides operations to manage the accessPackageResourceRole property of the microsoft.graph.accessPackageResourceRoleScope entity.
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesItemAccesspackageresourceroleAccessPackageResourceRoleRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) AccessPackageResourceRole()(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesItemAccesspackageresourceroleAccessPackageResourceRoleRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesItemAccesspackageresourceroleAccessPackageResourceRoleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScope provides operations to manage the accessPackageResourceScope property of the microsoft.graph.accessPackageResourceRoleScope entity.
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccessPackageResourceScopeRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) AccessPackageResourceScope()(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccessPackageResourceScopeRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesItemAccesspackageresourcescopeAccessPackageResourceScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResourceRoleScopes for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get accessPackageResourceRoleScopes from identityGovernance
// returns a AccessPackageResourceRoleScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable), nil
}
// Patch update the navigation property accessPackageResourceRoleScopes in identityGovernance
// returns a AccessPackageResourceRoleScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageResourceRoleScopes for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get accessPackageResourceRoleScopes from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property accessPackageResourceRoleScopes in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, requestConfiguration *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsItemAccesspackageAccesspackageresourcerolescopesAccessPackageResourceRoleScopeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
