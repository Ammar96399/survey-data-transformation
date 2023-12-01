package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/extractors/existing"
	"survey_data_transfomation/extractors/future"
	"survey_data_transfomation/extractors/observatory_data"
	"survey_data_transfomation/processors"
	"survey_data_transfomation/structs"
)

func main() {
	// Specify the input JSON file and output JSON file
	inputFile := "../result_files/results.json"
	inputFolder := "../result_files/extracted_results/"
	fullInputFile := "../result_files/full_results.json"
	inputFolderGeneratedByAi := "../result_files/generated_by_ai/"
	inputFolderTransormed := "../result_files/transformed_results/"

	// Call the extractor function
	if err := extractor.ExtractAndSaveEnergyData(inputFile, "observatories_energy_data.json"); err != nil {
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
}
