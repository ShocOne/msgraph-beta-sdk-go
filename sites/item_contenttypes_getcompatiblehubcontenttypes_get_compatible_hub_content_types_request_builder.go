package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder provides operations to call the getCompatibleHubContentTypes method.
type ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
type ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters struct {
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
// ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetQueryParameters
}
// NewItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal instantiates a new ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    m := &ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/contentTypes/getCompatibleHubContentTypes(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder instantiates a new ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder and sets the default values.
func NewItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
// Deprecated: This method is obsolete. Use GetAsGetCompatibleHubContentTypesGetResponse instead.
// returns a ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-getcompatiblehubcontenttypes?view=graph-rest-beta
func (m *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseable), nil
}
// GetAsGetCompatibleHubContentTypesGetResponse get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
// returns a ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-getcompatiblehubcontenttypes?view=graph-rest-beta
func (m *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) GetAsGetCompatibleHubContentTypesGetResponse(ctx context.Context, requestConfiguration *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable), nil
}
// ToGetRequestInformation get compatible content types in the content type hub that can be added to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see addCopyFromContentTypeHub and the blog post Syntex Product Updates – August 2021.
// returns a *RequestInformation when successful
func (m *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder when successful
func (m *ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) WithUrl(rawUrl string)(*ItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder) {
    return NewItemContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
