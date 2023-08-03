//
// Robert Cray <rdcray@pm.me> 
// May 2023

package main

import (
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

type gSheet struct {
	dt    string
	spent string
}

func readSheet(credFile, spreadSheetID, tabName string) []gSheet {
	email, privateKey := loadCreds(credFile)
	var results []gSheet

	conf := &jwt.Config{
		Email:      email,
		PrivateKey: []byte(privateKey),
		TokenURL:   "https://oauth2.googleapis.com/token",
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets.readonly",
		},
	}

	client := conf.Client(oauth2.NoContext)

	// Create service object for Google sheets
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Define the Sheet Name and fields to select
	readRange := tabName + "!" + "A2:C"
	//readRange := "Sheet1!A2:E"

	// Pull the data from the sheet
	resp, err := srv.Spreadsheets.Values.Get(spreadSheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Display pulled data
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			results = append(results, gSheet{
				dt:    fmt.Sprint(row[0]),
				spent: fmt.Sprint(row[2]),
			})
		}
	}
	return results
}
