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

// const myDatabase = new genezio.Database("MyDatabase", {
//   name: "my-database-4",
//   type: "postgres-neon",
//   region: "aws-us-east-1",
//   authToken: process.env.AUTH_TOKEN ?? "",
// });

// export const databaseOutput = {
//   id: myDatabase.databaseId,
// };

// console.log("the auth token is", process.env.AUTH_TOKEN);

const myFunction = new genezio.ServerlessFunction("MyFunction", {
  folderHash: sha256FromFolder("./function"),
  path: "./function",
  projectName: "project-function-pulumi-2",
  region: "us-east-1",
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
  authToken: process.env.AUTH_TOKEN ?? "",
  environmentVariables: [
    {
      name: "POSTGRES_URL",
      value: "sadas",
    },
  ],
});

export const functionOutput = {
  url: myFunction.url,
  id: myFunction.functionId,
  envId: myFunction.projectEnvId,
  projectId: myFunction.projectId,
};
