package event_detection

import (
	"fmt"
	"os"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractIntegrateExternalObservatories(outputFileName string) error {
	files, err := os.ReadDir(extractors.ObservatoriesPath)

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Results []struct {
			Observatory                    structs.Observatory `json:"observatory"`
			IntegrateExternalObservatories bool                `json:"integrateExternalObservatories"`
		} `json:"results"`
	}{}

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual(extractors.ObservatoriesPath + file.Name())
		var result = struct {
			Observatory                    structs.Observatory `json:"observatory"`
			IntegrateExternalObservatories bool                `json:"integrateExternalObservatories"`
		}{}

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		result.Observatory.Name = content.ObservatoryInformation.Observatory
		result.Observatory.Site = content.ObservatoryInformation.ObservatorySite
		result.IntegrateExternalObservatories = content.FutureSystem.EventDetectionSystem.IntegrateExternalObservatories

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
