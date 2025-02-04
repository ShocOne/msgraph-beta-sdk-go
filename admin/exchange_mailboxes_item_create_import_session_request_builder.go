package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemCreateImportSessionRequestBuilder provides operations to call the createImportSession method.
type ExchangeMailboxesItemCreateImportSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemCreateImportSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemCreateImportSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeMailboxesItemCreateImportSessionRequestBuilderInternal instantiates a new ExchangeMailboxesItemCreateImportSessionRequestBuilder and sets the default values.
func NewExchangeMailboxesItemCreateImportSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemCreateImportSessionRequestBuilder) {
    m := &ExchangeMailboxesItemCreateImportSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/createImportSession", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemCreateImportSessionRequestBuilder instantiates a new ExchangeMailboxesItemCreateImportSessionRequestBuilder and sets the default values.
func NewExchangeMailboxesItemCreateImportSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemCreateImportSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemCreateImportSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a session to import an Exchange mailbox item using its FastTransfer stream (FTS) format.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a MailboxItemImportSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailbox-createimportsession?view=graph-rest-beta
func (m *ExchangeMailboxesItemCreateImportSessionRequestBuilder) Post(ctx context.Context, requestConfiguration *ExchangeMailboxesItemCreateImportSessionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxItemImportSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxItemImportSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxItemImportSessionable), nil
}
// ToPostRequestInformation create a session to import an Exchange mailbox item using its FastTransfer stream (FTS) format.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemCreateImportSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ExchangeMailboxesItemCreateImportSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *ExchangeMailboxesItemCreateImportSessionRequestBuilder when successful
func (m *ExchangeMailboxesItemCreateImportSessionRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemCreateImportSessionRequestBuilder) {
    return NewExchangeMailboxesItemCreateImportSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
