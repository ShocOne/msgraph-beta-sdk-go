package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    icffa63881025e30e94ee621d689ad739df1d6c5f54801abebaf0e6571205dd01 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/publishedresources"
    ie7fe4ccfff052401443978ac091cbae88e1cf14e4ff2071c64d21c0152352ec0 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/agents"
    i2b3dc5b3beae25c25ad041275e164179676fcb837611ad79911642fba5228ebe "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/publishedresources/item"
    i5c382280fd562ccfaf5ac4114acc2547213ff1b2a4b04ae1dedc167accbe54b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/agents/item"
)

// OnPremisesAgentGroupItemRequestBuilder provides operations to manage the agentGroups property of the microsoft.graph.publishedResource entity.
type OnPremisesAgentGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnPremisesAgentGroupItemRequestBuilderDeleteOptions options for Delete
type OnPremisesAgentGroupItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnPremisesAgentGroupItemRequestBuilderGetOptions options for Get
type OnPremisesAgentGroupItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnPremisesAgentGroupItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnPremisesAgentGroupItemRequestBuilderGetQueryParameters list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
type OnPremisesAgentGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnPremisesAgentGroupItemRequestBuilderPatchOptions options for Patch
type OnPremisesAgentGroupItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesAgentGroupable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OnPremisesAgentGroupItemRequestBuilder) Agents()(*ie7fe4ccfff052401443978ac091cbae88e1cf14e4ff2071c64d21c0152352ec0.AgentsRequestBuilder) {
    return ie7fe4ccfff052401443978ac091cbae88e1cf14e4ff2071c64d21c0152352ec0.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.publishedResources.item.agentGroups.item.agents.item collection
func (m *OnPremisesAgentGroupItemRequestBuilder) AgentsById(id string)(*i5c382280fd562ccfaf5ac4114acc2547213ff1b2a4b04ae1dedc167accbe54b3.OnPremisesAgentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent_id"] = id
    }
    return i5c382280fd562ccfaf5ac4114acc2547213ff1b2a4b04ae1dedc167accbe54b3.NewOnPremisesAgentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesAgentGroupItemRequestBuilder) {
    m := &OnPremisesAgentGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile_id}/publishedResources/{publishedResource_id}/agentGroups/{onPremisesAgentGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesAgentGroupItemRequestBuilder instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property agentGroups for onPremisesPublishingProfiles
func (m *OnPremisesAgentGroupItemRequestBuilder) CreateDeleteRequestInformation(options *OnPremisesAgentGroupItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupItemRequestBuilder) CreateGetRequestInformation(options *OnPremisesAgentGroupItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property agentGroups in onPremisesPublishingProfiles
func (m *OnPremisesAgentGroupItemRequestBuilder) CreatePatchRequestInformation(options *OnPremisesAgentGroupItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property agentGroups for onPremisesPublishingProfiles
func (m *OnPremisesAgentGroupItemRequestBuilder) Delete(options *OnPremisesAgentGroupItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupItemRequestBuilder) Get(options *OnPremisesAgentGroupItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesAgentGroupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOnPremisesAgentGroupFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesAgentGroupable), nil
}
// Patch update the navigation property agentGroups in onPremisesPublishingProfiles
func (m *OnPremisesAgentGroupItemRequestBuilder) Patch(options *OnPremisesAgentGroupItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *OnPremisesAgentGroupItemRequestBuilder) PublishedResources()(*icffa63881025e30e94ee621d689ad739df1d6c5f54801abebaf0e6571205dd01.PublishedResourcesRequestBuilder) {
    return icffa63881025e30e94ee621d689ad739df1d6c5f54801abebaf0e6571205dd01.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.publishedResources.item.agentGroups.item.publishedResources.item collection
func (m *OnPremisesAgentGroupItemRequestBuilder) PublishedResourcesById(id string)(*i2b3dc5b3beae25c25ad041275e164179676fcb837611ad79911642fba5228ebe.PublishedResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource_id1"] = id
    }
    return i2b3dc5b3beae25c25ad041275e164179676fcb837611ad79911642fba5228ebe.NewPublishedResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}