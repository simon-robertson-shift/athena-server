package store

type Registration struct {
	Id                      int64
	PricePlanId             int64
	CompanyName             string
	CompanyCountryCode      string
	CompanyCountryName      string
	PrimaryContactNameFirst string
	PrimaryContactNameLast  string
	PrimaryContactEmail     string
	Token                   string
	CreatedAt               int64
	ExpiresAt               int64
	UpdatedAt               int64
}
