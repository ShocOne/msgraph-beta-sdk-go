package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder provides operations to manage the definition property of the microsoft.graph.groupPolicyPresentation entity.
type DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetQueryParameters the group policy definition associated with the presentation.
type DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetQueryParameters
}
// NewDeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal instantiates a new DefinitionRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    m := &DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition/previousVersionDefinition/presentations/{groupPolicyPresentation%2Did}/definition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder instantiates a new DefinitionRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the group policy definition associated with the presentation.
func (m *DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the group policy definition associated with the presentation.
func (m *DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable), nil
}
