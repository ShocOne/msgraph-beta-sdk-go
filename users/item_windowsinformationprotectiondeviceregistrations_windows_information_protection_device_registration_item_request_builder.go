package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
type ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters zero or more WIP device registrations that belong to the user.
type ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters
}
// NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal instantiates a new ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder and sets the default values.
func NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    m := &ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/windowsInformationProtectionDeviceRegistrations/{windowsInformationProtectionDeviceRegistration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder instantiates a new ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder and sets the default values.
func NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get zero or more WIP device registrations that belong to the user.
// returns a WindowsInformationProtectionDeviceRegistrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable), nil
}
// ToGetRequestInformation zero or more WIP device registrations that belong to the user.
// returns a *RequestInformation when successful
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder when successful
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) WithUrl(rawUrl string)(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    return NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
