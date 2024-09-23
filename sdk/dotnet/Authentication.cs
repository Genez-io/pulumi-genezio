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
    [GenezioResourceType("genezio:index:Authentication")]
    public partial class Authentication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of database to be used for authentication.
        /// 
        /// 	Supported database types are:
        /// 	- postgresql
        /// 	- mongodb
        /// </summary>
        [Output("databaseType")]
        public Output<string?> DatabaseType { get; private set; } = null!;

        /// <summary>
        /// The URL of the database to be used for authentication.
        /// </summary>
        [Output("databaseUrl")]
        public Output<string> DatabaseUrl { get; private set; } = null!;

        /// <summary>
        /// The project to which the authentication will be added.
        /// </summary>
        [Output("project")]
        public Output<Pulumi.Genezio.Domain.Outputs.Project> Project { get; private set; } = null!;

        /// <summary>
        /// The authentication providers to be enabled for the project.
        /// 
        /// 	You can enable the following providers:
        /// 	- email
        /// 	- web3
        /// 	- google
        /// </summary>
        [Output("provider")]
        public Output<Pulumi.Genezio.Domain.Outputs.AuthenticationProviders?> Provider { get; private set; } = null!;

        /// <summary>
        /// The region in which the authentication is deployed.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The token for the authentication. This token is used on the client side.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a Authentication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Authentication(string name, AuthenticationArgs args, CustomResourceOptions? options = null)
            : base("genezio:index:Authentication", name, args ?? new AuthenticationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Authentication(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("genezio:index:Authentication", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Authentication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Authentication Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Authentication(name, id, options);
        }
    }

    public sealed class AuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of database to be used for authentication.
        /// 
        /// 	Supported database types are:
        /// 	- postgresql
        /// 	- mongodb
        /// </summary>
        [Input("databaseType")]
        public Input<string>? DatabaseType { get; set; }

        /// <summary>
        /// The URL of the database to be used for authentication.
        /// </summary>
        [Input("databaseUrl", required: true)]
        public Input<string> DatabaseUrl { get; set; } = null!;

        /// <summary>
        /// The project to which the authentication will be added.
        /// </summary>
        [Input("project", required: true)]
        public Input<Pulumi.Genezio.Domain.Inputs.ProjectArgs> Project { get; set; } = null!;

        /// <summary>
        /// The authentication providers to be enabled for the project.
        /// 
        /// 	You can enable the following providers:
        /// 	- email
        /// 	- web3
        /// 	- google
        /// </summary>
        [Input("provider")]
        public Input<Pulumi.Genezio.Domain.Inputs.AuthenticationProvidersArgs>? Provider { get; set; }

        public AuthenticationArgs()
        {
            DatabaseType = "postgresql";
        }
        public static new AuthenticationArgs Empty => new AuthenticationArgs();
    }
}
