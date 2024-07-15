import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as crypto from "crypto";
import * as fs from "fs";
import * as path from "path";

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

const myDatabase = new genezio.Database("MyDatabase", {
  name: "my-database-3",
  type: "postgres-neon",
  region: "aws-us-east-1",
  authToken:
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZpcmdpbEBnZW5lei5pbyIsImV4cCI6MTcxODAyODY5NSwiaWF0IjoxNzE4MDIzMjk1LCJuYmYiOjE3MTgwMjI5OTV9.PzCMA7EGARhnyIgRozhGAaI2G0ITQlnmHKXAosQRvko",
});

export const databaseOutput = {
  id: myDatabase.databaseId,
};

const myFunction = new genezio.ServerlessFunction("MyFunction", {
  folderHash: sha256FromFolder("./function"),
  path: "./function",
  projectName: "project-function-pulumi-2",
  region: "us-east-1",
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
  authToken:
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZpcmdpbEBnZW5lei5pbyIsImV4cCI6MTcxODAyODY5NSwiaWF0IjoxNzE4MDIzMjk1LCJuYmYiOjE3MTgwMjI5OTV9.PzCMA7EGARhnyIgRozhGAaI2G0ITQlnmHKXAosQRvko",
  environmentVariables: [
    {
      name: "POSTGRES_URL",
      value: myDatabase.url,
    },
  ],
});

export const functionOutput = {
  url: myFunction.url,
  id: myFunction.functionId,
  envId: myFunction.projectEnvId,
  projectId: myFunction.projectId,
};
