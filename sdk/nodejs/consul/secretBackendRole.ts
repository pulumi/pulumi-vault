// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Consul secrets role for a Consul secrets engine in Vault. Consul secret backends can then issue Consul tokens.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = new vault.consul.SecretBackend("test", {
 *     path: "consul",
 *     description: "Manages the Consul backend",
 *     address: "127.0.0.1:8500",
 *     token: "4240861b-ce3d-8530-115a-521ff070dd29",
 * });
 * const example = new vault.consul.SecretBackendRole("example", {
 *     name: "test-role",
 *     backend: test.path,
 *     consulPolicies: ["example-policy"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Note About Required Arguments
 *
 * *At least one* of the four arguments `consulPolicies`, `consulRoles`, `serviceIdentities`, or
 * `nodeIdentities` is required for a token. If desired, any combination of the four arguments up-to and
 * including all four, is valid.
 *
 * ## Import
 *
 * Consul secret backend roles can be imported using the `backend`, `/roles/`, and the `name` e.g.
 *
 * ```sh
 * $ pulumi import vault:consul/secretBackendRole:SecretBackendRole example consul/roles/my-role
 * ```
 */
export class SecretBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRoleState, opts?: pulumi.CustomResourceOptions): SecretBackendRole {
        return new SecretBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:consul/secretBackendRole:SecretBackendRole';

    /**
     * Returns true if the given object is an instance of SecretBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRole.__pulumiType;
    }

    /**
     * The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The Consul namespace that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.7+".
     */
    public readonly consulNamespace!: pulumi.Output<string>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> The list of Consul ACL policies to associate with these roles.
     */
    public readonly consulPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul roles to attach to the token.
     * Applicable for Vault 1.10+ with Consul 1.5+.
     */
    public readonly consulRoles!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates that the token should not be replicated globally and instead be local to the current datacenter.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     */
    public readonly maxTtl!: pulumi.Output<number | undefined>;
    /**
     * The name of the Consul secrets engine role to create.
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
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul node
     * identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
     */
    public readonly nodeIdentities!: pulumi.Output<string[] | undefined>;
    /**
     * The admin partition that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.11+".
     */
    public readonly partition!: pulumi.Output<string>;
    /**
     * The list of Consul ACL policies to associate with these roles.
     * **NOTE:** The new parameter `consulPolicies` should be used in favor of this. This parameter,
     * `policies`, remains supported for legacy users, but Vault has deprecated this field.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul
     * service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
     */
    public readonly serviceIdentities!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the TTL for this role.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;

    /**
     * Create a SecretBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRoleArgs | SecretBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["consulNamespace"] = state ? state.consulNamespace : undefined;
            resourceInputs["consulPolicies"] = state ? state.consulPolicies : undefined;
            resourceInputs["consulRoles"] = state ? state.consulRoles : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["nodeIdentities"] = state ? state.nodeIdentities : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["serviceIdentities"] = state ? state.serviceIdentities : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["consulNamespace"] = args ? args.consulNamespace : undefined;
            resourceInputs["consulPolicies"] = args ? args.consulPolicies : undefined;
            resourceInputs["consulRoles"] = args ? args.consulRoles : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["nodeIdentities"] = args ? args.nodeIdentities : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["serviceIdentities"] = args ? args.serviceIdentities : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
     */
    backend?: pulumi.Input<string>;
    /**
     * The Consul namespace that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.7+".
     */
    consulNamespace?: pulumi.Input<string>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> The list of Consul ACL policies to associate with these roles.
     */
    consulPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul roles to attach to the token.
     * Applicable for Vault 1.10+ with Consul 1.5+.
     */
    consulRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates that the token should not be replicated globally and instead be local to the current datacenter.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * The name of the Consul secrets engine role to create.
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
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul node
     * identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
     */
    nodeIdentities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The admin partition that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.11+".
     */
    partition?: pulumi.Input<string>;
    /**
     * The list of Consul ACL policies to associate with these roles.
     * **NOTE:** The new parameter `consulPolicies` should be used in favor of this. This parameter,
     * `policies`, remains supported for legacy users, but Vault has deprecated this field.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul
     * service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
     */
    serviceIdentities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the TTL for this role.
     */
    ttl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
     */
    backend?: pulumi.Input<string>;
    /**
     * The Consul namespace that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.7+".
     */
    consulNamespace?: pulumi.Input<string>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> The list of Consul ACL policies to associate with these roles.
     */
    consulPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul roles to attach to the token.
     * Applicable for Vault 1.10+ with Consul 1.5+.
     */
    consulRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates that the token should not be replicated globally and instead be local to the current datacenter.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum TTL for leases associated with this role, in seconds.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * The name of the Consul secrets engine role to create.
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
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul node
     * identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.
     */
    nodeIdentities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The admin partition that the token will be created in.
     * Applicable for Vault 1.10+ and Consul 1.11+".
     */
    partition?: pulumi.Input<string>;
    /**
     * The list of Consul ACL policies to associate with these roles.
     * **NOTE:** The new parameter `consulPolicies` should be used in favor of this. This parameter,
     * `policies`, remains supported for legacy users, but Vault has deprecated this field.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * <sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul
     * service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.
     */
    serviceIdentities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the TTL for this role.
     */
    ttl?: pulumi.Input<number>;
}
