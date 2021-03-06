// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package linode

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/domain.html.markdown.
type Domain struct {
	pulumi.CustomResourceState

	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	AxfrIps pulumi.StringArrayOutput `pulumi:"axfrIps"`
	// A description for this Domain. This is for display purposes only.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	ExpireSec pulumi.IntPtrOutput `pulumi:"expireSec"`
	// The group this Domain belongs to. This is for display purposes only.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// The IP addresses representing the master DNS for this Domain.
	MasterIps pulumi.StringArrayOutput `pulumi:"masterIps"`
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RefreshSec pulumi.IntPtrOutput `pulumi:"refreshSec"`
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RetrySec pulumi.IntPtrOutput `pulumi:"retrySec"`
	// Start of Authority email address. This is required for master Domains.
	SoaEmail pulumi.StringPtrOutput `pulumi:"soaEmail"`
	// Used to control whether this Domain is currently being rendered (defaults to "active").
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec pulumi.IntPtrOutput `pulumi:"ttlSec"`
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &DomainArgs{}
	}
	var resource Domain
	err := ctx.RegisterResource("linode:index/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("linode:index/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	AxfrIps []string `pulumi:"axfrIps"`
	// A description for this Domain. This is for display purposes only.
	Description *string `pulumi:"description"`
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain *string `pulumi:"domain"`
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	ExpireSec *int `pulumi:"expireSec"`
	// The group this Domain belongs to. This is for display purposes only.
	Group *string `pulumi:"group"`
	// The IP addresses representing the master DNS for this Domain.
	MasterIps []string `pulumi:"masterIps"`
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RefreshSec *int `pulumi:"refreshSec"`
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RetrySec *int `pulumi:"retrySec"`
	// Start of Authority email address. This is required for master Domains.
	SoaEmail *string `pulumi:"soaEmail"`
	// Used to control whether this Domain is currently being rendered (defaults to "active").
	Status *string `pulumi:"status"`
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags []string `pulumi:"tags"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec *int `pulumi:"ttlSec"`
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type *string `pulumi:"type"`
}

type DomainState struct {
	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	AxfrIps pulumi.StringArrayInput
	// A description for this Domain. This is for display purposes only.
	Description pulumi.StringPtrInput
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain pulumi.StringPtrInput
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	ExpireSec pulumi.IntPtrInput
	// The group this Domain belongs to. This is for display purposes only.
	Group pulumi.StringPtrInput
	// The IP addresses representing the master DNS for this Domain.
	MasterIps pulumi.StringArrayInput
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RefreshSec pulumi.IntPtrInput
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RetrySec pulumi.IntPtrInput
	// Start of Authority email address. This is required for master Domains.
	SoaEmail pulumi.StringPtrInput
	// Used to control whether this Domain is currently being rendered (defaults to "active").
	Status pulumi.StringPtrInput
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags pulumi.StringArrayInput
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec pulumi.IntPtrInput
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	AxfrIps []string `pulumi:"axfrIps"`
	// A description for this Domain. This is for display purposes only.
	Description *string `pulumi:"description"`
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain string `pulumi:"domain"`
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	ExpireSec *int `pulumi:"expireSec"`
	// The group this Domain belongs to. This is for display purposes only.
	Group *string `pulumi:"group"`
	// The IP addresses representing the master DNS for this Domain.
	MasterIps []string `pulumi:"masterIps"`
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RefreshSec *int `pulumi:"refreshSec"`
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RetrySec *int `pulumi:"retrySec"`
	// Start of Authority email address. This is required for master Domains.
	SoaEmail *string `pulumi:"soaEmail"`
	// Used to control whether this Domain is currently being rendered (defaults to "active").
	Status *string `pulumi:"status"`
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags []string `pulumi:"tags"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec *int `pulumi:"ttlSec"`
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	AxfrIps pulumi.StringArrayInput
	// A description for this Domain. This is for display purposes only.
	Description pulumi.StringPtrInput
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain pulumi.StringInput
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	ExpireSec pulumi.IntPtrInput
	// The group this Domain belongs to. This is for display purposes only.
	Group pulumi.StringPtrInput
	// The IP addresses representing the master DNS for this Domain.
	MasterIps pulumi.StringArrayInput
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RefreshSec pulumi.IntPtrInput
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	RetrySec pulumi.IntPtrInput
	// Start of Authority email address. This is required for master Domains.
	SoaEmail pulumi.StringPtrInput
	// Used to control whether this Domain is currently being rendered (defaults to "active").
	Status pulumi.StringPtrInput
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags pulumi.StringArrayInput
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec pulumi.IntPtrInput
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

