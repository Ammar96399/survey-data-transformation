package future

import (
	"fmt"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractFutureEventDetectionSystem(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.FutureEventDetectionSystem `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		futureEventDetectionSystem := structs.FutureEventDetectionSystem{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			IntegrateExternalObservatories: result.FutureSystem.EventDetectionSystem.IntegrateExternalObservatories,
			IntegrateExternalSources:       result.FutureSystem.EventDetectionSystem.IntegrateExternalSources,
			ManualDataIntegration:          result.FutureSystem.EventDetectionSystem.ManualDataIntegration,
			HigherFrequencyData:            result.FutureSystem.EventDetectionSystem.HigherFrequencyData,
			LowerFrequencyData:             result.FutureSystem.EventDetectionSystem.LowerFrequencyData,
			ActivateNewSystems:             result.FutureSystem.EventDetectionSystem.ActivateNewSystems,
		}
		results.Results = append(results.Results, futureEventDetectionSystem)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
