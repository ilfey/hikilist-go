package postgres

type Code = string

const (
	UniqueViolation     Code = "23505"
	ForeignKeyViolation Code = "23503"
)
