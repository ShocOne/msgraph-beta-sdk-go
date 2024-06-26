package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder provides operations to manage the intelligenceProfileIndicators property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderGetQueryParameters read the properties and relationships of a intelligenceProfileIndicator object.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderGetQueryParameters
}
// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Artifact provides operations to manage the artifact property of the microsoft.graph.security.indicator entity.
// returns a *ThreatintelligenceIntelligenceprofileindicatorsItemArtifactRequestBuilder when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) Artifact()(*ThreatintelligenceIntelligenceprofileindicatorsItemArtifactRequestBuilder) {
    return NewThreatintelligenceIntelligenceprofileindicatorsItemArtifactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal instantiates a new ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder and sets the default values.
func NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) {
    m := &ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/intelligenceProfileIndicators/{intelligenceProfileIndicator%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder instantiates a new ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder and sets the default values.
func NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property intelligenceProfileIndicators for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a intelligenceProfileIndicator object.
// returns a IntelligenceProfileIndicatorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-intelligenceprofileindicator-get?view=graph-rest-beta
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IntelligenceProfileIndicatorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateIntelligenceProfileIndicatorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IntelligenceProfileIndicatorable), nil
}
// Patch update the navigation property intelligenceProfileIndicators in security
// returns a IntelligenceProfileIndicatorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IntelligenceProfileIndicatorable, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IntelligenceProfileIndicatorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateIntelligenceProfileIndicatorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IntelligenceProfileIndicatorable), nil
}
// ToDeleteRequestInformation delete navigation property intelligenceProfileIndicators for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a intelligenceProfileIndicator object.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property intelligenceProfileIndicators in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IntelligenceProfileIndicatorable, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) {
    return NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
