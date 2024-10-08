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
export class Authentication extends pulumi.CustomResource {
    /**
     * Get an existing Authentication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Authentication {
        return new Authentication(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'genezio:index:Authentication';

    /**
     * Returns true if the given object is an instance of Authentication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Authentication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Authentication.__pulumiType;
    }

    /**
     * The type of database to be used for authentication.
     *
     * 	Supported database types are:
     * 	- postgresql
     * 	- mongodb
     */
    public readonly databaseType!: pulumi.Output<string | undefined>;
    /**
     * The URL of the database to be used for authentication.
     */
    public readonly databaseUrl!: pulumi.Output<string>;
    /**
     * The project to which the authentication will be added.
     */
    public readonly project!: pulumi.Output<outputs.domain.Project>;
    /**
     * The authentication providers to be enabled for the project.
     *
     * 	You can enable the following providers:
     * 	- email
     * 	- web3
     * 	- google
     */
    public readonly provider!: pulumi.Output<outputs.domain.AuthenticationProviders | undefined>;
    /**
     * The region in which the authentication is deployed.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * The token for the authentication. This token is used on the client side.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;

    /**
     * Create a Authentication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthenticationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.databaseUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseUrl'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["databaseType"] = (args ? args.databaseType : undefined) ?? "postgresql";
            resourceInputs["databaseUrl"] = args ? args.databaseUrl : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["provider"] = args ? args.provider : undefined;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
        } else {
            resourceInputs["databaseType"] = undefined /*out*/;
            resourceInputs["databaseUrl"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["provider"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Authentication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Authentication resource.
 */
export interface AuthenticationArgs {
    /**
     * The type of database to be used for authentication.
     *
     * 	Supported database types are:
     * 	- postgresql
     * 	- mongodb
     */
    databaseType?: pulumi.Input<string>;
    /**
     * The URL of the database to be used for authentication.
     */
    databaseUrl: pulumi.Input<string>;
    /**
     * The project to which the authentication will be added.
     */
    project: pulumi.Input<inputs.domain.ProjectArgs>;
    /**
     * The authentication providers to be enabled for the project.
     *
     * 	You can enable the following providers:
     * 	- email
     * 	- web3
     * 	- google
     */
    provider?: pulumi.Input<inputs.domain.AuthenticationProvidersArgs>;
}
