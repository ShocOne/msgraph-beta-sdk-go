package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
type FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderGetQueryParameters analytics about the view activities that took place on this item.
type FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllTime provides operations to manage the allTime property of the microsoft.graph.itemAnalytics entity.
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsAlltimeAllTimeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) AllTime()(*FilestorageContainersItemDriveItemsItemAnalyticsAlltimeAllTimeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsAlltimeAllTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/analytics{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property analytics for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get analytics about the view activities that took place on this item.
// returns a ItemAnalyticsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemAnalyticsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemAnalyticsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemAnalyticsable), nil
}
// ItemActivityStats provides operations to manage the itemActivityStats property of the microsoft.graph.itemAnalytics entity.
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) ItemActivityStats()(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivityStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastSevenDays provides operations to manage the lastSevenDays property of the microsoft.graph.itemAnalytics entity.
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) LastSevenDays()(*FilestorageContainersItemDriveItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsLastsevendaysLastSevenDaysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property analytics in storage
// returns a ItemAnalyticsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemAnalyticsable, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemAnalyticsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemAnalyticsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemAnalyticsable), nil
}
// ToDeleteRequestInformation delete navigation property analytics for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation analytics about the view activities that took place on this item.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property analytics in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemAnalyticsable, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
