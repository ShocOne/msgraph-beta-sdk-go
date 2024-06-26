package networkaccess

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder provides operations to call the getDestinationSummaries method.
type ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters get counts of the visits to the top destination aggregations.
type ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters
}
// NewReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal instantiates a new ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder and sets the default values.
func NewReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, aggregatedBy *string, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    m := &ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/reports/microsoft.graph.networkaccess.getDestinationSummaries(startDateTime={startDateTime},endDateTime={endDateTime},aggregatedBy='{aggregatedBy}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if aggregatedBy != nil {
        m.BaseRequestBuilder.PathParameters["aggregatedBy"] = *aggregatedBy
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder instantiates a new ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder and sets the default values.
func NewReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get get counts of the visits to the top destination aggregations.
// Deprecated: This method is obsolete. Use GetAsGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse instead.
// returns a ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable), nil
}
// GetAsGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse get counts of the visits to the top destination aggregations.
// returns a ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) GetAsGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse(ctx context.Context, requestConfiguration *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable), nil
}
// ToGetRequestInformation get counts of the visits to the top destination aggregations.
// returns a *RequestInformation when successful
func (m *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder when successful
func (m *ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) WithUrl(rawUrl string)(*ReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    return NewReportsMicrosoftgraphnetworkaccessgetdestinationsummarieswithstartdatetimewithenddatetimewithaggregatedbyMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
