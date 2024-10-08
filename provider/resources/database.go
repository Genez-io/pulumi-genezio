package resources

import (
	"context"
	_ "embed"
	"log"
	"strings"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/requests"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Database struct{}

type DatabaseArgs struct {
	Name    string          `pulumi:"name"`
	Project *domain.Project `pulumi:"project,optional"`
	Type    *string         `pulumi:"type,optional"`
	Region  *string         `pulumi:"region,optional"`
}

type DatabaseState struct {
	DatabaseArgs

	URL        string `pulumi:"url" provider:"secret"`
	DatabaseId string `pulumi:"databaseId"`
}

//go:embed documentation/project.md
var databaseDocumentation string

func (r *Database) Annotate(a infer.Annotator) {
	a.Describe(&r, databaseDocumentation)
}

func (r *DatabaseArgs) Annotate(a infer.Annotator) {
	a.Describe(&r.Name, `The name of the database to be deployed.`)

	a.Describe(&r.Project, `A database can be used in a project by linking it.
	Linking the database will expose a connection URL as an environment variable for convenience.
	The same database can be linked to multiple projects.`)

	a.Describe(&r.Type, `The type of the database to be deployed.

	Supported types are:
	- postgres-neon`)
	a.SetDefault(&r.Type, "postgres-neon")

	a.Describe(&r.Region, `The region in which the database will be deployed.`)
	a.SetDefault(&r.Region, "us-east-1")
}

func (r *DatabaseState) Annotate(a infer.Annotator) {
	a.Describe(&r.URL, `The URL of the database.`)

	a.Describe(&r.DatabaseId, `The database ID.`)
}

func (*Database) Diff(ctx context.Context, id string, olds DatabaseState, news DatabaseArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if olds.Name != news.Name {
		diff["name"] = p.PropertyDiff{Kind: p.DeleteReplace}
	}

	if olds.Project == nil && news.Project != nil {
		diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
	} else if olds.Project != nil && news.Project == nil {
		diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
	} else if olds.Project != nil && news.Project != nil {
		areProjectsIdentical := utils.CompareProjects(*olds.Project, *news.Project)
		if !areProjectsIdentical {
			diff["project"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}

	}

	if olds.Type == nil {
		if news.Type != nil && *news.Type != "postgres-neon" {
			diff["type"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	} else {
		if news.Type != nil {
			if *olds.Type != *news.Type {
				diff["type"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		} else {
			if *olds.Type != "postgres-neon" {
				diff["type"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		}
	}

	if olds.Region == nil {
		if news.Region != nil && *news.Region != "us-east-1" {
			diff["region"] = p.PropertyDiff{Kind: p.DeleteReplace}
		}
	} else {
		if news.Region != nil {
			if *olds.Region != *news.Region {
				diff["region"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		} else {
			if *olds.Region != "us-east-1" {
				diff["region"] = p.PropertyDiff{Kind: p.DeleteReplace}
			}
		}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil

}

func (*Database) Read(ctx context.Context, id string, inputs DatabaseArgs, state DatabaseState) (string, DatabaseArgs, DatabaseState, error) {
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

func (*Database) Create(ctx context.Context, name string, input DatabaseArgs, preview bool) (string, DatabaseState, error) {
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

	// If a project name is provided, link the database to the project
	if input.Project != nil {
		projectDetails, err := requests.GetProjectDetails(ctx, input.Project.Name)
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
	}

	return name, state, nil
}

func (*Database) Delete(ctx context.Context, id string, state DatabaseState) error {

	err := requests.DeleteDatabase(ctx, state.DatabaseId)
	if err != nil {
		if strings.Contains(err.Error(), "database id does not exist") || strings.Contains(err.Error(), "405 Method Not Allowed") {
			log.Println("Database is already deleted")
			return nil
		}
		log.Println("Error deleting database", err.Error())
		return err
	}
	return nil
}
