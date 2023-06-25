package enum

type AccountStatus string

const (
	PENDING  AccountStatus = "PENDING"
	ACCEPTED AccountStatus = "ACCEPTED"
	REJECTED AccountStatus = "REJECTED"
)
