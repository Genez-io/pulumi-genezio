import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as crypto from "crypto";
import * as fs from "fs";
import * as path from "path";
import * as dotenv from "dotenv";

dotenv.config();

const myDatabase = new genezio.Database("MyDatabase", {
  name: "my-database-fullstack-pulumi-6",
  type: "postgres-neon",
});

const MyProject = new genezio.Project("MyProject", {
  name: "my-fullstack-pulumi",
  region: "us-east-1",
  environmentVariables: [
    {
      name: "DATABASE_URL",
      value: myDatabase.url,
    },
  ],
});

const frontendPublishPath = path.join(__dirname, "client", "dist");

const myFrontend = new genezio.Frontend("MyFrontend", {
  project: {
    name: MyProject.name,
    region: MyProject.region,
  },
  path: "./client",
  publish: new pulumi.asset.FileArchive(frontendPublishPath),
  subdomain: "my-frontend-pulumi-10",
});

// export const databaseOutput = {
//   id: myDatabase.databaseId,
//   name: myDatabase.name,
//   region: myDatabase.region,
//   type: myDatabase.type,
//   endpoint: myDatabase.url,
// };

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
