import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";

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

const myFunction = new genezio.ServerlessFunction("MyFunction", {
  path: new pulumi.asset.FileArchive("./function"),
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
  backendPath: ".",
});

const myFrontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  path: "./client",
  publish: new pulumi.asset.FileArchive("./client/dist"),
  subdomain: "my-frontend-pulumi-10",
});
