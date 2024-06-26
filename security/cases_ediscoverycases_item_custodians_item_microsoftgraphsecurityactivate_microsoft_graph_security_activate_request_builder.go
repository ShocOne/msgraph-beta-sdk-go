package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder provides operations to call the activate method.
type CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/microsoft.graph.security.activate", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate a custodian that has been released from a case to make them part of the case again. For details, see Manage custodians in an eDiscovery (Premium) case.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-activate?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation activate a custodian that has been released from a case to make them part of the case again. For details, see Manage custodians in an eDiscovery (Premium) case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
