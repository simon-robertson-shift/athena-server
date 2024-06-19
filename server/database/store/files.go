package store

type File struct {
	Id        int64
	AccountId int64
	Path      string
	Name      string
	Type      string
	Size      int64
	CreatedBy int64
	CreatedAt int64
	UpdatedAt int64
}
