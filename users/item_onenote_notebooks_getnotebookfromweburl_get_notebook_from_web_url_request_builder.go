package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder provides operations to call the getNotebookFromWebUrl method.
type ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilderInternal instantiates a new ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder) {
    m := &ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onenote/notebooks/getNotebookFromWebUrl", pathParameters),
    }
    return m
}
// NewItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder instantiates a new ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post retrieve the properties and relationships of a notebook object by using its URL path. The location can be user notebooks on Microsoft 365, group notebooks, or SharePoint site-hosted team notebooks on Microsoft 365.
// returns a CopyNotebookModelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/notebook-getnotebookfromweburl?view=graph-rest-beta
func (m *ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder) Post(ctx context.Context, body ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CopyNotebookModelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCopyNotebookModelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CopyNotebookModelable), nil
}
// ToPostRequestInformation retrieve the properties and relationships of a notebook object by using its URL path. The location can be user notebooks on Microsoft 365, group notebooks, or SharePoint site-hosted team notebooks on Microsoft 365.
// returns a *RequestInformation when successful
func (m *ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder when successful
func (m *ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder) {
    return NewItemOnenoteNotebooksGetnotebookfromweburlGetNotebookFromWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
