// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    public readonly alerts: pulumi.Output<{ cpu: number, io: number, networkIn: number, networkOut: number, transferQuota: number }>;
    /**
     * A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is
     * provided.
     */
    public readonly authorizedKeys: pulumi.Output<string[] | undefined>;
    /**
     * A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root`
     * user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.
     */
    public readonly authorizedUsers: pulumi.Output<string[] | undefined>;
    /**
     * A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the
     * Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See
     * /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually
     * exclusive.
     */
    public readonly backupId: pulumi.Output<number | undefined>;
    /**
     * Information about this Linode's backups status.
     */
    public /*out*/ readonly backups: pulumi.Output<{ enabled: boolean, schedule: { day: string, window: string } }>;
    /**
     * If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This
     * will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
     */
    public readonly backupsEnabled: pulumi.Output<boolean>;
    /**
     * The Label of the Instance Config that should be used to boot the Linode instance.
     */
    public readonly bootConfigLabel: pulumi.Output<string>;
    /**
     * Configuration profiles define the VM settings and boot behavior of the Linode Instance.
     */
    public readonly configs: pulumi.Output<{ comments?: string, devices: { sda: { diskId: number, diskLabel?: string, volumeId?: number }, sdb: { diskId: number, diskLabel?: string, volumeId?: number }, sdc: { diskId: number, diskLabel?: string, volumeId?: number }, sdd: { diskId: number, diskLabel?: string, volumeId?: number }, sde: { diskId: number, diskLabel?: string, volumeId?: number }, sdf: { diskId: number, diskLabel?: string, volumeId?: number }, sdg: { diskId: number, diskLabel?: string, volumeId?: number }, sdh: { diskId: number, diskLabel?: string, volumeId?: number } }, helpers: { devtmpfsAutomount?: boolean, distro?: boolean, modulesDep?: boolean, network?: boolean, updatedbDisabled?: boolean }, kernel?: string, label: string, memoryLimit?: number, rootDevice: string, runLevel?: string, virtMode?: string }[] | undefined>;
    public readonly disks: pulumi.Output<{ authorizedKeys?: string[], authorizedUsers?: string[], filesystem: string, id: number, image: string, label: string, readOnly: boolean, rootPass?: string, size: number, stackscriptData: {[key: string]: any}, stackscriptId: number }[] | undefined>;
    /**
     * The display group of the Linode instance.
     */
    public readonly group: pulumi.Output<string | undefined>;
    /**
     * An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with
     * private/. See /images for more information on the Images available for you to use.
     */
    public readonly image: pulumi.Output<string | undefined>;
    /**
     * This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary
     * address will be used for this field.
     */
    public /*out*/ readonly ipAddress: pulumi.Output<string>;
    /**
     * This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a
     * single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.
     */
    public /*out*/ readonly ipv4s: pulumi.Output<string[]>;
    /**
     * This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.
     */
    public /*out*/ readonly ipv6: pulumi.Output<string>;
    /**
     * The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned
     */
    public readonly label: pulumi.Output<string>;
    /**
     * If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network
     * within the Linode's region.
     */
    public readonly privateIp: pulumi.Output<boolean | undefined>;
    /**
     * This Linode's Private IPv4 Address. The regional private IP address range is 192.168.128/17 address shared by all
     * Linode Instances in a region.
     */
    public /*out*/ readonly privateIpAddress: pulumi.Output<string>;
    /**
     * This is the location where the Linode was deployed. This cannot be changed without opening a support ticket.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * The password that will be initialially assigned to the 'root' user account.
     */
    public readonly rootPass: pulumi.Output<string | undefined>;
    public /*out*/ readonly specs: pulumi.Output<{ disk: number, memory: number, transfer: number, vcpus: number }>;
    /**
     * An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode.
     * Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
     */
    public readonly stackscriptData: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an
     * Image that is compatible with this StackScript.
     */
    public readonly stackscriptId: pulumi.Output<number | undefined>;
    /**
     * The status of the instance, indicating the current readiness state.
     */
    public /*out*/ readonly status: pulumi.Output<string>;
    /**
     * When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored.
     * This is used to set the swap disk size for the newly-created Linode.
     */
    public readonly swapSize: pulumi.Output<number>;
    /**
     * An array of tags applied to this object. Tags are for organizational purposes only.
     */
    public readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * The type of instance to be deployed, determining the price and size.
     */
    public readonly type: pulumi.Output<string | undefined>;
    /**
     * The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off
     * unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible.
     * To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
     */
    public readonly watchdogEnabled: pulumi.Output<boolean | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceState = argsOrState as InstanceState | undefined;
            inputs["alerts"] = state ? state.alerts : undefined;
            inputs["authorizedKeys"] = state ? state.authorizedKeys : undefined;
            inputs["authorizedUsers"] = state ? state.authorizedUsers : undefined;
            inputs["backupId"] = state ? state.backupId : undefined;
            inputs["backups"] = state ? state.backups : undefined;
            inputs["backupsEnabled"] = state ? state.backupsEnabled : undefined;
            inputs["bootConfigLabel"] = state ? state.bootConfigLabel : undefined;
            inputs["configs"] = state ? state.configs : undefined;
            inputs["disks"] = state ? state.disks : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["image"] = state ? state.image : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["ipv4s"] = state ? state.ipv4s : undefined;
            inputs["ipv6"] = state ? state.ipv6 : undefined;
            inputs["label"] = state ? state.label : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["rootPass"] = state ? state.rootPass : undefined;
            inputs["specs"] = state ? state.specs : undefined;
            inputs["stackscriptData"] = state ? state.stackscriptData : undefined;
            inputs["stackscriptId"] = state ? state.stackscriptId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["swapSize"] = state ? state.swapSize : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["watchdogEnabled"] = state ? state.watchdogEnabled : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            inputs["alerts"] = args ? args.alerts : undefined;
            inputs["authorizedKeys"] = args ? args.authorizedKeys : undefined;
            inputs["authorizedUsers"] = args ? args.authorizedUsers : undefined;
            inputs["backupId"] = args ? args.backupId : undefined;
            inputs["backupsEnabled"] = args ? args.backupsEnabled : undefined;
            inputs["bootConfigLabel"] = args ? args.bootConfigLabel : undefined;
            inputs["configs"] = args ? args.configs : undefined;
            inputs["disks"] = args ? args.disks : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["image"] = args ? args.image : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["privateIp"] = args ? args.privateIp : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["rootPass"] = args ? args.rootPass : undefined;
            inputs["stackscriptData"] = args ? args.stackscriptData : undefined;
            inputs["stackscriptId"] = args ? args.stackscriptId : undefined;
            inputs["swapSize"] = args ? args.swapSize : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["watchdogEnabled"] = args ? args.watchdogEnabled : undefined;
            inputs["backups"] = undefined /*out*/;
            inputs["ipAddress"] = undefined /*out*/;
            inputs["ipv4s"] = undefined /*out*/;
            inputs["ipv6"] = undefined /*out*/;
            inputs["privateIpAddress"] = undefined /*out*/;
            inputs["specs"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        super("linode:index/instance:Instance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    readonly alerts?: pulumi.Input<{ cpu?: pulumi.Input<number>, io?: pulumi.Input<number>, networkIn?: pulumi.Input<number>, networkOut?: pulumi.Input<number>, transferQuota?: pulumi.Input<number> }>;
    /**
     * A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is
     * provided.
     */
    readonly authorizedKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root`
     * user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.
     */
    readonly authorizedUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the
     * Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See
     * /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually
     * exclusive.
     */
    readonly backupId?: pulumi.Input<number>;
    /**
     * Information about this Linode's backups status.
     */
    readonly backups?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, schedule?: pulumi.Input<{ day?: pulumi.Input<string>, window?: pulumi.Input<string> }> }>;
    /**
     * If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This
     * will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
     */
    readonly backupsEnabled?: pulumi.Input<boolean>;
    /**
     * The Label of the Instance Config that should be used to boot the Linode instance.
     */
    readonly bootConfigLabel?: pulumi.Input<string>;
    /**
     * Configuration profiles define the VM settings and boot behavior of the Linode Instance.
     */
    readonly configs?: pulumi.Input<pulumi.Input<{ comments?: pulumi.Input<string>, devices?: pulumi.Input<{ sda?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdb?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdc?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdd?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sde?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdf?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdg?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdh?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }> }>, helpers?: pulumi.Input<{ devtmpfsAutomount?: pulumi.Input<boolean>, distro?: pulumi.Input<boolean>, modulesDep?: pulumi.Input<boolean>, network?: pulumi.Input<boolean>, updatedbDisabled?: pulumi.Input<boolean> }>, kernel?: pulumi.Input<string>, label: pulumi.Input<string>, memoryLimit?: pulumi.Input<number>, rootDevice?: pulumi.Input<string>, runLevel?: pulumi.Input<string>, virtMode?: pulumi.Input<string> }>[]>;
    readonly disks?: pulumi.Input<pulumi.Input<{ authorizedKeys?: pulumi.Input<pulumi.Input<string>[]>, authorizedUsers?: pulumi.Input<pulumi.Input<string>[]>, filesystem?: pulumi.Input<string>, id?: pulumi.Input<number>, image?: pulumi.Input<string>, label: pulumi.Input<string>, readOnly?: pulumi.Input<boolean>, rootPass?: pulumi.Input<string>, size: pulumi.Input<number>, stackscriptData?: pulumi.Input<{[key: string]: any}>, stackscriptId?: pulumi.Input<number> }>[]>;
    /**
     * The display group of the Linode instance.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with
     * private/. See /images for more information on the Images available for you to use.
     */
    readonly image?: pulumi.Input<string>;
    /**
     * This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary
     * address will be used for this field.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a
     * single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.
     */
    readonly ipv4s?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.
     */
    readonly ipv6?: pulumi.Input<string>;
    /**
     * The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned
     */
    readonly label?: pulumi.Input<string>;
    /**
     * If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network
     * within the Linode's region.
     */
    readonly privateIp?: pulumi.Input<boolean>;
    /**
     * This Linode's Private IPv4 Address. The regional private IP address range is 192.168.128/17 address shared by all
     * Linode Instances in a region.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
    /**
     * This is the location where the Linode was deployed. This cannot be changed without opening a support ticket.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The password that will be initialially assigned to the 'root' user account.
     */
    readonly rootPass?: pulumi.Input<string>;
    readonly specs?: pulumi.Input<{ disk?: pulumi.Input<number>, memory?: pulumi.Input<number>, transfer?: pulumi.Input<number>, vcpus?: pulumi.Input<number> }>;
    /**
     * An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode.
     * Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
     */
    readonly stackscriptData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an
     * Image that is compatible with this StackScript.
     */
    readonly stackscriptId?: pulumi.Input<number>;
    /**
     * The status of the instance, indicating the current readiness state.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored.
     * This is used to set the swap disk size for the newly-created Linode.
     */
    readonly swapSize?: pulumi.Input<number>;
    /**
     * An array of tags applied to this object. Tags are for organizational purposes only.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of instance to be deployed, determining the price and size.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off
     * unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible.
     * To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
     */
    readonly watchdogEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    readonly alerts?: pulumi.Input<{ cpu?: pulumi.Input<number>, io?: pulumi.Input<number>, networkIn?: pulumi.Input<number>, networkOut?: pulumi.Input<number>, transferQuota?: pulumi.Input<number> }>;
    /**
     * A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if 'image' is
     * provided.
     */
    readonly authorizedKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root`
     * user's `~/.ssh/authorized_keys` file automatically. Only accepted if 'image' is provided.
     */
    readonly authorizedUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the
     * Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See
     * /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually
     * exclusive.
     */
    readonly backupId?: pulumi.Input<number>;
    /**
     * If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This
     * will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
     */
    readonly backupsEnabled?: pulumi.Input<boolean>;
    /**
     * The Label of the Instance Config that should be used to boot the Linode instance.
     */
    readonly bootConfigLabel?: pulumi.Input<string>;
    /**
     * Configuration profiles define the VM settings and boot behavior of the Linode Instance.
     */
    readonly configs?: pulumi.Input<pulumi.Input<{ comments?: pulumi.Input<string>, devices?: pulumi.Input<{ sda?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdb?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdc?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdd?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sde?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdf?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdg?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }>, sdh?: pulumi.Input<{ diskId?: pulumi.Input<number>, diskLabel?: pulumi.Input<string>, volumeId?: pulumi.Input<number> }> }>, helpers?: pulumi.Input<{ devtmpfsAutomount?: pulumi.Input<boolean>, distro?: pulumi.Input<boolean>, modulesDep?: pulumi.Input<boolean>, network?: pulumi.Input<boolean>, updatedbDisabled?: pulumi.Input<boolean> }>, kernel?: pulumi.Input<string>, label: pulumi.Input<string>, memoryLimit?: pulumi.Input<number>, rootDevice?: pulumi.Input<string>, runLevel?: pulumi.Input<string>, virtMode?: pulumi.Input<string> }>[]>;
    readonly disks?: pulumi.Input<pulumi.Input<{ authorizedKeys?: pulumi.Input<pulumi.Input<string>[]>, authorizedUsers?: pulumi.Input<pulumi.Input<string>[]>, filesystem?: pulumi.Input<string>, id?: pulumi.Input<number>, image?: pulumi.Input<string>, label: pulumi.Input<string>, readOnly?: pulumi.Input<boolean>, rootPass?: pulumi.Input<string>, size: pulumi.Input<number>, stackscriptData?: pulumi.Input<{[key: string]: any}>, stackscriptId?: pulumi.Input<number> }>[]>;
    /**
     * The display group of the Linode instance.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with
     * private/. See /images for more information on the Images available for you to use.
     */
    readonly image?: pulumi.Input<string>;
    /**
     * The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned
     */
    readonly label?: pulumi.Input<string>;
    /**
     * If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network
     * within the Linode's region.
     */
    readonly privateIp?: pulumi.Input<boolean>;
    /**
     * This is the location where the Linode was deployed. This cannot be changed without opening a support ticket.
     */
    readonly region: pulumi.Input<string>;
    /**
     * The password that will be initialially assigned to the 'root' user account.
     */
    readonly rootPass?: pulumi.Input<string>;
    /**
     * An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode.
     * Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
     */
    readonly stackscriptData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an
     * Image that is compatible with this StackScript.
     */
    readonly stackscriptId?: pulumi.Input<number>;
    /**
     * When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored.
     * This is used to set the swap disk size for the newly-created Linode.
     */
    readonly swapSize?: pulumi.Input<number>;
    /**
     * An array of tags applied to this object. Tags are for organizational purposes only.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of instance to be deployed, determining the price and size.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off
     * unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible.
     * To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
     */
    readonly watchdogEnabled?: pulumi.Input<boolean>;
}