package extractor

import (
	"fmt"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractObservatory(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	observatoriesData := struct {
		Results []structs.Observatory `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		observatoryData := structs.Observatory{
			Name: result.ObservatoryInformation.Observatory,
			Site: result.ObservatoryInformation.ObservatorySite,
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
