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
	"github.com/Genez-io/pulumi-genezio/provider/domain"
	r "github.com/Genez-io/pulumi-genezio/provider/resources"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	p "github.com/pulumi/pulumi-go-provider"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "genezio"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[*r.ServerlessFunction,r.ServerlessFunctionArgs, r.ServerlessFunctionState](),
			infer.Resource[*r.Database,r.DatabaseArgs,r.DatabaseState](),
			infer.Resource[*r.Project,r.ProjectArgs,r.ProjectState](),
			infer.Resource[*r.Frontend,r.FrontendArgs,r.FrontendState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"resources": "index",
		},
		Config: infer.Config[*domain.Config](),
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










