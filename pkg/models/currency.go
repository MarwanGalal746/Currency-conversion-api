package models

type Currency struct {
	Success    bool   `json:"success"`
	Timestamp  int    `json:"timestamp"`
	Historical bool   `json:"historical"`
	Base       string `json:"base"`
	Date       string `json:"date"`
	Rates      struct {
		USD float64 `json:"USD"`
	} `json:"rates"`
}
