package enum

type ContextKey string

const (
	RequestIDContextKey ContextKey = "UniqueRequestId"
	UserIDContextKey    ContextKey = "UserID"
)
