package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponse struct {
    FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountGetResponse
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponse instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponse and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponse()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponse) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponse{
        FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountGetResponse: *NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountResponseable interface {
    FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTasksCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
