// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Linode.Inputs
{

    public sealed class InstanceConfigDevicesGetArgs : Pulumi.ResourceArgs
    {
        [Input("sda")]
        public Input<Inputs.InstanceConfigDevicesSdaGetArgs>? Sda { get; set; }

        [Input("sdb")]
        public Input<Inputs.InstanceConfigDevicesSdbGetArgs>? Sdb { get; set; }

        [Input("sdc")]
        public Input<Inputs.InstanceConfigDevicesSdcGetArgs>? Sdc { get; set; }

        [Input("sdd")]
        public Input<Inputs.InstanceConfigDevicesSddGetArgs>? Sdd { get; set; }

        [Input("sde")]
        public Input<Inputs.InstanceConfigDevicesSdeGetArgs>? Sde { get; set; }

        [Input("sdf")]
        public Input<Inputs.InstanceConfigDevicesSdfGetArgs>? Sdf { get; set; }

        [Input("sdg")]
        public Input<Inputs.InstanceConfigDevicesSdgGetArgs>? Sdg { get; set; }

        [Input("sdh")]
        public Input<Inputs.InstanceConfigDevicesSdhGetArgs>? Sdh { get; set; }

        public InstanceConfigDevicesGetArgs()
        {
        }
    }
}
