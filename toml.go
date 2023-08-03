//
// Robert Cray <rdcray@pm.me> 
// May 2023

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Google google
}

type google struct {
	Credfile    string
	Spreadsheet string
	Rate        float64
	Tabname     string
}

// Parse TOML file
// Example
// [google]
// credfile = "/usr/local/etc/1099.json"
// spreadsheet = "<spreadsheet-id for google sheet>"
// rate = 200.0
func parseToml(configFile string) (string, string, float64, string) {
	var config config

	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	_ = err

	credfile := config.Google.Credfile
	spreadsheet := config.Google.Spreadsheet
	rate := config.Google.Rate
	tabName := config.Google.Tabname
	if tabName == "" {
		tabName = "Sheet1"
	}
	return credfile, spreadsheet, rate, tabName
}
