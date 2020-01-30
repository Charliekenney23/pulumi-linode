// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package linode

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The provider type for the linode package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/index.html.markdown.
type Provider struct {
	pulumi.ProviderResourceState

}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.StringPtr(getEnvOrDefault("", nil, "LINODE_API_VERSION").(string))
	}
	if args.Token == nil {
		args.Token = pulumi.StringPtr(getEnvOrDefault("", nil, "LINODE_TOKEN", "LINODE_API_TOKEN").(string))
	}
	if args.UaPrefix == nil {
		args.UaPrefix = pulumi.StringPtr(getEnvOrDefault("", nil, "LINODE_UA_PREFIX").(string))
	}
	if args.Url == nil {
		args.Url = pulumi.StringPtr(getEnvOrDefault("", nil, "LINODE_URL").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:linode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// An HTTP User-Agent Prefix to prepend in API requests.
	ApiVersion *string `pulumi:"apiVersion"`
	// The token that allows you access to your Linode account
	Token *string `pulumi:"token"`
	// An HTTP User-Agent Prefix to prepend in API requests.
	UaPrefix *string `pulumi:"uaPrefix"`
	// The HTTP(S) API address of the Linode API to use.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// An HTTP User-Agent Prefix to prepend in API requests.
	ApiVersion pulumi.StringPtrInput
	// The token that allows you access to your Linode account
	Token pulumi.StringPtrInput
	// An HTTP User-Agent Prefix to prepend in API requests.
	UaPrefix pulumi.StringPtrInput
	// The HTTP(S) API address of the Linode API to use.
	Url pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

