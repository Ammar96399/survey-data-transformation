package existing

import (
	"fmt"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractExistingDataClassification(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.ExistingDataClassification `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		existingDataClassification := structs.ExistingDataClassification{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			InvalidDataHandling: result.ExistingSystem.DataClassification.InvalidDataHandling,
			ClassificationProcess: structs.ClassificationProcess{
				Description:            result.ExistingSystem.DataClassification.ClassificationProcess.Description,
				TimeSpent:              result.ExistingSystem.DataClassification.ClassificationProcess.TimeSpent,
				ClassificationLocation: result.ExistingSystem.DataClassification.ClassificationProcess.ClassificationLocation,
			},
		}
		results.Results = append(results.Results, existingDataClassification)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
