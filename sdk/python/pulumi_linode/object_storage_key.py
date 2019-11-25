# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class ObjectStorageKey(pulumi.CustomResource):
    access_key: pulumi.Output[str]
    label: pulumi.Output[str]
    """
    The label given to this key. For display purposes only.
    """
    secret_key: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, label=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Linode Object Storage Key resource. This can be used to create, modify, and delete Linodes Object Storage Keys.
        
        ## Attributes
        
        This resource exports the following attributes:
        
        * `access_key` - This keypair's access key. This is not secret.
        
        * `secret_key` - This keypair's secret key.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] label: The label given to this key. For display purposes only.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/object_storage_key.html.markdown.
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

            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
            __props__['access_key'] = None
            __props__['secret_key'] = None
        super(ObjectStorageKey, __self__).__init__(
            'linode:index/objectStorageKey:ObjectStorageKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_key=None, label=None, secret_key=None):
        """
        Get an existing ObjectStorageKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] label: The label given to this key. For display purposes only.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/object_storage_key.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["access_key"] = access_key
        __props__["label"] = label
        __props__["secret_key"] = secret_key
        return ObjectStorageKey(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

