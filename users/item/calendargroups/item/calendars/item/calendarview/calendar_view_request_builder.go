package calendarview

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ib18a70bc5f4698904fd46c5bb453f8eb25200186dd29e13ba582753d0adddd81 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/delta"
    if288ff2ce9bf016877392b70529db5c8c023a226a65e0d9a2fa8106ca727797e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/count"
)

// CalendarViewRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type CalendarViewRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarViewRequestBuilderGetQueryParameters get the occurrences, exceptions and single instances of events in a calendar view defined by a time range,from a user's default calendar `(../me/calendarview)` or some other calendar of the user's.
type CalendarViewRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CalendarViewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarViewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarViewRequestBuilderGetQueryParameters
}
// NewCalendarViewRequestBuilderInternal instantiates a new CalendarViewRequestBuilder and sets the default values.
func NewCalendarViewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarViewRequestBuilder) {
    m := &CalendarViewRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarViewRequestBuilder instantiates a new CalendarViewRequestBuilder and sets the default values.
func NewCalendarViewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarViewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarViewRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *CalendarViewRequestBuilder) Count()(*if288ff2ce9bf016877392b70529db5c8c023a226a65e0d9a2fa8106ca727797e.CountRequestBuilder) {
    return if288ff2ce9bf016877392b70529db5c8c023a226a65e0d9a2fa8106ca727797e.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the occurrences, exceptions and single instances of events in a calendar view defined by a time range,from a user's default calendar `(../me/calendarview)` or some other calendar of the user's.
func (m *CalendarViewRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CalendarViewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *CalendarViewRequestBuilder) Delta()(*ib18a70bc5f4698904fd46c5bb453f8eb25200186dd29e13ba582753d0adddd81.DeltaRequestBuilder) {
    return ib18a70bc5f4698904fd46c5bb453f8eb25200186dd29e13ba582753d0adddd81.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the occurrences, exceptions and single instances of events in a calendar view defined by a time range,from a user's default calendar `(../me/calendarview)` or some other calendar of the user's.
func (m *CalendarViewRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarViewRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable), nil
}
