package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionChannelsItemMembersAddRequestBuilder provides operations to call the add method.
type ItemTeamdefinitionChannelsItemMembersAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamdefinitionChannelsItemMembersAddRequestBuilderInternal instantiates a new ItemTeamdefinitionChannelsItemMembersAddRequestBuilder and sets the default values.
func NewItemTeamdefinitionChannelsItemMembersAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionChannelsItemMembersAddRequestBuilder) {
    m := &ItemTeamdefinitionChannelsItemMembersAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/members/add", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionChannelsItemMembersAddRequestBuilder instantiates a new ItemTeamdefinitionChannelsItemMembersAddRequestBuilder and sets the default values.
func NewItemTeamdefinitionChannelsItemMembersAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionChannelsItemMembersAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionChannelsItemMembersAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// Deprecated: This method is obsolete. Use PostAsAddPostResponse instead.
// returns a ItemTeamdefinitionChannelsItemMembersAddResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *ItemTeamdefinitionChannelsItemMembersAddRequestBuilder) Post(ctx context.Context, body ItemTeamdefinitionChannelsItemMembersAddPostRequestBodyable, requestConfiguration *ItemTeamdefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration)(ItemTeamdefinitionChannelsItemMembersAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamdefinitionChannelsItemMembersAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamdefinitionChannelsItemMembersAddResponseable), nil
}
// PostAsAddPostResponse add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a ItemTeamdefinitionChannelsItemMembersAddPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-beta
func (m *ItemTeamdefinitionChannelsItemMembersAddRequestBuilder) PostAsAddPostResponse(ctx context.Context, body ItemTeamdefinitionChannelsItemMembersAddPostRequestBodyable, requestConfiguration *ItemTeamdefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration)(ItemTeamdefinitionChannelsItemMembersAddPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamdefinitionChannelsItemMembersAddPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamdefinitionChannelsItemMembersAddPostResponseable), nil
}
// ToPostRequestInformation add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionChannelsItemMembersAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamdefinitionChannelsItemMembersAddPostRequestBodyable, requestConfiguration *ItemTeamdefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamdefinitionChannelsItemMembersAddRequestBuilder when successful
func (m *ItemTeamdefinitionChannelsItemMembersAddRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionChannelsItemMembersAddRequestBuilder) {
    return NewItemTeamdefinitionChannelsItemMembersAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
