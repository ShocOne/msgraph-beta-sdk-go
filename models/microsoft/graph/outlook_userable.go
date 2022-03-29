package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookUserable 
type OutlookUserable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetMasterCategories()([]OutlookCategoryable)
    GetTaskFolders()([]OutlookTaskFolderable)
    GetTaskGroups()([]OutlookTaskGroupable)
    GetTasks()([]OutlookTaskable)
    SetMasterCategories(value []OutlookCategoryable)()
    SetTaskFolders(value []OutlookTaskFolderable)()
    SetTaskGroups(value []OutlookTaskGroupable)()
    SetTasks(value []OutlookTaskable)()
}