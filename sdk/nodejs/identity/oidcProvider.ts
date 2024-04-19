// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages OIDC Providers in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-an-assignment)
 * for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const testOidcKey = new vault.identity.OidcKey("testOidcKey", {
 *     allowedClientIds: ["*"],
 *     rotationPeriod: 3600,
 *     verificationTtl: 3600,
 * });
 * const testOidcAssignment = new vault.identity.OidcAssignment("testOidcAssignment", {
 *     entityIds: ["fake-ascbascas-2231a-sdfaa"],
 *     groupIds: ["fake-sajkdsad-32414-sfsada"],
 * });
 * const testOidcClient = new vault.identity.OidcClient("testOidcClient", {
 *     key: testOidcKey.name,
 *     redirectUris: [
 *         "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
 *         "http://127.0.0.1:8251/callback",
 *         "http://127.0.0.1:8080/callback",
 *     ],
 *     assignments: [testOidcAssignment.name],
 *     idTokenTtl: 2400,
 *     accessTokenTtl: 7200,
 * });
 * const testOidcScope = new vault.identity.OidcScope("testOidcScope", {
 *     template: JSON.stringify({
 *         groups: "{{identity.entity.groups.names}}",
 *     }),
 *     description: "Groups scope.",
 * });
 * const testOidcProvider = new vault.identity.OidcProvider("testOidcProvider", {
 *     httpsEnabled: false,
 *     issuerHost: "127.0.0.1:8200",
 *     allowedClientIds: [testOidcClient.clientId],
 *     scopesSupporteds: [testOidcScope.name],
 * });
 * ```
 *
 * ## Import
 *
 * OIDC Providers can be imported using the `name`, e.g.
 *
 * ```sh
 * $ pulumi import vault:identity/oidcProvider:OidcProvider test my-provider
 * ```
 */
export class OidcProvider extends pulumi.CustomResource {
    /**
     * Get an existing OidcProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OidcProviderState, opts?: pulumi.CustomResourceOptions): OidcProvider {
        return new OidcProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/oidcProvider:OidcProvider';

    /**
     * Returns true if the given object is an instance of OidcProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OidcProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OidcProvider.__pulumiType;
    }

    /**
     * The client IDs that are permitted to use the provider. 
     * If empty, no clients are allowed. If `*`, all clients are allowed.
     */
    public readonly allowedClientIds!: pulumi.Output<string[] | undefined>;
    /**
     * Set to true if the issuer endpoint uses HTTPS.
     */
    public readonly httpsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies what will be used as the `scheme://host:port`
     * component for the `iss` claim of ID tokens. This value is computed using the
     * `issuerHost` and `httpsEnabled` fields.
     */
    public /*out*/ readonly issuer!: pulumi.Output<string>;
    /**
     * The host for the issuer. Can be either host or host:port.
     */
    public readonly issuerHost!: pulumi.Output<string | undefined>;
    /**
     * The name of the provider.
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
     * The scopes available for requesting on the provider.
     */
    public readonly scopesSupporteds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a OidcProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OidcProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OidcProviderArgs | OidcProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OidcProviderState | undefined;
            resourceInputs["allowedClientIds"] = state ? state.allowedClientIds : undefined;
            resourceInputs["httpsEnabled"] = state ? state.httpsEnabled : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["issuerHost"] = state ? state.issuerHost : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["scopesSupporteds"] = state ? state.scopesSupporteds : undefined;
        } else {
            const args = argsOrState as OidcProviderArgs | undefined;
            resourceInputs["allowedClientIds"] = args ? args.allowedClientIds : undefined;
            resourceInputs["httpsEnabled"] = args ? args.httpsEnabled : undefined;
            resourceInputs["issuerHost"] = args ? args.issuerHost : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["scopesSupporteds"] = args ? args.scopesSupporteds : undefined;
            resourceInputs["issuer"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OidcProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OidcProvider resources.
 */
export interface OidcProviderState {
    /**
     * The client IDs that are permitted to use the provider. 
     * If empty, no clients are allowed. If `*`, all clients are allowed.
     */
    allowedClientIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set to true if the issuer endpoint uses HTTPS.
     */
    httpsEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies what will be used as the `scheme://host:port`
     * component for the `iss` claim of ID tokens. This value is computed using the
     * `issuerHost` and `httpsEnabled` fields.
     */
    issuer?: pulumi.Input<string>;
    /**
     * The host for the issuer. Can be either host or host:port.
     */
    issuerHost?: pulumi.Input<string>;
    /**
     * The name of the provider.
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
     * The scopes available for requesting on the provider.
     */
    scopesSupporteds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OidcProvider resource.
 */
export interface OidcProviderArgs {
    /**
     * The client IDs that are permitted to use the provider. 
     * If empty, no clients are allowed. If `*`, all clients are allowed.
     */
    allowedClientIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set to true if the issuer endpoint uses HTTPS.
     */
    httpsEnabled?: pulumi.Input<boolean>;
    /**
     * The host for the issuer. Can be either host or host:port.
     */
    issuerHost?: pulumi.Input<string>;
    /**
     * The name of the provider.
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
     * The scopes available for requesting on the provider.
     */
    scopesSupporteds?: pulumi.Input<pulumi.Input<string>[]>;
}
