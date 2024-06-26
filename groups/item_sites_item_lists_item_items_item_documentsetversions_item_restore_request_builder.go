package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder provides operations to call the restore method.
type ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderInternal instantiates a new ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) {
    m := &ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/documentSetVersions/{documentSetVersion%2Did}/restore", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder instantiates a new ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a document set version.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/documentsetversion-restore?view=graph-rest-beta
func (m *ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation restore a document set version.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder when successful
func (m *ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemDocumentsetversionsItemRestoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
