package pre_processing

import (
	"fmt"
	"os"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractMinMaxPublicationDelays(outputFileName string) error {
	files, err := os.ReadDir("../result_files/manual_data/manual_observatories/")

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Results []struct {
			Observatory       structs.Observatory `json:"observatory"`
			PublicationDelays struct {
				Min float64 `json:"min"`
				Max float64 `json:"max"`
			} `json:"publicationDelays"`
		} `json:"results"`
	}{}

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual("../result_files/manual_data/manual_observatories/" + file.Name())
		var result = struct {
			Observatory       structs.Observatory `json:"observatory"`
			PublicationDelays struct {
				Min float64 `json:"min"`
				Max float64 `json:"max"`
			} `json:"publicationDelays"`
		}{}

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		result.Observatory.Name = content.ObservatoryInformation.Observatory
		result.Observatory.Site = content.ObservatoryInformation.ObservatorySite
		result.PublicationDelays.Min = content.ExistingSystem.DataProcessing.MinimalPublicationDelay
		result.PublicationDelays.Max = content.ExistingSystem.DataProcessing.MaximalPublicationDelay

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
