package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponse struct {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponse
}
// NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponse instantiates a new FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponse()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponse) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponse{
        FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponse: *NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountResponseable interface {
    FilestorageDeletedcontainersItemDriveItemsItemWorkbookCommentsItemRepliesItemTaskChangesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
