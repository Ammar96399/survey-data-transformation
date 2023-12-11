package processors

import (
	"fmt"
	"os"
	"survey_data_transfomation/utils"
)

func ProcessAverageMinMaxTransmissionDelay() error {
	files, err := os.ReadDir("../result_files/manual_data/manual_observatories/")

	if err != nil {
		fmt.Printf("Error during reading directory: %v", err)
		return err
	}

	var minTransmissionTrue = 0.0
	var minTransmissionFalse = 0.0
	var maxTransmissionTrue = 0.0
	var maxTransmissionFalse = 0.0

	var minPublicationTrue = 0.0
	var minPublicationFalse = 0.0
	var maxPublicationFalse = 0.0
	var maxPublicationTrue = 0.0
	for _, f := range files {
		var results, err = utils.LoadSurveyResultManual("../result_files/manual_data/manual_observatories/" + f.Name())

		if err != nil {
			fmt.Printf(f.Name())
			fmt.Printf("Error during reading file: %v", err)
			return err
		}
		if results.ExistingSystem.RealTimeSystem.UseOfObservatoryData {
			minTransmissionTrue += float64(results.ExistingSystem.DataProcessing.MinimalTransferDelay)
			maxTransmissionTrue += float64(results.ExistingSystem.DataProcessing.MaximalTransferDelay)
			maxPublicationTrue += float64(results.ExistingSystem.DataProcessing.MaximalPublicationDelay)
			minPublicationTrue += float64(results.ExistingSystem.DataProcessing.MinimalPublicationDelay)
		} else {
			minTransmissionFalse += float64(results.ExistingSystem.DataProcessing.MinimalTransferDelay)
			maxTransmissionFalse += float64(results.ExistingSystem.DataProcessing.MaximalTransferDelay)
			maxPublicationFalse += float64(results.ExistingSystem.DataProcessing.MaximalPublicationDelay)
			minPublicationFalse += float64(results.ExistingSystem.DataProcessing.MinimalPublicationDelay)
		}
	}

	fmt.Printf("Average minimal transmission delay for true: %v\n", minTransmissionTrue/float64(len(files)))
	fmt.Printf("Average minimal transmission delay for false: %v\n", minTransmissionFalse/float64(len(files)))
	fmt.Printf("Average maximal transmission delay for true: %v\n", maxTransmissionTrue/float64(len(files)))
	fmt.Printf("Average maximal transmission delay for false: %v\n", maxTransmissionFalse/float64(len(files)))
	fmt.Printf("\n\n")
	fmt.Printf("Average minimal publication delay for true: %v\n", minPublicationTrue/float64(len(files)))
	fmt.Printf("Average minimal publication delay for false: %v\n", minPublicationFalse/float64(len(files)))
	fmt.Printf("Average maximal publication delay for true: %v\n", maxPublicationTrue/float64(len(files)))
	fmt.Printf("Average maximal publication delay for false: %v\n", maxPublicationFalse/float64(len(files)))

	return nil
}
