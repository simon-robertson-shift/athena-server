package store

type MilestoneReport struct {
	Id                int64
	AccountId         int64
	MilestoneId       int64
	Content           string
	ContentVector     []float32
	DurationAccurancy int64
	CreatedAt         int64
	ExpiresAt         int64
	UpdatedAt         int64
}
