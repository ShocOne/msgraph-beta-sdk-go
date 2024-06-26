package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder provides operations to manage the whoisRecords property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderGetQueryParameters get the specified whoisRecord resource.  Specify the desired whoisRecord in one of the following two ways:- Identify a host and get its current whoisRecord. - Specify an id value to get the corresponding whoisRecord.
type ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderGetQueryParameters
}
// ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderInternal instantiates a new ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder and sets the default values.
func NewThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) {
    m := &ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/whoisRecords/{whoisRecord%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder instantiates a new ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder and sets the default values.
func NewThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property whoisRecords for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the specified whoisRecord resource.  Specify the desired whoisRecord in one of the following two ways:- Identify a host and get its current whoisRecord. - Specify an id value to get the corresponding whoisRecord.
// returns a WhoisRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-whoisrecord-get?view=graph-rest-beta
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.WhoisRecordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateWhoisRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.WhoisRecordable), nil
}
// History provides operations to manage the history property of the microsoft.graph.security.whoisRecord entity.
// returns a *ThreatintelligenceWhoisrecordsItemHistoryRequestBuilder when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) History()(*ThreatintelligenceWhoisrecordsItemHistoryRequestBuilder) {
    return NewThreatintelligenceWhoisrecordsItemHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Host provides operations to manage the host property of the microsoft.graph.security.whoisBaseRecord entity.
// returns a *ThreatintelligenceWhoisrecordsItemHostRequestBuilder when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) Host()(*ThreatintelligenceWhoisrecordsItemHostRequestBuilder) {
    return NewThreatintelligenceWhoisrecordsItemHostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property whoisRecords in security
// returns a WhoisRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.WhoisRecordable, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.WhoisRecordable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateWhoisRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.WhoisRecordable), nil
}
// ToDeleteRequestInformation delete navigation property whoisRecords for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the specified whoisRecord resource.  Specify the desired whoisRecord in one of the following two ways:- Identify a host and get its current whoisRecord. - Specify an id value to get the corresponding whoisRecord.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property whoisRecords in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.WhoisRecordable, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) {
    return NewThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
