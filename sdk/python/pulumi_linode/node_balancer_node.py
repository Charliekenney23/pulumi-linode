# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class NodeBalancerNode(pulumi.CustomResource):
    address: pulumi.Output[str]
    """
    The private IP Address where this backend can be reached. This must be a private IP address.
    """
    config_id: pulumi.Output[float]
    """
    The ID of the NodeBalancerConfig to access.
    """
    label: pulumi.Output[str]
    """
    The label of the Linode NodeBalancer Node. This is for display purposes only.
    """
    mode: pulumi.Output[str]
    """
    The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it
    """
    nodebalancer_id: pulumi.Output[float]
    """
    The ID of the NodeBalancer to access.
    """
    status: pulumi.Output[str]
    weight: pulumi.Output[float]
    """
    Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).
    """
    def __init__(__self__, resource_name, opts=None, address=None, config_id=None, label=None, mode=None, nodebalancer_id=None, weight=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a NodeBalancerNode resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The private IP Address where this backend can be reached. This must be a private IP address.
        :param pulumi.Input[float] config_id: The ID of the NodeBalancerConfig to access.
        :param pulumi.Input[str] label: The label of the Linode NodeBalancer Node. This is for display purposes only.
        :param pulumi.Input[str] mode: The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it
        :param pulumi.Input[float] nodebalancer_id: The ID of the NodeBalancer to access.
        :param pulumi.Input[float] weight: Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/nodebalancer_node.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if address is None:
                raise TypeError("Missing required property 'address'")
            __props__['address'] = address
            if config_id is None:
                raise TypeError("Missing required property 'config_id'")
            __props__['config_id'] = config_id
            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
            __props__['mode'] = mode
            if nodebalancer_id is None:
                raise TypeError("Missing required property 'nodebalancer_id'")
            __props__['nodebalancer_id'] = nodebalancer_id
            __props__['weight'] = weight
            __props__['status'] = None
        super(NodeBalancerNode, __self__).__init__(
            'linode:index/nodeBalancerNode:NodeBalancerNode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address=None, config_id=None, label=None, mode=None, nodebalancer_id=None, status=None, weight=None):
        """
        Get an existing NodeBalancerNode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The private IP Address where this backend can be reached. This must be a private IP address.
        :param pulumi.Input[float] config_id: The ID of the NodeBalancerConfig to access.
        :param pulumi.Input[str] label: The label of the Linode NodeBalancer Node. This is for display purposes only.
        :param pulumi.Input[str] mode: The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it
        :param pulumi.Input[float] nodebalancer_id: The ID of the NodeBalancer to access.
        :param pulumi.Input[float] weight: Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/nodebalancer_node.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["address"] = address
        __props__["config_id"] = config_id
        __props__["label"] = label
        __props__["mode"] = mode
        __props__["nodebalancer_id"] = nodebalancer_id
        __props__["status"] = status
        __props__["weight"] = weight
        return NodeBalancerNode(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

