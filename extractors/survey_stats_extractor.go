package extractor

import (
	"fmt"
	"strings"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractSurveyStats(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	var numberOfFullResponses int
	var numberOfPartialResponses int

	var participatingObservatories []structs.Observatory

	for _, result := range surveyResult.Results {
		if result.GeneralInformation.LastPage == "7" {
			numberOfFullResponses++
		} else {
			numberOfPartialResponses++
		}
		if result.ObservatoryInformation.Observatory != "" &&
			result.ObservatoryInformation.Observatory != " " &&
			result.ObservatoryInformation.Observatory != "X" &&
			containsIgnoreCase(participatingObservatories, structs.Observatory{Name: result.ObservatoryInformation.Observatory, Site: result.ObservatoryInformation.ObservatorySite}) == false {
			participatingObservatories = append(participatingObservatories, structs.Observatory{Name: result.ObservatoryInformation.Observatory, Site: result.ObservatoryInformation.ObservatorySite})
		}
	}

	err = utils.WriteToJsonFile(structs.SurveyStats{
		NumberOfFullResponses:      numberOfFullResponses,
		NumberOfPartialResponses:   numberOfPartialResponses,
		ParticipatingObservatories: participatingObservatories,
	}, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}

func containsIgnoreCase(slice []structs.Observatory, target structs.Observatory) bool {
	for _, value := range slice {
		if strings.EqualFold(value.Name, target.Name) && strings.EqualFold(value.Site, target.Site) {
			return true
		}
	}
	return false
}
