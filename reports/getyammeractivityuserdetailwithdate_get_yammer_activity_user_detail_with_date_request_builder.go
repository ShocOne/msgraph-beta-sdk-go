package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder provides operations to call the getYammerActivityUserDetail method.
type GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilderInternal instantiates a new GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder) {
    m := &GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getYammerActivityUserDetail(date={date})", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder instantiates a new GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder and sets the default values.
func NewGetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getYammerActivityUserDetail
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getYammerActivityUserDetail
// returns a *RequestInformation when successful
func (m *GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder when successful
func (m *GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*GetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder) {
    return NewGetyammeractivityuserdetailwithdateGetYammerActivityUserDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
