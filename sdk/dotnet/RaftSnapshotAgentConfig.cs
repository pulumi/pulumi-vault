// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// #### Local Storage
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var localBackups = new Vault.RaftSnapshotAgentConfig("local_backups", new()
    ///     {
    ///         Name = "local",
    ///         IntervalSeconds = 86400,
    ///         Retain = 7,
    ///         PathPrefix = "/opt/vault/snapshots/",
    ///         StorageType = "local",
    ///         LocalMaxSpace = 10000000,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// #### AWS S3
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var awsAccessKeyId = config.RequireObject&lt;dynamic&gt;("awsAccessKeyId");
    ///     var awsSecretAccessKey = config.RequireObject&lt;dynamic&gt;("awsSecretAccessKey");
    ///     var current = Aws.GetRegion.Invoke();
    /// 
    ///     var s3Backups = new Vault.RaftSnapshotAgentConfig("s3_backups", new()
    ///     {
    ///         Name = "s3",
    ///         IntervalSeconds = 86400,
    ///         Retain = 7,
    ///         PathPrefix = "/path/in/bucket",
    ///         StorageType = "aws-s3",
    ///         AwsS3Bucket = "my-bucket",
    ///         AwsS3Region = current.Apply(getRegionResult =&gt; getRegionResult.Name),
    ///         AwsAccessKeyId = awsAccessKeyId,
    ///         AwsSecretAccessKey = awsSecretAccessKey,
    ///         AwsS3EnableKms = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// #### Azure BLOB
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var azureAccountName = config.RequireObject&lt;dynamic&gt;("azureAccountName");
    ///     var azureAccountKey = config.RequireObject&lt;dynamic&gt;("azureAccountKey");
    ///     var azureBackups = new Vault.RaftSnapshotAgentConfig("azure_backups", new()
    ///     {
    ///         Name = "azure_backup",
    ///         IntervalSeconds = 86400,
    ///         Retain = 7,
    ///         PathPrefix = "/",
    ///         StorageType = "azure-blob",
    ///         AzureContainerName = "vault-blob",
    ///         AzureAccountName = azureAccountName,
    ///         AzureAccountKey = azureAccountKey,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Raft Snapshot Agent Configurations can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig local local
    /// ```
    /// </summary>
    [VaultResourceType("vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig")]
    public partial class RaftSnapshotAgentConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AWS access key ID.
        /// </summary>
        [Output("awsAccessKeyId")]
        public Output<string?> AwsAccessKeyId { get; private set; } = null!;

        /// <summary>
        /// S3 bucket to write snapshots to.
        /// </summary>
        [Output("awsS3Bucket")]
        public Output<string?> AwsS3Bucket { get; private set; } = null!;

        /// <summary>
        /// Disable TLS for the S3 endpoint. This should only be used for testing purposes.
        /// </summary>
        [Output("awsS3DisableTls")]
        public Output<bool?> AwsS3DisableTls { get; private set; } = null!;

        /// <summary>
        /// Use KMS to encrypt bucket contents.
        /// </summary>
        [Output("awsS3EnableKms")]
        public Output<bool?> AwsS3EnableKms { get; private set; } = null!;

        /// <summary>
        /// AWS endpoint. This is typically only set when using a non-AWS S3 implementation like Minio.
        /// </summary>
        [Output("awsS3Endpoint")]
        public Output<string?> AwsS3Endpoint { get; private set; } = null!;

        /// <summary>
        /// Use the endpoint/bucket URL style instead of bucket.endpoint.
        /// </summary>
        [Output("awsS3ForcePathStyle")]
        public Output<bool?> AwsS3ForcePathStyle { get; private set; } = null!;

        /// <summary>
        /// Use named KMS key, when aws_s3_enable_kms=true
        /// </summary>
        [Output("awsS3KmsKey")]
        public Output<string?> AwsS3KmsKey { get; private set; } = null!;

        /// <summary>
        /// AWS region bucket is in.
        /// </summary>
        [Output("awsS3Region")]
        public Output<string?> AwsS3Region { get; private set; } = null!;

        /// <summary>
        /// Use AES256 to encrypt bucket contents.
        /// </summary>
        [Output("awsS3ServerSideEncryption")]
        public Output<bool?> AwsS3ServerSideEncryption { get; private set; } = null!;

        /// <summary>
        /// AWS secret access key.
        /// </summary>
        [Output("awsSecretAccessKey")]
        public Output<string?> AwsSecretAccessKey { get; private set; } = null!;

        /// <summary>
        /// AWS session token.
        /// </summary>
        [Output("awsSessionToken")]
        public Output<string?> AwsSessionToken { get; private set; } = null!;

        /// <summary>
        /// Azure account key.
        /// </summary>
        [Output("azureAccountKey")]
        public Output<string?> AzureAccountKey { get; private set; } = null!;

        /// <summary>
        /// Azure account name.
        /// </summary>
        [Output("azureAccountName")]
        public Output<string?> AzureAccountName { get; private set; } = null!;

        /// <summary>
        /// Azure blob environment.
        /// </summary>
        [Output("azureBlobEnvironment")]
        public Output<string?> AzureBlobEnvironment { get; private set; } = null!;

        /// <summary>
        /// Azure container name to write snapshots to.
        /// </summary>
        [Output("azureContainerName")]
        public Output<string?> AzureContainerName { get; private set; } = null!;

        /// <summary>
        /// Azure blob storage endpoint. This is typically only set when using a non-Azure implementation like Azurite.
        /// </summary>
        [Output("azureEndpoint")]
        public Output<string?> AzureEndpoint { get; private set; } = null!;

        /// <summary>
        /// Within the directory or bucket
        /// prefix given by `path_prefix`, the file or object name of snapshot files
        /// will start with this string.
        /// </summary>
        [Output("filePrefix")]
        public Output<string?> FilePrefix { get; private set; } = null!;

        /// <summary>
        /// Disable TLS for the GCS endpoint.
        /// </summary>
        [Output("googleDisableTls")]
        public Output<bool?> GoogleDisableTls { get; private set; } = null!;

        /// <summary>
        /// GCS endpoint. This is typically only set when using a non-Google GCS implementation like fake-gcs-server.
        /// </summary>
        [Output("googleEndpoint")]
        public Output<string?> GoogleEndpoint { get; private set; } = null!;

        /// <summary>
        /// GCS bucket to write snapshots to.
        /// </summary>
        [Output("googleGcsBucket")]
        public Output<string?> GoogleGcsBucket { get; private set; } = null!;

        /// <summary>
        /// Google service account key in JSON format.
        /// </summary>
        [Output("googleServiceAccountKey")]
        public Output<string?> GoogleServiceAccountKey { get; private set; } = null!;

        /// <summary>
        /// `&lt;required&gt;` - Time (in seconds) between snapshots.
        /// </summary>
        [Output("intervalSeconds")]
        public Output<int> IntervalSeconds { get; private set; } = null!;

        /// <summary>
        /// The maximum space, in bytes, to use for snapshots.
        /// </summary>
        [Output("localMaxSpace")]
        public Output<int?> LocalMaxSpace { get; private set; } = null!;

        /// <summary>
        /// `&lt;required&gt;` – Name of the configuration to modify.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// `&lt;required&gt;` - For `storage_type = "local"`, the directory to
        /// write the snapshots in. For cloud storage types, the bucket prefix to use.
        /// Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
        /// Types `local` and `aws-s3` the trailing `/` is optional.
        /// </summary>
        [Output("pathPrefix")]
        public Output<string> PathPrefix { get; private set; } = null!;

        /// <summary>
        /// How many snapshots are to be kept; when writing a
        /// snapshot, if there are more snapshots already stored than this number, the
        /// oldest ones will be deleted.
        /// </summary>
        [Output("retain")]
        public Output<int?> Retain { get; private set; } = null!;

        /// <summary>
        /// `&lt;required&gt;` - One of "local", "azure-blob", "aws-s3",
        /// or "google-gcs". The remaining parameters described below are all specific to
        /// the selected `storage_type` and prefixed accordingly.
        /// </summary>
        [Output("storageType")]
        public Output<string> StorageType { get; private set; } = null!;


        /// <summary>
        /// Create a RaftSnapshotAgentConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RaftSnapshotAgentConfig(string name, RaftSnapshotAgentConfigArgs args, CustomResourceOptions? options = null)
            : base("vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig", name, args ?? new RaftSnapshotAgentConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RaftSnapshotAgentConfig(string name, Input<string> id, RaftSnapshotAgentConfigState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RaftSnapshotAgentConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RaftSnapshotAgentConfig Get(string name, Input<string> id, RaftSnapshotAgentConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new RaftSnapshotAgentConfig(name, id, state, options);
        }
    }

    public sealed class RaftSnapshotAgentConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS access key ID.
        /// </summary>
        [Input("awsAccessKeyId")]
        public Input<string>? AwsAccessKeyId { get; set; }

        /// <summary>
        /// S3 bucket to write snapshots to.
        /// </summary>
        [Input("awsS3Bucket")]
        public Input<string>? AwsS3Bucket { get; set; }

        /// <summary>
        /// Disable TLS for the S3 endpoint. This should only be used for testing purposes.
        /// </summary>
        [Input("awsS3DisableTls")]
        public Input<bool>? AwsS3DisableTls { get; set; }

        /// <summary>
        /// Use KMS to encrypt bucket contents.
        /// </summary>
        [Input("awsS3EnableKms")]
        public Input<bool>? AwsS3EnableKms { get; set; }

        /// <summary>
        /// AWS endpoint. This is typically only set when using a non-AWS S3 implementation like Minio.
        /// </summary>
        [Input("awsS3Endpoint")]
        public Input<string>? AwsS3Endpoint { get; set; }

        /// <summary>
        /// Use the endpoint/bucket URL style instead of bucket.endpoint.
        /// </summary>
        [Input("awsS3ForcePathStyle")]
        public Input<bool>? AwsS3ForcePathStyle { get; set; }

        /// <summary>
        /// Use named KMS key, when aws_s3_enable_kms=true
        /// </summary>
        [Input("awsS3KmsKey")]
        public Input<string>? AwsS3KmsKey { get; set; }

        /// <summary>
        /// AWS region bucket is in.
        /// </summary>
        [Input("awsS3Region")]
        public Input<string>? AwsS3Region { get; set; }

        /// <summary>
        /// Use AES256 to encrypt bucket contents.
        /// </summary>
        [Input("awsS3ServerSideEncryption")]
        public Input<bool>? AwsS3ServerSideEncryption { get; set; }

        /// <summary>
        /// AWS secret access key.
        /// </summary>
        [Input("awsSecretAccessKey")]
        public Input<string>? AwsSecretAccessKey { get; set; }

        /// <summary>
        /// AWS session token.
        /// </summary>
        [Input("awsSessionToken")]
        public Input<string>? AwsSessionToken { get; set; }

        /// <summary>
        /// Azure account key.
        /// </summary>
        [Input("azureAccountKey")]
        public Input<string>? AzureAccountKey { get; set; }

        /// <summary>
        /// Azure account name.
        /// </summary>
        [Input("azureAccountName")]
        public Input<string>? AzureAccountName { get; set; }

        /// <summary>
        /// Azure blob environment.
        /// </summary>
        [Input("azureBlobEnvironment")]
        public Input<string>? AzureBlobEnvironment { get; set; }

        /// <summary>
        /// Azure container name to write snapshots to.
        /// </summary>
        [Input("azureContainerName")]
        public Input<string>? AzureContainerName { get; set; }

        /// <summary>
        /// Azure blob storage endpoint. This is typically only set when using a non-Azure implementation like Azurite.
        /// </summary>
        [Input("azureEndpoint")]
        public Input<string>? AzureEndpoint { get; set; }

        /// <summary>
        /// Within the directory or bucket
        /// prefix given by `path_prefix`, the file or object name of snapshot files
        /// will start with this string.
        /// </summary>
        [Input("filePrefix")]
        public Input<string>? FilePrefix { get; set; }

        /// <summary>
        /// Disable TLS for the GCS endpoint.
        /// </summary>
        [Input("googleDisableTls")]
        public Input<bool>? GoogleDisableTls { get; set; }

        /// <summary>
        /// GCS endpoint. This is typically only set when using a non-Google GCS implementation like fake-gcs-server.
        /// </summary>
        [Input("googleEndpoint")]
        public Input<string>? GoogleEndpoint { get; set; }

        /// <summary>
        /// GCS bucket to write snapshots to.
        /// </summary>
        [Input("googleGcsBucket")]
        public Input<string>? GoogleGcsBucket { get; set; }

        /// <summary>
        /// Google service account key in JSON format.
        /// </summary>
        [Input("googleServiceAccountKey")]
        public Input<string>? GoogleServiceAccountKey { get; set; }

        /// <summary>
        /// `&lt;required&gt;` - Time (in seconds) between snapshots.
        /// </summary>
        [Input("intervalSeconds", required: true)]
        public Input<int> IntervalSeconds { get; set; } = null!;

        /// <summary>
        /// The maximum space, in bytes, to use for snapshots.
        /// </summary>
        [Input("localMaxSpace")]
        public Input<int>? LocalMaxSpace { get; set; }

        /// <summary>
        /// `&lt;required&gt;` – Name of the configuration to modify.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// `&lt;required&gt;` - For `storage_type = "local"`, the directory to
        /// write the snapshots in. For cloud storage types, the bucket prefix to use.
        /// Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
        /// Types `local` and `aws-s3` the trailing `/` is optional.
        /// </summary>
        [Input("pathPrefix", required: true)]
        public Input<string> PathPrefix { get; set; } = null!;

        /// <summary>
        /// How many snapshots are to be kept; when writing a
        /// snapshot, if there are more snapshots already stored than this number, the
        /// oldest ones will be deleted.
        /// </summary>
        [Input("retain")]
        public Input<int>? Retain { get; set; }

        /// <summary>
        /// `&lt;required&gt;` - One of "local", "azure-blob", "aws-s3",
        /// or "google-gcs". The remaining parameters described below are all specific to
        /// the selected `storage_type` and prefixed accordingly.
        /// </summary>
        [Input("storageType", required: true)]
        public Input<string> StorageType { get; set; } = null!;

        public RaftSnapshotAgentConfigArgs()
        {
        }
        public static new RaftSnapshotAgentConfigArgs Empty => new RaftSnapshotAgentConfigArgs();
    }

    public sealed class RaftSnapshotAgentConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS access key ID.
        /// </summary>
        [Input("awsAccessKeyId")]
        public Input<string>? AwsAccessKeyId { get; set; }

        /// <summary>
        /// S3 bucket to write snapshots to.
        /// </summary>
        [Input("awsS3Bucket")]
        public Input<string>? AwsS3Bucket { get; set; }

        /// <summary>
        /// Disable TLS for the S3 endpoint. This should only be used for testing purposes.
        /// </summary>
        [Input("awsS3DisableTls")]
        public Input<bool>? AwsS3DisableTls { get; set; }

        /// <summary>
        /// Use KMS to encrypt bucket contents.
        /// </summary>
        [Input("awsS3EnableKms")]
        public Input<bool>? AwsS3EnableKms { get; set; }

        /// <summary>
        /// AWS endpoint. This is typically only set when using a non-AWS S3 implementation like Minio.
        /// </summary>
        [Input("awsS3Endpoint")]
        public Input<string>? AwsS3Endpoint { get; set; }

        /// <summary>
        /// Use the endpoint/bucket URL style instead of bucket.endpoint.
        /// </summary>
        [Input("awsS3ForcePathStyle")]
        public Input<bool>? AwsS3ForcePathStyle { get; set; }

        /// <summary>
        /// Use named KMS key, when aws_s3_enable_kms=true
        /// </summary>
        [Input("awsS3KmsKey")]
        public Input<string>? AwsS3KmsKey { get; set; }

        /// <summary>
        /// AWS region bucket is in.
        /// </summary>
        [Input("awsS3Region")]
        public Input<string>? AwsS3Region { get; set; }

        /// <summary>
        /// Use AES256 to encrypt bucket contents.
        /// </summary>
        [Input("awsS3ServerSideEncryption")]
        public Input<bool>? AwsS3ServerSideEncryption { get; set; }

        /// <summary>
        /// AWS secret access key.
        /// </summary>
        [Input("awsSecretAccessKey")]
        public Input<string>? AwsSecretAccessKey { get; set; }

        /// <summary>
        /// AWS session token.
        /// </summary>
        [Input("awsSessionToken")]
        public Input<string>? AwsSessionToken { get; set; }

        /// <summary>
        /// Azure account key.
        /// </summary>
        [Input("azureAccountKey")]
        public Input<string>? AzureAccountKey { get; set; }

        /// <summary>
        /// Azure account name.
        /// </summary>
        [Input("azureAccountName")]
        public Input<string>? AzureAccountName { get; set; }

        /// <summary>
        /// Azure blob environment.
        /// </summary>
        [Input("azureBlobEnvironment")]
        public Input<string>? AzureBlobEnvironment { get; set; }

        /// <summary>
        /// Azure container name to write snapshots to.
        /// </summary>
        [Input("azureContainerName")]
        public Input<string>? AzureContainerName { get; set; }

        /// <summary>
        /// Azure blob storage endpoint. This is typically only set when using a non-Azure implementation like Azurite.
        /// </summary>
        [Input("azureEndpoint")]
        public Input<string>? AzureEndpoint { get; set; }

        /// <summary>
        /// Within the directory or bucket
        /// prefix given by `path_prefix`, the file or object name of snapshot files
        /// will start with this string.
        /// </summary>
        [Input("filePrefix")]
        public Input<string>? FilePrefix { get; set; }

        /// <summary>
        /// Disable TLS for the GCS endpoint.
        /// </summary>
        [Input("googleDisableTls")]
        public Input<bool>? GoogleDisableTls { get; set; }

        /// <summary>
        /// GCS endpoint. This is typically only set when using a non-Google GCS implementation like fake-gcs-server.
        /// </summary>
        [Input("googleEndpoint")]
        public Input<string>? GoogleEndpoint { get; set; }

        /// <summary>
        /// GCS bucket to write snapshots to.
        /// </summary>
        [Input("googleGcsBucket")]
        public Input<string>? GoogleGcsBucket { get; set; }

        /// <summary>
        /// Google service account key in JSON format.
        /// </summary>
        [Input("googleServiceAccountKey")]
        public Input<string>? GoogleServiceAccountKey { get; set; }

        /// <summary>
        /// `&lt;required&gt;` - Time (in seconds) between snapshots.
        /// </summary>
        [Input("intervalSeconds")]
        public Input<int>? IntervalSeconds { get; set; }

        /// <summary>
        /// The maximum space, in bytes, to use for snapshots.
        /// </summary>
        [Input("localMaxSpace")]
        public Input<int>? LocalMaxSpace { get; set; }

        /// <summary>
        /// `&lt;required&gt;` – Name of the configuration to modify.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// `&lt;required&gt;` - For `storage_type = "local"`, the directory to
        /// write the snapshots in. For cloud storage types, the bucket prefix to use.
        /// Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
        /// Types `local` and `aws-s3` the trailing `/` is optional.
        /// </summary>
        [Input("pathPrefix")]
        public Input<string>? PathPrefix { get; set; }

        /// <summary>
        /// How many snapshots are to be kept; when writing a
        /// snapshot, if there are more snapshots already stored than this number, the
        /// oldest ones will be deleted.
        /// </summary>
        [Input("retain")]
        public Input<int>? Retain { get; set; }

        /// <summary>
        /// `&lt;required&gt;` - One of "local", "azure-blob", "aws-s3",
        /// or "google-gcs". The remaining parameters described below are all specific to
        /// the selected `storage_type` and prefixed accordingly.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        public RaftSnapshotAgentConfigState()
        {
        }
        public static new RaftSnapshotAgentConfigState Empty => new RaftSnapshotAgentConfigState();
    }
}
