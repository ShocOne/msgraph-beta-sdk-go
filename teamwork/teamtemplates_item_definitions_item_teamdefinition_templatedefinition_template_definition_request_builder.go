package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderGetQueryParameters generic representation of a team template definition for a team with a specific structure and configuration.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderGetQueryParameters
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/templateDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get generic representation of a team template definition for a team with a specific structure and configuration.
// returns a TeamTemplateDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamTemplateDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamTemplateDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamTemplateDefinitionable), nil
}
// ToGetRequestInformation generic representation of a team template definition for a team with a specific structure and configuration.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
