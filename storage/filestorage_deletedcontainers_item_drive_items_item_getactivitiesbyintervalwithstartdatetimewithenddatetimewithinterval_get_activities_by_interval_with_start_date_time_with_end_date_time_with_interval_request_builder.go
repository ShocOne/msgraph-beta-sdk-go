package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters get itemActivityStats for the activities that took place under this resource within the specified time interval. Analytics aggregates might not be available for all action types.
type FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters struct {
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
// FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters
}
// NewFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *string, interval *string, startDateTime *string)(*FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = *endDateTime
    }
    if interval != nil {
        m.BaseRequestBuilder.PathParameters["interval"] = *interval
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = *startDateTime
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get get itemActivityStats for the activities that took place under this resource within the specified time interval. Analytics aggregates might not be available for all action types.
// Deprecated: This method is obsolete. Use GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instead.
// returns a FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/itemactivity-getbyinterval?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse get itemActivityStats for the activities that took place under this resource within the specified time interval. Analytics aggregates might not be available for all action types.
// returns a FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/itemactivity-getbyinterval?view=graph-rest-beta
func (m *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable), nil
}
// ToGetRequestInformation get itemActivityStats for the activities that took place under this resource within the specified time interval. Analytics aggregates might not be available for all action types.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
