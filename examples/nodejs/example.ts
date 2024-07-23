import * as genezio from "@pulumi/genezio";

import * as archive from "@pulumi/archive";
import * as pulumi from "@pulumi/pulumi";

const code = archive.getFile({
  type: "zip",
  sourceDir: "./function",
  outputPath: "./genezioDeploy.zip",
})

const myFunction = new genezio.ServerlessFunction("MyFunction", {
  code: new pulumi.asset.FileArchive("./genezioDeploy.zip"),
  name: "my-function",
  path: ".",
  entry: "app.mjs",
  handler: "handler",
});

const myClass = new genezio.Class("MyClass", {
  path: "./backend.ts",
}

const myBackend = new genezio.Backend("MyBackend", {
  path: "./server",
  language: {
    name: "ts",
    packageManager: "npm",
  },
  environmentVariables: [
    {
      name: "TEST_VAR",
      value: "variable"
    },
  ],
  functions: [myFunction],
  classes: [myClass],
})

const myDatabase = new genezio.Database("MyDatabase", {
  name: "my-database",
  region: "us-east-1",
});

const myFrontend = new genezio.Frontend("MyFrontend", {
  path: "./client",
  sdk: {
    langugae: "ts",
  },
  publish: "./dist",
  subdomain: "pulumi-test",
  functions: [myFunction.url],
})

const myStage = genezio.Stage("MyStage", {
  stageName: "dev",
  database: [myDatabase],
  backend: [myBackend],
  frontend: [myFrontend],
})

const myProject = genezio.Project("MyProject", {
  projectName: "genezio-project-pulumi",
  region: "us-east-1",
  stages: [myStage],
}
