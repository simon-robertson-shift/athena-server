package store

type FeatureComment struct {
	Id            int64
	AccountId     int64
	FeatureId     int64
	Content       string
	ContentVector []float32
	CreatedBy     int64
	CreatedAt     int64
	UpdatedAt     int64
}
