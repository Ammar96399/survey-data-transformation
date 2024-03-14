package evaluation

import (
	"fmt"
	"os"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

const FolderPath = "../result_files/manual_data/extracted_data_from_manual_observatories/evaluation/"
const SvelteFolderPath = "../../echarts-report/src/lib/json_files/evaluation/"

var firstHeaders = []string{
	"Observatory",
	"Site",
	"Contains data loggers",
	"Monitors events",
	"More than 20 sensors",
	"Existing event detection",
	"Existing event response",
	"Interested in health monitoring",
	"Score (out of 10)",
}
var scores = []float64{
	2.25,
	2.25,
	2.25,
	1,
	1,
	1.25,
}

func ExtractFirstEvaluation(outputFileName string) error {
	files, err := os.ReadDir(extractors.ObservatoriesPath)

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Headers []string `json:"headers"`
		Rows    []struct {
			Observatory structs.Observatory `json:"observatory"`
			Evaluation  []bool              `json:"evaluation"`
			Score       float64             `json:"score"`
		} `json:"rows"`
	}{}

	results.Headers = firstHeaders

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual(extractors.ObservatoriesPath + file.Name())

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		var result = struct {
			Observatory structs.Observatory `json:"observatory"`
			Evaluation  []bool              `json:"evaluation"`
			Score       float64             `json:"score"`
		}{}

		result.Evaluation = append(
			result.Evaluation,
			len(content.ExistingSystem.InfrastructureDetails.InfrastructureComposition) != 0)

		result.Evaluation = append(
			result.Evaluation,
			len(content.ExistingSystem.EventDetectionSystem.Events) != 0 &&
				content.ExistingSystem.EventDetectionSystem.Events != nil)

		result.Evaluation = append(
			result.Evaluation,
			content.ExistingSystem.SensorDetails.NumberOfSensors != "1 - 10" &&
				content.ExistingSystem.SensorDetails.NumberOfSensors != "10 - 20")

		result.Evaluation = append(
			result.Evaluation,
			content.ExistingSystem.EventDetectionSystem.Automation.AutomatedDetection)

		result.Evaluation = append(
			result.Evaluation,
			content.ExistingSystem.EventResponseSystem.ActivateAdditionalSensors)

		result.Evaluation = append(
			result.Evaluation,
			content.FutureSystem.SensorHealthMonitoring.IntegrateAdditionalData)

		for i, r := range result.Evaluation {
			if r {
				result.Score += scores[i]
			}
		}

		result.Observatory.Name = content.ObservatoryInformation.Observatory
		result.Observatory.Site = content.ObservatoryInformation.ObservatorySite

		results.Rows = append(results.Rows, result)
	}

	err = utils.WriteToJsonFile(results, FolderPath+outputFileName)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}

	if os.Getenv("SAVE_TO_SVELTE") == "true" {
		err = utils.WriteToJsonFile(results, SvelteFolderPath+outputFileName)
		if err != nil {
			fmt.Printf("Error during writing to file: %v", err)
			return err
		}
	}
	return nil
}
