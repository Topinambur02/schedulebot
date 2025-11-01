package model

type Lesson struct {
	Name         string   `json:"name"`
	DateInterval string   `json:"dateInterval"`
	Link         *string  `json:"link"`
	Place        string   `json:"place"`
	Rooms        []string `json:"rooms"`
	Teachers     []string `json:"teachers"`
	TimeInterval string   `json:"timeInterval"`
}
