package store

type Feature struct {
	Id            int64
	AccountId     int64
	ProductId     int64
	Name          string
	Description   string
	PriorityName  string
	PriorityValue int64
	CreatedBy     int64
	CreatedAt     int64
	UpdatedAt     int64
}
