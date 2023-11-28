package structs

type Delays struct {
	Name                     string `json:"observatory"`
	Site                     string `json:"observatorySite"`
	MinimalTransmissionDelay string `json:"minimalTransmissionDelay"`
	MaximalTransmissionDelay string `json:"maximalTransmissionDelay"`
	MinimalPublishingDelay   string `json:"minimalPublishingDelay"`
	MaximalPublishingDelay   string `json:"maximalPublishingDelay"`
}
