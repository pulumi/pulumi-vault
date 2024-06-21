// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing an
 * [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.okta.AuthBackend("example", {
 *     description: "Demonstration of the Terraform Okta auth backend",
 *     organization: "example",
 *     token: "something that should be kept secret",
 *     groups: [{
 *         groupName: "foo",
 *         policies: [
 *             "one",
 *             "two",
 *         ],
 *     }],
 *     users: [{
 *         username: "bar",
 *         groups: ["foo"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Okta authentication backends can be imported using its `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:okta/authBackend:AuthBackend example okta
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
    public static readonly __pulumiType = 'vault:okta/authBackend:AuthBackend';

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
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * The Okta url. Examples: oktapreview.com, okta.com
     */
    public readonly baseUrl!: pulumi.Output<string | undefined>;
    /**
     * When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
     */
    public readonly bypassOktaMfa!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the auth backend
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * Associate Okta groups with policies within Vault.
     * See below for more details.
     */
    public readonly groups!: pulumi.Output<outputs.okta.AuthBackendGroup[]>;
    /**
     * Maximum duration after which authentication will be expired
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated Deprecated. Please use `tokenMaxTtl` instead.
     */
    public readonly maxTtl!: pulumi.Output<string | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The Okta organization. This will be the first part of the url `https://XXX.okta.com`
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * Path to mount the Okta auth backend. Default to path `okta`.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The Okta API token. This is required to query Okta for user group membership.
     * If this is not supplied only locally configured groups will be enabled.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime of the generated token
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Period
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Policies
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token to generate, service or batch
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;
    /**
     * Duration after which authentication will be expired.
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated Deprecated. Please use `tokenTtl` instead.
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
    /**
     * Associate Okta users with groups or policies within Vault.
     * See below for more details.
     */
    public readonly users!: pulumi.Output<outputs.okta.AuthBackendUser[]>;

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
            resourceInputs["baseUrl"] = state ? state.baseUrl : undefined;
            resourceInputs["bypassOktaMfa"] = state ? state.bypassOktaMfa : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            resourceInputs["baseUrl"] = args ? args.baseUrl : undefined;
            resourceInputs["bypassOktaMfa"] = args ? args.bypassOktaMfa : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            resourceInputs["tokenType"] = args ? args.tokenType : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["accessor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     */
    accessor?: pulumi.Input<string>;
    /**
     * The Okta url. Examples: oktapreview.com, okta.com
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
     */
    bypassOktaMfa?: pulumi.Input<boolean>;
    /**
     * The description of the auth backend
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Associate Okta groups with policies within Vault.
     * See below for more details.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.okta.AuthBackendGroup>[]>;
    /**
     * Maximum duration after which authentication will be expired
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated Deprecated. Please use `tokenMaxTtl` instead.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The Okta organization. This will be the first part of the url `https://XXX.okta.com`
     */
    organization?: pulumi.Input<string>;
    /**
     * Path to mount the Okta auth backend. Default to path `okta`.
     */
    path?: pulumi.Input<string>;
    /**
     * The Okta API token. This is required to query Okta for user group membership.
     * If this is not supplied only locally configured groups will be enabled.
     */
    token?: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
    /**
     * Duration after which authentication will be expired.
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated Deprecated. Please use `tokenTtl` instead.
     */
    ttl?: pulumi.Input<string>;
    /**
     * Associate Okta users with groups or policies within Vault.
     * See below for more details.
     */
    users?: pulumi.Input<pulumi.Input<inputs.okta.AuthBackendUser>[]>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * The Okta url. Examples: oktapreview.com, okta.com
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
     */
    bypassOktaMfa?: pulumi.Input<boolean>;
    /**
     * The description of the auth backend
     */
    description?: pulumi.Input<string>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Associate Okta groups with policies within Vault.
     * See below for more details.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.okta.AuthBackendGroup>[]>;
    /**
     * Maximum duration after which authentication will be expired
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated Deprecated. Please use `tokenMaxTtl` instead.
     */
    maxTtl?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The Okta organization. This will be the first part of the url `https://XXX.okta.com`
     */
    organization: pulumi.Input<string>;
    /**
     * Path to mount the Okta auth backend. Default to path `okta`.
     */
    path?: pulumi.Input<string>;
    /**
     * The Okta API token. This is required to query Okta for user group membership.
     * If this is not supplied only locally configured groups will be enabled.
     */
    token?: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    tokenType?: pulumi.Input<string>;
    /**
     * Duration after which authentication will be expired.
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     *
     * @deprecated Deprecated. Please use `tokenTtl` instead.
     */
    ttl?: pulumi.Input<string>;
    /**
     * Associate Okta users with groups or policies within Vault.
     * See below for more details.
     */
    users?: pulumi.Input<pulumi.Input<inputs.okta.AuthBackendUser>[]>;
}
