package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponse struct {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountGetResponse
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponse instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponse()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponse) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponse{
        FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountGetResponse: *NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountResponseable interface {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookTablesItemColumnsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
