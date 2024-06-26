package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder casts the previous resource to device.
type ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderGetQueryParameters get the items of type microsoft.graph.device in the microsoft.graph.directoryObject collection
type ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderGetQueryParameters struct {
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
// ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderGetQueryParameters
}
// NewItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderInternal instantiates a new ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder and sets the default values.
func NewItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) {
    m := &ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/transitiveMembers/graph.device{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder instantiates a new ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder and sets the default values.
func NewItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTransitivemembersGraphdeviceCountRequestBuilder when successful
func (m *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) Count()(*ItemTransitivemembersGraphdeviceCountRequestBuilder) {
    return NewItemTransitivemembersGraphdeviceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.device in the microsoft.graph.directoryObject collection
// returns a DeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.device in the microsoft.graph.directoryObject collection
// returns a *RequestInformation when successful
func (m *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder when successful
func (m *ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) WithUrl(rawUrl string)(*ItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder) {
    return NewItemTransitivemembersGraphdeviceGraphDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
