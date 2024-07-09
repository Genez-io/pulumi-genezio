import * as genezio from "@pulumi/genezio";
import * as pulumi from "@pulumi/pulumi";

const myRandomResource = new genezio.Random("myRandomResource", { length: 26 });
export const output = {
  value: myRandomResource.result,
};
const myFunction = new genezio.Function("MyFunction", {
  path: "function/path",
  projectName: "projectName",
  region: "region",
  entry: "entry",
  handler: "handler",
  name: "name",
});

export const functionOutput = {
  url: myFunction.url,
  id: myFunction.functionId,
};
