// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Reads the Role of an Kubernetes from a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
 * information.
 */
export function getAuthBackendRole(args: GetAuthBackendRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthBackendRoleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:kubernetes/getAuthBackendRole:getAuthBackendRole", {
        "audience": args.audience,
        "backend": args.backend,
        "namespace": args.namespace,
        "roleName": args.roleName,
        "tokenBoundCidrs": args.tokenBoundCidrs,
        "tokenExplicitMaxTtl": args.tokenExplicitMaxTtl,
        "tokenMaxTtl": args.tokenMaxTtl,
        "tokenNoDefaultPolicy": args.tokenNoDefaultPolicy,
        "tokenNumUses": args.tokenNumUses,
        "tokenPeriod": args.tokenPeriod,
        "tokenPolicies": args.tokenPolicies,
        "tokenTtl": args.tokenTtl,
        "tokenType": args.tokenType,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthBackendRole.
 */
export interface GetAuthBackendRoleArgs {
    /**
     * Audience claim to verify in the JWT.
     */
    audience?: string;
    /**
     * The unique name for the Kubernetes backend the role to
     * retrieve Role attributes for resides in. Defaults to "kubernetes".
     */
    backend?: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
    /**
     * The name of the role to retrieve the Role attributes for.
     */
    roleName: string;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: string[];
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: number;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: number;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: boolean;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    tokenNumUses?: number;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: number;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    tokenPolicies?: string[];
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: number;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    tokenType?: string;
}

/**
 * A collection of values returned by getAuthBackendRole.
 */
export interface GetAuthBackendRoleResult {
    /**
     * Method used for generating identity aliases. (vault-1.9+)
     */
    readonly aliasNameSource: string;
    /**
     * Audience claim to verify in the JWT.
     */
    readonly audience?: string;
    readonly backend?: string;
    /**
     * List of service account names able to access this role. If set to "*" all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
     */
    readonly boundServiceAccountNames: string[];
    /**
     * List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
     */
    readonly boundServiceAccountNamespaces: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namespace?: string;
    readonly roleName: string;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: string[];
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: number;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: number;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: boolean;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: number;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: number;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: string[];
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: number;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: string;
}
/**
 * Reads the Role of an Kubernetes from a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
 * information.
 */
export function getAuthBackendRoleOutput(args: GetAuthBackendRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAuthBackendRoleResult> {
    return pulumi.output(args).apply((a: any) => getAuthBackendRole(a, opts))
}

/**
 * A collection of arguments for invoking getAuthBackendRole.
 */
export interface GetAuthBackendRoleOutputArgs {
    /**
     * Audience claim to verify in the JWT.
     */
    audience?: pulumi.Input<string>;
    /**
     * The unique name for the Kubernetes backend the role to
     * retrieve Role attributes for resides in. Defaults to "kubernetes".
     */
    backend?: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured namespace.
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The name of the role to retrieve the Role attributes for.
     */
    roleName: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    tokenType?: pulumi.Input<string>;
}
