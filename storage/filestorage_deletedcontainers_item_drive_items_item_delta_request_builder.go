package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder provides operations to call the delta method.
type FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetQueryParameters track changes in a driveItem and its children over time. Your app begins by calling delta without any parameters.The service starts enumerating the drive's hierarchy, returning pages of items and either an @odata.nextLink or an @odata.deltaLink.Your app should continue calling with the @odata.nextLink until you no longer see an @odata.nextLink returned, or you see a response with an empty set of changes. After you finish receiving all the changes, you may apply them to your local state.To check for changes in the future, call delta again with the @odata.deltaLink from the previous response. Deleted items are returned with the deleted facet.Items with this property set should be removed from your local state. Note: you should only delete a folder locally if it's empty after syncing all the changes.
type FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetQueryParameters struct {
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
// FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetQueryParameters
}
// NewFilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get track changes in a driveItem and its children over time. Your app begins by calling delta without any parameters.The service starts enumerating the drive's hierarchy, returning pages of items and either an @odata.nextLink or an @odata.deltaLink.Your app should continue calling with the @odata.nextLink until you no longer see an @odata.nextLink returned, or you see a response with an empty set of changes. After you finish receiving all the changes, you may apply them to your local state.To check for changes in the future, call delta again with the @odata.deltaLink from the previous response. Deleted items are returned with the deleted facet.Items with this property set should be removed from your local state. Note: you should only delete a folder locally if it's empty after syncing all the changes.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a FilestorageDeletedcontainersItemDriveItemsItemDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-delta?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveItemsItemDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveItemsItemDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveItemsItemDeltaResponseable), nil
}
// GetAsDeltaGetResponse track changes in a driveItem and its children over time. Your app begins by calling delta without any parameters.The service starts enumerating the drive's hierarchy, returning pages of items and either an @odata.nextLink or an @odata.deltaLink.Your app should continue calling with the @odata.nextLink until you no longer see an @odata.nextLink returned, or you see a response with an empty set of changes. After you finish receiving all the changes, you may apply them to your local state.To check for changes in the future, call delta again with the @odata.deltaLink from the previous response. Deleted items are returned with the deleted facet.Items with this property set should be removed from your local state. Note: you should only delete a folder locally if it's empty after syncing all the changes.
// returns a FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-delta?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponseable), nil
}
// ToGetRequestInformation track changes in a driveItem and its children over time. Your app begins by calling delta without any parameters.The service starts enumerating the drive's hierarchy, returning pages of items and either an @odata.nextLink or an @odata.deltaLink.Your app should continue calling with the @odata.nextLink until you no longer see an @odata.nextLink returned, or you see a response with an empty set of changes. After you finish receiving all the changes, you may apply them to your local state.To check for changes in the future, call delta again with the @odata.deltaLink from the previous response. Deleted items are returned with the deleted facet.Items with this property set should be removed from your local state. Note: you should only delete a folder locally if it's empty after syncing all the changes.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
