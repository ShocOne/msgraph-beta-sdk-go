package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponse struct {
    FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponse
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponse instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponse and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponse()(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponse) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponse{
        FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponse: *NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountResponseable interface {
    FilestorageContainersItemDriveItemsItemWorkbookTablesItemRowsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
