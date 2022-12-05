package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
type DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderGetQueryParameters the device management intents
type DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderGetQueryParameters
}
// DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) Assign()(*DeviceManagementIntentsItemAssignRequestBuilder) {
    return NewDeviceManagementIntentsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) Assignments()(*DeviceManagementIntentsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementIntentsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementIntentsItemAssignmentsDeviceManagementIntentAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentAssignment%2Did"] = id
    }
    return NewDeviceManagementIntentsItemAssignmentsDeviceManagementIntentAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) Categories()(*DeviceManagementIntentsItemCategoriesRequestBuilder) {
    return NewDeviceManagementIntentsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) CategoriesById(id string)(*DeviceManagementIntentsItemCategoriesDeviceManagementIntentSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentSettingCategory%2Did"] = id
    }
    return NewDeviceManagementIntentsItemCategoriesDeviceManagementIntentSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) CompareWithTemplateId(templateId *string)(*DeviceManagementIntentsItemCompareWithTemplateIdRequestBuilder) {
    return NewDeviceManagementIntentsItemCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementIntentsDeviceManagementIntentItemRequestBuilderInternal instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewDeviceManagementIntentsDeviceManagementIntentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) {
    m := &DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementIntentsDeviceManagementIntentItemRequestBuilder instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewDeviceManagementIntentsDeviceManagementIntentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementIntentsDeviceManagementIntentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy provides operations to call the createCopy method.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) CreateCopy()(*DeviceManagementIntentsItemCreateCopyRequestBuilder) {
    return NewDeviceManagementIntentsItemCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property intents for deviceManagement
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the device management intents
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property intents in deviceManagement
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property intents for deviceManagement
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) DeviceSettingStateSummaries()(*DeviceManagementIntentsItemDeviceSettingStateSummariesRequestBuilder) {
    return NewDeviceManagementIntentsItemDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*DeviceManagementIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentDeviceSettingStateSummary%2Did"] = id
    }
    return NewDeviceManagementIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) DeviceStates()(*DeviceManagementIntentsItemDeviceStatesRequestBuilder) {
    return NewDeviceManagementIntentsItemDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById provides operations to manage the deviceStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) DeviceStatesById(id string)(*DeviceManagementIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentDeviceState%2Did"] = id
    }
    return NewDeviceManagementIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStateSummary provides operations to manage the deviceStateSummary property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) DeviceStateSummary()(*DeviceManagementIntentsItemDeviceStateSummaryRequestBuilder) {
    return NewDeviceManagementIntentsItemDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device management intents
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// MigrateToTemplate provides operations to call the migrateToTemplate method.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) MigrateToTemplate()(*DeviceManagementIntentsItemMigrateToTemplateRequestBuilder) {
    return NewDeviceManagementIntentsItemMigrateToTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property intents in deviceManagement
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) Settings()(*DeviceManagementIntentsItemSettingsRequestBuilder) {
    return NewDeviceManagementIntentsItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) SettingsById(id string)(*DeviceManagementIntentsItemSettingsDeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance%2Did"] = id
    }
    return NewDeviceManagementIntentsItemSettingsDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdateSettings provides operations to call the updateSettings method.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) UpdateSettings()(*DeviceManagementIntentsItemUpdateSettingsRequestBuilder) {
    return NewDeviceManagementIntentsItemUpdateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStates provides operations to manage the userStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) UserStates()(*DeviceManagementIntentsItemUserStatesRequestBuilder) {
    return NewDeviceManagementIntentsItemUserStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatesById provides operations to manage the userStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) UserStatesById(id string)(*DeviceManagementIntentsItemUserStatesDeviceManagementIntentUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentUserState%2Did"] = id
    }
    return NewDeviceManagementIntentsItemUserStatesDeviceManagementIntentUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStateSummary provides operations to manage the userStateSummary property of the microsoft.graph.deviceManagementIntent entity.
func (m *DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) UserStateSummary()(*DeviceManagementIntentsItemUserStateSummaryRequestBuilder) {
    return NewDeviceManagementIntentsItemUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
