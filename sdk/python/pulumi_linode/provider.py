# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, api_version=None, token=None, ua_prefix=None, url=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the linode package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: An HTTP User-Agent Prefix to prepend in API requests.
        :param pulumi.Input[str] token: The token that allows you access to your Linode account
        :param pulumi.Input[str] ua_prefix: An HTTP User-Agent Prefix to prepend in API requests.
        :param pulumi.Input[str] url: The HTTP(S) API address of the Linode API to use.
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

            if api_version is None:
                api_version = utilities.get_env('LINODE_API_VERSION')
            __props__['api_version'] = api_version
            if token is None:
                token = utilities.get_env('LINODE_TOKEN', 'LINODE_API_TOKEN')
            __props__['token'] = token
            if ua_prefix is None:
                ua_prefix = utilities.get_env('LINODE_UA_PREFIX')
            __props__['ua_prefix'] = ua_prefix
            if url is None:
                url = utilities.get_env('LINODE_URL')
            __props__['url'] = url
        super(Provider, __self__).__init__(
            'linode',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

