package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder provides operations to call the retrieveReviewStatus method.
type VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/retrieveReviewStatus()", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder instantiates a new VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveReviewStatus
// returns a CloudPcReviewStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReviewStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable), nil
}
// ToGetRequestInformation invoke function retrieveReviewStatus
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder when successful
func (m *VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemRetrieveReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
