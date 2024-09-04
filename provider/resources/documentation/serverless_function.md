The Serverless Function Resource allows you to deploy and manage serverless functions on the Genezio platform. Serverless functions are ideal for executing backend logic without the need for managing servers, enabling scalable and cost-effective event-driven architectures.

When creating a Serverless Function Resource, you define the function’s code path, entry file, and handler function, which specifies the core logic to be executed. Additionally, you can configure attributes like the function’s name and associate it with a specific project and region. This allows you to easily integrate serverless functions into your project for tasks such as API endpoints, background processing, or event handling.

## Example Usage

```typescript
import * as genezio from "@pulumi/genezio";

const functionPath = path.join(__dirname, "function");

const serverlessFunction = new genezio.ServerlessFunction("MyFunction", {
  path: new pulumi.asset.FileArchive(functionPath),
  project: {
    name: project.name,
    region: project.region,
  },
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
});
```

## Pulumi Output Reference

Once the serverless function is created, the `functionUrl` and `functionId` are available as outputs.

```typescript
import * as genezio from "@pulumi/genezio";

const functionPath = path.join(__dirname, "function");

const serverlessFunction = new genezio.ServerlessFunction("MyFunction", {
  path: new pulumi.asset.FileArchive(functionPath),
  project: {
    name: project.name,
    region: project.region,
  },
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
});

export const functionUrl = serverlessFunction.functionUrl;
```
