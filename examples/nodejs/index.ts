import * as pulumi from "@pulumi/pulumi";
import * as genezio from "@pulumi/genezio";

const myRandomResource = new genezio.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
