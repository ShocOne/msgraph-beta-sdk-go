package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder provides operations to call the sendReminder method.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal instantiates a new SendReminderRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    m := &PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/microsoft.graph.sendReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder instantiates a new SendReminderRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send a reminder to the reviewers of a currently active accessReviewInstance.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreviewinstance-sendreminder?view=graph-rest-1.0
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation send a reminder to the reviewers of a currently active accessReviewInstance.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
