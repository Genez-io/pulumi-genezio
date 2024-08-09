package domain

type ResponseStatus string

const (
	Success ResponseStatus = "ok"
	Failure ResponseStatus = "error"
)
