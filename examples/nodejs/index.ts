import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";

const myRandomResource = new genezio.Random("myRandomResource", { length: 26 });
export const output = {
  value: myRandomResource.result,
};
const myFunction = new genezio.ServerlessFunction("MyFunction", {
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
