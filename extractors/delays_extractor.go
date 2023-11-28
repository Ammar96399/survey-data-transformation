package extractor

import (
	"fmt"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractAndSaveDelaysData(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	observatoriesData := struct {
		Results []structs.Delays `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		observatoryData := structs.Delays{
			Name:                     result.ObservatoryInformation.Observatory,
			Site:                     result.ObservatoryInformation.ObservatorySite,
			MinimalTransmissionDelay: result.ExistingSystem.DataProcessing.MinimalTransferDelay,
			MaximalTransmissionDelay: result.ExistingSystem.DataProcessing.MaximalTransferDelay,
			MinimalPublishingDelay:   result.ExistingSystem.DataProcessing.MinimalPublicationDelay,
			MaximalPublishingDelay:   result.ExistingSystem.DataProcessing.MaximalPublicationDelay,
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
