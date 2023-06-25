package enum

type AccountStatus string

const (
	ACCOUNT_PENDING  AccountStatus = "PENDING"
	ACCOUNT_ACCEPTED AccountStatus = "ACCEPTED"
	ACCOUNT_REJECTED AccountStatus = "REJECTED"
)
