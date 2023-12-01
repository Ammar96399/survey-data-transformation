package structs

// ObservatoryCommunicationData represents the structure for each observatory's communication data
type ObservatoryCommunicationData struct {
	Observatory       Observatory             `json:"observatory"`
	CommunicationData CommunicationTechniques `json:"communicationData"`
}

// ObservatoriesCommunicationData represents the structure for aggregating communication data for all observatories
type ObservatoriesCommunicationData struct {
	Observatories []ObservatoryCommunicationData `json:"observatories"`
}
