// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-linode/blob/master/website/docs/r/stackscript.html.markdown.
 */
export class StackScript extends pulumi.CustomResource {
    /**
     * Get an existing StackScript resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackScriptState, opts?: pulumi.CustomResourceOptions): StackScript {
        return new StackScript(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'linode:index/stackScript:StackScript';

    /**
     * Returns true if the given object is an instance of StackScript.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StackScript {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StackScript.__pulumiType;
    }

    /**
     * The date this StackScript was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Count of currently active, deployed Linodes created from this StackScript.
     */
    public /*out*/ readonly deploymentsActive!: pulumi.Output<number>;
    /**
     * The total number of times this StackScript has been deployed.
     */
    public /*out*/ readonly deploymentsTotal!: pulumi.Output<number>;
    /**
     * A description for the StackScript.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
     */
    public readonly images!: pulumi.Output<string[]>;
    /**
     * This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private. *Changing `isPublic` forces the creation of a new StackScript*
     */
    public readonly isPublic!: pulumi.Output<boolean | undefined>;
    /**
     * The StackScript's label is for display purposes only.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * This field allows you to add notes for the set of revisions made to this StackScript.
     */
    public readonly revNote!: pulumi.Output<string | undefined>;
    /**
     * The script to execute when provisioning a new Linode with this StackScript.
     */
    public readonly script!: pulumi.Output<string>;
    /**
     * The date this StackScript was updated.
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;
    /**
     * This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized
     * parameters during deployment.
     */
    public readonly userDefinedFields!: pulumi.Output<{ default: string, example: string, label: string, manyOf: string, name: string, oneOf: string }[]>;
    /**
     * The Gravatar ID for the User who created the StackScript.
     */
    public /*out*/ readonly userGravatarId!: pulumi.Output<string>;
    /**
     * The User who created the StackScript.
     */
    public /*out*/ readonly username!: pulumi.Output<string>;

    /**
     * Create a StackScript resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackScriptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackScriptArgs | StackScriptState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StackScriptState | undefined;
            inputs["created"] = state ? state.created : undefined;
            inputs["deploymentsActive"] = state ? state.deploymentsActive : undefined;
            inputs["deploymentsTotal"] = state ? state.deploymentsTotal : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["images"] = state ? state.images : undefined;
            inputs["isPublic"] = state ? state.isPublic : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["revNote"] = state ? state.revNote : undefined;
            inputs["script"] = state ? state.script : undefined;
            inputs["updated"] = state ? state.updated : undefined;
            inputs["userDefinedFields"] = state ? state.userDefinedFields : undefined;
            inputs["userGravatarId"] = state ? state.userGravatarId : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as StackScriptArgs | undefined;
            if (!args || args.description === undefined) {
                throw new Error("Missing required property 'description'");
            }
            if (!args || args.images === undefined) {
                throw new Error("Missing required property 'images'");
            }
            if (!args || args.label === undefined) {
                throw new Error("Missing required property 'label'");
            }
            if (!args || args.script === undefined) {
                throw new Error("Missing required property 'script'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["images"] = args ? args.images : undefined;
            inputs["isPublic"] = args ? args.isPublic : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["revNote"] = args ? args.revNote : undefined;
            inputs["script"] = args ? args.script : undefined;
            inputs["userDefinedFields"] = args ? args.userDefinedFields : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["deploymentsActive"] = undefined /*out*/;
            inputs["deploymentsTotal"] = undefined /*out*/;
            inputs["updated"] = undefined /*out*/;
            inputs["userGravatarId"] = undefined /*out*/;
            inputs["username"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StackScript.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StackScript resources.
 */
export interface StackScriptState {
    /**
     * The date this StackScript was created.
     */
    readonly created?: pulumi.Input<string>;
    /**
     * Count of currently active, deployed Linodes created from this StackScript.
     */
    readonly deploymentsActive?: pulumi.Input<number>;
    /**
     * The total number of times this StackScript has been deployed.
     */
    readonly deploymentsTotal?: pulumi.Input<number>;
    /**
     * A description for the StackScript.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
     */
    readonly images?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private. *Changing `isPublic` forces the creation of a new StackScript*
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * The StackScript's label is for display purposes only.
     */
    readonly label?: pulumi.Input<string>;
    /**
     * This field allows you to add notes for the set of revisions made to this StackScript.
     */
    readonly revNote?: pulumi.Input<string>;
    /**
     * The script to execute when provisioning a new Linode with this StackScript.
     */
    readonly script?: pulumi.Input<string>;
    /**
     * The date this StackScript was updated.
     */
    readonly updated?: pulumi.Input<string>;
    /**
     * This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized
     * parameters during deployment.
     */
    readonly userDefinedFields?: pulumi.Input<pulumi.Input<{ default?: pulumi.Input<string>, example?: pulumi.Input<string>, label?: pulumi.Input<string>, manyOf?: pulumi.Input<string>, name?: pulumi.Input<string>, oneOf?: pulumi.Input<string> }>[]>;
    /**
     * The Gravatar ID for the User who created the StackScript.
     */
    readonly userGravatarId?: pulumi.Input<string>;
    /**
     * The User who created the StackScript.
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StackScript resource.
 */
export interface StackScriptArgs {
    /**
     * A description for the StackScript.
     */
    readonly description: pulumi.Input<string>;
    /**
     * An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
     */
    readonly images: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private. *Changing `isPublic` forces the creation of a new StackScript*
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * The StackScript's label is for display purposes only.
     */
    readonly label: pulumi.Input<string>;
    /**
     * This field allows you to add notes for the set of revisions made to this StackScript.
     */
    readonly revNote?: pulumi.Input<string>;
    /**
     * The script to execute when provisioning a new Linode with this StackScript.
     */
    readonly script: pulumi.Input<string>;
    /**
     * This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized
     * parameters during deployment.
     */
    readonly userDefinedFields?: pulumi.Input<pulumi.Input<{ default?: pulumi.Input<string>, example?: pulumi.Input<string>, label?: pulumi.Input<string>, manyOf?: pulumi.Input<string>, name?: pulumi.Input<string>, oneOf?: pulumi.Input<string> }>[]>;
}
