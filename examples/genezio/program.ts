import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";

// This example deploys a fullstack project with a frontend, serverless functions, and a database

const MyProject = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
  environmentVariables: [
    {
      name: "TEST_ENV_VAR",
      value: "test",
    },
  ],
});

const myDatabase = new genezio.Database("MyDatabase", {
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  name: "my-database",
});

const helloWorldFunction = new genezio.ServerlessFunction("MyFunction", {
  path: new pulumi.asset.FileArchive("./server"),
  backendPath: ".",
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  entry: "hello.mjs",
  handler: "handler",
  name: "hello-world",
});

const goodbyeFunction = new genezio.ServerlessFunction("Goodbye", {
  path: new pulumi.asset.FileArchive("./server"),
  backendPath: ".",
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  entry: "goodbye.mjs",
  handler: "handler",
  name: "goodbye",
});

const myFrontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  path: "./client",
  publish: new pulumi.asset.FileArchive("./client/dist"),
  subdomain: "amazing-purple-capybara",
  // Needs to be implemented
  // environment: {
  //   VITE_HELLO_WORLD_FUNCTION_URL: helloWorldFunction.url,
  // },
});
