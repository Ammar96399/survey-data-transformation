package structs

type ExistingSensorHealthMonitoring struct {
	Observatory                       Observatory `json:"observatory"`
	DefectiveSensorDetectionTime      string      `json:"defectiveSensorDetectionTime"`
	AutomatedDefectiveSensorDetection string      `json:"automatedDefectiveSensorDetection"`
	ExistingAutomatedDetectionSystem  struct {
		Details                         string `json:"details"`
		Notification                    string `json:"notification"`
		DataHandlingForDefectiveSensors string `json:"dataHandlingForDefectiveSensors"`
	}
}
