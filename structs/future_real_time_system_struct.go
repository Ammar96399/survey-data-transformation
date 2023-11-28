package structs

type FutureRealTimeSystem struct {
	Observatory                 Observatory `json:"observatory"`
	EnableRealTimeAccess        string      `json:"enable_real_time_access"`
	CollectionFrequency         string      `json:"collection_frequency"`
	DataTransmissionRate        string      `json:"data_transmission_rate"`
	AcceptableTransmissionDelay string      `json:"acceptable_transmission_delay"`
}
