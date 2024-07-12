// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
	"github.com/Genez-io/pulumi-genezio/provider/utils"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "genezio"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[Random, RandomArgs, RandomState](),
			infer.Resource[ServerlessFunction, ServerlessFunctionArgs, ServerlessFunctionState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
	})
}

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Random struct{}

// Each resource has an input struct, defining what arguments it accepts.
type RandomArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but it's generally a
	// good idea.
	Length int `pulumi:"length"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type RandomState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	RandomArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minimum.
func (Random) Create(ctx p.Context, name string, input RandomArgs, preview bool) (string, RandomState, error) {
	state := RandomState{RandomArgs: input}
	if preview {
		return name, state, nil
	}
	state.Result = makeRandom(input.Length)
	return name, state, nil
}

func makeRandom(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") // SED_SKIP

	result := make([]rune, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return "1234"
}

type ServerlessFunction struct{}

type ServerlessFunctionArgs struct {
	Path string `pulumi:"path"` 
	ProjectName  string `pulumi:"projectName"`
	Name string `pulumi:"name"`
	Region	   string `pulumi:"region"`
	Entry string `pulumi:"entry"`
	Handler string `pulumi:"handler"`
	AuthToken string `pulumi:"authToken"`
}

type ServerlessFunctionState struct {
	ServerlessFunctionArgs

	ID string `pulumi:"functionId"`
	URL string `pulumi:"url"`
}

func (ServerlessFunction) Create(ctx p.Context, name string, input ServerlessFunctionArgs, preview bool) (string, ServerlessFunctionState, error) {
	state := ServerlessFunctionState{ServerlessFunctionArgs: input}
	if preview {
		return name, state, nil
	}


	backendPath := "."

	projectConfiguration := domain.ProjectConfiguration{
		Name: input.ProjectName,
		Region: input.Region,
		Options: domain.Options{
			NodeRuntime: "nodejs20.x",
			Architecture: "arm64",
		},
		CloudProvider: "genezio-cloud",
		Workspace: domain.Workspace{
			Backend: backendPath,
		},
		AstSummary: domain.AstSummary{
			Version: "2",
			Classes: []string{},
		},
		Classes: []string{},
		Functions: []domain.FunctionConfiguration{
			{
				Name: input.Name,
				Path: input.Path,
				Language: "ts",
				Handler: input.Handler,
				Entry: input.Entry,
				Type: "aws",
			},
		},
	}

	cloudInput, err := utils.FunctionToCloudInput(projectConfiguration.Functions[0], backendPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to convert the function to cloud input %v", err)
		return "", ServerlessFunctionState{}, err
	}
	cloudInputs := []domain.GenezioCloudInput{cloudInput}

	cloudAdapter := utils.NewGenezioCloudAdapter()

	response, err := cloudAdapter.Deploy(cloudInputs, projectConfiguration, utils.CloudAdapterOptions{Stage: nil}, nil, input.AuthToken)
	if err != nil {
		fmt.Printf("An error occurred while trying to deploy the function %v", err)
		return "", ServerlessFunctionState{}, err
	}


	state.ID = response.Functions[0].ID	
	state.URL = response.Functions[0].CloudUrl
	return name, state, nil
}







