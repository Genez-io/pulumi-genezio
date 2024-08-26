import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import { existsSync, mkdirSync } from "fs";
import path = require("path");

// This example deploys a fullstack project with a frontend, serverless functions, and a database

const project = new genezio.Project("MyProject", {
  name: "my-fullstack-project",
  region: "us-east-1",
  environment: [
    {
      name: "MY_ENV_VAR",
      value: "my-env-var",
    },
  ],
});

const database = new genezio.Database("MyDatabase", {
  project: {
    name: project.name,
    region: project.region,
  },
  name: "core",
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

const goodbyeFunction = new genezio.ServerlessFunction("Goodbye", {
  path: new pulumi.asset.FileArchive(serverPath),
  project: {
    name: project.name,
    region: project.region,
  },
  entry: "goodbye.mjs",
  handler: "handler",
  name: "goodbye",
});

const addUserFunction = new genezio.ServerlessFunction("AddUser", {
  path: new pulumi.asset.FileArchive(serverPath),
  project: {
    name: project.name,
    region: project.region,
  },
  entry: "addUser.mjs",
  handler: "handler",
  name: "add-user",
});

const frontendPublishPath = path.join(__dirname, "client", "dist");

// if the publish directory does not exist, create it
if (!existsSync(frontendPublishPath)) {
  mkdirSync(frontendPublishPath, { recursive: true });
}

const myFrontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: project.name,
    region: project.region,
  },
  path: "./client",
  publish: new pulumi.asset.FileArchive(frontendPublishPath),
  buildCommands: ["npm install", "npm run build"],
  environment: [
    {
      name: "VITE_HELLO_WORLD_FUNCTION_URL",
      value: helloWorldFunction.url,
    },
    {
      name: "VITE_GOODBYE_FUNCTION_URL",
      value: goodbyeFunction.url,
    },
    {
      name: "VITE_ADD_USER_FUNCTION_URL",
      value: addUserFunction.url,
    },
  ],
});
