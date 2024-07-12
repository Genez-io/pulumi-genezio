import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";
import * as crypto from "crypto";
import * as fs from "fs";
import * as path from "path";

const myRandomResource = new genezio.Random("myRandomResource", { length: 26 });
export const output = {
  value: myRandomResource.result,
};

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

const myFunction = new genezio.ServerlessFunction("MyFunction", {
  folderHash: sha256FromFolder("./function"),
  path: "./function",
  projectName: "project-function-pulumi",
  region: "us-east-1",
  entry: "app.mjs",
  handler: "handler",
  name: "my-function",
  authToken: "",
});

export const functionOutput = {
  url: myFunction.url,
  id: myFunction.functionId,
};
