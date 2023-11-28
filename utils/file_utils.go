package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"survey_data_transfomation/structs"
)

func LoadSurveyResult(filePath string) (*structs.SurveyResult, error) {
	data, err := ReadFileData(filePath)
	if err != nil {
		return nil, err
	}

	surveyResult, err := UnmarshalData(data)
	if err != nil {
		return nil, err
	}

	return surveyResult, nil
}

func ReadFileData(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %v", err)
	}

	return data, nil
}

func UnmarshalData(data []byte) (*structs.SurveyResult, error) {
	var surveyResult structs.SurveyResult
	if err := json.Unmarshal(data, &surveyResult); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &surveyResult, nil
}

func WriteToJsonFile(data interface{}, outputFilePath string) error {
	resultJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	if err := os.WriteFile(outputFilePath, resultJSON, 0644); err != nil {
		return fmt.Errorf("error writing JSON file: %v", err)
	}

	fmt.Printf("Data written to %s\n", outputFilePath)
	return nil
}
