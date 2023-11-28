package structs

// SurveyResult represents the structure of the survey data
type SurveyResult struct {
	Results []struct {
		GeneralInformation struct {
			ResponseID     string `json:"responseId"`
			SubmissionDate string `json:"submissionDate"`
			LastPage       string `json:"lastPage"`
			Language       string `json:"language"`
			SerialNumber   string `json:"serialNumber"`
		} `json:"generalInformation"`
		ObservatoryInformation struct {
			Observatory     string `json:"observatory"`
			ObservatorySite string `json:"observatorySite"`
			JobFunction     string `json:"jobFunction"`
			Question        string `json:"question"`
		} `json:"observatoryInformation"`
		Personnel struct {
			DedicatedPersonnel string `json:"dedicatedPersonnel"`
			OngoingProjects    string `json:"ongoingProjects"`
		} `json:"personnel"`
		ExistingSystem struct {
			SensorDetails struct {
				NumberOfSensors string `json:"numberOfSensors"`
				EnergySources   struct {
					ElectricalNetwork string `json:"electricalNetwork"`
					SolarPower        string `json:"solarPower"`
					InternalBatteries string `json:"internalBatteries"`
					OtherSources      string `json:"otherSources"`
				} `json:"energySources"`
				CommunicationTechniques struct {
					LoRa                    string `json:"LoRa"`
					LoRaWAN                 string `json:"LoRaWAN"`
					Zigbee                  string `json:"Zigbee"`
					Bluetooth               string `json:"Bluetooth"`
					Wifi                    string `json:"Wifi"`
					G2                      string `json:"G2"`
					G3                      string `json:"G3"`
					G4                      string `json:"G4"`
					G5                      string `json:"G5"`
					SDCard                  string `json:"SDCard"`
					InternalMemory          string `json:"InternalMemory"`
					WiredConnection         string `json:"WiredConnection"`
					OtherWirelessConnection string `json:"OtherWirelessConnection"`
				} `json:"communicationTechniques"`
			} `json:"sensorDetails"`
			InfrastructureDetails struct {
				OnsiteInfrastructure      string `json:"onsiteInfrastructure"`
				InfrastructureComposition string `json:"infrastructureComposition"`
				EnergySource              string `json:"energySource"`
				InternetConnexion         string `json:"internetConnexion"`
				DataTransmission          string `json:"dataTransmission"`
			} `json:"infrastructureDetails"`
			DataProcessing struct {
				InSituDataProcessing    string `json:"inSituDataProcessing"`
				MinimalTransferDelay    string `json:"minimalTransferDelay"`
				MaximalTransferDelay    string `json:"maximalTransferDelay"`
				MinimalPublicationDelay string `json:"minimalPublicationDelay"`
				MaximalPublicationDelay string `json:"maximalPublicationDelay"`
			} `json:"dataProcessing"`
			SensorHealthMonitoring struct {
				DefectiveSensorDetectionTime      string `json:"defectiveSensorDetectionTime"`
				AutomatedDefectiveSensorDetection string `json:"automatedDefectiveSensorDetection"`
				ExistingAutomatedDetectionSystem  struct {
					Details                         string `json:"details"`
					Notification                    string `json:"notification"`
					DataHandlingForDefectiveSensors string `json:"dataHandlingForDefectiveSensors"`
				} `json:"existingAutomatedDetectionSystem"`
			} `json:"sensorHealthMonitoring"`
			DataClassification struct {
				InvalidDataHandling   string `json:"invalidDataHandling"`
				ClassificationProcess struct {
					Description            string `json:"description"`
					TimeSpent              string `json:"timeSpent"`
					ClassificationLocation string `json:"classificationLocation"`
				} `json:"classificationProcess"`
			} `json:"dataClassification"`
			EventDetectionSystem struct {
				MonitoredEvents      string `json:"monitoredEvents"`
				EventFrequency       string `json:"eventFrequency"`
				AverageEventDuration string `json:"averageEventDuration"`
				EventTermination     string `json:"eventTermination"`
				RequiredData         string `json:"requiredData"`
				Automation           struct {
					AutomatedDetection string `json:"automatedDetection"`
					ManualIntegration  string `json:"manualIntegration"`
				} `json:"automation"`
			} `json:"eventDetectionSystem"`
			EventResponseSystem struct {
				ActivateAdditionalSensors string `json:"activateAdditionalSensors"`
			} `json:"eventResponseSystem"`
			RealTimeSystem struct {
				UseOfObservatoryData string `json:"useOfObservatoryData"`
			} `json:"realTimeSystem"`
		} `json:"existingSystem"`
		FutureSystem struct {
			SensorHealthMonitoring struct {
				IntegrateAdditionalData     string                 `json:"integrateAdditionalData"`
				AcceptableNotificationDelay string                 `json:"acceptableNotificationDelay"`
				NotificationPreferences     NotificationPreference `json:"notificationPreferences"`
			} `json:"sensorHealthMonitoring"`
			EventDetectionSystem struct {
				IntegrateExternalObservatories string `json:"integrateExternalObservatories"`
				IntegrateExternalSources       string `json:"integrateExternalSources"`
				ManualDataIntegration          string `json:"manualDataIntegration"`
				HigherFrequencyData            string `json:"higherFrequencyData"`
				LowerFrequencyData             string `json:"lowerFrequencyData"`
				ActivateNewSystems             string `json:"activateNewSystems"`
			} `json:"eventDetectionSystem"`
			RealTimeSystem struct {
				EnableRealTimeAccess        string `json:"enableRealTimeAccess"`
				CollectionFrequency         string `json:"collectionFrequency"`
				DataTransmissionRate        string `json:"dataTransmissionRate"`
				AcceptableTransmissionDelay string `json:"acceptableTransmissionDelay"`
			} `json:"realTimeSystem"`
			DataClassification struct {
				IdentifyOutOfRangeValues       string `json:"identifyOutOfRangeValues"`
				IdentifyPeaksJumps             string `json:"identifyPeaksJumps"`
				PreQualificationForPublication string `json:"preQualificationForPublication"`
				ProximityToModel               string `json:"proximityToModel"`
				ProvideModelForDoubtfulData    string `json:"provideModelForDoubtfulData"`
				ModelFormat                    string `json:"modelFormat"`
				OperationForIncorrectData      string `json:"operationForIncorrectData"`
				PrioritizeDataInPreProcessing  string `json:"prioritizeDataInPreProcessing"`
			} `json:"dataClassification"`
		} `json:"futureSystem"`
		ContactInformation struct {
			EmailAddress        string `json:"emailAddress"`
			ContactForInterview string `json:"contactForInterview"`
			UsefulLinks         string `json:"usefulLinks"`
		} `json:"contactInformation"`
	} `json:"results"`
}

type NotificationPreference struct {
	SMS                string `json:"sms"`
	PhoneCall          string `json:"phoneCall"`
	Email              string `json:"email"`
	MobileNotification string `json:"mobileNotification"`
}
