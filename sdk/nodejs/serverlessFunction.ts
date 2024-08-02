// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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

    public readonly backendPath!: pulumi.Output<string | undefined>;
    public readonly cloudProvider!: pulumi.Output<string | undefined>;
    public readonly entry!: pulumi.Output<string>;
    public /*out*/ readonly functionId!: pulumi.Output<string>;
    public readonly handler!: pulumi.Output<string>;
    public readonly language!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly pathAsset!: pulumi.Output<pulumi.asset.Archive>;
    public /*out*/ readonly projectEnvId!: pulumi.Output<string>;
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    public readonly projectName!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly stage!: pulumi.Output<string | undefined>;
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
            if ((!args || args.pathAsset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pathAsset'");
            }
            if ((!args || args.projectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectName'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["backendPath"] = args ? args.backendPath : undefined;
            resourceInputs["cloudProvider"] = args ? args.cloudProvider : undefined;
            resourceInputs["entry"] = args ? args.entry : undefined;
            resourceInputs["handler"] = args ? args.handler : undefined;
            resourceInputs["language"] = args ? args.language : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pathAsset"] = args ? args.pathAsset : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["stage"] = args ? args.stage : undefined;
            resourceInputs["functionId"] = undefined /*out*/;
            resourceInputs["projectEnvId"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["backendPath"] = undefined /*out*/;
            resourceInputs["cloudProvider"] = undefined /*out*/;
            resourceInputs["entry"] = undefined /*out*/;
            resourceInputs["functionId"] = undefined /*out*/;
            resourceInputs["handler"] = undefined /*out*/;
            resourceInputs["language"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pathAsset"] = undefined /*out*/;
            resourceInputs["projectEnvId"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["projectName"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["stage"] = undefined /*out*/;
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
    backendPath?: pulumi.Input<string>;
    cloudProvider?: pulumi.Input<string>;
    entry: pulumi.Input<string>;
    handler: pulumi.Input<string>;
    language?: pulumi.Input<string>;
    name: pulumi.Input<string>;
    pathAsset: pulumi.Input<pulumi.asset.Archive>;
    projectName: pulumi.Input<string>;
    region: pulumi.Input<string>;
    stage?: pulumi.Input<string>;
}
