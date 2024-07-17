package anime

type ListItemModel struct {
	ID      *uint64   `json:"id"`
	Name    *string `json:"name"`
	Russian *string `json:"russian"`
	Image   *struct {
		Original *string `json:"original"`
		Preview  *string `json:"preview"`
		X96      *string `json:"x96"`
		X48      *string `json:"x48"`
	} `json:"image"`
	URL           *string `json:"url"`
	Kind          *string `json:"kind"`
	Score         *string `json:"score"`
	Status        *string `json:"status"`
	Episodes      *uint   `json:"episodes"`
	EpisodesAired *uint   `json:"episodes_aired"`
	AiredOn       *string `json:"aired_on"`
	ReleasedOn    *string `json:"released_on"`
}

// Сравнить статус
func (m *ListItemModel) CompareStatus(status string) bool {
	return *m.Status == status
}