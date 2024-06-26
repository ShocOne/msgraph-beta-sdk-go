package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder provides operations to call the unenrollAssets method.
type WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    m := &WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/members/microsoft.graph.windowsUpdates.unenrollAssets", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder instantiates a new WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unenrollAssets
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsUnenrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action unenrollAssets
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsUnenrollAssetsPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder when successful
func (m *WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder) {
    return NewWindowsUpdatesDeploymentaudiencesItemMembersMicrosoftgraphwindowsupdatesunenrollassetsMicrosoftGraphWindowsUpdatesUnenrollAssetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
