package health_monitoring

import (
	"fmt"
	"os"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractAutomatedDefectiveSensorDetection(outputFileName string) error {
	files, err := os.ReadDir("../result_files/manual_data/manual_observatories/")

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Results []struct {
			Observatory                       structs.Observatory `json:"observatory"`
			AutomatedDefectiveSensorDetection bool                `json:"automatedDefectiveSensorDetection"`
		} `json:"results"`
	}{}

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual("../result_files/manual_data/manual_observatories/" + file.Name())
		var result = struct {
			Observatory                       structs.Observatory `json:"observatory"`
			AutomatedDefectiveSensorDetection bool                `json:"automatedDefectiveSensorDetection"`
		}{}

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		result.Observatory.Name = content.ObservatoryInformation.Observatory
		result.Observatory.Site = content.ObservatoryInformation.ObservatorySite
		result.AutomatedDefectiveSensorDetection = content.ExistingSystem.SensorHealthMonitoring.AutomatedDefectiveSensorDetection

		results.Results = append(results.Results, result)
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
