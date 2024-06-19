package store

type TaskChange struct {
	Id           int64
	AccountId    int64
	TaskId       int64
	Reason       string
	ReasonVector []float32
	CreatedBy    int64
	CreatedAt    int64
	UpdatedAt    int64
}
