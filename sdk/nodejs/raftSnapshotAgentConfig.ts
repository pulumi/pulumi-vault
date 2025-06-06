// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ### Local Storage
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const localBackups = new vault.RaftSnapshotAgentConfig("local_backups", {
 *     name: "local",
 *     intervalSeconds: 86400,
 *     retain: 7,
 *     pathPrefix: "/opt/vault/snapshots/",
 *     storageType: "local",
 *     localMaxSpace: 10000000,
 * });
 * ```
 *
 * ### AWS S3
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new pulumi.Config();
 * const awsAccessKeyId = config.requireObject<any>("awsAccessKeyId");
 * const awsSecretAccessKey = config.requireObject<any>("awsSecretAccessKey");
 * const current = aws.getRegion({});
 * const s3Backups = new vault.RaftSnapshotAgentConfig("s3_backups", {
 *     name: "s3",
 *     intervalSeconds: 86400,
 *     retain: 7,
 *     pathPrefix: "/path/in/bucket",
 *     storageType: "aws-s3",
 *     awsS3Bucket: "my-bucket",
 *     awsS3Region: current.then(current => current.name),
 *     awsAccessKeyId: awsAccessKeyId,
 *     awsSecretAccessKey: awsSecretAccessKey,
 *     awsS3EnableKms: true,
 * });
 * ```
 *
 * ### Azure BLOB
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const config = new pulumi.Config();
 * const azureAccountName = config.requireObject<any>("azureAccountName");
 * const azureAccountKey = config.requireObject<any>("azureAccountKey");
 * const azureBackups = new vault.RaftSnapshotAgentConfig("azure_backups", {
 *     name: "azure_backup",
 *     intervalSeconds: 86400,
 *     retain: 7,
 *     pathPrefix: "/",
 *     storageType: "azure-blob",
 *     azureContainerName: "vault-blob",
 *     azureAccountName: azureAccountName,
 *     azureAccountKey: azureAccountKey,
 * });
 * ```
 *
 * ## Import
 *
 * Raft Snapshot Agent Configurations can be imported using the `name`, e.g.
 *
 * ```sh
 * $ pulumi import vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig local local
 * ```
 */
export class RaftSnapshotAgentConfig extends pulumi.CustomResource {
    /**
     * Get an existing RaftSnapshotAgentConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RaftSnapshotAgentConfigState, opts?: pulumi.CustomResourceOptions): RaftSnapshotAgentConfig {
        return new RaftSnapshotAgentConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig';

    /**
     * Returns true if the given object is an instance of RaftSnapshotAgentConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RaftSnapshotAgentConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RaftSnapshotAgentConfig.__pulumiType;
    }

    /**
     * AWS access key ID.
     */
    public readonly awsAccessKeyId!: pulumi.Output<string | undefined>;
    /**
     * S3 bucket to write snapshots to.
     */
    public readonly awsS3Bucket!: pulumi.Output<string | undefined>;
    /**
     * Disable TLS for the S3 endpoint. This should only be used for testing purposes.
     */
    public readonly awsS3DisableTls!: pulumi.Output<boolean | undefined>;
    /**
     * Use KMS to encrypt bucket contents.
     */
    public readonly awsS3EnableKms!: pulumi.Output<boolean | undefined>;
    /**
     * AWS endpoint. This is typically only set when using a non-AWS S3 implementation like Minio.
     */
    public readonly awsS3Endpoint!: pulumi.Output<string | undefined>;
    /**
     * Use the endpoint/bucket URL style instead of bucket.endpoint.
     */
    public readonly awsS3ForcePathStyle!: pulumi.Output<boolean | undefined>;
    /**
     * Use named KMS key, when aws_s3_enable_kms=true
     */
    public readonly awsS3KmsKey!: pulumi.Output<string | undefined>;
    /**
     * AWS region bucket is in.
     */
    public readonly awsS3Region!: pulumi.Output<string | undefined>;
    /**
     * Use AES256 to encrypt bucket contents.
     */
    public readonly awsS3ServerSideEncryption!: pulumi.Output<boolean | undefined>;
    /**
     * AWS secret access key.
     */
    public readonly awsSecretAccessKey!: pulumi.Output<string | undefined>;
    /**
     * AWS session token.
     */
    public readonly awsSessionToken!: pulumi.Output<string | undefined>;
    /**
     * Azure account key.
     */
    public readonly azureAccountKey!: pulumi.Output<string | undefined>;
    /**
     * Azure account name.
     */
    public readonly azureAccountName!: pulumi.Output<string | undefined>;
    /**
     * Azure blob environment.
     */
    public readonly azureBlobEnvironment!: pulumi.Output<string | undefined>;
    /**
     * Azure container name to write snapshots to.
     */
    public readonly azureContainerName!: pulumi.Output<string | undefined>;
    /**
     * Azure blob storage endpoint. This is typically only set when using a non-Azure implementation like Azurite.
     */
    public readonly azureEndpoint!: pulumi.Output<string | undefined>;
    /**
     * Within the directory or bucket
     * prefix given by `pathPrefix`, the file or object name of snapshot files
     * will start with this string.
     */
    public readonly filePrefix!: pulumi.Output<string | undefined>;
    /**
     * Disable TLS for the GCS endpoint.
     */
    public readonly googleDisableTls!: pulumi.Output<boolean | undefined>;
    /**
     * GCS endpoint. This is typically only set when using a non-Google GCS implementation like fake-gcs-server.
     */
    public readonly googleEndpoint!: pulumi.Output<string | undefined>;
    /**
     * GCS bucket to write snapshots to.
     */
    public readonly googleGcsBucket!: pulumi.Output<string | undefined>;
    /**
     * Google service account key in JSON format.
     */
    public readonly googleServiceAccountKey!: pulumi.Output<string | undefined>;
    /**
     * `<required>` - Time (in seconds) between snapshots.
     */
    public readonly intervalSeconds!: pulumi.Output<number>;
    /**
     * The maximum space, in bytes, to use for snapshots.
     */
    public readonly localMaxSpace!: pulumi.Output<number | undefined>;
    /**
     * `<required>` – Name of the configuration to modify.
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
     * `<required>` - For `storageType = "local"`, the directory to
     * write the snapshots in. For cloud storage types, the bucket prefix to use.
     * Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
     * Types `local` and `aws-s3` the trailing `/` is optional.
     */
    public readonly pathPrefix!: pulumi.Output<string>;
    /**
     * How many snapshots are to be kept; when writing a
     * snapshot, if there are more snapshots already stored than this number, the
     * oldest ones will be deleted.
     */
    public readonly retain!: pulumi.Output<number | undefined>;
    /**
     * `<required>` - One of "local", "azure-blob", "aws-s3",
     * or "google-gcs". The remaining parameters described below are all specific to
     * the selected `storageType` and prefixed accordingly.
     */
    public readonly storageType!: pulumi.Output<string>;

    /**
     * Create a RaftSnapshotAgentConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RaftSnapshotAgentConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RaftSnapshotAgentConfigArgs | RaftSnapshotAgentConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RaftSnapshotAgentConfigState | undefined;
            resourceInputs["awsAccessKeyId"] = state ? state.awsAccessKeyId : undefined;
            resourceInputs["awsS3Bucket"] = state ? state.awsS3Bucket : undefined;
            resourceInputs["awsS3DisableTls"] = state ? state.awsS3DisableTls : undefined;
            resourceInputs["awsS3EnableKms"] = state ? state.awsS3EnableKms : undefined;
            resourceInputs["awsS3Endpoint"] = state ? state.awsS3Endpoint : undefined;
            resourceInputs["awsS3ForcePathStyle"] = state ? state.awsS3ForcePathStyle : undefined;
            resourceInputs["awsS3KmsKey"] = state ? state.awsS3KmsKey : undefined;
            resourceInputs["awsS3Region"] = state ? state.awsS3Region : undefined;
            resourceInputs["awsS3ServerSideEncryption"] = state ? state.awsS3ServerSideEncryption : undefined;
            resourceInputs["awsSecretAccessKey"] = state ? state.awsSecretAccessKey : undefined;
            resourceInputs["awsSessionToken"] = state ? state.awsSessionToken : undefined;
            resourceInputs["azureAccountKey"] = state ? state.azureAccountKey : undefined;
            resourceInputs["azureAccountName"] = state ? state.azureAccountName : undefined;
            resourceInputs["azureBlobEnvironment"] = state ? state.azureBlobEnvironment : undefined;
            resourceInputs["azureContainerName"] = state ? state.azureContainerName : undefined;
            resourceInputs["azureEndpoint"] = state ? state.azureEndpoint : undefined;
            resourceInputs["filePrefix"] = state ? state.filePrefix : undefined;
            resourceInputs["googleDisableTls"] = state ? state.googleDisableTls : undefined;
            resourceInputs["googleEndpoint"] = state ? state.googleEndpoint : undefined;
            resourceInputs["googleGcsBucket"] = state ? state.googleGcsBucket : undefined;
            resourceInputs["googleServiceAccountKey"] = state ? state.googleServiceAccountKey : undefined;
            resourceInputs["intervalSeconds"] = state ? state.intervalSeconds : undefined;
            resourceInputs["localMaxSpace"] = state ? state.localMaxSpace : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["pathPrefix"] = state ? state.pathPrefix : undefined;
            resourceInputs["retain"] = state ? state.retain : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
        } else {
            const args = argsOrState as RaftSnapshotAgentConfigArgs | undefined;
            if ((!args || args.intervalSeconds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'intervalSeconds'");
            }
            if ((!args || args.pathPrefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pathPrefix'");
            }
            if ((!args || args.storageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageType'");
            }
            resourceInputs["awsAccessKeyId"] = args ? args.awsAccessKeyId : undefined;
            resourceInputs["awsS3Bucket"] = args ? args.awsS3Bucket : undefined;
            resourceInputs["awsS3DisableTls"] = args ? args.awsS3DisableTls : undefined;
            resourceInputs["awsS3EnableKms"] = args ? args.awsS3EnableKms : undefined;
            resourceInputs["awsS3Endpoint"] = args ? args.awsS3Endpoint : undefined;
            resourceInputs["awsS3ForcePathStyle"] = args ? args.awsS3ForcePathStyle : undefined;
            resourceInputs["awsS3KmsKey"] = args ? args.awsS3KmsKey : undefined;
            resourceInputs["awsS3Region"] = args ? args.awsS3Region : undefined;
            resourceInputs["awsS3ServerSideEncryption"] = args ? args.awsS3ServerSideEncryption : undefined;
            resourceInputs["awsSecretAccessKey"] = args ? args.awsSecretAccessKey : undefined;
            resourceInputs["awsSessionToken"] = args ? args.awsSessionToken : undefined;
            resourceInputs["azureAccountKey"] = args ? args.azureAccountKey : undefined;
            resourceInputs["azureAccountName"] = args ? args.azureAccountName : undefined;
            resourceInputs["azureBlobEnvironment"] = args ? args.azureBlobEnvironment : undefined;
            resourceInputs["azureContainerName"] = args ? args.azureContainerName : undefined;
            resourceInputs["azureEndpoint"] = args ? args.azureEndpoint : undefined;
            resourceInputs["filePrefix"] = args ? args.filePrefix : undefined;
            resourceInputs["googleDisableTls"] = args ? args.googleDisableTls : undefined;
            resourceInputs["googleEndpoint"] = args ? args.googleEndpoint : undefined;
            resourceInputs["googleGcsBucket"] = args ? args.googleGcsBucket : undefined;
            resourceInputs["googleServiceAccountKey"] = args ? args.googleServiceAccountKey : undefined;
            resourceInputs["intervalSeconds"] = args ? args.intervalSeconds : undefined;
            resourceInputs["localMaxSpace"] = args ? args.localMaxSpace : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["pathPrefix"] = args ? args.pathPrefix : undefined;
            resourceInputs["retain"] = args ? args.retain : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RaftSnapshotAgentConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RaftSnapshotAgentConfig resources.
 */
export interface RaftSnapshotAgentConfigState {
    /**
     * AWS access key ID.
     */
    awsAccessKeyId?: pulumi.Input<string>;
    /**
     * S3 bucket to write snapshots to.
     */
    awsS3Bucket?: pulumi.Input<string>;
    /**
     * Disable TLS for the S3 endpoint. This should only be used for testing purposes.
     */
    awsS3DisableTls?: pulumi.Input<boolean>;
    /**
     * Use KMS to encrypt bucket contents.
     */
    awsS3EnableKms?: pulumi.Input<boolean>;
    /**
     * AWS endpoint. This is typically only set when using a non-AWS S3 implementation like Minio.
     */
    awsS3Endpoint?: pulumi.Input<string>;
    /**
     * Use the endpoint/bucket URL style instead of bucket.endpoint.
     */
    awsS3ForcePathStyle?: pulumi.Input<boolean>;
    /**
     * Use named KMS key, when aws_s3_enable_kms=true
     */
    awsS3KmsKey?: pulumi.Input<string>;
    /**
     * AWS region bucket is in.
     */
    awsS3Region?: pulumi.Input<string>;
    /**
     * Use AES256 to encrypt bucket contents.
     */
    awsS3ServerSideEncryption?: pulumi.Input<boolean>;
    /**
     * AWS secret access key.
     */
    awsSecretAccessKey?: pulumi.Input<string>;
    /**
     * AWS session token.
     */
    awsSessionToken?: pulumi.Input<string>;
    /**
     * Azure account key.
     */
    azureAccountKey?: pulumi.Input<string>;
    /**
     * Azure account name.
     */
    azureAccountName?: pulumi.Input<string>;
    /**
     * Azure blob environment.
     */
    azureBlobEnvironment?: pulumi.Input<string>;
    /**
     * Azure container name to write snapshots to.
     */
    azureContainerName?: pulumi.Input<string>;
    /**
     * Azure blob storage endpoint. This is typically only set when using a non-Azure implementation like Azurite.
     */
    azureEndpoint?: pulumi.Input<string>;
    /**
     * Within the directory or bucket
     * prefix given by `pathPrefix`, the file or object name of snapshot files
     * will start with this string.
     */
    filePrefix?: pulumi.Input<string>;
    /**
     * Disable TLS for the GCS endpoint.
     */
    googleDisableTls?: pulumi.Input<boolean>;
    /**
     * GCS endpoint. This is typically only set when using a non-Google GCS implementation like fake-gcs-server.
     */
    googleEndpoint?: pulumi.Input<string>;
    /**
     * GCS bucket to write snapshots to.
     */
    googleGcsBucket?: pulumi.Input<string>;
    /**
     * Google service account key in JSON format.
     */
    googleServiceAccountKey?: pulumi.Input<string>;
    /**
     * `<required>` - Time (in seconds) between snapshots.
     */
    intervalSeconds?: pulumi.Input<number>;
    /**
     * The maximum space, in bytes, to use for snapshots.
     */
    localMaxSpace?: pulumi.Input<number>;
    /**
     * `<required>` – Name of the configuration to modify.
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
     * `<required>` - For `storageType = "local"`, the directory to
     * write the snapshots in. For cloud storage types, the bucket prefix to use.
     * Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
     * Types `local` and `aws-s3` the trailing `/` is optional.
     */
    pathPrefix?: pulumi.Input<string>;
    /**
     * How many snapshots are to be kept; when writing a
     * snapshot, if there are more snapshots already stored than this number, the
     * oldest ones will be deleted.
     */
    retain?: pulumi.Input<number>;
    /**
     * `<required>` - One of "local", "azure-blob", "aws-s3",
     * or "google-gcs". The remaining parameters described below are all specific to
     * the selected `storageType` and prefixed accordingly.
     */
    storageType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RaftSnapshotAgentConfig resource.
 */
export interface RaftSnapshotAgentConfigArgs {
    /**
     * AWS access key ID.
     */
    awsAccessKeyId?: pulumi.Input<string>;
    /**
     * S3 bucket to write snapshots to.
     */
    awsS3Bucket?: pulumi.Input<string>;
    /**
     * Disable TLS for the S3 endpoint. This should only be used for testing purposes.
     */
    awsS3DisableTls?: pulumi.Input<boolean>;
    /**
     * Use KMS to encrypt bucket contents.
     */
    awsS3EnableKms?: pulumi.Input<boolean>;
    /**
     * AWS endpoint. This is typically only set when using a non-AWS S3 implementation like Minio.
     */
    awsS3Endpoint?: pulumi.Input<string>;
    /**
     * Use the endpoint/bucket URL style instead of bucket.endpoint.
     */
    awsS3ForcePathStyle?: pulumi.Input<boolean>;
    /**
     * Use named KMS key, when aws_s3_enable_kms=true
     */
    awsS3KmsKey?: pulumi.Input<string>;
    /**
     * AWS region bucket is in.
     */
    awsS3Region?: pulumi.Input<string>;
    /**
     * Use AES256 to encrypt bucket contents.
     */
    awsS3ServerSideEncryption?: pulumi.Input<boolean>;
    /**
     * AWS secret access key.
     */
    awsSecretAccessKey?: pulumi.Input<string>;
    /**
     * AWS session token.
     */
    awsSessionToken?: pulumi.Input<string>;
    /**
     * Azure account key.
     */
    azureAccountKey?: pulumi.Input<string>;
    /**
     * Azure account name.
     */
    azureAccountName?: pulumi.Input<string>;
    /**
     * Azure blob environment.
     */
    azureBlobEnvironment?: pulumi.Input<string>;
    /**
     * Azure container name to write snapshots to.
     */
    azureContainerName?: pulumi.Input<string>;
    /**
     * Azure blob storage endpoint. This is typically only set when using a non-Azure implementation like Azurite.
     */
    azureEndpoint?: pulumi.Input<string>;
    /**
     * Within the directory or bucket
     * prefix given by `pathPrefix`, the file or object name of snapshot files
     * will start with this string.
     */
    filePrefix?: pulumi.Input<string>;
    /**
     * Disable TLS for the GCS endpoint.
     */
    googleDisableTls?: pulumi.Input<boolean>;
    /**
     * GCS endpoint. This is typically only set when using a non-Google GCS implementation like fake-gcs-server.
     */
    googleEndpoint?: pulumi.Input<string>;
    /**
     * GCS bucket to write snapshots to.
     */
    googleGcsBucket?: pulumi.Input<string>;
    /**
     * Google service account key in JSON format.
     */
    googleServiceAccountKey?: pulumi.Input<string>;
    /**
     * `<required>` - Time (in seconds) between snapshots.
     */
    intervalSeconds: pulumi.Input<number>;
    /**
     * The maximum space, in bytes, to use for snapshots.
     */
    localMaxSpace?: pulumi.Input<number>;
    /**
     * `<required>` – Name of the configuration to modify.
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
     * `<required>` - For `storageType = "local"`, the directory to
     * write the snapshots in. For cloud storage types, the bucket prefix to use.
     * Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
     * Types `local` and `aws-s3` the trailing `/` is optional.
     */
    pathPrefix: pulumi.Input<string>;
    /**
     * How many snapshots are to be kept; when writing a
     * snapshot, if there are more snapshots already stored than this number, the
     * oldest ones will be deleted.
     */
    retain?: pulumi.Input<number>;
    /**
     * `<required>` - One of "local", "azure-blob", "aws-s3",
     * or "google-gcs". The remaining parameters described below are all specific to
     * the selected `storageType` and prefixed accordingly.
     */
    storageType: pulumi.Input<string>;
}
