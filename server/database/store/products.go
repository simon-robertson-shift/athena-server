package store

type Product struct {
	Id              int64
	AccountId       int64
	Name            string
	Description     string
	DescriptionHTML string
	CreatedBy       int64
	CreatedAt       int64
	UpdatedAt       int64
}
