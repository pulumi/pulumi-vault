// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages KMIP Secret roles in a Vault server. This feature requires
 * Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
 * for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const _default = new vault.kmip.SecretBackend("default", {
 *     path: "kmip",
 *     description: "Vault KMIP backend",
 * });
 * const dev = new vault.kmip.SecretScope("dev", {
 *     path: _default.path,
 *     scope: "dev",
 *     force: true,
 * });
 * const admin = new vault.kmip.SecretRole("admin", {
 *     path: dev.path,
 *     scope: dev.scope,
 *     role: "admin",
 *     tlsClientKeyType: "ec",
 *     tlsClientKeyBits: 256,
 *     operationActivate: true,
 *     operationGet: true,
 *     operationGetAttributes: true,
 *     operationCreate: true,
 *     operationDestroy: true,
 * });
 * ```
 *
 * ## Import
 *
 * KMIP Secret role can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:kmip/secretRole:SecretRole admin kmip
 * ```
 */
export class SecretRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretRoleState, opts?: pulumi.CustomResourceOptions): SecretRole {
        return new SecretRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:kmip/secretRole:SecretRole';

    /**
     * Returns true if the given object is an instance of SecretRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretRole.__pulumiType;
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Grant permission to use the KMIP Activate operation.
     */
    public readonly operationActivate!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Add Attribute operation.
     */
    public readonly operationAddAttribute!: pulumi.Output<boolean>;
    /**
     * Grant all permissions to this role. May not be specified with any other `operation_*` params.
     */
    public readonly operationAll!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Create operation.
     */
    public readonly operationCreate!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Destroy operation.
     */
    public readonly operationDestroy!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Discover Version operation.
     */
    public readonly operationDiscoverVersions!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Get operation.
     */
    public readonly operationGet!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Get Atrribute List operation.
     */
    public readonly operationGetAttributeList!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Get Atrributes operation.
     */
    public readonly operationGetAttributes!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Get Locate operation.
     */
    public readonly operationLocate!: pulumi.Output<boolean>;
    /**
     * Remove all permissions from this role. May not be specified with any other `operation_*` params.
     */
    public readonly operationNone!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Register operation.
     */
    public readonly operationRegister!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Rekey operation.
     */
    public readonly operationRekey!: pulumi.Output<boolean>;
    /**
     * Grant permission to use the KMIP Revoke operation.
     */
    public readonly operationRevoke!: pulumi.Output<boolean>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Name of the role.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Name of the scope.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * Client certificate key bits, valid values depend on key type.
     */
    public readonly tlsClientKeyBits!: pulumi.Output<number | undefined>;
    /**
     * Client certificate key type, `rsa` or `ec`.
     */
    public readonly tlsClientKeyType!: pulumi.Output<string | undefined>;
    /**
     * Client certificate TTL in seconds.
     */
    public readonly tlsClientTtl!: pulumi.Output<number | undefined>;

    /**
     * Create a SecretRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretRoleArgs | SecretRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretRoleState | undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["operationActivate"] = state ? state.operationActivate : undefined;
            resourceInputs["operationAddAttribute"] = state ? state.operationAddAttribute : undefined;
            resourceInputs["operationAll"] = state ? state.operationAll : undefined;
            resourceInputs["operationCreate"] = state ? state.operationCreate : undefined;
            resourceInputs["operationDestroy"] = state ? state.operationDestroy : undefined;
            resourceInputs["operationDiscoverVersions"] = state ? state.operationDiscoverVersions : undefined;
            resourceInputs["operationGet"] = state ? state.operationGet : undefined;
            resourceInputs["operationGetAttributeList"] = state ? state.operationGetAttributeList : undefined;
            resourceInputs["operationGetAttributes"] = state ? state.operationGetAttributes : undefined;
            resourceInputs["operationLocate"] = state ? state.operationLocate : undefined;
            resourceInputs["operationNone"] = state ? state.operationNone : undefined;
            resourceInputs["operationRegister"] = state ? state.operationRegister : undefined;
            resourceInputs["operationRekey"] = state ? state.operationRekey : undefined;
            resourceInputs["operationRevoke"] = state ? state.operationRevoke : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["tlsClientKeyBits"] = state ? state.tlsClientKeyBits : undefined;
            resourceInputs["tlsClientKeyType"] = state ? state.tlsClientKeyType : undefined;
            resourceInputs["tlsClientTtl"] = state ? state.tlsClientTtl : undefined;
        } else {
            const args = argsOrState as SecretRoleArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["operationActivate"] = args ? args.operationActivate : undefined;
            resourceInputs["operationAddAttribute"] = args ? args.operationAddAttribute : undefined;
            resourceInputs["operationAll"] = args ? args.operationAll : undefined;
            resourceInputs["operationCreate"] = args ? args.operationCreate : undefined;
            resourceInputs["operationDestroy"] = args ? args.operationDestroy : undefined;
            resourceInputs["operationDiscoverVersions"] = args ? args.operationDiscoverVersions : undefined;
            resourceInputs["operationGet"] = args ? args.operationGet : undefined;
            resourceInputs["operationGetAttributeList"] = args ? args.operationGetAttributeList : undefined;
            resourceInputs["operationGetAttributes"] = args ? args.operationGetAttributes : undefined;
            resourceInputs["operationLocate"] = args ? args.operationLocate : undefined;
            resourceInputs["operationNone"] = args ? args.operationNone : undefined;
            resourceInputs["operationRegister"] = args ? args.operationRegister : undefined;
            resourceInputs["operationRekey"] = args ? args.operationRekey : undefined;
            resourceInputs["operationRevoke"] = args ? args.operationRevoke : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["tlsClientKeyBits"] = args ? args.tlsClientKeyBits : undefined;
            resourceInputs["tlsClientKeyType"] = args ? args.tlsClientKeyType : undefined;
            resourceInputs["tlsClientTtl"] = args ? args.tlsClientTtl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretRole resources.
 */
export interface SecretRoleState {
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Grant permission to use the KMIP Activate operation.
     */
    operationActivate?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Add Attribute operation.
     */
    operationAddAttribute?: pulumi.Input<boolean>;
    /**
     * Grant all permissions to this role. May not be specified with any other `operation_*` params.
     */
    operationAll?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Create operation.
     */
    operationCreate?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Destroy operation.
     */
    operationDestroy?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Discover Version operation.
     */
    operationDiscoverVersions?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get operation.
     */
    operationGet?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get Atrribute List operation.
     */
    operationGetAttributeList?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get Atrributes operation.
     */
    operationGetAttributes?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get Locate operation.
     */
    operationLocate?: pulumi.Input<boolean>;
    /**
     * Remove all permissions from this role. May not be specified with any other `operation_*` params.
     */
    operationNone?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Register operation.
     */
    operationRegister?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Rekey operation.
     */
    operationRekey?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Revoke operation.
     */
    operationRevoke?: pulumi.Input<boolean>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     */
    path?: pulumi.Input<string>;
    /**
     * Name of the role.
     */
    role?: pulumi.Input<string>;
    /**
     * Name of the scope.
     */
    scope?: pulumi.Input<string>;
    /**
     * Client certificate key bits, valid values depend on key type.
     */
    tlsClientKeyBits?: pulumi.Input<number>;
    /**
     * Client certificate key type, `rsa` or `ec`.
     */
    tlsClientKeyType?: pulumi.Input<string>;
    /**
     * Client certificate TTL in seconds.
     */
    tlsClientTtl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SecretRole resource.
 */
export interface SecretRoleArgs {
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Grant permission to use the KMIP Activate operation.
     */
    operationActivate?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Add Attribute operation.
     */
    operationAddAttribute?: pulumi.Input<boolean>;
    /**
     * Grant all permissions to this role. May not be specified with any other `operation_*` params.
     */
    operationAll?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Create operation.
     */
    operationCreate?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Destroy operation.
     */
    operationDestroy?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Discover Version operation.
     */
    operationDiscoverVersions?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get operation.
     */
    operationGet?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get Atrribute List operation.
     */
    operationGetAttributeList?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get Atrributes operation.
     */
    operationGetAttributes?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Get Locate operation.
     */
    operationLocate?: pulumi.Input<boolean>;
    /**
     * Remove all permissions from this role. May not be specified with any other `operation_*` params.
     */
    operationNone?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Register operation.
     */
    operationRegister?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Rekey operation.
     */
    operationRekey?: pulumi.Input<boolean>;
    /**
     * Grant permission to use the KMIP Revoke operation.
     */
    operationRevoke?: pulumi.Input<boolean>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     */
    path: pulumi.Input<string>;
    /**
     * Name of the role.
     */
    role: pulumi.Input<string>;
    /**
     * Name of the scope.
     */
    scope: pulumi.Input<string>;
    /**
     * Client certificate key bits, valid values depend on key type.
     */
    tlsClientKeyBits?: pulumi.Input<number>;
    /**
     * Client certificate key type, `rsa` or `ec`.
     */
    tlsClientKeyType?: pulumi.Input<string>;
    /**
     * Client certificate TTL in seconds.
     */
    tlsClientTtl?: pulumi.Input<number>;
}
