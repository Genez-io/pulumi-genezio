// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { AuthenticationArgs } from "./authentication";
export type Authentication = import("./authentication").Authentication;
export const Authentication: typeof import("./authentication").Authentication = null as any;
utilities.lazyLoad(exports, ["Authentication"], () => require("./authentication"));

export { DatabaseArgs } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { FrontendArgs } from "./frontend";
export type Frontend = import("./frontend").Frontend;
export const Frontend: typeof import("./frontend").Frontend = null as any;
utilities.lazyLoad(exports, ["Frontend"], () => require("./frontend"));

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ServerlessFunctionArgs } from "./serverlessFunction";
export type ServerlessFunction = import("./serverlessFunction").ServerlessFunction;
export const ServerlessFunction: typeof import("./serverlessFunction").ServerlessFunction = null as any;
utilities.lazyLoad(exports, ["ServerlessFunction"], () => require("./serverlessFunction"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "genezio:index:Authentication":
                return new Authentication(name, <any>undefined, { urn })
            case "genezio:index:Database":
                return new Database(name, <any>undefined, { urn })
            case "genezio:index:Frontend":
                return new Frontend(name, <any>undefined, { urn })
            case "genezio:index:Project":
                return new Project(name, <any>undefined, { urn })
            case "genezio:index:ServerlessFunction":
                return new ServerlessFunction(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("genezio", "index", _module)
pulumi.runtime.registerResourcePackage("genezio", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:genezio") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
