package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder provides operations to manage the sourceSystem property of the microsoft.graph.industryData.industryDataConnector entity.
type IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderGetQueryParameters the sourceSystemDefinition this connector is connected to.
type IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderGetQueryParameters
}
// NewIndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderInternal instantiates a new IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder and sets the default values.
func NewIndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder) {
    m := &IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/dataConnectors/{industryDataConnector%2Did}/sourceSystem{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder instantiates a new IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder and sets the default values.
func NewIndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the sourceSystemDefinition this connector is connected to.
// returns a SourceSystemDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.SourceSystemDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateSourceSystemDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.SourceSystemDefinitionable), nil
}
// ToGetRequestInformation the sourceSystemDefinition this connector is connected to.
// returns a *RequestInformation when successful
func (m *IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder when successful
func (m *IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder) WithUrl(rawUrl string)(*IndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder) {
    return NewIndustrydataDataconnectorsItemSourcesystemSourceSystemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
