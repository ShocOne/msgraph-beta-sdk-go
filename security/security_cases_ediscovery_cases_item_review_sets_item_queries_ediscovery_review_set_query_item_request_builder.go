package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder provides operations to manage the queries property of the microsoft.graph.security.ediscoveryReviewSet entity.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters represents queries within the review set.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters
}
// SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyTags provides operations to call the applyTags method.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) ApplyTags()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemApplyTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderInternal instantiates a new EdiscoveryReviewSetQueryItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) {
    m := &SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder instantiates a new EdiscoveryReviewSetQueryItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property queries for security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents queries within the review set.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property queries in security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property queries for security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Export provides operations to call the export method.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Export()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents queries within the review set.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewSetQueryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable), nil
}
// Patch update the navigation property queries in security
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, requestConfiguration *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewSetQueryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable), nil
}
// Run provides operations to call the run method.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Run()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemRunRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemRunRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
