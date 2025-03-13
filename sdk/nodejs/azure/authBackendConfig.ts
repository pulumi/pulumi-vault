// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * You can setup the Azure auth engine with Workload Identity Federation (WIF) for a secret-less configuration:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.AuthBackend("example", {
 *     type: "azure",
 *     identityTokenKey: "example-key",
 * });
 * const exampleAuthBackendConfig = new vault.azure.AuthBackendConfig("example", {
 *     backend: example.path,
 *     tenantId: "11111111-2222-3333-4444-555555555555",
 *     clientId: "11111111-2222-3333-4444-555555555555",
 *     identityTokenAudience: "<TOKEN_AUDIENCE>",
 *     identityTokenTtl: "<TOKEN_TTL>",
 *     rotationSchedule: "0 * * * SAT",
 *     rotationWindow: 3600,
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.AuthBackend("example", {type: "azure"});
 * const exampleAuthBackendConfig = new vault.azure.AuthBackendConfig("example", {
 *     backend: example.path,
 *     tenantId: "11111111-2222-3333-4444-555555555555",
 *     clientId: "11111111-2222-3333-4444-555555555555",
 *     clientSecret: "01234567890123456789",
 *     resource: "https://vault.hashicorp.com",
 *     rotationSchedule: "0 * * * SAT",
 *     rotationWindow: 3600,
 * });
 * ```
 *
 * ## Import
 *
 * Azure auth backends can be imported using `auth/`, the `backend` path, and `/config` e.g.
 *
 * ```sh
 * $ pulumi import vault:azure/authBackendConfig:AuthBackendConfig example auth/azure/config
 * ```
 */
export class AuthBackendConfig extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendConfigState, opts?: pulumi.CustomResourceOptions): AuthBackendConfig {
        return new AuthBackendConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:azure/authBackendConfig:AuthBackendConfig';

    /**
     * Returns true if the given object is an instance of AuthBackendConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendConfig.__pulumiType;
    }

    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client secret for credentials to query the
     * Azure APIs.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    public readonly disableAutomatedRotation!: pulumi.Output<boolean | undefined>;
    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * The audience claim value for plugin identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     */
    public readonly identityTokenAudience!: pulumi.Output<string | undefined>;
    /**
     * The TTL of generated identity tokens in seconds.
     */
    public readonly identityTokenTtl!: pulumi.Output<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The configured URL for the application registered in
     * Azure Active Directory.
     */
    public readonly resource!: pulumi.Output<string>;
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    public readonly rotationPeriod!: pulumi.Output<number | undefined>;
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    public readonly rotationSchedule!: pulumi.Output<string | undefined>;
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    public readonly rotationWindow!: pulumi.Output<number | undefined>;
    /**
     * The tenant id for the Azure Active Directory
     * organization.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a AuthBackendConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendConfigArgs | AuthBackendConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendConfigState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["disableAutomatedRotation"] = state ? state.disableAutomatedRotation : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["identityTokenAudience"] = state ? state.identityTokenAudience : undefined;
            resourceInputs["identityTokenTtl"] = state ? state.identityTokenTtl : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["resource"] = state ? state.resource : undefined;
            resourceInputs["rotationPeriod"] = state ? state.rotationPeriod : undefined;
            resourceInputs["rotationSchedule"] = state ? state.rotationSchedule : undefined;
            resourceInputs["rotationWindow"] = state ? state.rotationWindow : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as AuthBackendConfigArgs | undefined;
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["clientId"] = args?.clientId ? pulumi.secret(args.clientId) : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["disableAutomatedRotation"] = args ? args.disableAutomatedRotation : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["identityTokenAudience"] = args ? args.identityTokenAudience : undefined;
            resourceInputs["identityTokenTtl"] = args ? args.identityTokenTtl : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["resource"] = args ? args.resource : undefined;
            resourceInputs["rotationPeriod"] = args ? args.rotationPeriod : undefined;
            resourceInputs["rotationSchedule"] = args ? args.rotationSchedule : undefined;
            resourceInputs["rotationWindow"] = args ? args.rotationWindow : undefined;
            resourceInputs["tenantId"] = args?.tenantId ? pulumi.secret(args.tenantId) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientId", "clientSecret", "tenantId"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AuthBackendConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendConfig resources.
 */
export interface AuthBackendConfigState {
    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     */
    backend?: pulumi.Input<string>;
    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client secret for credentials to query the
     * Azure APIs.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    disableAutomatedRotation?: pulumi.Input<boolean>;
    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     */
    environment?: pulumi.Input<string>;
    /**
     * The audience claim value for plugin identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     */
    identityTokenAudience?: pulumi.Input<string>;
    /**
     * The TTL of generated identity tokens in seconds.
     */
    identityTokenTtl?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The configured URL for the application registered in
     * Azure Active Directory.
     */
    resource?: pulumi.Input<string>;
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    rotationPeriod?: pulumi.Input<number>;
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    rotationSchedule?: pulumi.Input<string>;
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    rotationWindow?: pulumi.Input<number>;
    /**
     * The tenant id for the Azure Active Directory
     * organization.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendConfig resource.
 */
export interface AuthBackendConfigArgs {
    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     */
    backend?: pulumi.Input<string>;
    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client secret for credentials to query the
     * Azure APIs.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    disableAutomatedRotation?: pulumi.Input<boolean>;
    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     */
    environment?: pulumi.Input<string>;
    /**
     * The audience claim value for plugin identity tokens. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     */
    identityTokenAudience?: pulumi.Input<string>;
    /**
     * The TTL of generated identity tokens in seconds.
     */
    identityTokenTtl?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The configured URL for the application registered in
     * Azure Active Directory.
     */
    resource: pulumi.Input<string>;
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    rotationPeriod?: pulumi.Input<number>;
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    rotationSchedule?: pulumi.Input<string>;
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * *Available only for Vault Enterprise*
     */
    rotationWindow?: pulumi.Input<number>;
    /**
     * The tenant id for the Azure Active Directory
     * organization.
     */
    tenantId: pulumi.Input<string>;
}
