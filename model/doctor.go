package model

type Doctor struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Specialty  string  `json:"specialty"`
	Education  string  `json:"education"`
	Experience int     `json:"experience"`
	Rating     float64 `json:"rating"`
}

type Specialty struct {
	Name string `json:"name"`
}
