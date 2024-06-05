package model

type Items_t struct {
	Chrt_id      uint    `json:"chrt_id"`
	Track_number string  `json:"track_number"`
	Price        float64 `json:"price"`
	Rid          string  `json:"rid"`
	Name         string  `json:"name"`
	Sale         uint    `json:"sale"`
	Size         string  `json:"size"`
	Total_price  float64 `json:"total_price"`
	Nm_id        uint    `json:"nm_id"`
	Brand        string  `json:"brand"`
	Status       uint    `json:"status"`
}
