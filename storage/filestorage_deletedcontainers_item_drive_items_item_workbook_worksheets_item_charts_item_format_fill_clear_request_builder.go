package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder provides operations to call the clear method.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/format/fill/clear", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the fill color of a chart element.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartfill-clear?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clear the fill color of a chart element.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemFormatFillClearRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
