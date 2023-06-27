package enum

type RequestStatus string

const (
	REQUEST_PENDING  RequestStatus = "PENDING"
	REQUEST_ACCEPTED RequestStatus = "ACCEPTED"
	REQUEST_REJECTED RequestStatus = "REJECTED"
)
