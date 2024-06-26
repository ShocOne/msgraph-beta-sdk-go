package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder provides operations to manage the learningCourseActivities property of the microsoft.graph.learningProvider entity.
type LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters get learningCourseActivities from employeeExperience
type LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters
}
// LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal instantiates a new LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, externalcourseActivityId *string)(*LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    m := &LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/learningProviders/{learningProvider%2Did}/learningCourseActivities(externalcourseActivityId='{externalcourseActivityId}'){?%24expand,%24select}", pathParameters),
    }
    if externalcourseActivityId != nil {
        m.BaseRequestBuilder.PathParameters["externalcourseActivityId"] = *externalcourseActivityId
    }
    return m
}
// NewLearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder instantiates a new LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete a learningCourseActivity object by using the course activity ID of either an assignment or a self-initiated activity.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningcourseactivity-delete?view=graph-rest-beta
func (m *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get learningCourseActivities from employeeExperience
// returns a LearningCourseActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) Get(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningCourseActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLearningCourseActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningCourseActivityable), nil
}
// Patch update the properties of a learningCourseActivity object. 
// returns a LearningCourseActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningcourseactivity-update?view=graph-rest-beta
func (m *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningCourseActivityable, requestConfiguration *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningCourseActivityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLearningCourseActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningCourseActivityable), nil
}
// ToDeleteRequestInformation delete a learningCourseActivity object by using the course activity ID of either an assignment or a self-initiated activity.
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get learningCourseActivities from employeeExperience
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a learningCourseActivity object. 
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LearningCourseActivityable, requestConfiguration *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder when successful
func (m *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) WithUrl(rawUrl string)(*LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    return NewLearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
