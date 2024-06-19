package store

type TaskReport struct {
	Id            int64
	AccountId     int64
	TaskId        int64
	Content       string
	ContentVector []float32
	CreatedAt     int64
	ExpiresAt     int64
	UpdatedAt     int64
}
