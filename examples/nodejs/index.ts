import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import path = require("path");

const myDatabase = new genezio.Database("MyDatabase", {
  name: "my-database-fullstack-pulumi-6",
});

const MyProject = new genezio.Project("MyProject", {
  name: "test",
  region: "us-east-1",
  environmentVariables: [
    {
      name: "CUSTOM_DATABASE_URL",
      value: myDatabase.url,
    },
  ],
});

const functionPath = path.join(__dirname, "function");

const myFunction = new genezio.ServerlessFunction("MyFunction", {
  path: new pulumi.asset.FileArchive(functionPath),
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
  backendPath: ".",
});

const frontendPublishPath = path.join(__dirname, "client", "dist");

const myFrontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  path: "./client",
  publish: new pulumi.asset.FileArchive(frontendPublishPath),
  subdomain: "my-frontend-pulumi",
  buildCommands: ["npm install", "npm run build"],
  environmentVariables: [
    {
      name: "VITE_HELLO_WORLD_FUNCTION_URL",
      value: myFunction.url,
    },
  ],
});

const myAuth = new genezio.Authentication("MyAuth", {
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  databaseType: "postgresql",
  databaseUrl: myDatabase.url,
  provider: {
    email: true,
  },
});
