package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponse struct {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageGetResponse
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponse instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponse()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponse) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponse{
        FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageGetResponse: *NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageResponseable interface {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImageGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
