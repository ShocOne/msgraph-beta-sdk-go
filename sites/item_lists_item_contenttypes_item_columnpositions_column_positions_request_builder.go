package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
type ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderGetQueryParameters column order information in a content type.
type ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderGetQueryParameters struct {
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
// ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderGetQueryParameters
}
// ByColumnDefinitionId provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
// returns a *ItemListsItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder when successful
func (m *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) ByColumnDefinitionId(columnDefinitionId string)(*ItemListsItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if columnDefinitionId != "" {
        urlTplParams["columnDefinition%2Did"] = columnDefinitionId
    }
    return NewItemListsItemContenttypesItemColumnpositionsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderInternal instantiates a new ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder and sets the default values.
func NewItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) {
    m := &ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/columnPositions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder instantiates a new ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder and sets the default values.
func NewItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemListsItemContenttypesItemColumnpositionsCountRequestBuilder when successful
func (m *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) Count()(*ItemListsItemContenttypesItemColumnpositionsCountRequestBuilder) {
    return NewItemListsItemContenttypesItemColumnpositionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get column order information in a content type.
// returns a ColumnDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ColumnDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateColumnDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ColumnDefinitionCollectionResponseable), nil
}
// ToGetRequestInformation column order information in a content type.
// returns a *RequestInformation when successful
func (m *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder when successful
func (m *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) {
    return NewItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
