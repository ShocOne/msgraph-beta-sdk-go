package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestPostResponseable instead.
type ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponse struct {
    ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestPostResponse
}
// NewApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponse instantiates a new ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponse and sets the default values.
func NewApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponse()(*ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponse) {
    m := &ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponse{
        ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestPostResponse: *NewApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestPostResponse(),
    }
    return m
}
// CreateApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponse(), nil
}
// Deprecated: This class is obsolete. Use ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestPostResponseable instead.
type ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestResponseable interface {
    ApplepushnotificationcertificateGenerateapplepushnotificationcertificatesigningrequestGenerateApplePushNotificationCertificateSigningRequestPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
