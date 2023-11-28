package extractor

import (
	"fmt"
	"strings"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractPreProcessingInSitu(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	results := struct {
		Results []structs.PreProcessingInSitu `json:"results"`
	}{}
	for _, result := range surveyResult.Results {
		observatoryData := structs.PreProcessingInSitu{
			Name:                result.ObservatoryInformation.Observatory,
			Site:                result.ObservatoryInformation.ObservatorySite,
			PreProcessingInSitu: getIfPreProcessing(result.ExistingSystem.DataProcessing.InSituDataProcessing),
			PreProcessingAnswer: result.ExistingSystem.DataProcessing.InSituDataProcessing,
		}
		results.Results = append(results.Results, observatoryData)
	}

	err = utils.WriteToJsonFile(results, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}

func getIfPreProcessing(composition string) bool {
	// Add your logic to determine the type of infrastructure based on composition
	if strings.Contains(strings.ToLower(composition), "oui") {
		return true
	}
	if strings.Contains(strings.ToLower(composition), "non") {
		return false
	}
	if composition == "" || composition == " " || strings.ToLower(composition) == "n/a" || strings.Contains(strings.ToLower(composition), "pas de traitement") {
		return false
	}
	if len(composition) > 5 {
		return true
	}
	return false
}
