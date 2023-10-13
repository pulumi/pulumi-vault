// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * AWS secret backends can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:aws/secretBackend:SecretBackend aws aws
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
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    public readonly disableRemount!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies a custom HTTP IAM endpoint to use.
     */
    public readonly iamEndpoint!: pulumi.Output<string | undefined>;
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
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The AWS region for API calls. Defaults to `us-east-1`.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The AWS Secret Key this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * Specifies a custom HTTP STS endpoint to use.
     */
    public readonly stsEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
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
            resourceInputs["disableRemount"] = state ? state.disableRemount : undefined;
            resourceInputs["iamEndpoint"] = state ? state.iamEndpoint : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["stsEndpoint"] = state ? state.stsEndpoint : undefined;
            resourceInputs["usernameTemplate"] = state ? state.usernameTemplate : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            resourceInputs["accessKey"] = args?.accessKey ? pulumi.secret(args.accessKey) : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRemount"] = args ? args.disableRemount : undefined;
            resourceInputs["iamEndpoint"] = args ? args.iamEndpoint : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["stsEndpoint"] = args ? args.stsEndpoint : undefined;
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
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Specifies a custom HTTP IAM endpoint to use.
     */
    iamEndpoint?: pulumi.Input<string>;
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
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     */
    path?: pulumi.Input<string>;
    /**
     * The AWS region for API calls. Defaults to `us-east-1`.
     */
    region?: pulumi.Input<string>;
    /**
     * The AWS Secret Key this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * Specifies a custom HTTP STS endpoint to use.
     */
    stsEndpoint?: pulumi.Input<string>;
    /**
     * Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
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
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     */
    disableRemount?: pulumi.Input<boolean>;
    /**
     * Specifies a custom HTTP IAM endpoint to use.
     */
    iamEndpoint?: pulumi.Input<string>;
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
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     */
    path?: pulumi.Input<string>;
    /**
     * The AWS region for API calls. Defaults to `us-east-1`.
     */
    region?: pulumi.Input<string>;
    /**
     * The AWS Secret Key this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * Specifies a custom HTTP STS endpoint to use.
     */
    stsEndpoint?: pulumi.Input<string>;
    /**
     * Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
     */
    usernameTemplate?: pulumi.Input<string>;
}
