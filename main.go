package main

import (
	"fmt"
	"os"
	"survey_data_transfomation/extractors/new_extractors/evaluation"
	"survey_data_transfomation/extractors/new_extractors/existing/event_detection"
	"survey_data_transfomation/extractors/new_extractors/existing/event_response"
	"survey_data_transfomation/extractors/new_extractors/existing/health_monitoring"
	"survey_data_transfomation/extractors/new_extractors/existing/pre_processing"
	"survey_data_transfomation/extractors/new_extractors/existing/sensors_overview"
	event_detection2 "survey_data_transfomation/extractors/new_extractors/future/event_detection"
	event_response2 "survey_data_transfomation/extractors/new_extractors/future/event_response"
	health_monitoring2 "survey_data_transfomation/extractors/new_extractors/future/health_monitoring"
	pre_processing2 "survey_data_transfomation/extractors/new_extractors/future/pre_processing"
	real_time2 "survey_data_transfomation/extractors/new_extractors/future/real_time"
	"survey_data_transfomation/extractors/new_extractors/general_information"
)

func main() {

	err := os.Setenv("SAVE_TO_SVELTE", "true")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	/*
		Extractors for existing sensors overview
	*/
	if err := sensors_overview.ExtractSensorsEnergySources("sensors_energy_sources.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := sensors_overview.ExtractSensorsCommunicationTechniques("sensors_communication_techniques.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := sensors_overview.ExtractItInfrastructureInSitu("it_infrastructure_in_situ.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := sensors_overview.ExtractNumberOfSensors("number_of_sensors.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := sensors_overview.ExtractUsageByOperationalPlayers("usage_by_operational_players.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for existing event detection
	*/
	if err := event_detection.ExtractAutomation("automation.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for existing event response
	*/
	if err := event_response.ExtractActivateAdditionalSensors("activate_additional_sensors.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for pre-processing
	*/
	if err := pre_processing.ExtractMinMaxTransmissionDelays("minimal_maximal_transmission_delay.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := pre_processing.ExtractMinMaxPublicationDelays("minimal_maximal_publication_delay.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := pre_processing.ExtractInSituDataProcessing("in_situ_data_processing.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := pre_processing.ExtractInvalidDataHandling("invalid_data_handling.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for health monitoring
	*/
	if err := health_monitoring.ExtractDefectiveSensorDetectionTime("defective_sensor_detection_time.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := health_monitoring.ExtractAutomatedDefectiveSensorDetection("automated_defective_sensor_detection.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for future event detection
	*/
	if err := event_detection2.ExtractIntegrateExternalSources("integrate_external_sources.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := event_detection2.ExtractIntegrateExternalObservatories("integrate_external_observatories.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := event_detection2.ExtractManualDataIntegration("manual_data_integration.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for future event response
	*/
	if err := event_response2.ExtractLowerFrequencyData("lower_frequency_data.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := event_response2.ExtractActivateNewSystems("activate_new_systems.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := event_response2.ExtractHigherFrequencyData("higher_frequency_data.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for future health monitoring
	*/
	if err := health_monitoring2.ExtractIntegrateAdditionalData("integrate_additional_data.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := health_monitoring2.ExtractNotificationPreferences("notification_preferences.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := health_monitoring2.ExtractAcceptableNotificationDelay("acceptable_notification_delay.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for future pre-processing
	*/
	if err := pre_processing2.ExtractAcceptableNotificationDelay("acceptable_notification_delay.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for future real time
	*/
	if err := real_time2.ExtractEnableRealTimeAccess("enable_real_time_access.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Extractors for general information
	*/
	if err := general_information.ExtractObservatoriesGeolocationGeoJson("test.geojson.json"); err != nil {
		fmt.Print("Error", err)
	}

	/*
		Evaluations
	*/
	if err := evaluation.ExtractFirstEvaluation("first_evaluation.json"); err != nil {
		fmt.Print("Error", err)
	}
	if err := evaluation.ExtractSecondEvaluation("second_evaluation.json"); err != nil {
		fmt.Print("Error", err)
	}
}
