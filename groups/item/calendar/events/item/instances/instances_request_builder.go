package instances

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic628c21df8325e0560a058660018b708ef3e6a45eb28b137a19d537c9c59e6d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/delta"
    idee23541c117d5d9212906623c90e629a1effb2d86b105001d1e1da25fb99184 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/count"
)

// InstancesRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type InstancesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InstancesRequestBuilderGetQueryParameters get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns theoccurrences and exceptions of the event in the specified time range.
type InstancesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// InstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstancesRequestBuilderGetQueryParameters
}
// NewInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    m := &InstancesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/instances{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *InstancesRequestBuilder) Count()(*idee23541c117d5d9212906623c90e629a1effb2d86b105001d1e1da25fb99184.CountRequestBuilder) {
    return idee23541c117d5d9212906623c90e629a1effb2d86b105001d1e1da25fb99184.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns theoccurrences and exceptions of the event in the specified time range.
func (m *InstancesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *InstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *InstancesRequestBuilder) Delta()(*ic628c21df8325e0560a058660018b708ef3e6a45eb28b137a19d537c9c59e6d0.DeltaRequestBuilder) {
    return ic628c21df8325e0560a058660018b708ef3e6a45eb28b137a19d537c9c59e6d0.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns theoccurrences and exceptions of the event in the specified time range.
func (m *InstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *InstancesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable), nil
}
