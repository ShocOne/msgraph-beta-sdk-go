package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters in Microsoft Entra Entitlement Management, retrieve a list of accessPackageAssignmentRequest objects filtered on the signed-in user.
type EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    m := &EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/filterByCurrentUser(on='{on}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if on != nil {
        m.BaseRequestBuilder.PathParameters["on"] = *on
    }
    return m
}
// NewEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder instantiates a new EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get in Microsoft Entra Entitlement Management, retrieve a list of accessPackageAssignmentRequest objects filtered on the signed-in user.
// Deprecated: This method is obsolete. Use GetAsFilterByCurrentUserWithOnGetResponse instead.
// returns a EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequest-filterbycurrentuser?view=graph-rest-beta
func (m *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable), nil
}
// GetAsFilterByCurrentUserWithOnGetResponse in Microsoft Entra Entitlement Management, retrieve a list of accessPackageAssignmentRequest objects filtered on the signed-in user.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequest-filterbycurrentuser?view=graph-rest-beta
func (m *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) GetAsFilterByCurrentUserWithOnGetResponse(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable), nil
}
// ToGetRequestInformation in Microsoft Entra Entitlement Management, retrieve a list of accessPackageAssignmentRequest objects filtered on the signed-in user.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
