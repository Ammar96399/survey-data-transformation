package structs

type Delays struct {
	Observatory              Observatory `json:"observatory"`
	MinimalTransmissionDelay string      `json:"minimalTransmissionDelay"`
	MaximalTransmissionDelay string      `json:"maximalTransmissionDelay"`
	MinimalPublishingDelay   string      `json:"minimalPublishingDelay"`
	MaximalPublishingDelay   string      `json:"maximalPublishingDelay"`
}
