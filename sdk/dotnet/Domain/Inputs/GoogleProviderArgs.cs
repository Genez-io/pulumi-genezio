// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Genezio.Domain.Inputs
{

    public sealed class GoogleProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        public GoogleProviderArgs()
        {
        }
        public static new GoogleProviderArgs Empty => new GoogleProviderArgs();
    }
}
