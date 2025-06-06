// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * AWS secret backends can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:aws/secretBackend:SecretBackend aws aws
 * ```
 */
export class SecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendState, opts?: pulumi.CustomResourceOptions): SecretBackend {
        return new SecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/secretBackend:SecretBackend';

    /**
     * Returns true if the given object is an instance of SecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackend.__pulumiType;
    }

    /**
     * The AWS Access Key ID this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * The default TTL for credentials
     * issued by this backend.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * A human-friendly description for this backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     */
    public readonly disableAutomatedRotation!: pulumi.Output<boolean | undefined>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies a custom HTTP IAM endpoint to use.
     */
    public readonly iamEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The audience claim value. Requires Vault 1.16+.
     */
    public readonly identityTokenAudience!: pulumi.Output<string | undefined>;
    /**
     * The key to use for signing identity tokens. Requires Vault 1.16+.
     */
    public readonly identityTokenKey!: pulumi.Output<string | undefined>;
    /**
     * The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
     */
    public readonly identityTokenTtl!: pulumi.Output<number>;
    /**
     * Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The AWS region to make API calls against. Defaults to us-east-1.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential. 
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     */
    public readonly rotationPeriod!: pulumi.Output<number | undefined>;
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     */
    public readonly rotationSchedule!: pulumi.Output<string | undefined>;
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     */
    public readonly rotationWindow!: pulumi.Output<number | undefined>;
    /**
     * The AWS Secret Access Key to use when generating new credentials.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * Specifies a custom HTTP STS endpoint to use.
     */
    public readonly stsEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Ordered list of `stsEndpoint`s to try if the defined one fails. Requires Vault 1.19+
     */
    public readonly stsFallbackEndpoints!: pulumi.Output<string[] | undefined>;
    /**
     * Ordered list of `stsRegion`s matching the fallback endpoints. Should correspond in order with those endpoints. Requires Vault 1.19+
     */
    public readonly stsFallbackRegions!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the region of the STS endpoint. Should be included if `stsEndpoint` is supplied. Requires Vault 1.19+
     */
    public readonly stsRegion!: pulumi.Output<string | undefined>;
    /**
     * Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
     *
     * ```
     * {{ if (eq .Type "STS") }}
     * {{ printf "vault-%s-%s" (unix_time) (random 20) | truncate 32 }}
     * {{ else }}
     * {{ printf "vault-%s-%s-%s" (printf "%s-%s" (.DisplayName) (.PolicyName) | truncate 42) (unix_time) (random 20) | truncate 64 }}
     * {{ end }}
     *
     * ```
     */
    public readonly usernameTemplate!: pulumi.Output<string>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableAutomatedRotation"] = state ? state.disableAutomatedRotation : undefined;
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["iamEndpoint"] = state ? state.iamEndpoint : undefined;
            resourceInputs["identityTokenAudience"] = state ? state.identityTokenAudience : undefined;
            resourceInputs["identityTokenKey"] = state ? state.identityTokenKey : undefined;
            resourceInputs["identityTokenTtl"] = state ? state.identityTokenTtl : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["rotationPeriod"] = state ? state.rotationPeriod : undefined;
            resourceInputs["rotationSchedule"] = state ? state.rotationSchedule : undefined;
            resourceInputs["rotationWindow"] = state ? state.rotationWindow : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["stsEndpoint"] = state ? state.stsEndpoint : undefined;
            resourceInputs["stsFallbackEndpoints"] = state ? state.stsFallbackEndpoints : undefined;
            resourceInputs["stsFallbackRegions"] = state ? state.stsFallbackRegions : undefined;
            resourceInputs["stsRegion"] = state ? state.stsRegion : undefined;
            resourceInputs["usernameTemplate"] = state ? state.usernameTemplate : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            resourceInputs["accessKey"] = args?.accessKey ? pulumi.secret(args.accessKey) : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableAutomatedRotation"] = args ? args.disableAutomatedRotation : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["iamEndpoint"] = args ? args.iamEndpoint : undefined;
            resourceInputs["identityTokenAudience"] = args ? args.identityTokenAudience : undefined;
            resourceInputs["identityTokenKey"] = args ? args.identityTokenKey : undefined;
            resourceInputs["identityTokenTtl"] = args ? args.identityTokenTtl : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["rotationPeriod"] = args ? args.rotationPeriod : undefined;
            resourceInputs["rotationSchedule"] = args ? args.rotationSchedule : undefined;
            resourceInputs["rotationWindow"] = args ? args.rotationWindow : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["stsEndpoint"] = args ? args.stsEndpoint : undefined;
            resourceInputs["stsFallbackEndpoints"] = args ? args.stsFallbackEndpoints : undefined;
            resourceInputs["stsFallbackRegions"] = args ? args.stsFallbackRegions : undefined;
            resourceInputs["stsRegion"] = args ? args.stsRegion : undefined;
            resourceInputs["usernameTemplate"] = args ? args.usernameTemplate : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecretBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * The AWS Access Key ID this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The default TTL for credentials
     * issued by this backend.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     */
    disableAutomatedRotation?: pulumi.Input<boolean>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Specifies a custom HTTP IAM endpoint to use.
     */
    iamEndpoint?: pulumi.Input<string>;
    /**
     * The audience claim value. Requires Vault 1.16+.
     */
    identityTokenAudience?: pulumi.Input<string>;
    /**
     * The key to use for signing identity tokens. Requires Vault 1.16+.
     */
    identityTokenKey?: pulumi.Input<string>;
    /**
     * The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
     */
    identityTokenTtl?: pulumi.Input<number>;
    /**
     * Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     */
    path?: pulumi.Input<string>;
    /**
     * The AWS region to make API calls against. Defaults to us-east-1.
     */
    region?: pulumi.Input<string>;
    /**
     * Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential. 
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     */
    rotationPeriod?: pulumi.Input<number>;
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     */
    rotationSchedule?: pulumi.Input<string>;
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     */
    rotationWindow?: pulumi.Input<number>;
    /**
     * The AWS Secret Access Key to use when generating new credentials.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * Specifies a custom HTTP STS endpoint to use.
     */
    stsEndpoint?: pulumi.Input<string>;
    /**
     * Ordered list of `stsEndpoint`s to try if the defined one fails. Requires Vault 1.19+
     */
    stsFallbackEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ordered list of `stsRegion`s matching the fallback endpoints. Should correspond in order with those endpoints. Requires Vault 1.19+
     */
    stsFallbackRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the region of the STS endpoint. Should be included if `stsEndpoint` is supplied. Requires Vault 1.19+
     */
    stsRegion?: pulumi.Input<string>;
    /**
     * Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
     *
     * ```
     * {{ if (eq .Type "STS") }}
     * {{ printf "vault-%s-%s" (unix_time) (random 20) | truncate 32 }}
     * {{ else }}
     * {{ printf "vault-%s-%s-%s" (printf "%s-%s" (.DisplayName) (.PolicyName) | truncate 42) (unix_time) (random 20) | truncate 64 }}
     * {{ end }}
     *
     * ```
     */
    usernameTemplate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * The AWS Access Key ID this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The default TTL for credentials
     * issued by this backend.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     */
    disableAutomatedRotation?: pulumi.Input<boolean>;
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Specifies a custom HTTP IAM endpoint to use.
     */
    iamEndpoint?: pulumi.Input<string>;
    /**
     * The audience claim value. Requires Vault 1.16+.
     */
    identityTokenAudience?: pulumi.Input<string>;
    /**
     * The key to use for signing identity tokens. Requires Vault 1.16+.
     */
    identityTokenKey?: pulumi.Input<string>;
    /**
     * The TTL of generated identity tokens in seconds. Requires Vault 1.16+.
     */
    identityTokenTtl?: pulumi.Input<number>;
    /**
     * Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     */
    path?: pulumi.Input<string>;
    /**
     * The AWS region to make API calls against. Defaults to us-east-1.
     */
    region?: pulumi.Input<string>;
    /**
     * Role ARN to assume for plugin identity token federation. Requires Vault 1.16+.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The amount of time in seconds Vault should wait before rotating the root credential. 
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     */
    rotationPeriod?: pulumi.Input<number>;
    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     */
    rotationSchedule?: pulumi.Input<string>;
    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     */
    rotationWindow?: pulumi.Input<number>;
    /**
     * The AWS Secret Access Key to use when generating new credentials.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * Specifies a custom HTTP STS endpoint to use.
     */
    stsEndpoint?: pulumi.Input<string>;
    /**
     * Ordered list of `stsEndpoint`s to try if the defined one fails. Requires Vault 1.19+
     */
    stsFallbackEndpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ordered list of `stsRegion`s matching the fallback endpoints. Should correspond in order with those endpoints. Requires Vault 1.19+
     */
    stsFallbackRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the region of the STS endpoint. Should be included if `stsEndpoint` is supplied. Requires Vault 1.19+
     */
    stsRegion?: pulumi.Input<string>;
    /**
     * Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
     *
     * ```
     * {{ if (eq .Type "STS") }}
     * {{ printf "vault-%s-%s" (unix_time) (random 20) | truncate 32 }}
     * {{ else }}
     * {{ printf "vault-%s-%s-%s" (printf "%s-%s" (.DisplayName) (.PolicyName) | truncate 42) (unix_time) (random 20) | truncate 64 }}
     * {{ end }}
     *
     * ```
     */
    usernameTemplate?: pulumi.Input<string>;
}
