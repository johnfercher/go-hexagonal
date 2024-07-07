package userstatus

type Status string

const (
	Allowed    Status = "allowed"
	Restricted Status = "restricted"
	Denied     Status = "denied"
	Pending    Status = "pending"
)
