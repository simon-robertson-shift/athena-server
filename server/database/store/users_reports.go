package store

type UserReport struct {
	Id        int64
	AccountId int64
	UserId    int64
	Content   string
	CreatedAt int64
	ExpiresAt int64
	UpdatedAt int64
}
