package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Axes provides operations to manage the axes property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexAxesRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Axes()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexAxesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexAxesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// DataLabels provides operations to manage the dataLabels property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) DataLabels()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexDatalabelsDataLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Format provides operations to manage the format property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexFormatRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Format()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexFormatRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get gets a chart based on its position in the collection.
// returns a WorkbookChartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartcollection-itemat?view=graph-rest-beta
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookChartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookChartable), nil
}
// Image provides operations to call the image method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Image()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImageWithWidth provides operations to call the image method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) ImageWithWidth(width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, width)
}
// ImageWithWidthWithHeight provides operations to call the image method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) ImageWithWidthWithHeight(height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, height, width)
}
// ImageWithWidthWithHeightWithFittingMode provides operations to call the image method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) ImageWithWidthWithHeightWithFittingMode(fittingMode *string, height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, fittingMode, height, width)
}
// Legend provides operations to manage the legend property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexLegendRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Legend()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexLegendRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexLegendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Series provides operations to manage the series property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSeriesRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Series()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSeriesRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSeriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetData provides operations to call the setData method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSetdataSetDataRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) SetData()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSetdataSetDataRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSetdataSetDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetPosition provides operations to call the setPosition method.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSetpositionSetPositionRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) SetPosition()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSetpositionSetPositionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexSetpositionSetPositionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Title provides operations to manage the title property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexTitleRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Title()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexTitleRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexTitleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a chart based on its position in the collection.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookChart entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexWorksheetRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexItemAtWithIndexRequestBuilder) Worksheet()(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexWorksheetRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
