package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder casts the previous resource to endpoint.
type ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.endpoint
type ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderGetQueryParameters
}
// NewItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderInternal instantiates a new ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) {
    m := &ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}/graph.endpoint{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder instantiates a new ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.endpoint
// returns a Endpointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Endpointable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Endpointable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.endpoint
// returns a *RequestInformation when successful
func (m *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) WithUrl(rawUrl string)(*ItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) {
    return NewItemDevicesItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
