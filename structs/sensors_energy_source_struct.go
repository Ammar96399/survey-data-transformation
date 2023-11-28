package structs

type EnergySources struct {
	ElectricalNetwork string `json:"electricalNetwork"`
	SolarPower        string `json:"solarPower"`
	InternalBatteries string `json:"internalBatteries"`
	OtherSources      string `json:"otherSources"`
}
