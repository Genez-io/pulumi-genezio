package constants

import "os"

var environment = os.Getenv("ENVIRONMENT")

var API_URL string

func init() {
	if environment == "dev" {
		API_URL = "https://dev.api.genez.io"
	} else {
		API_URL = "https://api.genez.io"
	}
}