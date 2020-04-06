// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package linode

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Linode domain.
//
// ## Attributes
//
// The Linode Domain resource exports the following attributes:
//
// * `id` - The unique ID of this Domain.
//
// * `domain` - The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain
//
// * `type` - If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave)
//
// * `group` - The group this Domain belongs to.
//
// * `status` - Used to control whether this Domain is currently being rendered.
//
// * `description` - A description for this Domain.
//
// * `masterIps` - The IP addresses representing the master DNS for this Domain.
//
// * `axfrIps` - The list of IPs that may perform a zone transfer for this Domain.
//
// * `ttlSec` - 'Time to Live'-the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.
//
// * `retrySec` - The interval, in seconds, at which a failed refresh should be retried.
// *
// * `expireSec` - The amount of time in seconds that may pass before this Domain is no longer authoritative.
//
// * `refreshSec` - The amount of time in seconds before this Domain should be refreshed.
//
// * `soaEmail` - Start of Authority email address.
//
// * `tags` - An array of tags applied to this object.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/domain.html.md.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("linode:index/getDomain:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomain.
type LookupDomainArgs struct {
	// The unique domain name of the Domain record to query.
	Domain *string `pulumi:"domain"`
	// The unique numeric ID of the Domain record to query.
	Id *string `pulumi:"id"`
}

// A collection of values returned by getDomain.
type LookupDomainResult struct {
	AxfrIps     []string `pulumi:"axfrIps"`
	Description string   `pulumi:"description"`
	Domain      *string  `pulumi:"domain"`
	ExpireSec   int      `pulumi:"expireSec"`
	Group       string   `pulumi:"group"`
	Id          *string  `pulumi:"id"`
	MasterIps   []string `pulumi:"masterIps"`
	RefreshSec  int      `pulumi:"refreshSec"`
	RetrySec    int      `pulumi:"retrySec"`
	SoaEmail    string   `pulumi:"soaEmail"`
	Status      string   `pulumi:"status"`
	Tags        []string `pulumi:"tags"`
	TtlSec      int      `pulumi:"ttlSec"`
	Type        string   `pulumi:"type"`
}
