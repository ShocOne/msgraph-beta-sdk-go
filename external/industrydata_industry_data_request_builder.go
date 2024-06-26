package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataIndustryDataRequestBuilder provides operations to manage the industryData property of the microsoft.graph.externalConnectors.external entity.
type IndustrydataIndustryDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataIndustryDataRequestBuilderGetQueryParameters get industryData from external
type IndustrydataIndustryDataRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataIndustryDataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataIndustryDataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataIndustryDataRequestBuilderGetQueryParameters
}
// NewIndustrydataIndustryDataRequestBuilderInternal instantiates a new IndustrydataIndustryDataRequestBuilder and sets the default values.
func NewIndustrydataIndustryDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataIndustryDataRequestBuilder) {
    m := &IndustrydataIndustryDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataIndustryDataRequestBuilder instantiates a new IndustrydataIndustryDataRequestBuilder and sets the default values.
func NewIndustrydataIndustryDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataIndustryDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataIndustryDataRequestBuilderInternal(urlParams, requestAdapter)
}
// DataConnectors provides operations to manage the dataConnectors property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataDataconnectorsDataConnectorsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) DataConnectors()(*IndustrydataDataconnectorsDataConnectorsRequestBuilder) {
    return NewIndustrydataDataconnectorsDataConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get industryData from external
// returns a IndustryDataRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataIndustryDataRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataIndustryDataRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRootable), nil
}
// InboundFlows provides operations to manage the inboundFlows property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataInboundflowsInboundFlowsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) InboundFlows()(*IndustrydataInboundflowsInboundFlowsRequestBuilder) {
    return NewIndustrydataInboundflowsInboundFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataOperationsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) Operations()(*IndustrydataOperationsRequestBuilder) {
    return NewIndustrydataOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OutboundProvisioningFlowSets provides operations to manage the outboundProvisioningFlowSets property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) OutboundProvisioningFlowSets()(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReferenceDefinitions provides operations to manage the referenceDefinitions property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) ReferenceDefinitions()(*IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) {
    return NewIndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleGroups provides operations to manage the roleGroups property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataRolegroupsRoleGroupsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) RoleGroups()(*IndustrydataRolegroupsRoleGroupsRequestBuilder) {
    return NewIndustrydataRolegroupsRoleGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runs provides operations to manage the runs property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataRunsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) Runs()(*IndustrydataRunsRequestBuilder) {
    return NewIndustrydataRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SourceSystems provides operations to manage the sourceSystems property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataSourcesystemsSourceSystemsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) SourceSystems()(*IndustrydataSourcesystemsSourceSystemsRequestBuilder) {
    return NewIndustrydataSourcesystemsSourceSystemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get industryData from external
// returns a *RequestInformation when successful
func (m *IndustrydataIndustryDataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataIndustryDataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataIndustryDataRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) WithUrl(rawUrl string)(*IndustrydataIndustryDataRequestBuilder) {
    return NewIndustrydataIndustryDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Years provides operations to manage the years property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataYearsRequestBuilder when successful
func (m *IndustrydataIndustryDataRequestBuilder) Years()(*IndustrydataYearsRequestBuilder) {
    return NewIndustrydataYearsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
