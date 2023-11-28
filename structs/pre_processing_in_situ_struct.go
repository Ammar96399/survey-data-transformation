package structs

type PreProcessingInSitu struct {
	Name                string `json:"observatory"`
	Site                string `json:"observatorySite"`
	PreProcessingInSitu bool   `json:"preProcessingInSitu"`
	PreProcessingAnswer string `json:"preProcessingAnswer"`
}
