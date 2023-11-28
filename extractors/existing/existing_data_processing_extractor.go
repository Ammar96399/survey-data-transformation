package existing

import (
	"fmt"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractExistingDataProcessing(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.ExistingDataProcessing `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		existingDataProcessing := structs.ExistingDataProcessing{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			InSituDataProcessing:    result.ExistingSystem.DataProcessing.InSituDataProcessing,
			MinimalTransferDelay:    result.ExistingSystem.DataProcessing.MinimalTransferDelay,
			MaximalTransferDelay:    result.ExistingSystem.DataProcessing.MaximalTransferDelay,
			MinimalPublicationDelay: result.ExistingSystem.DataProcessing.MinimalPublicationDelay,
			MaximalPublicationDelay: result.ExistingSystem.DataProcessing.MaximalPublicationDelay,
		}
		results.Results = append(results.Results, existingDataProcessing)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
