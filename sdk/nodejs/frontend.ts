// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Frontend extends pulumi.CustomResource {
    /**
     * Get an existing Frontend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Frontend {
        return new Frontend(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'genezio:index:Frontend';

    /**
     * Returns true if the given object is an instance of Frontend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Frontend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Frontend.__pulumiType;
    }

    public readonly buildCommand!: pulumi.Output<string>;
    public readonly environmentVariables!: pulumi.Output<outputs.domain.EnvironmentVariable[] | undefined>;
    public readonly path!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<outputs.domain.Project>;
    public readonly publish!: pulumi.Output<pulumi.asset.Archive>;
    public readonly subdomain!: pulumi.Output<string | undefined>;
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Frontend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FrontendArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.buildCommand === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buildCommand'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.publish === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publish'");
            }
            resourceInputs["buildCommand"] = args ? args.buildCommand : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["publish"] = args ? args.publish : undefined;
            resourceInputs["subdomain"] = args ? args.subdomain : undefined;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["buildCommand"] = undefined /*out*/;
            resourceInputs["environmentVariables"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["publish"] = undefined /*out*/;
            resourceInputs["subdomain"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Frontend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Frontend resource.
 */
export interface FrontendArgs {
    buildCommand: pulumi.Input<string>;
    environmentVariables?: pulumi.Input<pulumi.Input<inputs.domain.EnvironmentVariableArgs>[]>;
    path: pulumi.Input<string>;
    project: pulumi.Input<inputs.domain.ProjectArgs>;
    publish: pulumi.Input<pulumi.asset.Archive>;
    subdomain?: pulumi.Input<string>;
}
