package structs

type ManualObservatoryResults struct {
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
		GeoLocation     struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geoLocation"`
		JobFunction string `json:"jobFunction"`
		Question    string `json:"question"`
	} `json:"observatoryInformation"`
	Personnel struct {
		DedicatedPersonnel        float64 `json:"dedicatedPersonnel"`
		DedicatedPersonnelComment string  `json:"dedicatedPersonnelComment"`
		OngoingProjects           bool    `json:"ongoingProjects"`
		OngoingProjectsComment    string  `json:"ongoingProjectsComment"`
	} `json:"personnel"`
	ExistingSystem struct {
		SensorDetails struct {
			NumberOfSensors         string                  `json:"numberOfSensors"`
			EnergySources           EnergySources           `json:"energySources"`
			CommunicationTechniques CommunicationTechniques `json:"communicationTechniques"`
		} `json:"sensorDetails"`
		InfrastructureDetails InfrastructureDetails `json:"infrastructureDetails"`
		DataProcessing        struct {
			InSituDataProcessing           bool    `json:"inSituDataProcessing"`
			InSituDataProcessingComment    string  `json:"inSituDataProcessingComment"`
			MinimalTransferDelay           float64 `json:"minimalTransferDelay"`
			MinimalTransferDelayComment    string  `json:"minimalTransferDelayComment"`
			MaximalTransferDelay           float64 `json:"maximalTransferDelay"`
			MaximalTransferDelayComment    string  `json:"maximalTransferDelayComment"`
			MinimalPublicationDelay        float64 `json:"minimalPublicationDelay"`
			MinimalPublicationDelayComment string  `json:"minimalPublicationDelayComment"`
			MaximalPublicationDelay        float64 `json:"maximalPublicationDelay"`
			MaximalPublicationDelayComment string  `json:"maximalPublicationDelayComment"`
		} `json:"dataProcessing"`
		SensorHealthMonitoring struct {
			DefectiveSensorDetectionTime        int    `json:"defectiveSensorDetectionTime"`
			DefectiveSensorDetectionTimeComment string `json:"defectiveSensorDetectionTimeComment"`
			AutomatedDefectiveSensorDetection   bool   `json:"automatedDefectiveSensorDetection"`
			ExistingAutomatedDetectionSystem    struct {
				Details                         any `json:"details"`
				Notification                    any `json:"notification"`
				DataHandlingForDefectiveSensors any `json:"dataHandlingForDefectiveSensors"`
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
			Events []struct {
				MonitoredEvents             string `json:"monitoredEvents"`
				MonitoredEventsComment      string `json:"monitoredEventsComment"`
				EventFrequency              int    `json:"eventFrequency"`
				EventFrequencyComment       string `json:"eventFrequencyComment"`
				AverageEventDuration        int    `json:"averageEventDuration"`
				AverageEventDurationComment string `json:"averageEventDurationComment"`
				EventTermination            string `json:"eventTermination"`
				RequiredData                string `json:"requiredData"`
			} `json:"events"`
			Automation Automation `json:"automation"`
		} `json:"eventDetectionSystem"`
		EventResponseSystem struct {
			ActivateAdditionalSensors        bool   `json:"activateAdditionalSensors"`
			ActivateAdditionalSensorsComment string `json:"activateAdditionalSensorsComment"`
		} `json:"eventResponseSystem"`
		RealTimeSystem struct {
			UseOfObservatoryData        bool   `json:"useOfObservatoryData"`
			UseOfObservatoryDataComment string `json:"useOfObservatoryDataComment"`
		} `json:"realTimeSystem"`
	} `json:"existingSystem"`
	FutureSystem struct {
		SensorHealthMonitoring struct {
			IntegrateAdditionalData            bool   `json:"integrateAdditionalData"`
			IntegrateAdditionalDataComment     string `json:"integrateAdditionalDataComment"`
			AcceptableNotificationDelay        int    `json:"acceptableNotificationDelay"`
			AcceptableNotificationDelayComment string `json:"acceptableNotificationDelayComment"`
			NotificationPreferences            struct {
				Sms                bool `json:"sms"`
				PhoneCall          bool `json:"phoneCall"`
				Email              bool `json:"email"`
				MobileNotification bool `json:"mobileNotification"`
			} `json:"notificationPreferences"`
		} `json:"sensorHealthMonitoring"`
		EventDetectionSystem struct {
			IntegrateExternalObservatories        bool   `json:"integrateExternalObservatories"`
			IntegrateExternalObservatoriesComment string `json:"integrateExternalObservatoriesComment"`
			IntegrateExternalSources              bool   `json:"integrateExternalSources"`
			IntegrateExternalSourcesComment       string `json:"integrateExternalSourcesComment"`
			ManualDataIntegration                 bool   `json:"manualDataIntegration"`
			ManualDataIntegrationComment          string `json:"manualDataIntegrationComment"`
			HigherFrequencyData                   bool   `json:"higherFrequencyData"`
			HigherFrequencyDataComment            string `json:"higherFrequencyDataComment"`
			LowerFrequencyData                    bool   `json:"lowerFrequencyData"`
			LowerFrequencyDataComment             string `json:"lowerFrequencyDataComment"`
			ActivateNewSystems                    bool   `json:"activateNewSystems"`
			ActivateNewSystemsComment             string `json:"activateNewSystemsComment"`
		} `json:"eventDetectionSystem"`
		RealTimeSystem struct {
			EnableRealTimeAccess               bool   `json:"enableRealTimeAccess"`
			EnableRealTimeAccessComment        string `json:"enableRealTimeAccessComment"`
			CollectionFrequency                any    `json:"collectionFrequency"`
			CollectionFrequencyComment         string `json:"collectionFrequencyComment"`
			DataTransmissionRate               any    `json:"dataTransmissionRate"`
			DataTransmissionRateComment        string `json:"dataTransmissionRateComment"`
			AcceptableTransmissionDelay        any    `json:"acceptableTransmissionDelay"`
			AcceptableTransmissionDelayComment string `json:"acceptableTransmissionDelayComment"`
		} `json:"realTimeSystem"`
		DataClassification struct {
			IdentifyOutOfRangeValues           bool   `json:"identifyOutOfRangeValues"`
			IdentifyPeaksJumps                 bool   `json:"identifyPeaksJumps"`
			PreQualificationForPublication     bool   `json:"preQualificationForPublication"`
			ProximityToModel                   bool   `json:"proximityToModel"`
			ProvideModelForDoubtfulData        bool   `json:"provideModelForDoubtfulData"`
			ProvideModelForDoubtfulDataComment string `json:"provideModelForDoubtfulDataComment"`
			ModelFormat                        string `json:"modelFormat"`
			OperationForIncorrectData          string `json:"operationForIncorrectData"`
			PrioritizeDataInPreProcessing      string `json:"prioritizeDataInPreProcessing"`
		} `json:"dataClassification"`
	} `json:"futureSystem"`
	ContactInformation struct {
		EmailAddress        string `json:"emailAddress"`
		ContactForInterview string `json:"contactForInterview"`
		UsefulLinks         string `json:"usefulLinks"`
	} `json:"contactInformation"`
}

type EnergySources struct {
	ElectricalNetwork int `json:"electricalNetwork"`
	SolarPower        int `json:"solarPower"`
	InternalBatteries int `json:"internalBatteries"`
	OtherSources      int `json:"otherSources"`
}

type CommunicationTechniques struct {
	LoRa                    int `json:"LoRa"`
	LoRaWAN                 int `json:"LoRaWAN"`
	Zigbee                  int `json:"Zigbee"`
	Bluetooth               int `json:"Bluetooth"`
	Wifi                    int `json:"Wifi"`
	G2                      int `json:"G2"`
	G3                      int `json:"G3"`
	G4                      int `json:"G4"`
	G5                      int `json:"G5"`
	SDCard                  int `json:"SDCard"`
	InternalMemory          int `json:"InternalMemory"`
	WiredConnection         int `json:"WiredConnection"`
	OtherWirelessConnection int `json:"OtherWirelessConnection"`
}

type InfrastructureDetails struct {
	OnsiteInfrastructure             bool     `json:"onsiteInfrastructure"`
	OnsiteInfrastructureComment      string   `json:"onsiteInfrastructureComment"`
	InfrastructureComposition        []string `json:"infrastructureComposition"`
	InfrastructureCompositionComment string   `json:"infrastructureCompositionComment"`
	EnergySource                     string   `json:"energySource"`
	InternetConnexion                string   `json:"internetConnexion"`
	DataTransmission                 string   `json:"dataTransmission"`
}

type Automation struct {
	AutomatedDetection        bool   `json:"automatedDetection"`
	AutomatedDetectionComment string `json:"automatedDetectionComment"`
	ManualIntegration         bool   `json:"manualIntegration"`
	ManualIntegrationComment  string `json:"manualIntegrationComment"`
}
