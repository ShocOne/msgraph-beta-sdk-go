package identitygovernance

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the topTasksProcessedSummary method.
type LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters provide a summary of the most processed tasks, known as top tasks, for a specified time period in a tenant. The task definition is provided, along with numerical counts of total, successful, and failed runs. For information about workflows processed, see insights: topWorkflowsProcessedSummary.
type LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
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
// LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/insights/microsoft.graph.identityGovernance.topTasksProcessedSummary(startDateTime={startDateTime},endDateTime={endDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get provide a summary of the most processed tasks, known as top tasks, for a specified time period in a tenant. The task definition is provided, along with numerical counts of total, successful, and failed runs. For information about workflows processed, see insights: topWorkflowsProcessedSummary.
// Deprecated: This method is obsolete. Use GetAsTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-insights-toptasksprocessedsummary?view=graph-rest-beta
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse provide a summary of the most processed tasks, known as top tasks, for a specified time period in a tenant. The task definition is provided, along with numerical counts of total, successful, and failed runs. For information about workflows processed, see insights: topWorkflowsProcessedSummary.
// returns a LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-insights-toptasksprocessedsummary?view=graph-rest-beta
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation provide a summary of the most processed tasks, known as top tasks, for a specified time period in a tenant. The task definition is provided, along with numerical counts of total, successful, and failed runs. For information about workflows processed, see insights: topWorkflowsProcessedSummary.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewLifecycleworkflowsInsightsMicrosoftgraphidentitygovernancetoptasksprocessedsummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
