package store

type MilestoneReport struct {
	Id                int64
	AccountId         int64
	MilestoneId       int64
	Content           string
	ContentHTML       string
	DurationAccurancy int64
	CreatedAt         int64
	ExpiresAt         int64
	UpdatedAt         int64
}
