package store

type Task struct {
	Id               int64
	FeatureId        int64
	UserId           int64
	Name             string
	Description      string
	DescriptionHTML  string
	PriorityName     string
	PriorityValue    int64
	StatusName       string
	StatusValue      int64
	ExpectedStartOn  int64
	ExpectedDuration int64
	StartedOn        int64
	EndedOn          int64
	Duration         int64
	CreatedBy        int64
	CreatedAt        int64
	UpdatedAt        int64
}
