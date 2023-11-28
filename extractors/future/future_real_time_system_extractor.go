package future

import (
	"fmt"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractFutureRealTimeSystem(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.FutureRealTimeSystem `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		futureRealTimeSystem := structs.FutureRealTimeSystem{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			EnableRealTimeAccess:        result.FutureSystem.RealTimeSystem.EnableRealTimeAccess,
			CollectionFrequency:         result.FutureSystem.RealTimeSystem.CollectionFrequency,
			DataTransmissionRate:        result.FutureSystem.RealTimeSystem.DataTransmissionRate,
			AcceptableTransmissionDelay: result.FutureSystem.RealTimeSystem.AcceptableTransmissionDelay,
		}
		results.Results = append(results.Results, futureRealTimeSystem)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
