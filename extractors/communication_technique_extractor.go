package extractor

import (
	"fmt"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

// ExtractAndSaveCommunicationData extracts communication techniques data for each observatory
// and saves the aggregated result in a single JSON file.
func ExtractAndSaveCommunicationData(filePath string, outputFilePath string) error {

	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var observatoriesData structs.ObservatoriesCommunicationData

	for _, result := range surveyResult.Results {
		communicationData := structs.ObservatoryCommunicationData{
			Name:              result.ObservatoryInformation.Observatory,
			Site:              result.ObservatoryInformation.ObservatorySite, // Update this line with the correct field
			CommunicationData: result.ExistingSystem.SensorDetails.CommunicationTechniques,
		}
		observatoriesData.Observatories = append(observatoriesData.Observatories, communicationData)
	}

	err = utils.WriteToJsonFile(observatoriesData, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
