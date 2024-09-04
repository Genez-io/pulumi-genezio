A project resource that will be deployed on the Genezio platform.The project resource is used to group resources together and manage them as a single unit.

The project resource will deploy an empty project on the Genezio platform.

It is recommended to create a Project Resource as the first step in your deployment workflow. The output from this resource can then be utilized to provision and configure other resources within the project, ensuring they are properly associated and managed under a unified project.

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

## Pulumi Output Reference

Once the project is created, the `projectId` and `projectUrl` are available as outputs.

```typescript

const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

export const projectId = project.projectId;
export const projectUrl = project.projectUrl;
```
