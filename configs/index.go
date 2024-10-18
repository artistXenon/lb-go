package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port        uint16 `json:"port"`
	SecurePort  uint16 `json:"secure_port"`
	DomainRoot  string `json:"domain_root"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

var config *Config = nil

func Load() (cfg *Config) {
	if config != nil {
		return config
	}
	config = &Config{}
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("opening config file:\n %v", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(config); err != nil {
		log.Fatalf("parsing config file:\n %v", err.Error())
	}

	return config
}
