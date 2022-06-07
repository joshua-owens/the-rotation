package data

type Recipe struct {
	Title        string
	Time         int32
	Servings     int32
	Ingredients  []string
	Instructions []string
	Image        string
	Url          string
}
