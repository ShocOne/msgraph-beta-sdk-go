package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder casts the previous resource to appRoleAssignment.
type ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.appRoleAssignment
type ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderGetQueryParameters
}
// NewItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderInternal instantiates a new ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder and sets the default values.
func NewItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder) {
    m := &ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/registeredDevices/{directoryObject%2Did}/graph.appRoleAssignment{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder instantiates a new ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder and sets the default values.
func NewItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.appRoleAssignment
// returns a AppRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppRoleAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppRoleAssignmentable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.appRoleAssignment
// returns a *RequestInformation when successful
func (m *ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder) WithUrl(rawUrl string)(*ItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredDevicesItemGraphAppRoleAssignmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
