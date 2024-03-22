// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Auth methods can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:index/authBackend:AuthBackend example github
 * ```
 */
export class AuthBackend extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendState, opts?: pulumi.CustomResourceOptions): AuthBackend {
        return new AuthBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/authBackend:AuthBackend';

    /**
     * Returns true if the given object is an instance of AuthBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackend.__pulumiType;
    }

    /**
     * The accessor for this auth method
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * A description of the auth method.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the auth method is local only.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The path to mount the auth method — this defaults to the name of the type.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Extra configuration block. Structure is documented below.
     *
     * The `tune` block is used to tune the auth backend:
     */
    public readonly tune!: pulumi.Output<outputs.AuthBackendTune>;
    /**
     * The name of the auth method type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a AuthBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendArgs | AuthBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            resourceInputs["accessor"] = state ? state.accessor : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["tune"] = state ? state.tune : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["tune"] = args ? args.tune : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The accessor for this auth method
     */
    accessor?: pulumi.Input<string>;
    /**
     * A description of the auth method.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Specifies if the auth method is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The path to mount the auth method — this defaults to the name of the type.
     */
    path?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     *
     * The `tune` block is used to tune the auth backend:
     */
    tune?: pulumi.Input<inputs.AuthBackendTune>;
    /**
     * The name of the auth method type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * A description of the auth method.
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Specifies if the auth method is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The path to mount the auth method — this defaults to the name of the type.
     */
    path?: pulumi.Input<string>;
    /**
     * Extra configuration block. Structure is documented below.
     *
     * The `tune` block is used to tune the auth backend:
     */
    tune?: pulumi.Input<inputs.AuthBackendTune>;
    /**
     * The name of the auth method type.
     */
    type: pulumi.Input<string>;
}
