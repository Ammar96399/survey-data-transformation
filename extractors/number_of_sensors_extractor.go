package extractor

import (
	"fmt"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractNumberOfSensors(filePath string, outputFile string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	numberOfSensorsByObservatory := struct {
		Results []structs.NumberOfSensors `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		if result.ExistingSystem.SensorDetails.NumberOfSensors == "" {
			continue
		}
		data := structs.NumberOfSensors{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			NumberOfSensors: result.ExistingSystem.SensorDetails.NumberOfSensors,
		}
		numberOfSensorsByObservatory.Results = append(numberOfSensorsByObservatory.Results, data)
	}

	err = utils.WriteToJsonFile(numberOfSensorsByObservatory, FolderPath+outputFile)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
