package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder provides operations to call the createSnapshot method.
type VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/createSnapshot", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder instantiates a new VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a snapshot for a specific Cloud PC device.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-createsnapshot?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder) Post(ctx context.Context, body VirtualEndpointCloudPCsItemCreateSnapshotPostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation create a snapshot for a specific Cloud PC device.
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointCloudPCsItemCreateSnapshotPostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder when successful
func (m *VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
