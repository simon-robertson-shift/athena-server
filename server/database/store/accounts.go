package store

type Account struct {
	Id                     int64
	CompanyName            string
	CompanyCountryCode     string
	CompanyCountryName     string
	PricePlanId            int64
	CurrentFileStorageSize int64
	CurrentUserCount       int64
	CreatedAt              int64
	ExpiresAt              int64
	UpdatedAt              int64
}
