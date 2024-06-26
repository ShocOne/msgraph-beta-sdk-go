package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder provides operations to manage the filePlanReferences property of the microsoft.graph.security.labelsRoot entity.
type LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderGetQueryParameters read the properties and relationships of a filePlanReferenceTemplate object.
type LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderGetQueryParameters
}
// LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderInternal instantiates a new LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder and sets the default values.
func NewLabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) {
    m := &LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/labels/filePlanReferences/{filePlanReferenceTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder instantiates a new LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder and sets the default values.
func NewLabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property filePlanReferences for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a filePlanReferenceTemplate object.
// returns a FilePlanReferenceTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-fileplanreferencetemplate-get?view=graph-rest-beta
func (m *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanReferenceTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateFilePlanReferenceTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanReferenceTemplateable), nil
}
// Patch update the navigation property filePlanReferences in security
// returns a FilePlanReferenceTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanReferenceTemplateable, requestConfiguration *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanReferenceTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateFilePlanReferenceTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanReferenceTemplateable), nil
}
// ToDeleteRequestInformation delete navigation property filePlanReferences for security
// returns a *RequestInformation when successful
func (m *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a filePlanReferenceTemplate object.
// returns a *RequestInformation when successful
func (m *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property filePlanReferences in security
// returns a *RequestInformation when successful
func (m *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanReferenceTemplateable, requestConfiguration *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder when successful
func (m *LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) WithUrl(rawUrl string)(*LabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder) {
    return NewLabelsFileplanreferencesFilePlanReferenceTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
