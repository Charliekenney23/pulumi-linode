// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Linode Object Storage Bucket resource. This can be used to create, modify, and delete Linodes Object Storage Buckets.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/object_storage_bucket.html.markdown.
type ObjectStorageBucket struct {
	s *pulumi.ResourceState
}

// NewObjectStorageBucket registers a new resource with the given unique name, arguments, and options.
func NewObjectStorageBucket(ctx *pulumi.Context,
	name string, args *ObjectStorageBucketArgs, opts ...pulumi.ResourceOpt) (*ObjectStorageBucket, error) {
	if args == nil || args.Cluster == nil {
		return nil, errors.New("missing required argument 'Cluster'")
	}
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cluster"] = nil
		inputs["label"] = nil
	} else {
		inputs["cluster"] = args.Cluster
		inputs["label"] = args.Label
	}
	s, err := ctx.RegisterResource("linode:index/objectStorageBucket:ObjectStorageBucket", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ObjectStorageBucket{s: s}, nil
}

// GetObjectStorageBucket gets an existing ObjectStorageBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectStorageBucket(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ObjectStorageBucketState, opts ...pulumi.ResourceOpt) (*ObjectStorageBucket, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cluster"] = state.Cluster
		inputs["label"] = state.Label
	}
	s, err := ctx.ReadResource("linode:index/objectStorageBucket:ObjectStorageBucket", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ObjectStorageBucket{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ObjectStorageBucket) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ObjectStorageBucket) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The cluster of the Linode Object Storage Bucket.
func (r *ObjectStorageBucket) Cluster() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cluster"])
}

// The label of the Linode Object Storage Bucket.
func (r *ObjectStorageBucket) Label() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["label"])
}

// Input properties used for looking up and filtering ObjectStorageBucket resources.
type ObjectStorageBucketState struct {
	// The cluster of the Linode Object Storage Bucket.
	Cluster interface{}
	// The label of the Linode Object Storage Bucket.
	Label interface{}
}

// The set of arguments for constructing a ObjectStorageBucket resource.
type ObjectStorageBucketArgs struct {
	// The cluster of the Linode Object Storage Bucket.
	Cluster interface{}
	// The label of the Linode Object Storage Bucket.
	Label interface{}
}