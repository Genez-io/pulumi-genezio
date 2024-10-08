import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as path from "path";

const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
  environment: [
    {
      name: "ENV_VAR",
      value: "my-env-var",
    },
  ],
});

const database = new genezio.Database("MyDatabase", {
  project: {
    name: project.name,
    region: project.region,
  },
  name: "my-database",
  region: "us-east-1",
});

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
      value: serverlessFunction.url,
    },
  ],
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: database.url,
  provider: {
    email: true,
    web3: true,
    google: {
      id: "my-client-id",
      secret: "my-client",
    }
  },
});
