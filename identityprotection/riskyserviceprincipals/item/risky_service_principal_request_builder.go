package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ic3a4e720ce4fba459cd710019c0568f6809117d535d02723acc89f2098374e1f "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/item/history"
    i1afb9bfc08016ee30523a279a0207fcdef5578927857e36dadd98df8eb7a4001 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/item/history/item"
)

// RiskyServicePrincipalRequestBuilder builds and executes requests for operations under \identityProtection\riskyServicePrincipals\{riskyServicePrincipal-id}
type RiskyServicePrincipalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RiskyServicePrincipalRequestBuilderDeleteOptions options for Delete
type RiskyServicePrincipalRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RiskyServicePrincipalRequestBuilderGetOptions options for Get
type RiskyServicePrincipalRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RiskyServicePrincipalRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RiskyServicePrincipalRequestBuilderGetQueryParameters azure AD service principals that are at risk.
type RiskyServicePrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RiskyServicePrincipalRequestBuilderPatchOptions options for Patch
type RiskyServicePrincipalRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipal;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRiskyServicePrincipalRequestBuilderInternal instantiates a new RiskyServicePrincipalRequestBuilder and sets the default values.
func NewRiskyServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyServicePrincipalRequestBuilder) {
    m := &RiskyServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyServicePrincipals/{riskyServicePrincipal_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRiskyServicePrincipalRequestBuilder instantiates a new RiskyServicePrincipalRequestBuilder and sets the default values.
func NewRiskyServicePrincipalRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation azure AD service principals that are at risk.
func (m *RiskyServicePrincipalRequestBuilder) CreateDeleteRequestInformation(options *RiskyServicePrincipalRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation azure AD service principals that are at risk.
func (m *RiskyServicePrincipalRequestBuilder) CreateGetRequestInformation(options *RiskyServicePrincipalRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation azure AD service principals that are at risk.
func (m *RiskyServicePrincipalRequestBuilder) CreatePatchRequestInformation(options *RiskyServicePrincipalRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete azure AD service principals that are at risk.
func (m *RiskyServicePrincipalRequestBuilder) Delete(options *RiskyServicePrincipalRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get azure AD service principals that are at risk.
func (m *RiskyServicePrincipalRequestBuilder) Get(options *RiskyServicePrincipalRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipal, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRiskyServicePrincipal() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RiskyServicePrincipal), nil
}
func (m *RiskyServicePrincipalRequestBuilder) History()(*ic3a4e720ce4fba459cd710019c0568f6809117d535d02723acc89f2098374e1f.HistoryRequestBuilder) {
    return ic3a4e720ce4fba459cd710019c0568f6809117d535d02723acc89f2098374e1f.NewHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HistoryById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProtection.riskyServicePrincipals.item.history.item collection
func (m *RiskyServicePrincipalRequestBuilder) HistoryById(id string)(*i1afb9bfc08016ee30523a279a0207fcdef5578927857e36dadd98df8eb7a4001.RiskyServicePrincipalHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyServicePrincipalHistoryItem_id"] = id
    }
    return i1afb9bfc08016ee30523a279a0207fcdef5578927857e36dadd98df8eb7a4001.NewRiskyServicePrincipalHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch azure AD service principals that are at risk.
func (m *RiskyServicePrincipalRequestBuilder) Patch(options *RiskyServicePrincipalRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}