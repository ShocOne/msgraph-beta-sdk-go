package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder provides operations to call the revokeAppleVppLicenses method.
type ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderInternal instantiates a new ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder and sets the default values.
func NewManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) {
    m := &ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/revokeAppleVppLicenses", pathParameters),
    }
    return m
}
// NewManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder instantiates a new ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder and sets the default values.
func NewManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post revoke all Apple Vpp licenses for a device
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) Post(ctx context.Context, requestConfiguration *ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation revoke all Apple Vpp licenses for a device
// returns a *RequestInformation when successful
func (m *ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder when successful
func (m *ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) {
    return NewManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
