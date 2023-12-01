package structs

type PreProcessingInSitu struct {
	Observatory         Observatory `json:"observatory"`
	PreProcessingInSitu bool        `json:"preProcessingInSitu"`
	PreProcessingAnswer string      `json:"preProcessingAnswer"`
}
