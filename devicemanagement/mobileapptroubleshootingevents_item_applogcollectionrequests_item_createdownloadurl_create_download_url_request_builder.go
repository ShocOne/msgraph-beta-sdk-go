package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder provides operations to call the createDownloadUrl method.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal instantiates a new MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    m := &MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}/createDownloadUrl", pathParameters),
    }
    return m
}
// NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder instantiates a new MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createDownloadUrl
// returns a AppLogCollectionDownloadDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionDownloadDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppLogCollectionDownloadDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionDownloadDetailsable), nil
}
// ToPostRequestInformation invoke action createDownloadUrl
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*MobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
