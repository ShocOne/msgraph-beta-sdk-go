package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder provides operations to manage the dataLabels property of the microsoft.graph.workbookChart entity.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderGetQueryParameters represents the datalabels on the chart. Read-only.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderGetQueryParameters
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/itemAt(index={index})/dataLabels{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property dataLabels for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the datalabels on the chart. Read-only.
// returns a WorkbookChartDataLabelsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartDataLabelsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartDataLabelsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartDataLabelsable), nil
}
// Patch update the navigation property dataLabels in storage
// returns a WorkbookChartDataLabelsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartDataLabelsable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartDataLabelsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartDataLabelsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartDataLabelsable), nil
}
// ToDeleteRequestInformation delete navigation property dataLabels for storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the datalabels on the chart. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property dataLabels in storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartDataLabelsable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
