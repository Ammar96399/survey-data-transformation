package health_monitoring

import (
	"fmt"
	"os"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

const FolderPath = "../result_files/manual_data/extracted_data_from_manual_observatories/existing/health_monitoring/"
const SvelteFolderPath = "../../echarts-report/src/lib/json_files/existing/health_monitoring/"

func ExtractDefectiveSensorDetectionTime(outputFileName string) error {
	files, err := os.ReadDir(extractors.ObservatoriesPath)

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Results []struct {
			Observatory                  structs.Observatory `json:"observatory"`
			DefectiveSensorDetectionTime int                 `json:"defectiveSensorDetectionTime"`
		} `json:"results"`
	}{}

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual(extractors.ObservatoriesPath + file.Name())
		var result = struct {
			Observatory                  structs.Observatory `json:"observatory"`
			DefectiveSensorDetectionTime int                 `json:"defectiveSensorDetectionTime"`
		}{}

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		result.Observatory.Name = content.ObservatoryInformation.Observatory
		result.Observatory.Site = content.ObservatoryInformation.ObservatorySite
		result.DefectiveSensorDetectionTime = content.ExistingSystem.SensorHealthMonitoring.DefectiveSensorDetectionTime

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
