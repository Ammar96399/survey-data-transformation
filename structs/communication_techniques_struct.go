package structs

// CommunicationTechniques represents the structure of communication techniques data
type CommunicationTechniques struct {
	LoRa                    string `json:"LoRa"`
	LoRaWAN                 string `json:"LoRaWAN"`
	Zigbee                  string `json:"Zigbee"`
	Bluetooth               string `json:"Bluetooth"`
	Wifi                    string `json:"Wifi"`
	G2                      string `json:"G2"`
	G3                      string `json:"G3"`
	G4                      string `json:"G4"`
	G5                      string `json:"G5"`
	SDCard                  string `json:"SDCard"`
	InternalMemory          string `json:"InternalMemory"`
	WiredConnection         string `json:"WiredConnection"`
	OtherWirelessConnection string `json:"OtherWirelessConnection"`
}
