package structs

type FutureDataClassification struct {
	Observatory                    Observatory `json:"observatory"`
	IdentifyOutOfRangeValues       string      `json:"identify_out_of_range_values"`
	IdentifyPeaksJumps             string      `json:"identify_peaks_jumps"`
	PreQualificationForPublication string      `json:"pre_qualification_for_publication"`
	ProximityToModel               string      `json:"proximity_models"`
	ProvideModelForDoubtfulData    string      `json:"provide_model_for_doubtful_data"`
	ModelFormat                    string      `json:"model_format"`
	OperationForIncorrectData      string      `json:"operation_for_incorrect_data"`
	PrioritizeDataInPreProcessing  string      `json:"prioritize_data_pre_processing"`
}
