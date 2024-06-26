package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder provides operations to call the getEncryptionPublicKey method.
type DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderInternal instantiates a new DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) {
    m := &DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/getEncryptionPublicKey()", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder instantiates a new DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a public key to use to encrypt the Apple device enrollment program token
// Deprecated: This method is obsolete. Use GetAsGetEncryptionPublicKeyGetResponse instead.
// returns a DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration)(DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyResponseable), nil
}
// GetAsGetEncryptionPublicKeyGetResponse get a public key to use to encrypt the Apple device enrollment program token
// returns a DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) GetAsGetEncryptionPublicKeyGetResponse(ctx context.Context, requestConfiguration *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration)(DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyGetResponseable), nil
}
// ToGetRequestInformation get a public key to use to encrypt the Apple device enrollment program token
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder when successful
func (m *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) {
    return NewDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
