package existing

import (
	"fmt"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractExistingEventResponseSystem(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.ExistingEventResponseSystem `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		existingEventResponseSystem := structs.ExistingEventResponseSystem{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			ActivateAdditionalSensors: result.ExistingSystem.EventResponseSystem.ActivateAdditionalSensors,
		}
		results.Results = append(results.Results, existingEventResponseSystem)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
