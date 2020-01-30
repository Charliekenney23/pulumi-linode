// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package linode

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Linode RDNS resource.  This can be used to create and modify RDNS records.
// 
// Linode RDNS names must have a matching address value in an A or AAAA record.  This A or AAAA name must be resolvable at the time the RDNS resource is being associated.
// 
// For more information, see the [Linode APIv4 docs](https://developers.linode.com/api/docs/v4#operation/updateIP) and the [Configure your Linode for Reverse DNS](https://www.linode.com/docs/networking/dns/configure-your-linode-for-reverse-dns-classic-manager/) guide.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/rdns.html.markdown.
type Rdns struct {
	pulumi.CustomResourceState

	// The Public IPv4 or IPv6 address that will receive the `PTR` record.  A matching `A` or `AAAA` record must exist.
	Address pulumi.StringOutput `pulumi:"address"`
	// The name of the RDNS address.
	Rdns pulumi.StringOutput `pulumi:"rdns"`
}

// NewRdns registers a new resource with the given unique name, arguments, and options.
func NewRdns(ctx *pulumi.Context,
	name string, args *RdnsArgs, opts ...pulumi.ResourceOption) (*Rdns, error) {
	if args == nil || args.Address == nil {
		return nil, errors.New("missing required argument 'Address'")
	}
	if args == nil || args.Rdns == nil {
		return nil, errors.New("missing required argument 'Rdns'")
	}
	if args == nil {
		args = &RdnsArgs{}
	}
	var resource Rdns
	err := ctx.RegisterResource("linode:index/rdns:Rdns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdns gets an existing Rdns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdnsState, opts ...pulumi.ResourceOption) (*Rdns, error) {
	var resource Rdns
	err := ctx.ReadResource("linode:index/rdns:Rdns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rdns resources.
type rdnsState struct {
	// The Public IPv4 or IPv6 address that will receive the `PTR` record.  A matching `A` or `AAAA` record must exist.
	Address *string `pulumi:"address"`
	// The name of the RDNS address.
	Rdns *string `pulumi:"rdns"`
}

type RdnsState struct {
	// The Public IPv4 or IPv6 address that will receive the `PTR` record.  A matching `A` or `AAAA` record must exist.
	Address pulumi.StringPtrInput
	// The name of the RDNS address.
	Rdns pulumi.StringPtrInput
}

func (RdnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdnsState)(nil)).Elem()
}

type rdnsArgs struct {
	// The Public IPv4 or IPv6 address that will receive the `PTR` record.  A matching `A` or `AAAA` record must exist.
	Address string `pulumi:"address"`
	// The name of the RDNS address.
	Rdns string `pulumi:"rdns"`
}

// The set of arguments for constructing a Rdns resource.
type RdnsArgs struct {
	// The Public IPv4 or IPv6 address that will receive the `PTR` record.  A matching `A` or `AAAA` record must exist.
	Address pulumi.StringInput
	// The name of the RDNS address.
	Rdns pulumi.StringInput
}

func (RdnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdnsArgs)(nil)).Elem()
}

