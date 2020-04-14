// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode
{
    public partial class NodeBalancerConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.
        /// </summary>
        [Output("check")]
        public Output<string> Check { get; private set; } = null!;

        /// <summary>
        /// How many times to attempt a check before considering a backend to be down. (1-30)
        /// </summary>
        [Output("checkAttempts")]
        public Output<int> CheckAttempts { get; private set; } = null!;

        /// <summary>
        /// This value must be present in the response body of the check in order for it to pass. If this value is not present in
        /// the response body of a check request, the backend is considered to be down
        /// </summary>
        [Output("checkBody")]
        public Output<string> CheckBody { get; private set; } = null!;

        /// <summary>
        /// How often, in seconds, to check that backends are up and serving requests.
        /// </summary>
        [Output("checkInterval")]
        public Output<int> CheckInterval { get; private set; } = null!;

        /// <summary>
        /// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
        /// </summary>
        [Output("checkPassive")]
        public Output<bool> CheckPassive { get; private set; } = null!;

        /// <summary>
        /// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
        /// </summary>
        [Output("checkPath")]
        public Output<string> CheckPath { get; private set; } = null!;

        /// <summary>
        /// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
        /// </summary>
        [Output("checkTimeout")]
        public Output<int> CheckTimeout { get; private set; } = null!;

        /// <summary>
        /// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
        /// </summary>
        [Output("cipherSuite")]
        public Output<string> CipherSuite { get; private set; } = null!;

        [Output("nodeStatus")]
        public Output<Outputs.NodeBalancerConfigNodeStatus> NodeStatus { get; private set; } = null!;

        /// <summary>
        /// The ID of the NodeBalancer to access.
        /// </summary>
        [Output("nodebalancerId")]
        public Output<int> NodebalancerId { get; private set; } = null!;

        /// <summary>
        /// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. (Defaults to 80)
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (Defaults to "http")
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// The certificate this port is serving. This is not returned. If set, this field will come back as `&lt;REDACTED&gt;`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
        /// </summary>
        [Output("sslCert")]
        public Output<string?> SslCert { get; private set; } = null!;

        /// <summary>
        /// The common name for the SSL certification this port is serving if this port is not configured to use SSL.
        /// </summary>
        [Output("sslCommonname")]
        public Output<string> SslCommonname { get; private set; } = null!;

        /// <summary>
        /// The fingerprint for the SSL certification this port is serving if this port is not configured to use SSL.
        /// </summary>
        [Output("sslFingerprint")]
        public Output<string> SslFingerprint { get; private set; } = null!;

        /// <summary>
        /// The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as `&lt;REDACTED&gt;`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
        /// </summary>
        [Output("sslKey")]
        public Output<string?> SslKey { get; private set; } = null!;

        /// <summary>
        /// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
        /// </summary>
        [Output("stickiness")]
        public Output<string> Stickiness { get; private set; } = null!;


        /// <summary>
        /// Create a NodeBalancerConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeBalancerConfig(string name, NodeBalancerConfigArgs args, CustomResourceOptions? options = null)
            : base("linode:index/nodeBalancerConfig:NodeBalancerConfig", name, args ?? new NodeBalancerConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeBalancerConfig(string name, Input<string> id, NodeBalancerConfigState? state = null, CustomResourceOptions? options = null)
            : base("linode:index/nodeBalancerConfig:NodeBalancerConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NodeBalancerConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeBalancerConfig Get(string name, Input<string> id, NodeBalancerConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new NodeBalancerConfig(name, id, state, options);
        }
    }

    public sealed class NodeBalancerConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.
        /// </summary>
        [Input("check")]
        public Input<string>? Check { get; set; }

        /// <summary>
        /// How many times to attempt a check before considering a backend to be down. (1-30)
        /// </summary>
        [Input("checkAttempts")]
        public Input<int>? CheckAttempts { get; set; }

        /// <summary>
        /// This value must be present in the response body of the check in order for it to pass. If this value is not present in
        /// the response body of a check request, the backend is considered to be down
        /// </summary>
        [Input("checkBody")]
        public Input<string>? CheckBody { get; set; }

        /// <summary>
        /// How often, in seconds, to check that backends are up and serving requests.
        /// </summary>
        [Input("checkInterval")]
        public Input<int>? CheckInterval { get; set; }

        /// <summary>
        /// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
        /// </summary>
        [Input("checkPassive")]
        public Input<bool>? CheckPassive { get; set; }

        /// <summary>
        /// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
        /// </summary>
        [Input("checkPath")]
        public Input<string>? CheckPath { get; set; }

        /// <summary>
        /// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
        /// </summary>
        [Input("checkTimeout")]
        public Input<int>? CheckTimeout { get; set; }

        /// <summary>
        /// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
        /// </summary>
        [Input("cipherSuite")]
        public Input<string>? CipherSuite { get; set; }

        /// <summary>
        /// The ID of the NodeBalancer to access.
        /// </summary>
        [Input("nodebalancerId", required: true)]
        public Input<int> NodebalancerId { get; set; } = null!;

        /// <summary>
        /// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. (Defaults to 80)
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (Defaults to "http")
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The certificate this port is serving. This is not returned. If set, this field will come back as `&lt;REDACTED&gt;`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
        /// </summary>
        [Input("sslCert")]
        public Input<string>? SslCert { get; set; }

        /// <summary>
        /// The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as `&lt;REDACTED&gt;`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
        /// </summary>
        [Input("sslKey")]
        public Input<string>? SslKey { get; set; }

        /// <summary>
        /// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
        /// </summary>
        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        public NodeBalancerConfigArgs()
        {
        }
    }

    public sealed class NodeBalancerConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.
        /// </summary>
        [Input("check")]
        public Input<string>? Check { get; set; }

        /// <summary>
        /// How many times to attempt a check before considering a backend to be down. (1-30)
        /// </summary>
        [Input("checkAttempts")]
        public Input<int>? CheckAttempts { get; set; }

        /// <summary>
        /// This value must be present in the response body of the check in order for it to pass. If this value is not present in
        /// the response body of a check request, the backend is considered to be down
        /// </summary>
        [Input("checkBody")]
        public Input<string>? CheckBody { get; set; }

        /// <summary>
        /// How often, in seconds, to check that backends are up and serving requests.
        /// </summary>
        [Input("checkInterval")]
        public Input<int>? CheckInterval { get; set; }

        /// <summary>
        /// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
        /// </summary>
        [Input("checkPassive")]
        public Input<bool>? CheckPassive { get; set; }

        /// <summary>
        /// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
        /// </summary>
        [Input("checkPath")]
        public Input<string>? CheckPath { get; set; }

        /// <summary>
        /// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
        /// </summary>
        [Input("checkTimeout")]
        public Input<int>? CheckTimeout { get; set; }

        /// <summary>
        /// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
        /// </summary>
        [Input("cipherSuite")]
        public Input<string>? CipherSuite { get; set; }

        [Input("nodeStatus")]
        public Input<Inputs.NodeBalancerConfigNodeStatusGetArgs>? NodeStatus { get; set; }

        /// <summary>
        /// The ID of the NodeBalancer to access.
        /// </summary>
        [Input("nodebalancerId")]
        public Input<int>? NodebalancerId { get; set; }

        /// <summary>
        /// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. (Defaults to 80)
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (Defaults to "http")
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The certificate this port is serving. This is not returned. If set, this field will come back as `&lt;REDACTED&gt;`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
        /// </summary>
        [Input("sslCert")]
        public Input<string>? SslCert { get; set; }

        /// <summary>
        /// The common name for the SSL certification this port is serving if this port is not configured to use SSL.
        /// </summary>
        [Input("sslCommonname")]
        public Input<string>? SslCommonname { get; set; }

        /// <summary>
        /// The fingerprint for the SSL certification this port is serving if this port is not configured to use SSL.
        /// </summary>
        [Input("sslFingerprint")]
        public Input<string>? SslFingerprint { get; set; }

        /// <summary>
        /// The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as `&lt;REDACTED&gt;`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
        /// </summary>
        [Input("sslKey")]
        public Input<string>? SslKey { get; set; }

        /// <summary>
        /// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
        /// </summary>
        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        public NodeBalancerConfigState()
        {
        }
    }
}
