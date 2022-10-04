package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Workflow 
type Workflow struct {
    WorkflowBase
    // The time and date a workflow is deleted. Supports $filter(lt,gt) and $orderby.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the Azure AD identity that last modified the workflow object..
    executionScope []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable
    // Identifier used for individually addressing a specific workflow. Supports $filter(eq).
    id *string
    // If true, the workflow engine creates and processes taskProcessingResults on the users scoped to the workflow. Supports $filter(eq,ne) and orderby.
    isEnabled *bool
    // If true, the workflow engine executes the workflow on the schedule defined by tenant settings.
    isSchedulingEnabled *bool
    // The date time when the workflow is expected to run next based on the schedule interval, if there are any users matching the execution conditions. Supports $filter(lt,gt) and $orderby.
    nextScheduleRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The runs property
    runs []Runable
    // Represents the aggregation of task execution data for tasks within a workflow object.
    taskReports []TaskReportable
    // The userProcessingResults property
    userProcessingResults []UserProcessingResultable
    // The current version number of the workflow. Value is 1 when the workflow is first created. Supports $filter(eq).
    version *int32
    // The workflow versions that are available.
    versions []WorkflowVersionable
}
// NewWorkflow instantiates a new Workflow and sets the default values.
func NewWorkflow()(*Workflow) {
    m := &Workflow{
        WorkflowBase: *NewWorkflowBase(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.workflow";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkflowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflow(), nil
}
// GetDeletedDateTime gets the deletedDateTime property value. The time and date a workflow is deleted. Supports $filter(lt,gt) and $orderby.
func (m *Workflow) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deletedDateTime
}
// GetExecutionScope gets the executionScope property value. The unique identifier of the Azure AD identity that last modified the workflow object..
func (m *Workflow) GetExecutionScope()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    return m.executionScope
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Workflow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowBase.GetFieldDeserializers()
    res["deletedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDeletedDateTime)
    res["executionScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue , m.SetExecutionScope)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["isSchedulingEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSchedulingEnabled)
    res["nextScheduleRunDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetNextScheduleRunDateTime)
    res["runs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRunFromDiscriminatorValue , m.SetRuns)
    res["taskReports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTaskReportFromDiscriminatorValue , m.SetTaskReports)
    res["userProcessingResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserProcessingResultFromDiscriminatorValue , m.SetUserProcessingResults)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    res["versions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkflowVersionFromDiscriminatorValue , m.SetVersions)
    return res
}
// GetId gets the id property value. Identifier used for individually addressing a specific workflow. Supports $filter(eq).
func (m *Workflow) GetId()(*string) {
    return m.id
}
// GetIsEnabled gets the isEnabled property value. If true, the workflow engine creates and processes taskProcessingResults on the users scoped to the workflow. Supports $filter(eq,ne) and orderby.
func (m *Workflow) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetIsSchedulingEnabled gets the isSchedulingEnabled property value. If true, the workflow engine executes the workflow on the schedule defined by tenant settings.
func (m *Workflow) GetIsSchedulingEnabled()(*bool) {
    return m.isSchedulingEnabled
}
// GetNextScheduleRunDateTime gets the nextScheduleRunDateTime property value. The date time when the workflow is expected to run next based on the schedule interval, if there are any users matching the execution conditions. Supports $filter(lt,gt) and $orderby.
func (m *Workflow) GetNextScheduleRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nextScheduleRunDateTime
}
// GetRuns gets the runs property value. The runs property
func (m *Workflow) GetRuns()([]Runable) {
    return m.runs
}
// GetTaskReports gets the taskReports property value. Represents the aggregation of task execution data for tasks within a workflow object.
func (m *Workflow) GetTaskReports()([]TaskReportable) {
    return m.taskReports
}
// GetUserProcessingResults gets the userProcessingResults property value. The userProcessingResults property
func (m *Workflow) GetUserProcessingResults()([]UserProcessingResultable) {
    return m.userProcessingResults
}
// GetVersion gets the version property value. The current version number of the workflow. Value is 1 when the workflow is first created. Supports $filter(eq).
func (m *Workflow) GetVersion()(*int32) {
    return m.version
}
// GetVersions gets the versions property value. The workflow versions that are available.
func (m *Workflow) GetVersions()([]WorkflowVersionable) {
    return m.versions
}
// Serialize serializes information the current object
func (m *Workflow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetExecutionScope() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExecutionScope())
        err = writer.WriteCollectionOfObjectValues("executionScope", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSchedulingEnabled", m.GetIsSchedulingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("nextScheduleRunDateTime", m.GetNextScheduleRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRuns() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRuns())
        err = writer.WriteCollectionOfObjectValues("runs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskReports() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTaskReports())
        err = writer.WriteCollectionOfObjectValues("taskReports", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserProcessingResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserProcessingResults())
        err = writer.WriteCollectionOfObjectValues("userProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    if m.GetVersions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetVersions())
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedDateTime sets the deletedDateTime property value. The time and date a workflow is deleted. Supports $filter(lt,gt) and $orderby.
func (m *Workflow) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// SetExecutionScope sets the executionScope property value. The unique identifier of the Azure AD identity that last modified the workflow object..
func (m *Workflow) SetExecutionScope(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    m.executionScope = value
}
// SetId sets the id property value. Identifier used for individually addressing a specific workflow. Supports $filter(eq).
func (m *Workflow) SetId(value *string)() {
    m.id = value
}
// SetIsEnabled sets the isEnabled property value. If true, the workflow engine creates and processes taskProcessingResults on the users scoped to the workflow. Supports $filter(eq,ne) and orderby.
func (m *Workflow) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetIsSchedulingEnabled sets the isSchedulingEnabled property value. If true, the workflow engine executes the workflow on the schedule defined by tenant settings.
func (m *Workflow) SetIsSchedulingEnabled(value *bool)() {
    m.isSchedulingEnabled = value
}
// SetNextScheduleRunDateTime sets the nextScheduleRunDateTime property value. The date time when the workflow is expected to run next based on the schedule interval, if there are any users matching the execution conditions. Supports $filter(lt,gt) and $orderby.
func (m *Workflow) SetNextScheduleRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextScheduleRunDateTime = value
}
// SetRuns sets the runs property value. The runs property
func (m *Workflow) SetRuns(value []Runable)() {
    m.runs = value
}
// SetTaskReports sets the taskReports property value. Represents the aggregation of task execution data for tasks within a workflow object.
func (m *Workflow) SetTaskReports(value []TaskReportable)() {
    m.taskReports = value
}
// SetUserProcessingResults sets the userProcessingResults property value. The userProcessingResults property
func (m *Workflow) SetUserProcessingResults(value []UserProcessingResultable)() {
    m.userProcessingResults = value
}
// SetVersion sets the version property value. The current version number of the workflow. Value is 1 when the workflow is first created. Supports $filter(eq).
func (m *Workflow) SetVersion(value *int32)() {
    m.version = value
}
// SetVersions sets the versions property value. The workflow versions that are available.
func (m *Workflow) SetVersions(value []WorkflowVersionable)() {
    m.versions = value
}