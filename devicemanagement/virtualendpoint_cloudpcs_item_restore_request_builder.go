package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemRestoreRequestBuilder provides operations to call the restore method.
type VirtualendpointCloudpcsItemRestoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemRestoreRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemRestoreRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemRestoreRequestBuilder) {
    m := &VirtualendpointCloudpcsItemRestoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/restore", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemRestoreRequestBuilder instantiates a new VirtualendpointCloudpcsItemRestoreRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore a specific Cloud PC. Use this API to trigger a remote action that restores a Cloud PC device to a previous state.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-restore?view=graph-rest-beta
func (m *VirtualendpointCloudpcsItemRestoreRequestBuilder) Post(ctx context.Context, body VirtualendpointCloudpcsItemRestorePostRequestBodyable, requestConfiguration *VirtualendpointCloudpcsItemRestoreRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation restore a specific Cloud PC. Use this API to trigger a remote action that restores a Cloud PC device to a previous state.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemRestoreRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointCloudpcsItemRestorePostRequestBodyable, requestConfiguration *VirtualendpointCloudpcsItemRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *VirtualendpointCloudpcsItemRestoreRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemRestoreRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemRestoreRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRestoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
