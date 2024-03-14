package general_information

import (
	"fmt"
	"os"
	"survey_data_transfomation/extractors"
	"survey_data_transfomation/utils"
)

func ExtractObservatoriesGeolocationGeoJson(outputFileName string) error {
	files, err := os.ReadDir(extractors.ObservatoriesPath)

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var results = struct {
		Type     string `json:"type"`
		Features []struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"features"`
	}{}

	results.Type = "FeatureCollection"

	for _, file := range files {
		var content, err = utils.LoadSurveyResultManual(extractors.ObservatoriesPath + file.Name())
		var result = struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
		}{}

		if err != nil {
			fmt.Printf(file.Name() + "\n")
			fmt.Printf("Error during reading file: %v", err)
			return err
		}

		if content.ObservatoryInformation.GeoLocation.Longitude == 0 && content.ObservatoryInformation.GeoLocation.Latitude == 0 {
			continue
		}

		result.Type = "Feature"
		result.Geometry.Type = "Point"
		result.Geometry.Coordinates = append(result.Geometry.Coordinates, content.ObservatoryInformation.GeoLocation.Longitude)
		result.Geometry.Coordinates = append(result.Geometry.Coordinates, content.ObservatoryInformation.GeoLocation.Latitude)

		results.Features = append(results.Features, result)
	}

	err = utils.WriteToJsonFile(results, FolderPath+outputFileName)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}

	if os.Getenv("SAVE_TO_SVELTE") == "true" {
		err = utils.WriteToJsonFile(results, SvelteFolderPath+outputFileName)
		if err != nil {
			fmt.Printf("Error during writing to file: %v", err)
			return err
		}
	}
	return nil
}
