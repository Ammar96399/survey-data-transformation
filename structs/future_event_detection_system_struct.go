package structs

type FutureEventDetectionSystem struct {
	Observatory                    Observatory `json:"observatory"`
	IntegrateExternalObservatories string      `json:"integrate_external_observatories"`
	IntegrateExternalSources       string      `json:"integrate_external_sources"`
	ManualDataIntegration          string      `json:"manual_data_integration"`
	HigherFrequencyData            string      `json:"higher_frequency_data"`
	LowerFrequencyData             string      `json:"lower_frequency_data"`
	ActivateNewSystems             string      `json:"activate_new_system"`
}
