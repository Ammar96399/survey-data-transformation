package structs

// ObservatoryCommunicationData represents the structure for each observatory's communication data
type ObservatoryCommunicationData struct {
	Name              string                  `json:"observatory"`
	Site              string                  `json:"observatorySite"`
	CommunicationData CommunicationTechniques `json:"communicationData"`
}

// ObservatoriesCommunicationData represents the structure for aggregating communication data for all observatories
type ObservatoriesCommunicationData struct {
	Observatories []ObservatoryCommunicationData `json:"observatories"`
}
