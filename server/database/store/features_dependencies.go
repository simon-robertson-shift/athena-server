package store

type FeatureDependency struct {
	Id        int64
	FeatureId int64
	DependsOn int64
	CreatedAt int64
	UpdatedAt int64
}
