package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder provides operations to call the postAndSend method.
type CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderInternal instantiates a new CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder and sets the default values.
func NewCompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder) {
    m := &CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}/postAndSend", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder instantiates a new CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder and sets the default values.
func NewCompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action postAndSend
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder) Post(ctx context.Context, requestConfiguration *CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action postAndSend
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
