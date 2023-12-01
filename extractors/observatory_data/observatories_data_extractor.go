package observatory_data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"survey_data_transfomation/structs"
)

func GetObservatoryDataFromFiles(filePaths []string, observatory structs.Observatory) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})

	for _, filePath := range filePaths {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		var data map[string]interface{}
		err = json.Unmarshal(fileContent, &data)
		if err != nil {
			return nil, err
		}

		var additionalData map[string]interface{}
		if resultArray, ok := data["results"].([]interface{}); ok {
			for _, item := range resultArray {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if observatoryData, ok := itemMap["observatory"].(map[string]interface{}); ok {
						if observatoryData["observatory"].(string) == observatory.Name &&
							observatoryData["observatorySite"].(string) == observatory.Site {
							additionalData = itemMap
							break
						}
					}
				}
			}
		}

		if additionalData != nil {
			delete(additionalData, "observatory")
			delete(additionalData, "observatorySite")

			fileSource := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

			resultMap[fileSource] = additionalData
		}
	}

	if len(resultMap) == 0 {
		return nil, fmt.Errorf("observatory '%s' not found in any file", observatory.Name)
	}

	return resultMap, nil
}

func WriteObservatoryResultsToFile(filePaths []string, observatory structs.Observatory) error {
	resultMap, err := GetObservatoryDataFromFiles(filePaths, observatory)
	if err != nil {
		return err
	}

	outputFileName := strings.ReplaceAll(fmt.Sprintf("%s_%s_all_data.json", strings.ReplaceAll(strings.ToLower(observatory.Name), " ", "_"), strings.ReplaceAll(strings.ToLower(observatory.Site), " ", "_")), "/", "")

	var finalData = map[string]interface{}{
		"observatory": observatory,
		"data":        resultMap,
	}

	resultBytes, err := json.Marshal(finalData)
	if err != nil {
		return err
	}

	err = os.WriteFile("../result_files/observatories/"+outputFileName, resultBytes, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Results written to %s\n", outputFileName)
	return nil
}
