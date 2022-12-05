package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder provides operations to manage the customer property of the microsoft.graph.customerPayment entity.
type FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetQueryParameters get customer from financials
type FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderInternal instantiates a new CustomerRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) {
    m := &FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}/customerPayments/{customerPayment%2Did}/customer{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder instantiates a new CustomerRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customer for financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get customer from financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property customer in financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency provides operations to manage the currency property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Currency()(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerCurrencyRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property customer for financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Delete(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get customer from financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// Patch update the navigation property customer in financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// PaymentMethod provides operations to manage the paymentMethod property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) PaymentMethod()(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentMethodRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) PaymentTerm()(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentTermRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Picture provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Picture()(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPictureRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) PictureById(id string)(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPicturePictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPicturePictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) ShipmentMethod()(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerShipmentMethodRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
