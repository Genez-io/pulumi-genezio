import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as crypto from "crypto";
import * as fs from "fs";
import * as path from "path";
import * as dotenv from "dotenv";

dotenv.config();

function sha256FromFolder(folderPath: string): string {
  const hash = crypto.createHash("sha256");
  const files = fs.readdirSync(folderPath);
  for (const file of files) {
    const filePath = path.join(folderPath, file);
    const stat = fs.statSync(filePath);
    if (stat.isDirectory()) {
      const dirHash = sha256FromFolder(filePath);
      hash.update(dirHash);
    } else {
      const content = fs.readFileSync(filePath);
      hash.update(content);
    }
  }

  return hash.digest("hex");
}

const MyProject = new genezio.Project("MyProject", {
  name: "my-fullstack-pulumi",
  region: "us-east-1",
  cloudProvider: "genezio-cloud",
  stage: "prod",
});

const myFrontend = new genezio.Frontend("MyFrontend", {
  projectName: "my-fullstack-pulumi",
  region: "us-east-1",
  path: "./client",
  publish: "./dist",
  subdomain: "my-frontend-pulumi-10",
});

const myDatabase = new genezio.Database("MyDatabase", {
  name: "my-database-fullstack-pulumi-4",
  type: "postgres-neon",
  region: "aws-us-east-1",
});

// export const databaseOutput = {
//   id: myDatabase.databaseId,
//   name: myDatabase.name,
//   region: myDatabase.region,
//   type: myDatabase.type,
//   endpoint: myDatabase.url,
// };

// export enum DatabaseType {
//   POSTGRES = "postgres-neon",
//   MYSQL = "mysql-neon",
// }

const myFunction = new genezio.ServerlessFunction("MyFunction", {
  folderHash: sha256FromFolder("./function"),
  path: "./function",
  projectName: "my-fullstack-pulumi",
  region: "us-east-1",
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
  environmentVariables: {
    POSTGRES_URL: myDatabase.url,
  },
});
