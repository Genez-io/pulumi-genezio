package resources

import (
	"fmt"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	p "github.com/pulumi/pulumi-go-provider"
)

type Database struct{}

type DatabaseArgs struct {
	Name      string  `pulumi:"name"`
	Type      *string `pulumi:"type,optional"`
	Region    *string `pulumi:"region,optional"`
}

type DatabaseState struct {
	DatabaseArgs

	URL 	 string `pulumi:"url"`
	DatabaseId string `pulumi:"databaseId"`
}

// TODO - Improve this to handle changes for region and type - now they are ignored
// func (*Database) Diff(ctx p.Context, id string, olds DatabaseState, news DatabaseArgs) (p.DiffResponse, error) {
// 	diff := map[string]p.PropertyDiff{}

// 	if olds.Name != news.Name {
// 		diff["name"] = p.PropertyDiff{Kind: p.Update}
// 	}

// 	return p.DiffResponse{
// 		DeleteBeforeReplace: false,
// 		HasChanges:          len(diff) > 0,
// 		DetailedDiff:        diff,
// 	}, nil

// }

func (*Database) Read(ctx p.Context, id string, inputs DatabaseArgs, state DatabaseState) (string, DatabaseArgs, DatabaseState, error) {
	databases, err := requests.ListDatabases(ctx)
	if err != nil {
		return id, inputs, state, err
	}

	for _, database := range databases {
		if database.Id == state.DatabaseId {
			state.Name = database.Name
			return id, inputs, state, nil
		}
	}

	return id, inputs, DatabaseState{}, nil
}

func (*Database) Create(ctx p.Context, name string, input DatabaseArgs, preview bool) (string, DatabaseState, error) {
	state := DatabaseState{DatabaseArgs: input}
	if preview {
		return name, state, nil
	}

	
	databaseType := "postgres-neon"
	if input.Type != nil {
		databaseType = *input.Type
	}
	region := "aws-us-east-1"
	if input.Region != nil {
		region = *input.Region
	}

	fmt.Println("Creating database")
	createDatabaseResponse,err := requests.CreateDatabase(ctx, domain.CreateDatabaseRequest{
		Name: input.Name,
		Type: databaseType,
		Region: region,
	})
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

	// TODO - Already link the database to the project

	return name, state, nil
}
