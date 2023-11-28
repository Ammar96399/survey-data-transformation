package structs

type ObservatoryEnergySources struct {
	Name          string        `json:"observatory"`
	Site          string        `json:"observatorySite"`
	EnergySources EnergySources `json:"energySources"`
}
