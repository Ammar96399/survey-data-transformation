package future

import (
	"fmt"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractFutureSensorHealthMonitoring(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.FutureSensorHealthMonitoring `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		futureSensorHealthMonitoring := structs.FutureSensorHealthMonitoring{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			IntegrateAdditionalData: result.FutureSystem.SensorHealthMonitoring.IntegrateAdditionalData,
			AcceptableDelay:         result.FutureSystem.SensorHealthMonitoring.AcceptableNotificationDelay,
			NotificationPreference: extractNotificationPreference(
				result.FutureSystem.SensorHealthMonitoring.NotificationPreferences),
		}
		results.Results = append(results.Results, futureSensorHealthMonitoring)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}

func extractNotificationPreference(notificationPreference structs.NotificationPreference) []string {
	var result = make([]string, 0)
	if notificationPreference.SMS == "Oui" {
		result = append(result, "sms")
	}
	if notificationPreference.Email == "Oui" {
		result = append(result, "email")
	}
	if notificationPreference.PhoneCall == "Oui" {
		result = append(result, "phone")
	}
	if notificationPreference.MobileNotification == "Oui" {
		result = append(result, "mobileNotification")
	}
	return result
}
