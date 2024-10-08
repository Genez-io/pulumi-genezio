// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Genezio
{
    /// <summary>
    /// A project resource that will be deployed on the Genezio platform.The project resource is used to group resources together and manage them as a single unit.
    /// 
    /// The project resource will deploy an empty project on the Genezio platform.
    /// 
    /// It is recommended to create a Project Resource as the first step in your deployment workflow. The output from this resource can then be utilized to provision and configure other resources within the project, ensuring they are properly associated and managed under a unified project.
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic Usage
    /// 
    /// ### Environment Variables
    /// 
    /// ## Pulumi Output Reference
    /// 
    /// Once the project is created, the `projectId` and `projectUrl` are available as outputs.
    /// </summary>
    [GenezioResourceType("genezio:index:Database")]
    public partial class Database : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The database ID.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// The name of the database to be deployed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A database can be used in a project by linking it.
        /// 	Linking the database will expose a connection URL as an environment variable for convenience.
        /// 	The same database can be linked to multiple projects.
        /// </summary>
        [Output("project")]
        public Output<Pulumi.Genezio.Domain.Outputs.Project?> Project { get; private set; } = null!;

        /// <summary>
        /// The region in which the database will be deployed.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// The type of the database to be deployed.
        /// 
        /// 	Supported types are:
        /// 	- postgres-neon
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The URL of the database.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("genezio:index:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("genezio:index:Database", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "url",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Database(name, id, options);
        }
    }

    public sealed class DatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the database to be deployed.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// A database can be used in a project by linking it.
        /// 	Linking the database will expose a connection URL as an environment variable for convenience.
        /// 	The same database can be linked to multiple projects.
        /// </summary>
        [Input("project")]
        public Input<Pulumi.Genezio.Domain.Inputs.ProjectArgs>? Project { get; set; }

        /// <summary>
        /// The region in which the database will be deployed.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The type of the database to be deployed.
        /// 
        /// 	Supported types are:
        /// 	- postgres-neon
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DatabaseArgs()
        {
            Region = "us-east-1";
            Type = "postgres-neon";
        }
        public static new DatabaseArgs Empty => new DatabaseArgs();
    }
}
