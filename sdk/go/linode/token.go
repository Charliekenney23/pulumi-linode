// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package linode

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Token struct {
	pulumi.CustomResourceState

	// The date and time this token was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// When this token will expire. Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated. Tokens may be created with 'null' as their expiry and will never expire unless revoked.
	Expiry pulumi.StringPtrOutput `pulumi:"expiry"`
	// A label for the Token.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the Linode CLI, require tokens with access to *. Tokens with more restrictive scopes are generally more secure.
	Scopes pulumi.StringOutput `pulumi:"scopes"`
	// The token used to access the API.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewToken registers a new resource with the given unique name, arguments, and options.
func NewToken(ctx *pulumi.Context,
	name string, args *TokenArgs, opts ...pulumi.ResourceOption) (*Token, error) {
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	if args == nil {
		args = &TokenArgs{}
	}
	var resource Token
	err := ctx.RegisterResource("linode:index/token:Token", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetToken gets an existing Token resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenState, opts ...pulumi.ResourceOption) (*Token, error) {
	var resource Token
	err := ctx.ReadResource("linode:index/token:Token", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Token resources.
type tokenState struct {
	// The date and time this token was created.
	Created *string `pulumi:"created"`
	// When this token will expire. Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated. Tokens may be created with 'null' as their expiry and will never expire unless revoked.
	Expiry *string `pulumi:"expiry"`
	// A label for the Token.
	Label *string `pulumi:"label"`
	// The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the Linode CLI, require tokens with access to *. Tokens with more restrictive scopes are generally more secure.
	Scopes *string `pulumi:"scopes"`
	// The token used to access the API.
	Token *string `pulumi:"token"`
}

type TokenState struct {
	// The date and time this token was created.
	Created pulumi.StringPtrInput
	// When this token will expire. Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated. Tokens may be created with 'null' as their expiry and will never expire unless revoked.
	Expiry pulumi.StringPtrInput
	// A label for the Token.
	Label pulumi.StringPtrInput
	// The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the Linode CLI, require tokens with access to *. Tokens with more restrictive scopes are generally more secure.
	Scopes pulumi.StringPtrInput
	// The token used to access the API.
	Token pulumi.StringPtrInput
}

func (TokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenState)(nil)).Elem()
}

type tokenArgs struct {
	// When this token will expire. Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated. Tokens may be created with 'null' as their expiry and will never expire unless revoked.
	Expiry *string `pulumi:"expiry"`
	// A label for the Token.
	Label *string `pulumi:"label"`
	// The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the Linode CLI, require tokens with access to *. Tokens with more restrictive scopes are generally more secure.
	Scopes string `pulumi:"scopes"`
}

// The set of arguments for constructing a Token resource.
type TokenArgs struct {
	// When this token will expire. Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated. Tokens may be created with 'null' as their expiry and will never expire unless revoked.
	Expiry pulumi.StringPtrInput
	// A label for the Token.
	Label pulumi.StringPtrInput
	// The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the Linode CLI, require tokens with access to *. Tokens with more restrictive scopes are generally more secure.
	Scopes pulumi.StringInput
}

func (TokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenArgs)(nil)).Elem()
}
