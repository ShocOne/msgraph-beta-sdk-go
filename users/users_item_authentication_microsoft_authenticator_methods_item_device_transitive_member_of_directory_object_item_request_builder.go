package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Application()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemApplicationRequestBuilder) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    m := &UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device casts the previous resource to device.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Device()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemDeviceRequestBuilder) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Group casts the previous resource to group.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Group()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemGroupRequestBuilder) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) OrgContact()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemOrgContactRequestBuilder) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) ServicePrincipal()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemServicePrincipalRequestBuilder) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) User()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemUserRequestBuilder) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
