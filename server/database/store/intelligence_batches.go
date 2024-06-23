package store

type IntelligenceBatch struct {
	Id           int64
	AccountId    int64
	RemoteId     int64
	Type         int64
	Token        string
	Status       string
	InputFileId  string
	OutputFileId string
	MilestoneId  int64
	TaskId       int64
	UserId       int64
	CreatedAt    int64
	UpdatesAt    int64
}

const (
	IntelligenceBatchType_MilestoneManagement = 10
	IntelligenceBatchType_MilestoneReport     = 11
	IntelligenceBatchType_TaskReport          = 20
	IntelligenceBatchType_UserReport          = 30
)

const (
	IntelligenceBatchStatus_Cancelled  = "cancelled"
	IntelligenceBatchStatus_Cancelling = "cancelling"
	IntelligenceBatchStatus_Completed  = "completed"
	IntelligenceBatchStatus_Expired    = "expires"
	IntelligenceBatchStatus_Failed     = "failed"
	IntelligenceBatchStatus_Finalizing = "finalizing"
	IntelligenceBatchStatus_InProgress = "in_progress"
	IntelligenceBatchStatus_Validating = "validating"
)
