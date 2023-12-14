package general_information

import (
	"fmt"
	"os"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

const FolderPath = "../result_files/manual_data/extracted_data_from_manual_observatories/general_information/"
const SvelteFolderPath = "../../echarts-report/src/lib/json_files/general_information/"

func ExtractObservatoriesGeolocation(outputFileName string) error {
	files, err := os.ReadDir("../result_files/manual_data/manual_observatories/")

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Results []struct {
			Observatory structs.Observatory `json:"observatory"`
			Geolocation structs.Geolocation `json:"geolocation"`
		} `json:"results"`
	}{}

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual("../result_files/manual_data/manual_observatories/" + file.Name())
		var result = struct {
			Observatory structs.Observatory `json:"observatory"`
			Geolocation structs.Geolocation `json:"geolocation"`
		}{}

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		result.Observatory.Name = content.ObservatoryInformation.Observatory
		result.Observatory.Site = content.ObservatoryInformation.ObservatorySite
		result.Geolocation = content.ObservatoryInformation.GeoLocation

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
