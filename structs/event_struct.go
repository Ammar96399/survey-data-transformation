package structs

type Event struct {
	Observatory        Observatory `json:"observatory"`
	MonitoredEvent     string      `json:"monitoredEvent"`
	EventFrequency     string      `json:"eventFrequency"`
	EventDuration      string      `json:"eventDuration"`
	EventTermination   string      `json:"eventTermination"`
	RequiredData       string      `json:"requiredResponse"`
	AutomatedDetection string      `json:"automatedDetection"`
	ManualIntegration  string      `json:"manualIntegration"`
}
