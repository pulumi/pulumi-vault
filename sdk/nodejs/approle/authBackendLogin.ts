// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Logs into Vault using the AppRole auth backend. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/approle) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const approle = new vault.AuthBackend("approle", {
 *     type: "approle",
 * });
 * const example = new vault.appRole.AuthBackendRole("example", {
 *     backend: approle.path,
 *     policies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 *     roleName: "test-role",
 * });
 * const id = new vault.appRole.AuthBackendRoleSecretID("id", {
 *     backend: approle.path,
 *     roleName: example.roleName,
 * });
 * const login = new vault.appRole.AuthBackendLogin("login", {
 *     backend: approle.path,
 *     roleId: example.roleId,
 *     secretId: id.secretId,
 * });
 * ```
 */
export class AuthBackendLogin extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendLogin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendLoginState, opts?: pulumi.CustomResourceOptions): AuthBackendLogin {
        return new AuthBackendLogin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:appRole/authBackendLogin:AuthBackendLogin';

    /**
     * Returns true if the given object is an instance of AuthBackendLogin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendLogin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendLogin.__pulumiType;
    }

    /**
     * The accessor for the token.
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * The unique path of the Vault backend to log in with.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The Vault token created.
     */
    public /*out*/ readonly clientToken!: pulumi.Output<string>;
    /**
     * How long the token is valid for, in seconds.
     */
    public /*out*/ readonly leaseDuration!: pulumi.Output<number>;
    /**
     * The date and time the lease started, in RFC 3339 format.
     */
    public /*out*/ readonly leaseStarted!: pulumi.Output<string>;
    /**
     * The metadata associated with the token.
     */
    public /*out*/ readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * A list of policies applied to the token.
     */
    public /*out*/ readonly policies!: pulumi.Output<string[]>;
    /**
     * Whether the token is renewable or not.
     */
    public /*out*/ readonly renewable!: pulumi.Output<boolean>;
    /**
     * The ID of the role to log in with.
     */
    public readonly roleId!: pulumi.Output<string>;
    /**
     * The secret ID of the role to log in with. Required
     * unless `bindSecretId` is set to false on the role.
     */
    public readonly secretId!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendLogin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendLoginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendLoginArgs | AuthBackendLoginState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendLoginState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["clientToken"] = state ? state.clientToken : undefined;
            inputs["leaseDuration"] = state ? state.leaseDuration : undefined;
            inputs["leaseStarted"] = state ? state.leaseStarted : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["renewable"] = state ? state.renewable : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
            inputs["secretId"] = state ? state.secretId : undefined;
        } else {
            const args = argsOrState as AuthBackendLoginArgs | undefined;
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["roleId"] = args ? args.roleId : undefined;
            inputs["secretId"] = args ? args.secretId : undefined;
            inputs["accessor"] = undefined /*out*/;
            inputs["clientToken"] = undefined /*out*/;
            inputs["leaseDuration"] = undefined /*out*/;
            inputs["leaseStarted"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["policies"] = undefined /*out*/;
            inputs["renewable"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackendLogin.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendLogin resources.
 */
export interface AuthBackendLoginState {
    /**
     * The accessor for the token.
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * The unique path of the Vault backend to log in with.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The Vault token created.
     */
    readonly clientToken?: pulumi.Input<string>;
    /**
     * How long the token is valid for, in seconds.
     */
    readonly leaseDuration?: pulumi.Input<number>;
    /**
     * The date and time the lease started, in RFC 3339 format.
     */
    readonly leaseStarted?: pulumi.Input<string>;
    /**
     * The metadata associated with the token.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of policies applied to the token.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the token is renewable or not.
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * The ID of the role to log in with.
     */
    readonly roleId?: pulumi.Input<string>;
    /**
     * The secret ID of the role to log in with. Required
     * unless `bindSecretId` is set to false on the role.
     */
    readonly secretId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendLogin resource.
 */
export interface AuthBackendLoginArgs {
    /**
     * The unique path of the Vault backend to log in with.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The ID of the role to log in with.
     */
    readonly roleId: pulumi.Input<string>;
    /**
     * The secret ID of the role to log in with. Required
     * unless `bindSecretId` is set to false on the role.
     */
    readonly secretId?: pulumi.Input<string>;
}
