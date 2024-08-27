package config

import "time"

type Config struct {
	/*
		Shiki API base url.

		Example:
			https://shikimori.one/
	*/
	BaseUrl string

	// Application name.
	UserAgent string

	// TickTimeout in milliseconds.
	TickTimeout time.Duration

	// RequestTimeout in milliseconds.
	RequestTimeout time.Duration

	/*
		Config for:
			https://shikimori.one/api/doc/1.0/animes
	*/
	Anime *AnimeConfig
}

type AnimeConfig struct {
	/*
		Must be one of: id, id_desc, ranked, kind,
		popularity, name, aired_on, episodes, status,
		random, ranked_random, ranked_shiki, created_at,
		created_at_desc.
	*/
	Order string
	/*
		Set to false to allow hentai, yaoi and yuri.
		Must be one of: true, false.
	*/
	Censored string

	// 50 maximum
	Limit uint
}
