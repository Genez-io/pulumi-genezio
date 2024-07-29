package resources

import (
	"log"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	p "github.com/pulumi/pulumi-go-provider"
)

type Database struct{}

type DatabaseArgs struct {
	ProjectName string  `pulumi:"projectName"`
	Name        string  `pulumi:"name"`
	Type        *string `pulumi:"type,optional"`
	Region      *string `pulumi:"region,optional"`
}

type DatabaseState struct {
	DatabaseArgs

	URL        string `pulumi:"url"`
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

func (*Database) Delete(ctx p.Context, id string, state DatabaseState) error {
	err := requests.DeleteDatabase(ctx, state.DatabaseId)
	if err != nil {
		log.Println("Error deleting database", err)
		return err
	}

	return nil
}

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
	region := "us-east-1"
	if input.Region != nil {
		region = *input.Region
	}

	createDatabaseResponse, err := requests.CreateDatabase(ctx, domain.CreateDatabaseRequest{
		Name:   input.Name,
		Type:   databaseType,
		Region: "aws-" + region,
	})
	if err != nil {
		log.Println("Error creating database", err)
		return name, state, err
	}
	state.DatabaseId = createDatabaseResponse.DatabaseId

	getDatabaseConnectionUrl, err := requests.GetDatabaseConnectionUrl(ctx, state.DatabaseId)
	if err != nil {
		return name, state, err
	}

	state.URL = getDatabaseConnectionUrl

	state.DatabaseId = createDatabaseResponse.DatabaseId

	projectDetails, err := requests.GetProjectDetails(ctx, input.ProjectName)
	if err != nil {
		log.Println("Error getting project details", err)
		return name, state, err
	}

	_, err = requests.LinkDatabaseToProject(ctx, domain.LinkDatabaseToProjectRequest{
		ProjectId:  projectDetails.Project.Id,
		StageId:    projectDetails.Project.ProjectEnvs[0].Id,
		DatabaseId: createDatabaseResponse.DatabaseId,
	})
	if err != nil {
		log.Println("Error linking database to project", err)
		return name, state, err
	}

	return name, state, nil
}
