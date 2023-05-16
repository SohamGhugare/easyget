package models

type Aliases struct {
	Aliases []Alias `json:"aliases"`
}

type Alias struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

