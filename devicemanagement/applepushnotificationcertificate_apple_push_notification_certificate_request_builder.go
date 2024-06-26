package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder provides operations to manage the applePushNotificationCertificate property of the microsoft.graph.deviceManagement entity.
type ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderGetQueryParameters apple push notification certificate.
type ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderGetQueryParameters
}
// ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderInternal instantiates a new ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder and sets the default values.
func NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) {
    m := &ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/applePushNotificationCertificate{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder instantiates a new ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder and sets the default values.
func NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property applePushNotificationCertificate for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// DownloadApplePushNotificationCertificateSigningRequest provides operations to call the downloadApplePushNotificationCertificateSigningRequest method.
// returns a *ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder when successful
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) DownloadApplePushNotificationCertificateSigningRequest()(*ApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    return NewApplepushnotificationcertificateDownloadapplepushnotificationcertificatesigningrequestDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateApplePushNotificationCertificateSigningRequest provides operations to call the generateApplePushNotificationCertificateSigningRequest method.
// returns a *ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestRequestBuilder when successful
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) GenerateApplePushNotificationCertificateSigningRequest()(*ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    return NewApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get apple push notification certificate.
// returns a ApplePushNotificationCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplePushNotificationCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplePushNotificationCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplePushNotificationCertificateable), nil
}
// Patch update the navigation property applePushNotificationCertificate in deviceManagement
// returns a ApplePushNotificationCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplePushNotificationCertificateable, requestConfiguration *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplePushNotificationCertificateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplePushNotificationCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplePushNotificationCertificateable), nil
}
// ToDeleteRequestInformation delete navigation property applePushNotificationCertificate for deviceManagement
// returns a *RequestInformation when successful
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation apple push notification certificate.
// returns a *RequestInformation when successful
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property applePushNotificationCertificate in deviceManagement
// returns a *RequestInformation when successful
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApplePushNotificationCertificateable, requestConfiguration *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder when successful
func (m *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) WithUrl(rawUrl string)(*ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) {
    return NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
