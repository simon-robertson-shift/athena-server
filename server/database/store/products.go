package store

type Product struct {
	Id                int64
	AccountId         int64
	Name              string
	Description       string
	DescriptionVector []float32
	CreatedBy         int64
	CreatedAt         int64
	UpdatedAt         int64
}
