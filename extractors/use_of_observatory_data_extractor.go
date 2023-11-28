package extractor

import (
	"fmt"
	"strings"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractUseOfObservatoryData(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.UseOfObservatoryData `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		useOfObservatoryData := structs.UseOfObservatoryData{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			UseOfObservatoryData: strings.ToLower(result.ExistingSystem.RealTimeSystem.UseOfObservatoryData),
		}
		results.Results = append(results.Results, useOfObservatoryData)
	}

	err = utils.WriteToJsonFile(results, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
