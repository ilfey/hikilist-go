package shiki

type ShikiSource struct {
	name       string
	columnName string
}

/* Impliment extractor.Source interface. */

func (s *ShikiSource) GetName() string {
	return s.name
}

func (s *ShikiSource) GetColumnName() string {
	return s.columnName
}
