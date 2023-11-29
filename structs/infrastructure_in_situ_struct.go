package structs

type ObservatoryInfrastructureInSitu struct {
	Observatory               Observatory `json:"observatory"`
	InfrastructureInSitu      bool        `json:"infrastructureInSitu"`
	InfrastructureComposition string      `json:"infrastructureComposition"`
	EnergySource              string      `json:"energySource"`
	InternetConnexion         string      `json:"internetConnexion"`
	DataTransmission          string      `json:"dataTransmission"`
	TypeOfInfrastructure      string      `json:"typeOfInfrastructure"`
}
