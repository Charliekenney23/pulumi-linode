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
        /// Provides information about a Linode Object Storage Cluster
        /// 
        /// ## Attributes
        /// 
        /// The Linode Object Storage Cluster resource exports the following attributes:
        /// 
        /// * `domain` - The base URL for this cluster.
        /// 
        /// * `status` - This cluster's status.
        /// 
        /// * `region` - The region this cluster is located in.
        /// 
        /// * `static_site_domain` - The base URL for this cluster used when hosting static sites.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/object_storage_cluster.html.markdown.
        /// </summary>
        public static Task<GetObjectStorageClusterResult> GetObjectStorageCluster(GetObjectStorageClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetObjectStorageClusterResult>("linode:index/getObjectStorageCluster:getObjectStorageCluster", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetObjectStorageClusterArgs : Pulumi.ResourceArgs
    {
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The unique ID of this cluster. 
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("staticSiteDomain")]
        public Input<string>? StaticSiteDomain { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetObjectStorageClusterArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetObjectStorageClusterResult
    {
        public readonly string Domain;
        public readonly string Id;
        public readonly string Region;
        public readonly string StaticSiteDomain;
        public readonly string Status;

        [OutputConstructor]
        private GetObjectStorageClusterResult(
            string domain,
            string id,
            string region,
            string staticSiteDomain,
            string status)
        {
            Domain = domain;
            Id = id;
            Region = region;
            StaticSiteDomain = staticSiteDomain;
            Status = status;
        }
    }
}
