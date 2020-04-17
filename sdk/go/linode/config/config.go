// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// An HTTP User-Agent Prefix to prepend in API requests.
func GetApiVersion(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "linode:apiVersion")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "LINODE_API_VERSION").(string)
}

// The token that allows you access to your Linode account
func GetToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "linode:token")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "LINODE_TOKEN", "LINODE_API_TOKEN").(string)
}

// An HTTP User-Agent Prefix to prepend in API requests.
func GetUaPrefix(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "linode:uaPrefix")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "LINODE_UA_PREFIX").(string)
}

// The HTTP(S) API address of the Linode API to use.
func GetUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "linode:url")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "LINODE_URL").(string)
}