package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Settings contains the config.json information for configuring the listening port, monitored application details, etc
type Settings struct {
	HTTPPort              string `json:"httpPort"`
	UseQueue              bool   `json:"useQueue"`
	QueueConnectionString string `json:"queueConnectionString"`
	QueueChannel          string `json:"queueChannel"`
	AutoAck               bool   `json:"autoAck"`
}

func (s *Settings) load() {
	// fill settings struct from file
	s.parseSettingsFile()
}

func (s *Settings) parseSettingsFile() {
	confFile := "skidpdf_settings.json"
	if len(os.Args) > 1 {
		confFile = os.Args[1]
	}

	file, err := os.Open(confFile)
	defer file.Close()
	if err != nil {
		log.Fatalf("Could not open settings file: %s", confFile)
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&s); err != nil {
		fmt.Println(err)
	}
}
