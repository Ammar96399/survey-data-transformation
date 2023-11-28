package structs

type ExistingEventResponseSystem struct {
	Observatory               Observatory `json:"observatory"`
	ActivateAdditionalSensors string      `json:"activateAdditionalSensors"`
}
