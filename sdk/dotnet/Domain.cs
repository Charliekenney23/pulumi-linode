// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/domain.html.markdown.
    /// </summary>
    public partial class Domain : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
        /// </summary>
        [Output("axfrIps")]
        public Output<ImmutableArray<string>> AxfrIps { get; private set; } = null!;

        /// <summary>
        /// A description for this Domain. This is for display purposes only.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
        /// </summary>
        [Output("domain")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Output("expireSec")]
        public Output<int?> ExpireSec { get; private set; } = null!;

        /// <summary>
        /// The group this Domain belongs to. This is for display purposes only.
        /// </summary>
        [Output("group")]
        public Output<string?> Group { get; private set; } = null!;

        /// <summary>
        /// The IP addresses representing the master DNS for this Domain.
        /// </summary>
        [Output("masterIps")]
        public Output<ImmutableArray<string>> MasterIps { get; private set; } = null!;

        /// <summary>
        /// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Output("refreshSec")]
        public Output<int?> RefreshSec { get; private set; } = null!;

        /// <summary>
        /// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Output("retrySec")]
        public Output<int?> RetrySec { get; private set; } = null!;

        /// <summary>
        /// Start of Authority email address. This is required for master Domains.
        /// </summary>
        [Output("soaEmail")]
        public Output<string?> SoaEmail { get; private set; } = null!;

        /// <summary>
        /// Used to control whether this Domain is currently being rendered (defaults to "active").
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A list of tags applied to this object. Tags are for organizational purposes only.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Output("ttlSec")]
        public Output<int?> TtlSec { get; private set; } = null!;

        /// <summary>
        /// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("linode:index/domain:Domain", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("linode:index/domain:Domain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : Pulumi.ResourceArgs
    {
        [Input("axfrIps")]
        private InputList<string>? _axfrIps;

        /// <summary>
        /// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
        /// </summary>
        public InputList<string> AxfrIps
        {
            get => _axfrIps ?? (_axfrIps = new InputList<string>());
            set => _axfrIps = value;
        }

        /// <summary>
        /// A description for this Domain. This is for display purposes only.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("expireSec")]
        public Input<int>? ExpireSec { get; set; }

        /// <summary>
        /// The group this Domain belongs to. This is for display purposes only.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("masterIps")]
        private InputList<string>? _masterIps;

        /// <summary>
        /// The IP addresses representing the master DNS for this Domain.
        /// </summary>
        public InputList<string> MasterIps
        {
            get => _masterIps ?? (_masterIps = new InputList<string>());
            set => _masterIps = value;
        }

        /// <summary>
        /// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("refreshSec")]
        public Input<int>? RefreshSec { get; set; }

        /// <summary>
        /// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("retrySec")]
        public Input<int>? RetrySec { get; set; }

        /// <summary>
        /// Start of Authority email address. This is required for master Domains.
        /// </summary>
        [Input("soaEmail")]
        public Input<string>? SoaEmail { get; set; }

        /// <summary>
        /// Used to control whether this Domain is currently being rendered (defaults to "active").
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags applied to this object. Tags are for organizational purposes only.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("ttlSec")]
        public Input<int>? TtlSec { get; set; }

        /// <summary>
        /// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DomainArgs()
        {
        }
    }

    public sealed class DomainState : Pulumi.ResourceArgs
    {
        [Input("axfrIps")]
        private InputList<string>? _axfrIps;

        /// <summary>
        /// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
        /// </summary>
        public InputList<string> AxfrIps
        {
            get => _axfrIps ?? (_axfrIps = new InputList<string>());
            set => _axfrIps = value;
        }

        /// <summary>
        /// A description for this Domain. This is for display purposes only.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("expireSec")]
        public Input<int>? ExpireSec { get; set; }

        /// <summary>
        /// The group this Domain belongs to. This is for display purposes only.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("masterIps")]
        private InputList<string>? _masterIps;

        /// <summary>
        /// The IP addresses representing the master DNS for this Domain.
        /// </summary>
        public InputList<string> MasterIps
        {
            get => _masterIps ?? (_masterIps = new InputList<string>());
            set => _masterIps = value;
        }

        /// <summary>
        /// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("refreshSec")]
        public Input<int>? RefreshSec { get; set; }

        /// <summary>
        /// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("retrySec")]
        public Input<int>? RetrySec { get; set; }

        /// <summary>
        /// Start of Authority email address. This is required for master Domains.
        /// </summary>
        [Input("soaEmail")]
        public Input<string>? SoaEmail { get; set; }

        /// <summary>
        /// Used to control whether this Domain is currently being rendered (defaults to "active").
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags applied to this object. Tags are for organizational purposes only.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        /// </summary>
        [Input("ttlSec")]
        public Input<int>? TtlSec { get; set; }

        /// <summary>
        /// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DomainState()
        {
        }
    }
}
