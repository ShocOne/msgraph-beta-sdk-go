package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
type ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetQueryParameters the section groups in the section. Read-only. Nullable.
type ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetQueryParameters struct {
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
// ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetQueryParameters
}
// BySectionGroupId1 provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
// returns a *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupItemRequestBuilder when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) BySectionGroupId1(sectionGroupId1 string)(*ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sectionGroupId1 != "" {
        urlTplParams["sectionGroup%2Did1"] = sectionGroupId1
    }
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal instantiates a new ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    m := &ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sectionGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder instantiates a new ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsCountRequestBuilder when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) Count()(*ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsCountRequestBuilder) {
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the section groups in the section. Read-only. Nullable.
// returns a SectionGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSectionGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupCollectionResponseable), nil
}
// ToGetRequestInformation the section groups in the section. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder when successful
func (m *ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    return NewItemOnenoteNotebooksItemSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
