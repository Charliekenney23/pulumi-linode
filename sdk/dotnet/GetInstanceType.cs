// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about a Linode instance type
        /// 
        /// ## Attributes
        /// 
        /// The Linode Instance Type resource exports the following attributes:
        /// 
        /// * `id` - The ID representing the Linode Type
        /// 
        /// * `label` - The Linode Type's label is for display purposes only
        /// 
        /// * `class` - The class of the Linode Type
        /// 
        /// * `disk` - The Disk size, in MB, of the Linode Type
        /// 
        /// * `price.0.hourly` -  Cost (in US dollars) per hour.
        /// 
        /// * `price.0.monthly` - Cost (in US dollars) per month.
        /// 
        /// * `addons.0.backups.0.price.0.hourly` - The cost (in US dollars) per hour to add Backups service.
        /// 
        /// * `addons.0.backups.0.price.0.monthly` - The cost (in US dollars) per month to add Backups service.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/instance_type.html.markdown.
        /// </summary>
        public static Task<GetInstanceTypeResult> GetInstanceType(GetInstanceTypeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTypeResult>("linode:index/getInstanceType:getInstanceType", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstanceTypeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Label used to identify instance type
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("label")]
        public string? Label { get; set; }

        public GetInstanceTypeArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstanceTypeResult
    {
        public readonly Outputs.GetInstanceTypeAddonsResult Addons;
        public readonly string Class;
        public readonly int Disk;
        public readonly string Id;
        public readonly string Label;
        public readonly int Memory;
        public readonly int NetworkOut;
        public readonly Outputs.GetInstanceTypePriceResult Price;
        public readonly int Transfer;
        public readonly int Vcpus;

        [OutputConstructor]
        private GetInstanceTypeResult(
            Outputs.GetInstanceTypeAddonsResult addons,
            string @class,
            int disk,
            string id,
            string label,
            int memory,
            int networkOut,
            Outputs.GetInstanceTypePriceResult price,
            int transfer,
            int vcpus)
        {
            Addons = addons;
            Class = @class;
            Disk = disk;
            Id = id;
            Label = label;
            Memory = memory;
            NetworkOut = networkOut;
            Price = price;
            Transfer = transfer;
            Vcpus = vcpus;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetInstanceTypeAddonsBackupsPriceResult
    {
        public readonly double Hourly;
        public readonly double Monthly;

        [OutputConstructor]
        private GetInstanceTypeAddonsBackupsPriceResult(
            double hourly,
            double monthly)
        {
            Hourly = hourly;
            Monthly = monthly;
        }
    }

    [OutputType]
    public sealed class GetInstanceTypeAddonsBackupsResult
    {
        public readonly GetInstanceTypeAddonsBackupsPriceResult Price;

        [OutputConstructor]
        private GetInstanceTypeAddonsBackupsResult(GetInstanceTypeAddonsBackupsPriceResult price)
        {
            Price = price;
        }
    }

    [OutputType]
    public sealed class GetInstanceTypeAddonsResult
    {
        public readonly GetInstanceTypeAddonsBackupsResult Backups;

        [OutputConstructor]
        private GetInstanceTypeAddonsResult(GetInstanceTypeAddonsBackupsResult backups)
        {
            Backups = backups;
        }
    }

    [OutputType]
    public sealed class GetInstanceTypePriceResult
    {
        public readonly double Hourly;
        public readonly double Monthly;

        [OutputConstructor]
        private GetInstanceTypePriceResult(
            double hourly,
            double monthly)
        {
            Hourly = hourly;
            Monthly = monthly;
        }
    }
    }
}
