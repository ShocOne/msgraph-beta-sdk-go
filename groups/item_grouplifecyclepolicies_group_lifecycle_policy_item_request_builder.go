package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
type ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderGetQueryParameters the collection of lifecycle policies for this group. Read-only. Nullable.
type ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderGetQueryParameters
}
// ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddGroup provides operations to call the addGroup method.
// returns a *ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) AddGroup()(*ItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilder) {
    return NewItemGrouplifecyclepoliciesItemAddgroupAddGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderInternal instantiates a new ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) {
    m := &ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/{groupLifecyclePolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder instantiates a new ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupLifecyclePolicies for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of lifecycle policies for this group. Read-only. Nullable.
// returns a GroupLifecyclePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupLifecyclePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable), nil
}
// Patch update the navigation property groupLifecyclePolicies in groups
// returns a GroupLifecyclePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupLifecyclePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable), nil
}
// RemoveGroup provides operations to call the removeGroup method.
// returns a *ItemGrouplifecyclepoliciesItemRemovegroupRemoveGroupRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) RemoveGroup()(*ItemGrouplifecyclepoliciesItemRemovegroupRemoveGroupRequestBuilder) {
    return NewItemGrouplifecyclepoliciesItemRemovegroupRemoveGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property groupLifecyclePolicies for groups
// returns a *RequestInformation when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of lifecycle policies for this group. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupLifecyclePolicies in groups
// returns a *RequestInformation when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) WithUrl(rawUrl string)(*ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) {
    return NewItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
