// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class NodeBalancerTransferGetArgs : Pulumi.ResourceArgs
    {
        [Input("in")]
        public Input<double>? In { get; set; }

        [Input("out")]
        public Input<double>? Out { get; set; }

        [Input("total")]
        public Input<double>? Total { get; set; }

        public NodeBalancerTransferGetArgs()
        {
        }
    }
}