package structs

type ObservatoryEnergySources struct {
	Observatory   Observatory   `json:"observatory"`
	EnergySources EnergySources `json:"energySources"`
}
