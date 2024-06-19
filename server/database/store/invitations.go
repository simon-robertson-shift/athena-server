package store

type Invitation struct {
	Id        int64
	AccountId int64
	RoleId    int64
	Email     string
	Token     string
	CreatedBy int64
	CreatedAt int64
	ExpiresAt int64
	UpdatedAt int64
}
