package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
type ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal instantiates a new GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    m := &ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getWindowsUpdateAlertsPerPolicyPerDeviceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder instantiates a new GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) Post(ctx context.Context, body ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportPostRequestBodyable, requestConfiguration *ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportPostRequestBodyable, requestConfiguration *ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
