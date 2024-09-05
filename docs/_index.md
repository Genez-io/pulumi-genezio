---
title: Genezio Provider
meta_desc: Provides an overview on how to configure the Genezio provider for Pulumi.
layout: package
---

The Genezio provider for Pulumi can provision various resources in [Genezio](https://genezio.com).
This document provides an overview of the Genezio provider and how to configure it for Pulumi.

The Genezio provider must be configured with credentials to deploy and update resources.
Read more in [Installation & Configuration](https://www.pulumi.com/registry/packages/genezio/installation-configuration/) for instructions.

## Example

Check out the code snippets below to see how to create a Genezio project featuring with a Genezio Function connected to a frontend client.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as path from "path";

const project = new genezio.Project("MyProject", {
  name: "my-fullstack-project",
  region: "us-east-1",
});

const database = new genezio.Database("MyDatabase", {
  project: {
    name: project.name,
    region: project.region,
  },
  name: "core-database",
});

const serverPath = path.join(__dirname, "server");

const helloWorldFunction = new genezio.ServerlessFunction("MyFunction", {
  path: new pulumi.asset.FileArchive(serverPath),
  project: {
    name: project.name,
    region: project.region,
  },
  entry: "hello.mjs",
  handler: "handler",
  name: "hello-world",
});

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

{{% /choosable %}}
{{< /chooser >}}
