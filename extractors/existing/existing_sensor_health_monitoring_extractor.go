package existing

import (
	"fmt"
	"strings"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractExistingHealthMonitoring(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.ExistingSensorHealthMonitoring `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		existingSensorHealthMonitoring := structs.ExistingSensorHealthMonitoring{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			DefectiveSensorDetectionTime:      result.ExistingSystem.SensorHealthMonitoring.DefectiveSensorDetectionTime,
			AutomatedDefectiveSensorDetection: strings.ToLower(result.ExistingSystem.SensorHealthMonitoring.AutomatedDefectiveSensorDetection),
			ExistingAutomatedDetectionSystem:  result.ExistingSystem.SensorHealthMonitoring.ExistingAutomatedDetectionSystem,
		}
		results.Results = append(results.Results, existingSensorHealthMonitoring)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
