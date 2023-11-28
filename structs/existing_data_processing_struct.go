package structs

type ExistingDataProcessing struct {
	Observatory             Observatory `json:"observatory"`
	InSituDataProcessing    string      `json:"inSituDataProcessing"`
	MinimalTransferDelay    string      `json:"minimalTransferDelay"`
	MaximalTransferDelay    string      `json:"maximalTransferDelay"`
	MinimalPublicationDelay string      `json:"minimalPublicationDelay"`
	MaximalPublicationDelay string      `json:"maximalPublicationDelay"`
}
