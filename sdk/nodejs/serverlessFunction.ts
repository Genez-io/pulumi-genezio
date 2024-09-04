// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * A project resource that will be deployed on the Genezio platform.The project resource is used to group resources together and manage them as a single unit.
 *
 * The project resource will deploy an empty project on the Genezio platform.
 *
 * It is recommended to create a Project Resource as the first step in your deployment workflow. The output from this resource can then be utilized to provision and configure other resources within the project, ensuring they are properly associated and managed under a unified project.
 *
 * ## Example Usage
 *
 * ### Basic Usage
 *
 * ```typescript
 * import * as genezio from "@pulumi/genezio";
 *
 * const project = new genezio.Project("project", {
 *   name: "my-project",
 *   region: "us-east-1",
 * });
 * ```
 *
 * ### Environment Variables
 *
 * ```typescript
 * import * as genezio from "@pulumi/genezio";
 *
 * const project = new genezio.Project("MyProject", {
 *   name: "my-project",
 *   region: "us-east-1",
 *   environmentVariables: [
 *     {
 *       name: "MY_ENV_VAR",
 *       value: "my-value",
 *     },
 *   ],
 * });
 * ```
 *
 * ## Pulumi Output Reference
 *
 * Once the project is created, the `projectId` and `projectUrl` are available as outputs.
 *
 * ```typescript
 *
 * const project = new genezio.Project("MyProject", {
 *   name: "my-project",
 *   region: "us-east-1",
 * });
 *
 * export const projectId = project.projectId;
 * export const projectUrl = project.projectUrl;
 * ```
 */
export class ServerlessFunction extends pulumi.CustomResource {
    /**
     * Get an existing ServerlessFunction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerlessFunction {
        return new ServerlessFunction(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'genezio:index:ServerlessFunction';

    /**
     * Returns true if the given object is an instance of ServerlessFunction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerlessFunction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerlessFunction.__pulumiType;
    }

    /**
     * The path to the backend folder where the function is located.
     */
    public readonly backendPath!: pulumi.Output<string | undefined>;
    /**
     * The entry file of the function.
     */
    public readonly entry!: pulumi.Output<string>;
    /**
     * The function ID.
     */
    public /*out*/ readonly functionId!: pulumi.Output<string>;
    /**
     * The handler of the function.
     */
    public readonly handler!: pulumi.Output<string>;
    /**
     * The language in which the function is written.
     */
    public readonly language!: pulumi.Output<string | undefined>;
    /**
     * The name of the function to be deployed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The path to the function code.
     */
    public readonly path!: pulumi.Output<pulumi.asset.Archive>;
    /**
     * The project to which the function will be deployed.
     */
    public readonly project!: pulumi.Output<outputs.domain.Project>;
    /**
     * The URL of the function.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a ServerlessFunction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerlessFunctionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.entry === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entry'");
            }
            if ((!args || args.handler === undefined) && !opts.urn) {
                throw new Error("Missing required property 'handler'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["backendPath"] = args ? args.backendPath : undefined;
            resourceInputs["entry"] = args ? args.entry : undefined;
            resourceInputs["handler"] = args ? args.handler : undefined;
            resourceInputs["language"] = (args ? args.language : undefined) ?? "js";
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["functionId"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["backendPath"] = undefined /*out*/;
            resourceInputs["entry"] = undefined /*out*/;
            resourceInputs["functionId"] = undefined /*out*/;
            resourceInputs["handler"] = undefined /*out*/;
            resourceInputs["language"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerlessFunction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerlessFunction resource.
 */
export interface ServerlessFunctionArgs {
    /**
     * The path to the backend folder where the function is located.
     */
    backendPath?: pulumi.Input<string>;
    /**
     * The entry file of the function.
     */
    entry: pulumi.Input<string>;
    /**
     * The handler of the function.
     */
    handler: pulumi.Input<string>;
    /**
     * The language in which the function is written.
     */
    language?: pulumi.Input<string>;
    /**
     * The name of the function to be deployed.
     */
    name: pulumi.Input<string>;
    /**
     * The path to the function code.
     */
    path: pulumi.Input<pulumi.asset.Archive>;
    /**
     * The project to which the function will be deployed.
     */
    project: pulumi.Input<inputs.domain.ProjectArgs>;
}
