// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Linode Domain Record resource.  This can be used to create, modify, and delete Linodes Domain Records.
// For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createDomainRecord).
//
//
// ## Attributes
//
// This resource exports no additional attributes.
type DomainRecord struct {
	pulumi.CustomResourceState

	// The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
	DomainId pulumi.IntOutput `pulumi:"domainId"`
	// The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port this Record points to.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The priority of the target host. Lower values are preferred.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The protocol this Record's service communicates with. Only valid for SRV records.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `recordType` forces the creation of a new Linode Domain Record.*.
	RecordType pulumi.StringOutput `pulumi:"recordType"`
	// The service this Record identified. Only valid for SRV records.
	Service pulumi.StringPtrOutput `pulumi:"service"`
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	Target pulumi.StringOutput `pulumi:"target"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec pulumi.IntPtrOutput `pulumi:"ttlSec"`
	// The relative weight of this Record. Higher values are preferred.
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewDomainRecord registers a new resource with the given unique name, arguments, and options.
func NewDomainRecord(ctx *pulumi.Context,
	name string, args *DomainRecordArgs, opts ...pulumi.ResourceOption) (*DomainRecord, error) {
	if args == nil || args.DomainId == nil {
		return nil, errors.New("missing required argument 'DomainId'")
	}
	if args == nil || args.RecordType == nil {
		return nil, errors.New("missing required argument 'RecordType'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	if args == nil {
		args = &DomainRecordArgs{}
	}
	var resource DomainRecord
	err := ctx.RegisterResource("linode:index/domainRecord:DomainRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainRecord gets an existing DomainRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainRecordState, opts ...pulumi.ResourceOption) (*DomainRecord, error) {
	var resource DomainRecord
	err := ctx.ReadResource("linode:index/domainRecord:DomainRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainRecord resources.
type domainRecordState struct {
	// The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
	DomainId *int `pulumi:"domainId"`
	// The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name *string `pulumi:"name"`
	// The port this Record points to.
	Port *int `pulumi:"port"`
	// The priority of the target host. Lower values are preferred.
	Priority *int `pulumi:"priority"`
	// The protocol this Record's service communicates with. Only valid for SRV records.
	Protocol *string `pulumi:"protocol"`
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `recordType` forces the creation of a new Linode Domain Record.*.
	RecordType *string `pulumi:"recordType"`
	// The service this Record identified. Only valid for SRV records.
	Service *string `pulumi:"service"`
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	Tag *string `pulumi:"tag"`
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	Target *string `pulumi:"target"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec *int `pulumi:"ttlSec"`
	// The relative weight of this Record. Higher values are preferred.
	Weight *int `pulumi:"weight"`
}

type DomainRecordState struct {
	// The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
	DomainId pulumi.IntPtrInput
	// The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name pulumi.StringPtrInput
	// The port this Record points to.
	Port pulumi.IntPtrInput
	// The priority of the target host. Lower values are preferred.
	Priority pulumi.IntPtrInput
	// The protocol this Record's service communicates with. Only valid for SRV records.
	Protocol pulumi.StringPtrInput
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `recordType` forces the creation of a new Linode Domain Record.*.
	RecordType pulumi.StringPtrInput
	// The service this Record identified. Only valid for SRV records.
	Service pulumi.StringPtrInput
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	Tag pulumi.StringPtrInput
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	Target pulumi.StringPtrInput
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec pulumi.IntPtrInput
	// The relative weight of this Record. Higher values are preferred.
	Weight pulumi.IntPtrInput
}

func (DomainRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRecordState)(nil)).Elem()
}

type domainRecordArgs struct {
	// The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
	DomainId int `pulumi:"domainId"`
	// The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name *string `pulumi:"name"`
	// The port this Record points to.
	Port *int `pulumi:"port"`
	// The priority of the target host. Lower values are preferred.
	Priority *int `pulumi:"priority"`
	// The protocol this Record's service communicates with. Only valid for SRV records.
	Protocol *string `pulumi:"protocol"`
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `recordType` forces the creation of a new Linode Domain Record.*.
	RecordType string `pulumi:"recordType"`
	// The service this Record identified. Only valid for SRV records.
	Service *string `pulumi:"service"`
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	Tag *string `pulumi:"tag"`
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	Target string `pulumi:"target"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec *int `pulumi:"ttlSec"`
	// The relative weight of this Record. Higher values are preferred.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a DomainRecord resource.
type DomainRecordArgs struct {
	// The ID of the Domain to access.  *Changing `domainId` forces the creation of a new Linode Domain Record.*.
	DomainId pulumi.IntInput
	// The name of this Record. Setting this is invalid for `SRV` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name pulumi.StringPtrInput
	// The port this Record points to.
	Port pulumi.IntPtrInput
	// The priority of the target host. Lower values are preferred.
	Priority pulumi.IntPtrInput
	// The protocol this Record's service communicates with. Only valid for SRV records.
	Protocol pulumi.StringPtrInput
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `recordType` forces the creation of a new Linode Domain Record.*.
	RecordType pulumi.StringInput
	// The service this Record identified. Only valid for SRV records.
	Service pulumi.StringPtrInput
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	Tag pulumi.StringPtrInput
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	Target pulumi.StringInput
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec pulumi.IntPtrInput
	// The relative weight of this Record. Higher values are preferred.
	Weight pulumi.IntPtrInput
}

func (DomainRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainRecordArgs)(nil)).Elem()
}
