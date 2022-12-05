package termstore

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
type TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderGetQueryParameters the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
type TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderGetQueryParameters
}
// NewTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderInternal instantiates a new FromTermRequestBuilder and sets the default values.
func NewTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder) {
    m := &TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}/fromTerm{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder instantiates a new FromTermRequestBuilder and sets the default values.
func NewTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilder) Get(ctx context.Context, requestConfiguration *TermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemRelationsItemFromTermRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable), nil
}
