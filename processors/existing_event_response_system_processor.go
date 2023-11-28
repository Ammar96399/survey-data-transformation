package processors

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ProcessExistingEventResponseSystem(filePath string, outputFilePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}

	surveyResults := struct {
		Results []structs.ExistingEventResponseSystem `json:"results"`
	}{}
	if err := json.Unmarshal(data, &surveyResults); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	type processedData struct {
		ActivateAdditionalSensors bool   `json:"activateAdditionalSensors"`
		Details                   string `json:"details"`
	}

	type result struct {
		UnprocessedData structs.ExistingEventResponseSystem `json:"unprocessedData"`
		ProcessedData   processedData                       `json:"processedData"`
	}

	processedEventResponseSystem := struct {
		Results []result `json:"results"`
	}{}

	for _, surveyResult := range surveyResults.Results {
		processedEventResponseSystem.Results = append(processedEventResponseSystem.Results, result{
			UnprocessedData: surveyResult,
			ProcessedData: processedData{
				ActivateAdditionalSensors: getIfActivateAdditionalSensors(surveyResult.ActivateAdditionalSensors),
				Details:                   getDetails(surveyResult.ActivateAdditionalSensors),
			},
		})
	}

	err = utils.WriteToJsonFile(processedEventResponseSystem, FOLDER_PATH+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}

func getIfActivateAdditionalSensors(eventResponseSystem string) bool {
	if eventResponseSystem == "" || eventResponseSystem == " " || strings.ToLower(eventResponseSystem) == "n/a" || strings.Contains(strings.ToLower(eventResponseSystem), "non") {
		return false
	}
	if strings.Contains(strings.ToLower(eventResponseSystem), "oui") || strings.Contains(strings.ToLower(eventResponseSystem), "dans mon cas") {
		return true
	}
	return false
}

func getDetails(eventResponseSystem string) string {
	eventResponseSystem = strings.ReplaceAll(eventResponseSystem, "Dans mon cas, ", "")
	eventResponseSystem = strings.ReplaceAll(eventResponseSystem, "dans mon cas, ", "")
	eventResponseSystem = strings.ReplaceAll(eventResponseSystem, "Oui, ", "")
	eventResponseSystem = strings.ReplaceAll(eventResponseSystem, "oui, ", "")
	eventResponseSystem = strings.ReplaceAll(eventResponseSystem, "oui ", "")
	eventResponseSystem = strings.ReplaceAll(eventResponseSystem, "Oui ", "")
	return eventResponseSystem
}
