package sensors_overview

import (
	"fmt"
	"os"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractSensorsCommunicationTechniques(outputFileName string) error {
	files, err := os.ReadDir("../../echarts-report/src/lib/json_files/observatories/")

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Results []struct {
			Observatory             structs.Observatory             `json:"observatory"`
			CommunicationTechniques structs.CommunicationTechniques `json:"communicationTechniques"`
		} `json:"results"`
	}{}

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual("../../echarts-report/src/lib/json_files/observatories/" + file.Name())
		var result = struct {
			Observatory             structs.Observatory             `json:"observatory"`
			CommunicationTechniques structs.CommunicationTechniques `json:"communicationTechniques"`
		}{}

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		result.Observatory.Name = content.ObservatoryInformation.Observatory
		result.Observatory.Site = content.ObservatoryInformation.ObservatorySite
		result.CommunicationTechniques = content.ExistingSystem.SensorDetails.CommunicationTechniques

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
