The Frontend Resource allows you to deploy frontend components to the Genezio platform.

When creating a Frontend Resource, you can specify the path to your frontend source files, configure build commands, and define environment variables. This resource also allows for flexible deployment options, such as specifying the subdomain where the frontend will be published, and using the output from other resources like serverless functions to dynamically set environment variables.

## Example Usage

### Basic Usage
```typescript
import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as path from "path";

const frontendPath = path.join(__dirname, "client");

const frontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: project.name,
    region: project.region,
  },
  path: new pulumi.asset.FileArchive(frontendPath),
  publish: "dist",
  subdomain: "my-frontend",
  buildCommands: ["npm install", "npm run build"],
});
```

### Inject Environment Variables

```typescript
import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as path from "path";

const frontendPath = path.join(__dirname, "client");

const frontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: project.name,
    region: project.region,
  },
  path: new pulumi.asset.FileArchive(frontendPath),
  publish: "dist",
  subdomain: "my-frontend",
  buildCommands: ["npm install", "npm run build"],
  environment: [
    {
      name: "VITE_HELLO_WORLD_FUNCTION_URL",
      value: helloWorldFunction.url,
    },
  ],
});
```

## Pulumi Output Reference

Once the frontend is created, the `frontendUrl` and `frontendId` are available as outputs.

```typescript
const frontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: project.name,
    region: project.region,
  },
  path: new pulumi.asset.FileArchive(frontendPath),
  publish: "dist",
  subdomain: "my-frontend",
  buildCommands: ["npm install", "npm run build"],
});

export const frontendUrl = frontend.url;
export const frontendId = frontend.frontendId;
```
