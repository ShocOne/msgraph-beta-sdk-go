package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder provides operations to call the applyTags method.
type EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilderInternal instantiates a new EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder) {
    m := &EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}/queries/{reviewSetQuery%2Did}/microsoft.graph.ediscovery.applyTags", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder instantiates a new EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post apply tags to documents that match the specified reviewSetQuery.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-reviewsetquery-applytags?view=graph-rest-beta
func (m *EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder) Post(ctx context.Context, body EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsApplyTagsPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation apply tags to documents that match the specified reviewSetQuery.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsApplyTagsPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder when successful
func (m *EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder) {
    return NewEdiscoveryCasesItemReviewsetsItemQueriesItemMicrosoftgraphediscoveryapplytagsMicrosoftGraphEdiscoveryApplyTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
