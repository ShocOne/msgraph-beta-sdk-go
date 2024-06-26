package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BookingbusinessesBookingBusinessItemRequestBuilder provides operations to manage the bookingBusinesses property of the microsoft.graph.solutionsRoot entity.
type BookingbusinessesBookingBusinessItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingbusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BookingbusinessesBookingBusinessItemRequestBuilderGetQueryParameters get bookingBusinesses from solutions
type BookingbusinessesBookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BookingbusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingbusinessesBookingBusinessItemRequestBuilderGetQueryParameters
}
// BookingbusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Appointments provides operations to manage the appointments property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingbusinessesItemAppointmentsRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Appointments()(*BookingbusinessesItemAppointmentsRequestBuilder) {
    return NewBookingbusinessesItemAppointmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingbusinessesItemCalendarviewCalendarViewRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) CalendarView()(*BookingbusinessesItemCalendarviewCalendarViewRequestBuilder) {
    return NewBookingbusinessesItemCalendarviewCalendarViewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewBookingbusinessesBookingBusinessItemRequestBuilderInternal instantiates a new BookingbusinessesBookingBusinessItemRequestBuilder and sets the default values.
func NewBookingbusinessesBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesBookingBusinessItemRequestBuilder) {
    m := &BookingbusinessesBookingBusinessItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBookingbusinessesBookingBusinessItemRequestBuilder instantiates a new BookingbusinessesBookingBusinessItemRequestBuilder and sets the default values.
func NewBookingbusinessesBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesBookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingbusinessesBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Customers provides operations to manage the customers property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingbusinessesItemCustomersRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Customers()(*BookingbusinessesItemCustomersRequestBuilder) {
    return NewBookingbusinessesItemCustomersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CustomQuestions provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingbusinessesItemCustomquestionsCustomQuestionsRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) CustomQuestions()(*BookingbusinessesItemCustomquestionsCustomQuestionsRequestBuilder) {
    return NewBookingbusinessesItemCustomquestionsCustomQuestionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a bookingBusiness object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-delete?view=graph-rest-beta
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BookingbusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get bookingBusinesses from solutions
// returns a BookingBusinessable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingbusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable), nil
}
// GetStaffAvailability provides operations to call the getStaffAvailability method.
// returns a *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) GetStaffAvailability()(*BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) {
    return NewBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property bookingBusinesses in solutions
// returns a BookingBusinessable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, requestConfiguration *BookingbusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable), nil
}
// Publish provides operations to call the publish method.
// returns a *BookingbusinessesItemPublishRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Publish()(*BookingbusinessesItemPublishRequestBuilder) {
    return NewBookingbusinessesItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Services provides operations to manage the services property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingbusinessesItemServicesRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Services()(*BookingbusinessesItemServicesRequestBuilder) {
    return NewBookingbusinessesItemServicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StaffMembers provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingbusinessesItemStaffmembersStaffMembersRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) StaffMembers()(*BookingbusinessesItemStaffmembersStaffMembersRequestBuilder) {
    return NewBookingbusinessesItemStaffmembersStaffMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a bookingBusiness object.
// returns a *RequestInformation when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BookingbusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get bookingBusinesses from solutions
// returns a *RequestInformation when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BookingbusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property bookingBusinesses in solutions
// returns a *RequestInformation when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, requestConfiguration *BookingbusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Unpublish provides operations to call the unpublish method.
// returns a *BookingbusinessesItemUnpublishRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) Unpublish()(*BookingbusinessesItemUnpublishRequestBuilder) {
    return NewBookingbusinessesItemUnpublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BookingbusinessesBookingBusinessItemRequestBuilder when successful
func (m *BookingbusinessesBookingBusinessItemRequestBuilder) WithUrl(rawUrl string)(*BookingbusinessesBookingBusinessItemRequestBuilder) {
    return NewBookingbusinessesBookingBusinessItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
