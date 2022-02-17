package applepushnotificationcertificate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i318cff614ec00a7f91c68735e78b67b4815e16831a97d11833f74f2ad4a015f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/applepushnotificationcertificate/generateapplepushnotificationcertificatesigningrequest"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i77364310cf65cdd95360b454958409da5533146c0b4983b7b0838edea305420c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/applepushnotificationcertificate/downloadapplepushnotificationcertificatesigningrequest"
)

// ApplePushNotificationCertificateRequestBuilder builds and executes requests for operations under \deviceManagement\applePushNotificationCertificate
type ApplePushNotificationCertificateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ApplePushNotificationCertificateRequestBuilderDeleteOptions options for Delete
type ApplePushNotificationCertificateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApplePushNotificationCertificateRequestBuilderGetOptions options for Get
type ApplePushNotificationCertificateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ApplePushNotificationCertificateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApplePushNotificationCertificateRequestBuilderGetQueryParameters apple push notification certificate.
type ApplePushNotificationCertificateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ApplePushNotificationCertificateRequestBuilderPatchOptions options for Patch
type ApplePushNotificationCertificateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApplePushNotificationCertificate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewApplePushNotificationCertificateRequestBuilderInternal instantiates a new ApplePushNotificationCertificateRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplePushNotificationCertificateRequestBuilder) {
    m := &ApplePushNotificationCertificateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/applePushNotificationCertificate{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplePushNotificationCertificateRequestBuilder instantiates a new ApplePushNotificationCertificateRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplePushNotificationCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplePushNotificationCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation apple push notification certificate.
func (m *ApplePushNotificationCertificateRequestBuilder) CreateDeleteRequestInformation(options *ApplePushNotificationCertificateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation apple push notification certificate.
func (m *ApplePushNotificationCertificateRequestBuilder) CreateGetRequestInformation(options *ApplePushNotificationCertificateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation apple push notification certificate.
func (m *ApplePushNotificationCertificateRequestBuilder) CreatePatchRequestInformation(options *ApplePushNotificationCertificateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete apple push notification certificate.
func (m *ApplePushNotificationCertificateRequestBuilder) Delete(options *ApplePushNotificationCertificateRequestBuilderDeleteOptions)(error) {
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
// DownloadApplePushNotificationCertificateSigningRequest builds and executes requests for operations under \deviceManagement\applePushNotificationCertificate\microsoft.graph.downloadApplePushNotificationCertificateSigningRequest()
func (m *ApplePushNotificationCertificateRequestBuilder) DownloadApplePushNotificationCertificateSigningRequest()(*i77364310cf65cdd95360b454958409da5533146c0b4983b7b0838edea305420c.DownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    return i77364310cf65cdd95360b454958409da5533146c0b4983b7b0838edea305420c.NewDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplePushNotificationCertificateRequestBuilder) GenerateApplePushNotificationCertificateSigningRequest()(*i318cff614ec00a7f91c68735e78b67b4815e16831a97d11833f74f2ad4a015f2.GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    return i318cff614ec00a7f91c68735e78b67b4815e16831a97d11833f74f2ad4a015f2.NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get apple push notification certificate.
func (m *ApplePushNotificationCertificateRequestBuilder) Get(options *ApplePushNotificationCertificateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApplePushNotificationCertificate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewApplePushNotificationCertificate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApplePushNotificationCertificate), nil
}
// Patch apple push notification certificate.
func (m *ApplePushNotificationCertificateRequestBuilder) Patch(options *ApplePushNotificationCertificateRequestBuilderPatchOptions)(error) {
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
