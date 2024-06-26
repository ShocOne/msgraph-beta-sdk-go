package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder provides operations to call the getShiftWorkCloudPcAccessState method.
type ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderInternal instantiates a new ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) {
    m := &ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/getShiftWorkCloudPcAccessState()", pathParameters),
    }
    return m
}
// NewItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder instantiates a new ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the shiftWorkCloudPcAccessState of a shift work Cloud PC.  This API only supports shared-use licenses. For more information, see cloudPcProvisioningPolicy. Shared-use licenses allow three users per license, with one user signed in at a time. Callers can get the latest shift work Cloud PC accessState and determine whether the shift work Cloud PC is accessible to the user.  If a web client needs to connect to a shift work Cloud PC, the sharedCloudPcAccessState validates the bookmark scenario. If sharedCloudPcAccessState is not active/activating/standbyMode, the web client shows a 'bad bookmark'.
// Deprecated: The getShiftWorkCloudPcAccessState API is deprecated and will stop returning data on Dec 31, 2023. Please use the new getFrontlineCloudPcAccessState API as of 2023-08/getShiftWorkCloudPcAccessState
// returns a *ShiftWorkCloudPcAccessState when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getshiftworkcloudpcaccessstate?view=graph-rest-beta
func (m *ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftWorkCloudPcAccessState, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendEnum(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseShiftWorkCloudPcAccessState, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ShiftWorkCloudPcAccessState), nil
}
// ToGetRequestInformation get the shiftWorkCloudPcAccessState of a shift work Cloud PC.  This API only supports shared-use licenses. For more information, see cloudPcProvisioningPolicy. Shared-use licenses allow three users per license, with one user signed in at a time. Callers can get the latest shift work Cloud PC accessState and determine whether the shift work Cloud PC is accessible to the user.  If a web client needs to connect to a shift work Cloud PC, the sharedCloudPcAccessState validates the bookmark scenario. If sharedCloudPcAccessState is not active/activating/standbyMode, the web client shows a 'bad bookmark'.
// Deprecated: The getShiftWorkCloudPcAccessState API is deprecated and will stop returning data on Dec 31, 2023. Please use the new getFrontlineCloudPcAccessState API as of 2023-08/getShiftWorkCloudPcAccessState
// returns a *RequestInformation when successful
func (m *ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The getShiftWorkCloudPcAccessState API is deprecated and will stop returning data on Dec 31, 2023. Please use the new getFrontlineCloudPcAccessState API as of 2023-08/getShiftWorkCloudPcAccessState
// returns a *ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder when successful
func (m *ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
