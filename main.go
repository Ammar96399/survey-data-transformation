package main

import (
	"fmt"
	"os"
	"survey_data_transfomation/extractors/new_extractors/existing/event_detection"
	"survey_data_transfomation/extractors/new_extractors/existing/event_response"
	"survey_data_transfomation/extractors/new_extractors/existing/health_monitoring"
	"survey_data_transfomation/extractors/new_extractors/existing/pre_processing"
	"survey_data_transfomation/extractors/new_extractors/existing/sensors_overview"
	event_detection2 "survey_data_transfomation/extractors/new_extractors/future/event_detection"
	event_response2 "survey_data_transfomation/extractors/new_extractors/future/event_response"
)

func main() {
	// Specify the input JSON file and output JSON file
	/*	inputFile := "../result_files/results.json"
		inputFolder := "../result_files/extracted_results/"
		fullInputFile := "../result_files/full_results.json"
		inputFolderGeneratedByAi := "../result_files/generated_by_ai/"
		inputFolderTransormed := "../result_files/transformed_results/"*/

	// Call the extractor function
	/*	if err := extractor.ExtractAndSaveEnergyData(inputFile, "observatories_energy_data.json"); err != nil {
			fmt.Println("Error:", err)
		}
		if err := extractor.ExtractAndSaveCommunicationData(inputFile, "observatories_communication_data.json"); err != nil {
			fmt.Print("Error:", err)
		}
		if err := extractor.ExtractPersonOnSite(inputFile, "observatories_person_on_site.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractAndSaveInfrastructureInSituData(inputFile, "observatories_infrastructure_in_situ.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractPreProcessingInSitu(inputFile, "observatories_pre_processing_in_situ.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractAndSaveDelaysData(inputFile, "observatories_delays_data.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractObservatory(inputFile, "observatories.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractSurveyStats(fullInputFile, "survey_stats.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractJobFunction(fullInputFile, "job_functions.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractNumberOfSensors(fullInputFile, "number_of_sensors.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := extractor.ExtractEvent(inputFile, "event.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := existing.ExtractExistingDataProcessing(inputFile, "existing_data_processing.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := processors.ProcessExistingDataProcessing(inputFolder+"existing_data_processing.json", "processed_existing_data_transformation.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := existing.ExtractExistingHealthMonitoring(inputFile, "existing_health_monitoring.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := existing.ExtractExistingDataClassification(inputFile, "existing_data_classification.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := extractor.ExtractUseOfObservatoryData(inputFile, "use_of_observatory_data.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := existing.ExtractExistingEventResponseSystem(inputFile, "existing_event_response_system.json"); err != nil {
			fmt.Print("Error", err)
		}
		if err := processors.ProcessExistingEventResponseSystem(inputFolder+"existing_event_response_system.json", "processed_existing_event_response_system.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := future.ExtractFutureSensorHealthMonitoring(inputFile, "future_sensor_health_monitoring.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := future.ExtractFutureEventDetectionSystem(inputFile, "future_event_detection_system.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := future.ExtractFutureRealTimeSystem(inputFile, "future_real_time_system.json"); err != nil {
			fmt.Print("Error", err)
		}

		if err := future.ExtractFutureDataClassification(inputFile, "future_data_classification.json"); err != nil {
			fmt.Print("Error", err)
		}

		filePaths := []string{
			inputFolder + "existing_data_classification.json",
			inputFolder + "existing_data_processing.json",
			inputFolder + "existing_event_response_system.json",
			inputFolder + "existing_health_monitoring.json",
			inputFolder + "number_of_sensors.json",
			inputFolder + "observatories_communication_data.json",
			inputFolder + "observatories_energy_data.json",
			inputFolder + "observatories_infrastructure_in_situ.json",
			inputFolder + "observatories_pre_processing_in_situ.json",
			inputFolder + "use_of_observatory_data.json",
			inputFolderGeneratedByAi + "future_data_classification.json",
			inputFolderGeneratedByAi + "future_event_detection_system.json",
			inputFolderGeneratedByAi + "future_real_time_system.json",
			inputFolderGeneratedByAi + "future_sensor_health_monitoring.json",
			inputFolderGeneratedByAi + "minimal_maximal_publication_delay.json",
			inputFolderGeneratedByAi + "minimal_maximal_transmission_delay.json",
			inputFolderGeneratedByAi + "observatories_delays_in_days.json",
			inputFolderGeneratedByAi + "observatories_person_on_site.json",
			inputFolderGeneratedByAi + "existing_event_detection_system.json",
			inputFolderTransormed + "processed_existing_data_transformation.json",
			inputFolderTransormed + "processed_existing_event_response_system.json",
		}

		fileContent, err := os.ReadFile(inputFolder + "observatories.json")
		if err != nil {
			log.Fatal(err)
		}

		type ObservatoryData struct {
			Results []structs.Observatory `json:"results"`
		}
		// Unmarshal JSON data into a struct
		var observatoryData ObservatoryData
		err = json.Unmarshal(fileContent, &observatoryData)
		if err != nil {
			log.Fatal(err)
		}

		// Run WriteObservatoryResultsToFile for each observatory
		for _, result := range observatoryData.Results {
			observatory := structs.Observatory{Name: result.Name, Site: result.Site}
			err := observatory_data.WriteObservatoryResultsToFile(filePaths, observatory)
			if err != nil {
				fmt.Println(err)
			}
		}

		if err := processors.ProcessAverageMinMaxTransmissionDelay(); err != nil {
			fmt.Print("Error", err)
		}*/

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
}
