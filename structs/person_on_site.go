package structs

type PersonOnSite struct {
	Observatory Observatory `json:"observatory"`
	Answer      string      `json:"answer"`
}
