package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder provides operations to call the allowNextEnrollment method.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderInternal instantiates a new WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder and sets the default values.
func NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) {
    m := &WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}/assignedDevices/{windowsAutopilotDeviceIdentity%2Did}/allowNextEnrollment", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder instantiates a new WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder and sets the default values.
func NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unblocks next autopilot enrollment.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) Post(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation unblocks next autopilot enrollment.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
