package model

type ResultPage struct {
	Ingredients []Ingredient `json:"ingredients"`
	Name        string       `json:"name"`
	Picture     string       `json:"picture"`
	Price       float64      `json:"price"`
	Size        string       `json:"size"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
