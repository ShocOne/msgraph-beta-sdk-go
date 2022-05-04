package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookUserable 
type OutlookUserable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMasterCategories()([]OutlookCategoryable)
    GetTaskFolders()([]OutlookTaskFolderable)
    GetTaskGroups()([]OutlookTaskGroupable)
    GetTasks()([]OutlookTaskable)
    SetMasterCategories(value []OutlookCategoryable)()
    SetTaskFolders(value []OutlookTaskFolderable)()
    SetTaskGroups(value []OutlookTaskGroupable)()
    SetTasks(value []OutlookTaskable)()
}