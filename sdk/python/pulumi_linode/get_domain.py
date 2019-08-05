# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetDomainResult:
    """
    A collection of values returned by getDomain.
    """
    def __init__(__self__, axfr_ips=None, description=None, domain=None, expire_sec=None, group=None, id=None, master_ips=None, refresh_sec=None, retry_sec=None, soa_email=None, status=None, tags=None, ttl_sec=None, type=None):
        if axfr_ips and not isinstance(axfr_ips, list):
            raise TypeError("Expected argument 'axfr_ips' to be a list")
        __self__.axfr_ips = axfr_ips
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        __self__.domain = domain
        if expire_sec and not isinstance(expire_sec, float):
            raise TypeError("Expected argument 'expire_sec' to be a float")
        __self__.expire_sec = expire_sec
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        __self__.group = group
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if master_ips and not isinstance(master_ips, list):
            raise TypeError("Expected argument 'master_ips' to be a list")
        __self__.master_ips = master_ips
        if refresh_sec and not isinstance(refresh_sec, float):
            raise TypeError("Expected argument 'refresh_sec' to be a float")
        __self__.refresh_sec = refresh_sec
        if retry_sec and not isinstance(retry_sec, float):
            raise TypeError("Expected argument 'retry_sec' to be a float")
        __self__.retry_sec = retry_sec
        if soa_email and not isinstance(soa_email, str):
            raise TypeError("Expected argument 'soa_email' to be a str")
        __self__.soa_email = soa_email
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        if ttl_sec and not isinstance(ttl_sec, float):
            raise TypeError("Expected argument 'ttl_sec' to be a float")
        __self__.ttl_sec = ttl_sec
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_domain(domain=None,id=None,opts=None):
    """
    Provides information about a Linode domain.
    
    ## Attributes
    
    The Linode Domain resource exports the following attributes:
    
    * `id` - The unique ID of this Domain.
    
    * `domain` - The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain
    
    * `type` - If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave)
    
    * `group` - The group this Domain belongs to.
    
    * `status` - Used to control whether this Domain is currently being rendered.
    
    * `description` - A description for this Domain.
    
    * `master_ips` - The IP addresses representing the master DNS for this Domain.
    
    * `axfr_ips` - The list of IPs that may perform a zone transfer for this Domain.
    
    * `ttl_sec` - 'Time to Live'-the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.
    
    * `retry_sec` - The interval, in seconds, at which a failed refresh should be retried.
    *
    * `expire_sec` - The amount of time in seconds that may pass before this Domain is no longer authoritative.
    
    * `refresh_sec` - The amount of time in seconds before this Domain should be refreshed.
    
    * `soa_email` - Start of Authority email address.
    
    * `tags` - An array of tags applied to this object.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/domain.html.markdown.
    """
    __args__ = dict()

    __args__['domain'] = domain
    __args__['id'] = id
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('linode:index/getDomain:getDomain', __args__, opts=opts).value

    return GetDomainResult(
        axfr_ips=__ret__.get('axfrIps'),
        description=__ret__.get('description'),
        domain=__ret__.get('domain'),
        expire_sec=__ret__.get('expireSec'),
        group=__ret__.get('group'),
        id=__ret__.get('id'),
        master_ips=__ret__.get('masterIps'),
        refresh_sec=__ret__.get('refreshSec'),
        retry_sec=__ret__.get('retrySec'),
        soa_email=__ret__.get('soaEmail'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        ttl_sec=__ret__.get('ttlSec'),
        type=__ret__.get('type'))
