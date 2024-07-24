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

func (*Database) Read(ctx p.Context, id string, inputs DatabaseArgs, state DatabaseState) (string, DatabaseArgs, DatabaseState , error) {

	finalState := DatabaseState{
		DatabaseArgs: inputs,
		DatabaseId: state.DatabaseId,
		URL: state.URL,
	}

	databases, err := requests.ListDatabases(inputs.AuthToken)
	if err != nil {
		return id, inputs, state, err
	}

	for _, database := range databases {
		if database.Id == state.DatabaseId {
			finalState.Name = database.Name
			finalState.Type = database.Type
			finalState.Region = database.Region
			return id, inputs, finalState, nil
		}
	}

	finalState = DatabaseState{}

	return id, inputs, finalState, nil
}

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
	
	fmt.Println("Getting database connection url")
	getDatabaseConnectionUrl, err := requests.GetDatabaseConnectionUrl(state.DatabaseId, input.AuthToken)
	if err != nil {
		return name, state, err
	}

	state.URL = getDatabaseConnectionUrl

	return name, state, nil
}
