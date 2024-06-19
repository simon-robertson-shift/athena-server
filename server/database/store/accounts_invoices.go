package store

type AccountInvoice struct {
	Id          int64
	AccountId   int64
	PricePlanId int64
	Amount      int64
	DueOn       int64
	SettledAt   int64
	CreatedAt   int64
	UpdatedAt   int64
}
