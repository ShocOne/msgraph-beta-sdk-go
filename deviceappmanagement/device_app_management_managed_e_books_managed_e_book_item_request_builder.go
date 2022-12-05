package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
type DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters the Managed eBook.
type DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) Assign()(*DeviceAppManagementManagedEBooksItemAssignRequestBuilder) {
    return NewDeviceAppManagementManagedEBooksItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) Assignments()(*DeviceAppManagementManagedEBooksItemAssignmentsRequestBuilder) {
    return NewDeviceAppManagementManagedEBooksItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) AssignmentsById(id string)(*DeviceAppManagementManagedEBooksItemAssignmentsManagedEBookAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment%2Did"] = id
    }
    return NewDeviceAppManagementManagedEBooksItemAssignmentsManagedEBookAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories provides operations to manage the categories property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) Categories()(*DeviceAppManagementManagedEBooksItemCategoriesRequestBuilder) {
    return NewDeviceAppManagementManagedEBooksItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) CategoriesById(id string)(*DeviceAppManagementManagedEBooksItemCategoriesManagedEBookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookCategory%2Did"] = id
    }
    return NewDeviceAppManagementManagedEBooksItemCategoriesManagedEBookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderInternal instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewDeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) {
    m := &DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewDeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedEBooks for deviceAppManagement
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the Managed eBook.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedEBooks in deviceAppManagement
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, requestConfiguration *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property managedEBooks for deviceAppManagement
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) DeviceStates()(*DeviceAppManagementManagedEBooksItemDeviceStatesRequestBuilder) {
    return NewDeviceAppManagementManagedEBooksItemDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById provides operations to manage the deviceStates property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) DeviceStatesById(id string)(*DeviceAppManagementManagedEBooksItemDeviceStatesDeviceInstallStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState%2Did"] = id
    }
    return NewDeviceAppManagementManagedEBooksItemDeviceStatesDeviceInstallStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the Managed eBook.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedEBookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable), nil
}
// InstallSummary provides operations to manage the installSummary property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) InstallSummary()(*DeviceAppManagementManagedEBooksItemInstallSummaryRequestBuilder) {
    return NewDeviceAppManagementManagedEBooksItemInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedEBooks in deviceAppManagement
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, requestConfiguration *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedEBookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable), nil
}
// UserStateSummary provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) UserStateSummary()(*DeviceAppManagementManagedEBooksItemUserStateSummaryRequestBuilder) {
    return NewDeviceAppManagementManagedEBooksItemUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStateSummaryById provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
func (m *DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) UserStateSummaryById(id string)(*DeviceAppManagementManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary%2Did"] = id
    }
    return NewDeviceAppManagementManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
