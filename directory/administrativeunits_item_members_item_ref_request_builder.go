package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdministrativeunitsItemMembersItemRefRequestBuilder provides operations to manage the collection of directory entities.
type AdministrativeunitsItemMembersItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdministrativeunitsItemMembersItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeunitsItemMembersItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdministrativeunitsItemMembersItemRefRequestBuilderInternal instantiates a new AdministrativeunitsItemMembersItemRefRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersItemRefRequestBuilder) {
    m := &AdministrativeunitsItemMembersItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members/{directoryObject%2Did}/$ref", pathParameters),
    }
    return m
}
// NewAdministrativeunitsItemMembersItemRefRequestBuilder instantiates a new AdministrativeunitsItemMembersItemRefRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeunitsItemMembersItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property members for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdministrativeunitsItemMembersItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdministrativeunitsItemMembersItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// ToDeleteRequestInformation delete ref of navigation property members for directory
// returns a *RequestInformation when successful
func (m *AdministrativeunitsItemMembersItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AdministrativeunitsItemMembersItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AdministrativeunitsItemMembersItemRefRequestBuilder when successful
func (m *AdministrativeunitsItemMembersItemRefRequestBuilder) WithUrl(rawUrl string)(*AdministrativeunitsItemMembersItemRefRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
