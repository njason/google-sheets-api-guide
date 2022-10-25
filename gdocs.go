package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

const Jwt = `{
  "type": "service_account",
  "project_id": "x",
  "private_key_id": "x",
  "private_key": "x",
  "client_email": "x",
  "client_id": "x",
  "auth_uri": "x",
  "token_uri": "x",
  "auth_provider_x509_cert_url": "x",
  "client_x509_cert_url": "x"
}
`

func gDocRun() {
	ctx := context.Background()

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.JWTConfigFromJSON([]byte(Jwt), "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "184UXClOR-R2QETUc6CC2cAac4ztlFvT3WcoQgCgAx3c"
	readRange := "Class Data!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s\n", row[0], row[4])
		}
	}
}
