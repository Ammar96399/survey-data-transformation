package future

import (
	"fmt"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractFutureDataClassification(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var results = struct {
		Results []structs.FutureDataClassification `json:"results"`
	}{}

	for _, result := range surveyResult.Results {
		futureDataClassification := structs.FutureDataClassification{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			IdentifyOutOfRangeValues:       result.FutureSystem.DataClassification.IdentifyOutOfRangeValues,
			IdentifyPeaksJumps:             result.FutureSystem.DataClassification.IdentifyPeaksJumps,
			PreQualificationForPublication: result.FutureSystem.DataClassification.PreQualificationForPublication,
			ProximityToModel:               result.FutureSystem.DataClassification.ProximityToModel,
			ProvideModelForDoubtfulData:    result.FutureSystem.DataClassification.ProvideModelForDoubtfulData,
			ModelFormat:                    result.FutureSystem.DataClassification.ModelFormat,
			OperationForIncorrectData:      result.FutureSystem.DataClassification.OperationForIncorrectData,
			PrioritizeDataInPreProcessing:  result.FutureSystem.DataClassification.PrioritizeDataInPreProcessing,
		}
		results.Results = append(results.Results, futureDataClassification)
	}

	err = utils.WriteToJsonFile(results, extractor.FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
