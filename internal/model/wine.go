package model

// Wine basic information about the wine
type Wine struct {
	Name    string `json:"name"`
	Winery  string `json:"winery"`
	Vintage int    `json:"vintage"`
	Review  string `json:"review"`
}
