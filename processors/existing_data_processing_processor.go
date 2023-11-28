package processors

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

const FOLDER_PATH = "../result_files/transformed_results/"

func ProcessExistingDataProcessing(filePath string, outputFilePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}

	surveyResult := struct {
		Results []structs.ExistingDataProcessing `json:"results"`
	}{}
	if err := json.Unmarshal(data, &surveyResult); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	observatoriesData := struct {
		Results []structs.PreProcessingInSitu `json:"results"`
	}{}
	for _, result := range surveyResult.Results {
		observatoryData := structs.PreProcessingInSitu{
			Name:                result.Observatory.Name,
			Site:                result.Observatory.Site,
			PreProcessingInSitu: getIfPreProcessing(result.InSituDataProcessing),
			PreProcessingAnswer: result.InSituDataProcessing,
		}
		observatoriesData.Results = append(observatoriesData.Results, observatoryData)
	}

	err = utils.WriteToJsonFile(observatoriesData, FOLDER_PATH+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
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
