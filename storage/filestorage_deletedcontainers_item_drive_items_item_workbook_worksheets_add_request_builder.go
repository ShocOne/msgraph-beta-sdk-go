package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder provides operations to call the add method.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/add", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add a new worksheet to the workbook. The worksheet is added at the end of existing worksheets. If you want to activate the newly added worksheet, call '.activate() on it.
// returns a WorkbookWorksheetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/worksheetcollection-add?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder) Post(ctx context.Context, body FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddPostRequestBodyable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookWorksheetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookWorksheetable), nil
}
// ToPostRequestInformation add a new worksheet to the workbook. The worksheet is added at the end of existing worksheets. If you want to activate the newly added worksheet, call '.activate() on it.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddPostRequestBodyable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
