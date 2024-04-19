// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * You need to create a role with a named key.
 * At creation time, the key can be created independently of the role. However, the key must
 * exist before the role can be used to issue tokens. You must also configure the key with the
 * role's Client ID to allow the role to use the key.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new pulumi.Config();
 * const key = config.get("key") || "key";
 * const role = new vault.identity.OidcRole("role", {key: key});
 * const keyOidcKey = new vault.identity.OidcKey("keyOidcKey", {
 *     algorithm: "RS256",
 *     allowedClientIds: [role.clientId],
 * });
 * ```
 *
 * If you want to create the key first before creating the role, you can use a separate
 * resource to configure the allowed Client ID on
 * the key.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const key = new vault.identity.OidcKey("key", {algorithm: "RS256"});
 * const roleOidcRole = new vault.identity.OidcRole("roleOidcRole", {key: key.name});
 * const roleOidcKeyAllowedClientID = new vault.identity.OidcKeyAllowedClientID("roleOidcKeyAllowedClientID", {
 *     keyName: key.name,
 *     allowedClientId: roleOidcRole.clientId,
 * });
 * ```
 *
 * ## Import
 *
 * The key can be imported with the role name, for example:
 *
 * ```sh
 * $ pulumi import vault:identity/oidcRole:OidcRole role role
 * ```
 */
export class OidcRole extends pulumi.CustomResource {
    /**
     * Get an existing OidcRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OidcRoleState, opts?: pulumi.CustomResourceOptions): OidcRole {
        return new OidcRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/oidcRole:OidcRole';

    /**
     * Returns true if the given object is an instance of OidcRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OidcRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OidcRole.__pulumiType;
    }

    /**
     * The value that will be included in the `aud` field of all the OIDC identity
     * tokens issued by this role
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * A configured named key, the key must already exist
     * before tokens can be issued.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Name of the OIDC Role to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The template string to use for generating tokens. This may be in
     * string-ified JSON or base64 format. See the
     * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
     * for the template format.
     */
    public readonly template!: pulumi.Output<string | undefined>;
    /**
     * TTL of the tokens generated against the role in number of seconds.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;

    /**
     * Create a OidcRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OidcRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OidcRoleArgs | OidcRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OidcRoleState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as OidcRoleArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OidcRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OidcRole resources.
 */
export interface OidcRoleState {
    /**
     * The value that will be included in the `aud` field of all the OIDC identity
     * tokens issued by this role
     */
    clientId?: pulumi.Input<string>;
    /**
     * A configured named key, the key must already exist
     * before tokens can be issued.
     */
    key?: pulumi.Input<string>;
    /**
     * Name of the OIDC Role to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The template string to use for generating tokens. This may be in
     * string-ified JSON or base64 format. See the
     * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
     * for the template format.
     */
    template?: pulumi.Input<string>;
    /**
     * TTL of the tokens generated against the role in number of seconds.
     */
    ttl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OidcRole resource.
 */
export interface OidcRoleArgs {
    /**
     * The value that will be included in the `aud` field of all the OIDC identity
     * tokens issued by this role
     */
    clientId?: pulumi.Input<string>;
    /**
     * A configured named key, the key must already exist
     * before tokens can be issued.
     */
    key: pulumi.Input<string>;
    /**
     * Name of the OIDC Role to create.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The template string to use for generating tokens. This may be in
     * string-ified JSON or base64 format. See the
     * [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
     * for the template format.
     */
    template?: pulumi.Input<string>;
    /**
     * TTL of the tokens generated against the role in number of seconds.
     */
    ttl?: pulumi.Input<number>;
}
