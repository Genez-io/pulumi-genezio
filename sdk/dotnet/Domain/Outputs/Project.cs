// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Genezio.Domain.Outputs
{

    [OutputType]
    public sealed class Project
    {
        public readonly string Name;
        public readonly string Region;

        [OutputConstructor]
        private Project(
            string name,

            string region)
        {
            Name = name;
            Region = region;
        }
    }
}
