package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApplicationItemRequestBuilder provides operations to manage the collection of application entities.
type ApplicationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApplicationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationItemRequestBuilderGetQueryParameters get the properties and relationships of an application object.
type ApplicationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApplicationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApplicationItemRequestBuilderGetQueryParameters
}
// ApplicationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddKey provides operations to call the addKey method.
// returns a *ItemAddkeyAddKeyRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) AddKey()(*ItemAddkeyAddKeyRequestBuilder) {
    return NewItemAddkeyAddKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddPassword provides operations to call the addPassword method.
// returns a *ItemAddpasswordAddPasswordRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) AddPassword()(*ItemAddpasswordAddPasswordRequestBuilder) {
    return NewItemAddpasswordAddPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManagementPolicies provides operations to manage the appManagementPolicies property of the microsoft.graph.application entity.
// returns a *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) AppManagementPolicies()(*ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) {
    return NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
// returns a *ItemCheckmembergroupsCheckMemberGroupsRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) CheckMemberGroups()(*ItemCheckmembergroupsCheckMemberGroupsRequestBuilder) {
    return NewItemCheckmembergroupsCheckMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
// returns a *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) CheckMemberObjects()(*ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    return NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConnectorGroup provides operations to manage the connectorGroup property of the microsoft.graph.application entity.
// returns a *ItemConnectorgroupConnectorGroupRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) ConnectorGroup()(*ItemConnectorgroupConnectorGroupRequestBuilder) {
    return NewItemConnectorgroupConnectorGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApplicationItemRequestBuilderInternal instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationItemRequestBuilder) {
    m := &ApplicationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewApplicationItemRequestBuilder instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedOnBehalfOf provides operations to manage the createdOnBehalfOf property of the microsoft.graph.application entity.
// returns a *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) CreatedOnBehalfOf()(*ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder) {
    return NewItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deletes an application. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/application-delete?view=graph-rest-beta
func (m *ApplicationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExtensionProperties provides operations to manage the extensionProperties property of the microsoft.graph.application entity.
// returns a *ItemExtensionpropertiesExtensionPropertiesRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) ExtensionProperties()(*ItemExtensionpropertiesExtensionPropertiesRequestBuilder) {
    return NewItemExtensionpropertiesExtensionPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.application entity.
// returns a *ItemFederatedidentitycredentialsFederatedIdentityCredentialsRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) FederatedIdentityCredentials()(*ItemFederatedidentitycredentialsFederatedIdentityCredentialsRequestBuilder) {
    return NewItemFederatedidentitycredentialsFederatedIdentityCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FederatedIdentityCredentialsWithName provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.application entity.
// returns a *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) FederatedIdentityCredentialsWithName(name *string)(*ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) {
    return NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, name)
}
// Get get the properties and relationships of an application object.
// returns a Applicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/application-get?view=graph-rest-beta
func (m *ApplicationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
// returns a *ItemGetmembergroupsGetMemberGroupsRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) GetMemberGroups()(*ItemGetmembergroupsGetMemberGroupsRequestBuilder) {
    return NewItemGetmembergroupsGetMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
// returns a *ItemGetmemberobjectsGetMemberObjectsRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) GetMemberObjects()(*ItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    return NewItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
// returns a *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) HomeRealmDiscoveryPolicies()(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logo provides operations to manage the media for the application entity.
// returns a *ItemLogoRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) Logo()(*ItemLogoRequestBuilder) {
    return NewItemLogoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.application entity.
// returns a *ItemOwnersRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) Owners()(*ItemOwnersRequestBuilder) {
    return NewItemOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch create a new application object if it doesn't exist, or update the properties of an existing application object.
// returns a Applicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/application-upsert?view=graph-rest-beta
func (m *ApplicationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable), nil
}
// RemoveKey provides operations to call the removeKey method.
// returns a *ItemRemovekeyRemoveKeyRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) RemoveKey()(*ItemRemovekeyRemoveKeyRequestBuilder) {
    return NewItemRemovekeyRemoveKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemovePassword provides operations to call the removePassword method.
// returns a *ItemRemovepasswordRemovePasswordRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) RemovePassword()(*ItemRemovepasswordRemovePasswordRequestBuilder) {
    return NewItemRemovepasswordRemovePasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *ItemRestoreRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) Restore()(*ItemRestoreRequestBuilder) {
    return NewItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetVerifiedPublisher provides operations to call the setVerifiedPublisher method.
// returns a *ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) SetVerifiedPublisher()(*ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder) {
    return NewItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Synchronization provides operations to manage the synchronization property of the microsoft.graph.application entity.
// returns a *ItemSynchronizationRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) Synchronization()(*ItemSynchronizationRequestBuilder) {
    return NewItemSynchronizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes an application. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
// returns a *RequestInformation when successful
func (m *ApplicationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of an application object.
// returns a *RequestInformation when successful
func (m *ApplicationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.application entity.
// returns a *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) TokenIssuancePolicies()(*ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    return NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.application entity.
// returns a *ItemTokenlifetimepoliciesTokenLifetimePoliciesRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) TokenLifetimePolicies()(*ItemTokenlifetimepoliciesTokenLifetimePoliciesRequestBuilder) {
    return NewItemTokenlifetimepoliciesTokenLifetimePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPatchRequestInformation create a new application object if it doesn't exist, or update the properties of an existing application object.
// returns a *RequestInformation when successful
func (m *ApplicationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnsetVerifiedPublisher provides operations to call the unsetVerifiedPublisher method.
// returns a *ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) UnsetVerifiedPublisher()(*ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder) {
    return NewItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApplicationItemRequestBuilder when successful
func (m *ApplicationItemRequestBuilder) WithUrl(rawUrl string)(*ApplicationItemRequestBuilder) {
    return NewApplicationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
