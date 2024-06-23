package store

type TaskReport struct {
	Id          int64
	AccountId   int64
	TaskId      int64
	Content     string
	ContentHTML string
	CreatedAt   int64
	ExpiresAt   int64
	UpdatedAt   int64
}
