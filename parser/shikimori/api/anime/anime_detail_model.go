package anime

import (
	"time"
)

type AnimeDetailModel struct {
	ID                *uint      `json:"id"`
	Name              *string    `json:"name"`
	Russian           *string    `json:"russian"`
	URL               *string    `json:"url"`
	Kind              *string    `json:"kind"`
	Score             *string    `json:"score"`
	Status            *string    `json:"status"`
	Episodes          *uint      `json:"episodes"`
	EpisodesAired     *uint      `json:"episodes_aired"`
	AiredOn           *string    `json:"aired_on"`
	ReleasedOn        *string    `json:"released_on"`
	Rating            *string    `json:"rating"`
	English           *[]string  `json:"english"`
	Japanese          *[]string  `json:"japanese"`
	Synonyms          *[]string  `json:"synonyms"`
	LicenseNameRu     *string    `json:"license_name_ru"`
	Duration          *int       `json:"duration"`
	Description       *string    `json:"description"`
	DescriptionHTML   *string    `json:"description_html"`
	DescriptionSource *string    `json:"description_source"`
	Franchise         *string    `json:"franchise"`
	Favoured          *bool      `json:"favoured"`
	Anons             *bool      `json:"anons"`
	Ongoing           *bool      `json:"ongoing"`
	ThreadID          *int       `json:"thread_id"`
	TopicID           *int       `json:"topic_id"`
	MyanimelistID     *uint      `json:"myanimelist_id"`
	UpdatedAt         *time.Time `json:"updated_at"`
	NextEpisodeAt     *time.Time `json:"next_episode_at"`
	Fansubbers        *[]string  `json:"fansubbers"`
	Fandubbers        *[]string  `json:"fandubbers"`
	Licensors         *[]string  `json:"licensors"`
	Image             *struct {
		Original *string `json:"original"`
		Preview  *string `json:"preview"`
		X96      *string `json:"x96"`
		X48      *string `json:"x48"`
	} `json:"image"`
	RatesScoresStats *[]struct {
		Name  *int `json:"name"`
		Value *int `json:"value"`
	} `json:"rates_scores_stats"`
	RatesStatusesStats *[]struct {
		Name  *string `json:"name"`
		Value *int    `json:"value"`
	} `json:"rates_statuses_stats"`
	Genres *[]struct {
		ID        *int    `json:"id"`
		Name      *string `json:"name"`
		Russian   *string `json:"russian"`
		Kind      *string `json:"kind"`
		EntryType *string `json:"entry_type"`
	} `json:"genres"`
	Studios *[]struct {
		ID           *int    `json:"id"`
		Name         *string `json:"name"`
		FilteredName *string `json:"filtered_name"`
		Real         *bool   `json:"real"`
		Image        *string `json:"image"`
	} `json:"studios"`
	Videos *[]struct {
		ID        *int    `json:"id"`
		URL       *string `json:"url"`
		ImageURL  *string `json:"image_url"`
		PlayerURL *string `json:"player_url"`
		Name      *string `json:"name"`
		Kind      *string `json:"kind"`
		Hosting   *string `json:"hosting"`
	} `json:"videos"`
	Screenshots *[]struct {
		Original *string `json:"original"`
		Preview  *string `json:"preview"`
	} `json:"screenshots"`
}

// Сравнить статус
func (m *AnimeDetailModel) CompareStatus(status string) bool {
	return *m.Status == status
}
