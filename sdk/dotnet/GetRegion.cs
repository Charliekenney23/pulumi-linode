// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public static partial class Invokes
    {
        /// <summary>
        /// `linode..getRegion` provides details about a specific Linode region.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/region.html.markdown.
        /// </summary>
        public static Task<GetRegionResult> GetRegion(GetRegionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionResult>("linode:index/getRegion:getRegion", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetRegionArgs : Pulumi.ResourceArgs
    {
        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetRegionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRegionResult
    {
        public readonly string Country;
        public readonly string Id;

        [OutputConstructor]
        private GetRegionResult(
            string country,
            string id)
        {
            Country = country;
            Id = id;
        }
    }
}
