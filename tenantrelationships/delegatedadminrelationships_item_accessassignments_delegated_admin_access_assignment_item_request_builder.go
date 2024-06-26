package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
type DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetQueryParameters read the properties of a delegatedAdminAccessAssignment object.
type DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetQueryParameters
}
// DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal instantiates a new DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) {
    m := &DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}/accessAssignments/{delegatedAdminAccessAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder instantiates a new DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a delegatedAdminAccessAssignment object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminaccessassignment-delete?view=graph-rest-beta
func (m *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties of a delegatedAdminAccessAssignment object.
// returns a DelegatedAdminAccessAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminaccessassignment-get?view=graph-rest-beta
func (m *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminAccessAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminAccessAssignmentable), nil
}
// Patch update the properties of a delegatedAdminAccessAssignment object.
// returns a DelegatedAdminAccessAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminaccessassignment-update?view=graph-rest-beta
func (m *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminAccessAssignmentable, requestConfiguration *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminAccessAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminAccessAssignmentable), nil
}
// ToDeleteRequestInformation delete a delegatedAdminAccessAssignment object.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties of a delegatedAdminAccessAssignment object.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a delegatedAdminAccessAssignment object.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminAccessAssignmentable, requestConfiguration *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder when successful
func (m *DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*DelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) {
    return NewDelegatedadminrelationshipsItemAccessassignmentsDelegatedAdminAccessAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
