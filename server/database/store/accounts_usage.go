package store

type AccountUsage struct {
	Id        int64
	AccountId int64
	FileBytes int64
	FileCount int64
	UserCount int64
	CreatedAt int64
	UpdatedAt int64
}
