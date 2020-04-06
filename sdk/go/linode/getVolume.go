// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package linode

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Linode Volume.
//
// ## Attributes
//
// The Linode Volume resource exports the following attributes:
//
// - `id` - The unique ID of this Volume.
//
// - `created` - When this Volume was created.
//
// - `status` - The current status of the Volume. Can be one of "creating", "active", "resizing", or "contactSupport".
//
// - `label` - This Volume's label is for display purposes only.
//
// - `tags` - An array of tags applied to this object.
//
// - `size` - The Volume's size, in GiB.
//
// - `region` - The datacenter in which this Volume is located.
//
// - `updated` - When this Volume was last updated.
//
// - `linodeId` - If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here. If the Volume is unattached, this value will be null.
//
// - `filesystemPath` - The full filesystem path for the Volume based on the Volume's label. Path is /dev/disk/by-id/scsi-0LinodeVolume + Volume label.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/d/volume.html.md.
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("linode:index/getVolume:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolume.
type LookupVolumeArgs struct {
	Id int `pulumi:"id"`
}

// A collection of values returned by getVolume.
type LookupVolumeResult struct {
	Created        string   `pulumi:"created"`
	FilesystemPath string   `pulumi:"filesystemPath"`
	Id             int      `pulumi:"id"`
	Label          string   `pulumi:"label"`
	LinodeId       int      `pulumi:"linodeId"`
	Region         string   `pulumi:"region"`
	Size           int      `pulumi:"size"`
	Status         string   `pulumi:"status"`
	Tags           []string `pulumi:"tags"`
	Updated        string   `pulumi:"updated"`
}
