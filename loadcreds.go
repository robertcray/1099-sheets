//
// Robert Cray <rdcray@pm.me> 
// May 2023

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Read JSON file with google service acount, return info needed to connect
// Returns the email and private key contained in 1099.json (or whatever
// File is specified as "credfile" in 1099.toml)
func loadCreds(credFile string) (string, string) {

	jsonFile, err := os.Open(credFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	email := fmt.Sprint(result["client_email"])
	privatekey := fmt.Sprint(result["private_key"])
	return email, privatekey
}
