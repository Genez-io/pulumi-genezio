A project resource that will be deployed on the Genezio platform.

The project resource is used to group resources together and manage them as a single unit. This resource should be used as the base resource for all other resources. The project resource will deploy an empty project on the Genezio platform. The project can then be used to deploy other resources. It is good practce to create a project resource first and then use the output of the project resource to deploy other resources.

## Example Usage

### Basic Usage

```typescript
import * as genezio from "@pulumi/genezio";

const project = new genezio.Project("project", {
  name: "my-project",
  region: "us-east-1",
});
```

### Environment Variables

```typescript
import * as genezio from "@pulumi/genezio";

const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
  environmentVariables: [
    {
      name: "MY_ENV_VAR",
      value: "my-value",
    },
  ],
});
```
