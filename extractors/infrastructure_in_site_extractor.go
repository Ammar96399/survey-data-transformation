package extractor

import (
	"fmt"
	"strings"
	"survey_data_transfomation/structs"
	"survey_data_transfomation/utils"
)

func ExtractAndSaveInfrastructureInSituData(filePath string, outputFilePath string) error {
	var surveyResult, err = utils.LoadSurveyResult(filePath)

	observatoriesData := struct {
		Results []structs.ObservatoryInfrastructureInSitu `json:"results"`
	}{}
	for _, result := range surveyResult.Results {
		observatoryData := structs.ObservatoryInfrastructureInSitu{
			Observatory: structs.Observatory{
				Name: result.ObservatoryInformation.Observatory,
				Site: result.ObservatoryInformation.ObservatorySite,
			},
			InfrastructureInSitu:      result.ExistingSystem.InfrastructureDetails.OnsiteInfrastructure == "Oui",
			EnergySource:              result.ExistingSystem.InfrastructureDetails.EnergySource,
			InternetConnexion:         result.ExistingSystem.InfrastructureDetails.InternetConnexion,
			DataTransmission:          result.ExistingSystem.InfrastructureDetails.DataTransmission,
			InfrastructureComposition: result.ExistingSystem.InfrastructureDetails.InfrastructureComposition,
			TypeOfInfrastructure:      getTypeOfInfrastructure(result.ExistingSystem.InfrastructureDetails.InfrastructureComposition),
		}
		observatoriesData.Results = append(observatoriesData.Results, observatoryData)
	}

	err = utils.WriteToJsonFile(observatoriesData, FolderPath+outputFilePath)
	if err != nil {
		fmt.Printf("Error during writing to file: %v", err)
		return err
	}
	return nil
}

func getTypeOfInfrastructure(composition string) string {
	// Add your logic to determine the type of infrastructure based on composition
	if strings.Contains(strings.ToLower(composition), "propriétaire") {
		return "Systeme propriétaire"
	}
	// Add more conditions as needed

	return "" // Default value if type is not determined
}
