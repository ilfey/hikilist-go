package extractor

type Source interface {
	// Get source name.
	GetName() string

	// Get column name with anime id in database.
	GetColumnName() string
}
