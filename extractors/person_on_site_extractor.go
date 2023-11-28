package extractor

import (
	"fmt"
	"survey_data_transfomation/structs" // Import the structs package
	"survey_data_transfomation/utils"
)

func ExtractPersonOnSite(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	// Aggregate energy data for all observatories
	observatoriesData := struct {
		Results []structs.PersonOnSite `json:"results"`
	}{}
	for _, result := range surveyResult.Results {
		observatoryData := structs.PersonOnSite{
			Name:   result.ObservatoryInformation.Observatory,
			Site:   result.ObservatoryInformation.ObservatorySite,
			Answer: result.Personnel.DedicatedPersonnel,
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
