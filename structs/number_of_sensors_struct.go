package structs

type NumberOfSensors struct {
	Observatory     Observatory `json:"observatory"`
	NumberOfSensors string      `json:"numberOfSensors"`
}
