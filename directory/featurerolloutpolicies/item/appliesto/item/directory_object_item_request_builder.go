package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia64be5ca1bd912a6010bc409a6bd6945cb9100072a81f2af7f0655c2bc88e88e "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto/item/ref"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \directory\featureRolloutPolicies\{featureRolloutPolicy-id}\appliesTo\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of directory entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*ia64be5ca1bd912a6010bc409a6bd6945cb9100072a81f2af7f0655c2bc88e88e.RefRequestBuilder) {
    return ia64be5ca1bd912a6010bc409a6bd6945cb9100072a81f2af7f0655c2bc88e88e.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
