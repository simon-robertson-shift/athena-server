package store

type Milestone struct {
	Id              int64
	AccountId       int64
	Name            string
	Description     string
	DescriptionHTML string
	Duration        int64
	EndsOn          int64
	CreatedAt       int64
	UpdatedAt       int64
}
