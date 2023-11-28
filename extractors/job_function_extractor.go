package extractor

import (
	"encoding/json"
	"fmt"
	"os"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractJobFunction(filePath string, outputFilePath string) error {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}

	var surveyResult structs.SurveyResult
	if err := json.Unmarshal(data, &surveyResult); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	var jobFunctions = struct {
		JobFunctions []string `json:"jobFunctions"`
	}{}

	for _, result := range surveyResult.Results {
		if result.ObservatoryInformation.JobFunction != "" &&
			result.ObservatoryInformation.JobFunction != " " &&
			result.ObservatoryInformation.JobFunction != "X" {
			jobFunctions.JobFunctions = append(jobFunctions.JobFunctions, result.ObservatoryInformation.JobFunction)
		}
	}

	err = utils.WriteToJsonFile(jobFunctions, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}
