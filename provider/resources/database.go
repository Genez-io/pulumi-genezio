package resources

import (
	"fmt"

	"github.com/Genez-io/pulumi-genezio/provider/requests"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
)

type Database struct{}

type DatabaseArgs struct {
	Name    string `pulumi:"name"`
	Type string `pulumi:"type"`
	Region string `pulumi:"region"`
}

type DatabaseState struct {

	DatabaseArgs

	
	DatabaseId string `pulumi:"databaseId"`
	URL string `pulumi:"url"`
}

func (*Database) Read(ctx p.Context, id string, inputs DatabaseArgs, state DatabaseState) (string, DatabaseArgs, DatabaseState , error) {
	authToken, err := utils.IsLoggedIn(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	ctx = p.CtxWithValue(ctx, "authToken", authToken)

	finalState := DatabaseState{
		DatabaseArgs: inputs,
		DatabaseId: state.DatabaseId,
		URL: state.URL,
	}

	databases, err := requests.ListDatabases(ctx)
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
	authToken, err := utils.IsLoggedIn(ctx)
	if err != nil {
		return name, DatabaseState{}, err
	}
	ctx = p.CtxWithValue(ctx, "authToken", authToken)

	state := DatabaseState{DatabaseArgs: input}
	if preview {
		return name, state, nil
	}

	

	fmt.Println("Creating database")
	createDatabaseResponse,err := requests.CreateDatabase(ctx, input.Type, input.Region, input.Name)
	if err != nil {
		return name, state, err
	}
	state.DatabaseId = createDatabaseResponse.DatabaseId
	
	fmt.Println("Getting database connection url")
	getDatabaseConnectionUrl, err := requests.GetDatabaseConnectionUrl(ctx, state.DatabaseId)
	if err != nil {
		return name, state, err
	}

	state.URL = getDatabaseConnectionUrl

	return name, state, nil
}
