package extractor

import (
	"fmt"
	"strings"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractEvent(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.Event `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		event := structs.Event{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			MonitoredEvent:     result.ExistingSystem.EventDetectionSystem.MonitoredEvents,
			EventFrequency:     result.ExistingSystem.EventDetectionSystem.EventFrequency,
			EventDuration:      result.ExistingSystem.EventDetectionSystem.AverageEventDuration,
			EventTermination:   result.ExistingSystem.EventDetectionSystem.EventTermination,
			RequiredData:       result.ExistingSystem.EventDetectionSystem.RequiredData,
			AutomatedDetection: strings.ToLower(result.ExistingSystem.EventDetectionSystem.Automation.AutomatedDetection),
			ManualIntegration:  result.ExistingSystem.EventDetectionSystem.Automation.ManualIntegration,
		}
		results.Results = append(results.Results, event)
	}

	err = utils.WriteToJsonFile(results, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}

func getMonitoredEvent(result string) string {
	result = strings.ToLower(result)
	if result == "" || result == " " || result == "X" {
		return "Non spécifié"
	}
	if strings.Contains(result, "crue") || strings.Contains(result, "crues") {
		return "Crue"
	}
	if strings.Contains(result, "aucun") || strings.Contains(result, "pas d'évènement") {
		return "Aucun événement"
	}
	return result
}
