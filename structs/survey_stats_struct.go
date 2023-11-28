package structs

type SurveyStats struct {
	NumberOfFullResponses      int           `json:"numberOfFullResponses"`
	NumberOfPartialResponses   int           `json:"numberOfPartialResponses"`
	ParticipatingObservatories []Observatory `json:"participatingObservatories"`
}
