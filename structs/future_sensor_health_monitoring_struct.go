package structs

type FutureSensorHealthMonitoring struct {
	Observatory             Observatory `json:"observatory"`
	IntegrateAdditionalData string      `json:"integrateAdditionalData"`
	AcceptableDelay         string      `json:"acceptableDelay"`
	NotificationPreference  []string    `json:"notificationPreference"`
}
