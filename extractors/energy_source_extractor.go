package extractor

import (
	"fmt"
	"survey_data_transfomation/structs" // Import the structs package
	"survey_data_transfomation/utils"
)

const FolderPath = "../result_files/extracted_results/"

func ExtractAndSaveEnergyData(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	observatoriesData := struct {
		Results []structs.ObservatoryEnergySources `json:"results"`
	}{}
	for _, result := range surveyResult.Results {
		observatoryData := structs.ObservatoryEnergySources{
			Name:          result.ObservatoryInformation.Observatory,
			Site:          result.ObservatoryInformation.ObservatorySite,
			EnergySources: result.ExistingSystem.SensorDetails.EnergySources,
		}
		observatoriesData.Results = append(observatoriesData.Results, observatoryData)
	}

	err = utils.WriteToJsonFile(observatoriesData, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
