package evaluation

import (
	"fmt"
	"os"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

var secondHeaders = []string{
	"Observatory",
	"Site",
	"Monitors events",
	"Event happens at least one a month",
	"Existing event detection",
	"Existing event response",
	"Interested in health monitoring",
	"Interested in event detection",
	"Interested in event response",
	"Score (out of 10)",
}
var secondScores = []float64{
	2,
	1.5,
	1,
	1,
	1.5,
	1.5,
	1.5,
}

func ExtractSecondEvaluation(outputFileName string) error {
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

	results.Headers = secondHeaders

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
			len(content.ExistingSystem.EventDetectionSystem.Events) != 0 &&
				content.ExistingSystem.EventDetectionSystem.Events != nil)

		var b = false
		for _, e := range content.ExistingSystem.EventDetectionSystem.Events {
			if !b {
				b = e.EventFrequency >= 12
			}
		}

		result.Evaluation = append(
			result.Evaluation,
			b)

		result.Evaluation = append(
			result.Evaluation,
			content.ExistingSystem.EventDetectionSystem.Automation.AutomatedDetection)

		result.Evaluation = append(
			result.Evaluation,
			content.ExistingSystem.EventResponseSystem.ActivateAdditionalSensors)

		result.Evaluation = append(
			result.Evaluation,
			content.FutureSystem.SensorHealthMonitoring.IntegrateAdditionalData)

		result.Evaluation = append(
			result.Evaluation,
			content.FutureSystem.EventDetectionSystem.IntegrateExternalObservatories ||
				content.FutureSystem.EventDetectionSystem.IntegrateExternalSources ||
				content.FutureSystem.EventDetectionSystem.ManualDataIntegration)

		result.Evaluation = append(
			result.Evaluation,
			content.FutureSystem.EventDetectionSystem.LowerFrequencyData ||
				content.FutureSystem.EventDetectionSystem.HigherFrequencyData ||
				content.FutureSystem.EventDetectionSystem.ActivateNewSystems)

		for i, r := range result.Evaluation {
			if r {
				result.Score += secondScores[i]
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
