package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder casts the previous resource to group.
type ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderGetQueryParameters get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
type ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderGetQueryParameters
}
// NewItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderInternal instantiates a new ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder and sets the default values.
func NewItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder) {
    m := &ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/transitiveMembers/{directoryObject%2Did}/graph.group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder instantiates a new ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder and sets the default values.
func NewItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-list-transitivemembers?view=graph-rest-beta
func (m *ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// ToGetRequestInformation get a list of the group's members. A group can have different object types as members. For more information about supported member types for different groups, see Group membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an OData cast that represents an unsupported member type returns a 400 Bad Request error with the Request_UnsupportedQuery code.
// returns a *RequestInformation when successful
func (m *ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder) WithUrl(rawUrl string)(*ItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder) {
    return NewItemTransitivemembersItemGraphgroupGraphGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
