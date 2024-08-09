import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import path = require("path");

const project = new genezio.Project("MyProject", {
  name: "my-project-fullstack-pulumi",
  region: "us-east-1",
  environmentVariables: [
    {
      name: "CUSTOM_ENV_VAR",
      value: "my-env-var",
    },
  ],
});

const database = new genezio.Database("MyDatabase", {
  project: {
    name: project.name,
    region: project.region,
  },
  name: "my-database-fullstack-pulumi",
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

const frontendPublishPath = path.join(__dirname, "client", "dist");

const frontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: project.name,
    region: project.region,
  },
  path: "./client",
  publish: new pulumi.asset.FileArchive(frontendPublishPath),
  subdomain: "my-frontend-pulumi",
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: database.url,
  provider: {
    email: true,
  },
});
