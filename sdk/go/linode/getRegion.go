// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package linode

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `.getRegion` provides details about a specific Linode region.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/region.html.markdown.
func GetRegion(ctx *pulumi.Context, args *GetRegionArgs, opts ...pulumi.InvokeOption) (*GetRegionResult, error) {
	var rv GetRegionResult
	err := ctx.Invoke("linode:index/getRegion:getRegion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegion.
type GetRegionArgs struct {
	Country *string `pulumi:"country"`
	Id string `pulumi:"id"`
}


// A collection of values returned by getRegion.
type GetRegionResult struct {
	Country string `pulumi:"country"`
	Id string `pulumi:"id"`
}

