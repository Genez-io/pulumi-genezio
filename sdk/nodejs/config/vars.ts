// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("genezio");

export declare const authToken: string | undefined;
Object.defineProperty(exports, "authToken", {
    get() {
        return __config.get("authToken");
    },
    enumerable: true,
});

