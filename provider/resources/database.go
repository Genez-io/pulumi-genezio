package resources

import (
	"fmt"

	"github.com/Genez-io/pulumi-genezio/provider/requests"
	p "github.com/pulumi/pulumi-go-provider"
)

type Database struct{}

type DatabaseArgs struct {
	Name    string `pulumi:"name"`
	Type string `pulumi:"type"`
	Region string `pulumi:"region"`
	AuthToken string `pulumi:"authToken"`
}

type DatabaseState struct {

	DatabaseArgs

	
	DatabaseId string `pulumi:"databaseId"`
	URL string `pulumi:"url"`
}

// func (*Database) Read(ctx p.Context, name string, input DatabaseArgs) (DatabaseState, error) {
// state := DatabaseState{DatabaseArgs: input}

// 	return state, nil
// }

func (*Database) Create(ctx p.Context, name string, input DatabaseArgs, preview bool) (string, DatabaseState, error) {
	state := DatabaseState{DatabaseArgs: input}
	if preview {
		return name, state, nil
	}

	fmt.Println("Creating database")
	createDatabaseResponse,err := requests.CreateDatabase(input.Type, input.Region, input.AuthToken, input.Name)
	if err != nil {
		return name, state, err
	}



	state.DatabaseId = createDatabaseResponse.DatabaseId
	getDatabaseConnectionUrl, err := requests.GetDatabaseConnectionUrl(state.DatabaseId, input.AuthToken)
	if err != nil {
		return name, state, err
	}
	fmt.Printf("Database URL: %s\n", getDatabaseConnectionUrl)

	state.URL = getDatabaseConnectionUrl

	return name, state, nil
}
