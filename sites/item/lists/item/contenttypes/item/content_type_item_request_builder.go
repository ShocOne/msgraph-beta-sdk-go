package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i0eb0ef44081f6129ed82f84b95a9d8901c5c53db2ff78cf336da3540e51902a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/copytodefaultcontentlocation"
    i24fd2f2396a7ad9f1c120d6ef5eea2bb0f52a3b0ab9176b7b84abea77552d44d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/basetypes"
    i2c3f11bea353e2053360132c240fdb0421d0e64bd89c28e21a2a5f54d072902c "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/associatewithhubsites"
    i7971ec5b6746bc19a515d87f0c804fe6b9801df230d1bd98528a5cc26b925b7e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/unpublish"
    i99732ebd53d03d959a23a85d7cef79aa48fc15faa7863489113f7954f63618a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base"
    iafaa9e8f8d30224a4b19bafd6f27a6391457da364f2742173947aa5f8e09182a "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/publish"
    ibc27ed7d114db34133b5c7801610abab919a9cc58319247f2378aacf6c52c697 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columnpositions"
    id9af0414991da05e0a4cf709357ff74c7cf2e09da2b0b24dd95a2a0359adf307 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columns"
    ifb49fb4626206326449bcf75747e5bc8277181f4515cf8f4cbc7f5aa84a04e4b "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks"
    ifdcf97bbb94f49d415a91428a9d320151872eecd6c1f08d3669c845444f4ec54 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/ispublished"
    i27cf64cc6580d9aab73576050bd5ff42d2a2ca98593e1793629372ac762dd1ec "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columns/item"
    i36040d34d2c33b07f1954fdcc30ea2401889b5f3514f4a0ace751d1b76656bae "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columnpositions/item"
    i3af95fe3549090b1bd570314fc75927d3d5156998623754c717e9f723fa6a0fb "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/basetypes/item"
    i800508a07b27c3d6db01cd892d2b2ecaff16c5d270423c5b1cc9e33635b0bf53 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i2c3f11bea353e2053360132c240fdb0421d0e64bd89c28e21a2a5f54d072902c.AssociateWithHubSitesRequestBuilder) {
    return i2c3f11bea353e2053360132c240fdb0421d0e64bd89c28e21a2a5f54d072902c.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i99732ebd53d03d959a23a85d7cef79aa48fc15faa7863489113f7954f63618a3.BaseRequestBuilder) {
    return i99732ebd53d03d959a23a85d7cef79aa48fc15faa7863489113f7954f63618a3.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i24fd2f2396a7ad9f1c120d6ef5eea2bb0f52a3b0ab9176b7b84abea77552d44d.BaseTypesRequestBuilder) {
    return i24fd2f2396a7ad9f1c120d6ef5eea2bb0f52a3b0ab9176b7b84abea77552d44d.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i3af95fe3549090b1bd570314fc75927d3d5156998623754c717e9f723fa6a0fb.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i3af95fe3549090b1bd570314fc75927d3d5156998623754c717e9f723fa6a0fb.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*ifb49fb4626206326449bcf75747e5bc8277181f4515cf8f4cbc7f5aa84a04e4b.ColumnLinksRequestBuilder) {
    return ifb49fb4626206326449bcf75747e5bc8277181f4515cf8f4cbc7f5aa84a04e4b.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i800508a07b27c3d6db01cd892d2b2ecaff16c5d270423c5b1cc9e33635b0bf53.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i800508a07b27c3d6db01cd892d2b2ecaff16c5d270423c5b1cc9e33635b0bf53.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ibc27ed7d114db34133b5c7801610abab919a9cc58319247f2378aacf6c52c697.ColumnPositionsRequestBuilder) {
    return ibc27ed7d114db34133b5c7801610abab919a9cc58319247f2378aacf6c52c697.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i36040d34d2c33b07f1954fdcc30ea2401889b5f3514f4a0ace751d1b76656bae.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i36040d34d2c33b07f1954fdcc30ea2401889b5f3514f4a0ace751d1b76656bae.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*id9af0414991da05e0a4cf709357ff74c7cf2e09da2b0b24dd95a2a0359adf307.ColumnsRequestBuilder) {
    return id9af0414991da05e0a4cf709357ff74c7cf2e09da2b0b24dd95a2a0359adf307.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i27cf64cc6580d9aab73576050bd5ff42d2a2ca98593e1793629372ac762dd1ec.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i27cf64cc6580d9aab73576050bd5ff42d2a2ca98593e1793629372ac762dd1ec.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i0eb0ef44081f6129ed82f84b95a9d8901c5c53db2ff78cf336da3540e51902a2.CopyToDefaultContentLocationRequestBuilder) {
    return i0eb0ef44081f6129ed82f84b95a9d8901c5c53db2ff78cf336da3540e51902a2.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ifdcf97bbb94f49d415a91428a9d320151872eecd6c1f08d3669c845444f4ec54.IsPublishedRequestBuilder) {
    return ifdcf97bbb94f49d415a91428a9d320151872eecd6c1f08d3669c845444f4ec54.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeItemRequestBuilder) Publish()(*iafaa9e8f8d30224a4b19bafd6f27a6391457da364f2742173947aa5f8e09182a.PublishRequestBuilder) {
    return iafaa9e8f8d30224a4b19bafd6f27a6391457da364f2742173947aa5f8e09182a.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i7971ec5b6746bc19a515d87f0c804fe6b9801df230d1bd98528a5cc26b925b7e.UnpublishRequestBuilder) {
    return i7971ec5b6746bc19a515d87f0c804fe6b9801df230d1bd98528a5cc26b925b7e.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}