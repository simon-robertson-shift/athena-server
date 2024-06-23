package store

type IntelligenceFile struct {
	Id        int64
	AccountId int64
	RemoteId  string
	Purpose   string
	Name      string
	Size      int64
	CreatedAt int64
	UpdatesAt int64
}

const (
	IntelligenceFilePurpose_Assistants = "assistants"
	IntelligenceFilePurpose_Batch      = "batch"
	IntelligenceFilePurpose_Vision     = "vision"
)
