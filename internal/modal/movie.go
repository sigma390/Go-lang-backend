package modal

import "time"

type Movies struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	RunTime int `json:"run_time"`
	MPAARating string `json:"mpaa_rading"`
	Description string `json:"description"`
	Image string `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}