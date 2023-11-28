package structs

type ExistingDataClassification struct {
	Observatory           Observatory           `json:"observatory"`
	InvalidDataHandling   string                `json:"invalidDataHandling"`
	ClassificationProcess ClassificationProcess `json:"classificationProcess"`
}

type ClassificationProcess struct {
	Description            string `json:"description"`
	TimeSpent              string `json:"timeSpent"`
	ClassificationLocation string `json:"classificationLocation"`
}
