package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//
//// Shieldbox struct representing each shieldbox configuration
//type Shieldbox struct {
//	AMARISOFT_CALLBOX_ULTIMATE []int  `yaml:"AMARISOFT_CALLBOX_ULTIMATE"`
//	NOKIA                      []int  `yaml:"NOKIA"`
//	RADISYS                     []int  `yaml:"RADISYS"`
//	ManagementSwitch           string `yaml:"ManagementSwitch"`
//}

// Shieldbox struct representing each shieldbox configuration
type Shieldbox struct {
	AMARISOFT_CALLBOX_ULTIMATE []int
	NOKIA                      []int
	RADISYS                    []int
	ManagementSwitch           string
}

// Config struct to hold all shieldboxes from YAML
type Config struct {
	SHIELDBOX_1        Shieldbox `yaml:"SHIELDBOX_1"`
	SHIELDBOX_2        Shieldbox `yaml:"SHIELDBOX_2"`
	SHIELDBOX_3        Shieldbox `yaml:"SHIELDBOX_3"`
	SHIELDBOX_4        Shieldbox `yaml:"SHIELDBOX_4"`
	SHIELDBOX_5        Shieldbox `yaml:"SHIELDBOX_5"`
	AMARISOFT_SIMBOX_1 Shieldbox `yaml:"AMARISOFT_SIMBOX_1"`
}

var gConfig *Config

// readConfigFile reads the specified smsConfig file to setup some env variables
func readConfigFile(file string) (*Config, error) {

	f, err := os.Open(file)
	if err != nil {
		return defaultConfiguration(), err
	}
	defer f.Close()

	// Setup some defaults here
	// If the json file has values in it, the defaults will be overwritten
	conf := defaultConfiguration()

	// Read the configuration from json file
	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(conf)
	if err != nil {
		return conf, err
	}

	return conf, nil

}

func defaultConfiguration() *Config {

	return &Config{}

}

// GetConfiguration returns the configuration for the app.
// It will try to load it if it is not already loaded.
func GetConfiguration() *Config {
	if gConfig == nil {
		conf, err := readConfigFile("config.json")
		if err != nil {
			fmt.Println("Error loading config file: \n", err)
			fmt.Println("Using defaults...\n")
		}
		gConfig = conf

	}

	return gConfig
}
