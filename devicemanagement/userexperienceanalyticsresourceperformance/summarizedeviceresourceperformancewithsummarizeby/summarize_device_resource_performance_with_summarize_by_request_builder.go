package summarizedeviceresourceperformancewithsummarizeby

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsResourcePerformance\microsoft.graph.summarizeDeviceResourcePerformance(summarizeBy={summarizeBy})
type SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions options for Get
type SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal instantiates a new SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, summarizeBy *string)(*SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    m := &SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsResourcePerformance/microsoft.graph.summarizeDeviceResourcePerformance(summarizeBy={summarizeBy})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if summarizeBy != nil {
        urlTplParams["summarizeBy"] = *summarizeBy
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder instantiates a new SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function summarizeDeviceResourcePerformance
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) CreateGetRequestInformation(options *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function summarizeDeviceResourcePerformance
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) Get(options *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions)([]SummarizeDeviceResourcePerformanceWithSummarizeBy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSummarizeDeviceResourcePerformanceWithSummarizeBy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]SummarizeDeviceResourcePerformanceWithSummarizeBy, len(res))
    for i, v := range res {
        val[i] = *(v.(*SummarizeDeviceResourcePerformanceWithSummarizeBy))
    }
    return val, nil
}
