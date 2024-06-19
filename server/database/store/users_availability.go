package store

type UserAvailability struct {
	Id                int64
	UserId            int64
	DurationMonday    int64
	DurationTuesday   int64
	DurationWednesday int64
	DurationThursday  int64
	DurationFriday    int64
	IsAvailable       bool
	WeekStartingOn    int64
	WeekEndingOn      int64
	CreatedAt         int64
	ExpiresAt         int64
	UpdatedAt         int64
}
