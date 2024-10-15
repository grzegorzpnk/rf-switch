package main

import "fmt"
import "rf-switch/pkg/src/config"

func main() {
	config := config.GetConfiguration()

	// Example: Find channels for multiple RAN providers and Shieldboxes
	ranProviders := []string{"RADISYS"}
	shieldboxes := []string{"SHIELDBOX_4"}

	channels, managementSwitches := findChannels(config, ranProviders, shieldboxes)

	// Output the results
	for i, managementSwitch := range managementSwitches {
		fmt.Printf("Management Switch: %s, Channels: %v\n", managementSwitch, channels[i])
	}
}

// Function to find channels based on provided RAN providers and Shieldboxes
func findChannels(config *config.Config, ranProviders []string, shieldboxes []string) ([][]int, []string) {
	var resultChannels [][]int
	var resultManagementSwitches []string

	for _, shieldboxName := range shieldboxes {
		// Fetch the Shieldbox based on name
		box := getShieldbox(*config, shieldboxName)

		// If shieldbox is not found, skip
		if box == nil {
			continue
		}

		// Accumulate channels from all specified RAN providers for the current shieldbox
		var channels []int
		for _, provider := range ranProviders {
			switch provider {
			case "AMARISOFT_CALLBOX_ULTIMATE":
				channels = append(channels, box.AMARISOFT_CALLBOX_ULTIMATE...)
			case "NOKIA":
				channels = append(channels, box.NOKIA...)
			case "RADISYS":
				channels = append(channels, box.RADISYS...)
			}
		}

		// Add the collected channels and management switch to the result
		resultChannels = append(resultChannels, channels)
		resultManagementSwitches = append(resultManagementSwitches, box.ManagementSwitch)
	}

	return resultChannels, resultManagementSwitches
}

// Helper function to get a shieldbox from the config based on the name
func getShieldbox(config config.Config, shieldboxName string) *config.Shieldbox {
	switch shieldboxName {
	case "SHIELDBOX_1":
		return &config.SHIELDBOX_1
	case "SHIELDBOX_2":
		return &config.SHIELDBOX_2
	case "SHIELDBOX_3":
		return &config.SHIELDBOX_3
	case "SHIELDBOX_4":
		return &config.SHIELDBOX_4
	case "SHIELDBOX_5":
		return &config.SHIELDBOX_5
	case "AMARISOFT_SIMBOX_1":
		return &config.AMARISOFT_SIMBOX_1
	default:
		// Return nil if shieldbox is not found
		return nil
	}
}
