// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.RaftSnapshotAgentConfigArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.RaftSnapshotAgentConfigState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Local Storage
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.RaftSnapshotAgentConfig;
 * import com.pulumi.vault.RaftSnapshotAgentConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var localBackups = new RaftSnapshotAgentConfig(&#34;localBackups&#34;, RaftSnapshotAgentConfigArgs.builder()        
 *             .name(&#34;local&#34;)
 *             .intervalSeconds(86400)
 *             .retain(7)
 *             .pathPrefix(&#34;/opt/vault/snapshots/&#34;)
 *             .storageType(&#34;local&#34;)
 *             .localMaxSpace(10000000)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### AWS S3
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetRegionArgs;
 * import com.pulumi.vault.RaftSnapshotAgentConfig;
 * import com.pulumi.vault.RaftSnapshotAgentConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var awsAccessKeyId = config.get(&#34;awsAccessKeyId&#34;);
 *         final var awsSecretAccessKey = config.get(&#34;awsSecretAccessKey&#34;);
 *         final var current = AwsFunctions.getRegion();
 * 
 *         var s3Backups = new RaftSnapshotAgentConfig(&#34;s3Backups&#34;, RaftSnapshotAgentConfigArgs.builder()        
 *             .name(&#34;s3&#34;)
 *             .intervalSeconds(86400)
 *             .retain(7)
 *             .pathPrefix(&#34;/path/in/bucket&#34;)
 *             .storageType(&#34;aws-s3&#34;)
 *             .awsS3Bucket(&#34;my-bucket&#34;)
 *             .awsS3Region(current.applyValue(getRegionResult -&gt; getRegionResult.name()))
 *             .awsAccessKeyId(awsAccessKeyId)
 *             .awsSecretAccessKey(awsSecretAccessKey)
 *             .awsS3EnableKms(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Azure BLOB
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.RaftSnapshotAgentConfig;
 * import com.pulumi.vault.RaftSnapshotAgentConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var azureAccountName = config.get(&#34;azureAccountName&#34;);
 *         final var azureAccountKey = config.get(&#34;azureAccountKey&#34;);
 *         var azureBackups = new RaftSnapshotAgentConfig(&#34;azureBackups&#34;, RaftSnapshotAgentConfigArgs.builder()        
 *             .name(&#34;azure_backup&#34;)
 *             .intervalSeconds(86400)
 *             .retain(7)
 *             .pathPrefix(&#34;/&#34;)
 *             .storageType(&#34;azure-blob&#34;)
 *             .azureContainerName(&#34;vault-blob&#34;)
 *             .azureAccountName(azureAccountName)
 *             .azureAccountKey(azureAccountKey)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Raft Snapshot Agent Configurations can be imported using the `name`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig local local
 * ```
 * 
 */
@ResourceType(type="vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig")
public class RaftSnapshotAgentConfig extends com.pulumi.resources.CustomResource {
    /**
     * AWS access key ID.
     * 
     */
    @Export(name="awsAccessKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> awsAccessKeyId;

    /**
     * @return AWS access key ID.
     * 
     */
    public Output<Optional<String>> awsAccessKeyId() {
        return Codegen.optional(this.awsAccessKeyId);
    }
    /**
     * `&lt;required&gt;` - S3 bucket to write snapshots to.
     * 
     */
    @Export(name="awsS3Bucket", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> awsS3Bucket;

    /**
     * @return `&lt;required&gt;` - S3 bucket to write snapshots to.
     * 
     */
    public Output<Optional<String>> awsS3Bucket() {
        return Codegen.optional(this.awsS3Bucket);
    }
    /**
     * Disable TLS for the S3 endpoint. This
     * should only be used for testing purposes, typically in conjunction with
     * `aws_s3_endpoint`.
     * 
     */
    @Export(name="awsS3DisableTls", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> awsS3DisableTls;

    /**
     * @return Disable TLS for the S3 endpoint. This
     * should only be used for testing purposes, typically in conjunction with
     * `aws_s3_endpoint`.
     * 
     */
    public Output<Optional<Boolean>> awsS3DisableTls() {
        return Codegen.optional(this.awsS3DisableTls);
    }
    /**
     * Use KMS to encrypt bucket contents.
     * 
     */
    @Export(name="awsS3EnableKms", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> awsS3EnableKms;

    /**
     * @return Use KMS to encrypt bucket contents.
     * 
     */
    public Output<Optional<Boolean>> awsS3EnableKms() {
        return Codegen.optional(this.awsS3EnableKms);
    }
    /**
     * AWS endpoint. This is typically only set when
     * using a non-AWS S3 implementation like Minio.
     * 
     */
    @Export(name="awsS3Endpoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> awsS3Endpoint;

    /**
     * @return AWS endpoint. This is typically only set when
     * using a non-AWS S3 implementation like Minio.
     * 
     */
    public Output<Optional<String>> awsS3Endpoint() {
        return Codegen.optional(this.awsS3Endpoint);
    }
    /**
     * Use the endpoint/bucket URL style
     * instead of bucket.endpoint. May be needed when setting `aws_s3_endpoint`.
     * 
     */
    @Export(name="awsS3ForcePathStyle", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> awsS3ForcePathStyle;

    /**
     * @return Use the endpoint/bucket URL style
     * instead of bucket.endpoint. May be needed when setting `aws_s3_endpoint`.
     * 
     */
    public Output<Optional<Boolean>> awsS3ForcePathStyle() {
        return Codegen.optional(this.awsS3ForcePathStyle);
    }
    /**
     * Use named KMS key, when `aws_s3_enable_kms = true`
     * 
     */
    @Export(name="awsS3KmsKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> awsS3KmsKey;

    /**
     * @return Use named KMS key, when `aws_s3_enable_kms = true`
     * 
     */
    public Output<Optional<String>> awsS3KmsKey() {
        return Codegen.optional(this.awsS3KmsKey);
    }
    /**
     * `&lt;required&gt;` - AWS region bucket is in.
     * 
     */
    @Export(name="awsS3Region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> awsS3Region;

    /**
     * @return `&lt;required&gt;` - AWS region bucket is in.
     * 
     */
    public Output<Optional<String>> awsS3Region() {
        return Codegen.optional(this.awsS3Region);
    }
    /**
     * Use AES256 to encrypt bucket contents.
     * 
     */
    @Export(name="awsS3ServerSideEncryption", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> awsS3ServerSideEncryption;

    /**
     * @return Use AES256 to encrypt bucket contents.
     * 
     */
    public Output<Optional<Boolean>> awsS3ServerSideEncryption() {
        return Codegen.optional(this.awsS3ServerSideEncryption);
    }
    /**
     * AWS secret access key.
     * 
     */
    @Export(name="awsSecretAccessKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> awsSecretAccessKey;

    /**
     * @return AWS secret access key.
     * 
     */
    public Output<Optional<String>> awsSecretAccessKey() {
        return Codegen.optional(this.awsSecretAccessKey);
    }
    /**
     * AWS session token.
     * 
     */
    @Export(name="awsSessionToken", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> awsSessionToken;

    /**
     * @return AWS session token.
     * 
     */
    public Output<Optional<String>> awsSessionToken() {
        return Codegen.optional(this.awsSessionToken);
    }
    /**
     * Azure account key.
     * 
     */
    @Export(name="azureAccountKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azureAccountKey;

    /**
     * @return Azure account key.
     * 
     */
    public Output<Optional<String>> azureAccountKey() {
        return Codegen.optional(this.azureAccountKey);
    }
    /**
     * Azure account name.
     * 
     */
    @Export(name="azureAccountName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azureAccountName;

    /**
     * @return Azure account name.
     * 
     */
    public Output<Optional<String>> azureAccountName() {
        return Codegen.optional(this.azureAccountName);
    }
    /**
     * Azure blob environment.
     * 
     */
    @Export(name="azureBlobEnvironment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azureBlobEnvironment;

    /**
     * @return Azure blob environment.
     * 
     */
    public Output<Optional<String>> azureBlobEnvironment() {
        return Codegen.optional(this.azureBlobEnvironment);
    }
    /**
     * `&lt;required&gt;` - Azure container name to write
     * snapshots to.
     * 
     */
    @Export(name="azureContainerName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azureContainerName;

    /**
     * @return `&lt;required&gt;` - Azure container name to write
     * snapshots to.
     * 
     */
    public Output<Optional<String>> azureContainerName() {
        return Codegen.optional(this.azureContainerName);
    }
    /**
     * Azure blob storage endpoint. This is typically
     * only set when using a non-Azure implementation like Azurite.
     * 
     */
    @Export(name="azureEndpoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> azureEndpoint;

    /**
     * @return Azure blob storage endpoint. This is typically
     * only set when using a non-Azure implementation like Azurite.
     * 
     */
    public Output<Optional<String>> azureEndpoint() {
        return Codegen.optional(this.azureEndpoint);
    }
    /**
     * Within the directory or bucket
     * prefix given by `path_prefix`, the file or object name of snapshot files
     * will start with this string.
     * 
     */
    @Export(name="filePrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> filePrefix;

    /**
     * @return Within the directory or bucket
     * prefix given by `path_prefix`, the file or object name of snapshot files
     * will start with this string.
     * 
     */
    public Output<Optional<String>> filePrefix() {
        return Codegen.optional(this.filePrefix);
    }
    /**
     * Disable TLS for the GCS endpoint. This
     * should only be used for testing purposes, typically in conjunction with
     * `google_endpoint`.
     * 
     */
    @Export(name="googleDisableTls", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> googleDisableTls;

    /**
     * @return Disable TLS for the GCS endpoint. This
     * should only be used for testing purposes, typically in conjunction with
     * `google_endpoint`.
     * 
     */
    public Output<Optional<Boolean>> googleDisableTls() {
        return Codegen.optional(this.googleDisableTls);
    }
    /**
     * GCS endpoint. This is typically only set when
     * using a non-Google GCS implementation like fake-gcs-server.
     * 
     */
    @Export(name="googleEndpoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> googleEndpoint;

    /**
     * @return GCS endpoint. This is typically only set when
     * using a non-Google GCS implementation like fake-gcs-server.
     * 
     */
    public Output<Optional<String>> googleEndpoint() {
        return Codegen.optional(this.googleEndpoint);
    }
    /**
     * `&lt;required&gt;` - GCS bucket to write snapshots to.
     * 
     */
    @Export(name="googleGcsBucket", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> googleGcsBucket;

    /**
     * @return `&lt;required&gt;` - GCS bucket to write snapshots to.
     * 
     */
    public Output<Optional<String>> googleGcsBucket() {
        return Codegen.optional(this.googleGcsBucket);
    }
    /**
     * Google service account key in JSON format.
     * The raw value looks like this:
     * 
     */
    @Export(name="googleServiceAccountKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> googleServiceAccountKey;

    /**
     * @return Google service account key in JSON format.
     * The raw value looks like this:
     * 
     */
    public Output<Optional<String>> googleServiceAccountKey() {
        return Codegen.optional(this.googleServiceAccountKey);
    }
    /**
     * `&lt;required&gt;` - Time (in seconds) between snapshots.
     * 
     */
    @Export(name="intervalSeconds", refs={Integer.class}, tree="[0]")
    private Output<Integer> intervalSeconds;

    /**
     * @return `&lt;required&gt;` - Time (in seconds) between snapshots.
     * 
     */
    public Output<Integer> intervalSeconds() {
        return this.intervalSeconds;
    }
    /**
     * For `storage_type = local`, the maximum
     * space, in bytes, to use for snapshots. Snapshot attempts will fail if there is not enough
     * space left in this allowance.
     * 
     */
    @Export(name="localMaxSpace", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> localMaxSpace;

    /**
     * @return For `storage_type = local`, the maximum
     * space, in bytes, to use for snapshots. Snapshot attempts will fail if there is not enough
     * space left in this allowance.
     * 
     */
    public Output<Optional<Integer>> localMaxSpace() {
        return Codegen.optional(this.localMaxSpace);
    }
    /**
     * `&lt;required&gt;` – Name of the configuration to modify.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return `&lt;required&gt;` – Name of the configuration to modify.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * `&lt;required&gt;` - For `storage_type = &#34;local&#34;`, the directory to
     * write the snapshots in. For cloud storage types, the bucket prefix to use.
     * Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
     * Types `local` and `aws-s3` the trailing `/` is optional.
     * 
     */
    @Export(name="pathPrefix", refs={String.class}, tree="[0]")
    private Output<String> pathPrefix;

    /**
     * @return `&lt;required&gt;` - For `storage_type = &#34;local&#34;`, the directory to
     * write the snapshots in. For cloud storage types, the bucket prefix to use.
     * Types `azure-s3` and `google-gcs` require a trailing `/` (slash).
     * Types `local` and `aws-s3` the trailing `/` is optional.
     * 
     */
    public Output<String> pathPrefix() {
        return this.pathPrefix;
    }
    /**
     * How many snapshots are to be kept; when writing a
     * snapshot, if there are more snapshots already stored than this number, the
     * oldest ones will be deleted.
     * 
     */
    @Export(name="retain", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> retain;

    /**
     * @return How many snapshots are to be kept; when writing a
     * snapshot, if there are more snapshots already stored than this number, the
     * oldest ones will be deleted.
     * 
     */
    public Output<Optional<Integer>> retain() {
        return Codegen.optional(this.retain);
    }
    /**
     * `&lt;required&gt;` - One of &#34;local&#34;, &#34;azure-blob&#34;, &#34;aws-s3&#34;,
     * or &#34;google-gcs&#34;. The remaining parameters described below are all specific to
     * the selected `storage_type` and prefixed accordingly.
     * 
     */
    @Export(name="storageType", refs={String.class}, tree="[0]")
    private Output<String> storageType;

    /**
     * @return `&lt;required&gt;` - One of &#34;local&#34;, &#34;azure-blob&#34;, &#34;aws-s3&#34;,
     * or &#34;google-gcs&#34;. The remaining parameters described below are all specific to
     * the selected `storage_type` and prefixed accordingly.
     * 
     */
    public Output<String> storageType() {
        return this.storageType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RaftSnapshotAgentConfig(String name) {
        this(name, RaftSnapshotAgentConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RaftSnapshotAgentConfig(String name, RaftSnapshotAgentConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RaftSnapshotAgentConfig(String name, RaftSnapshotAgentConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig", name, args == null ? RaftSnapshotAgentConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RaftSnapshotAgentConfig(String name, Output<String> id, @Nullable RaftSnapshotAgentConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RaftSnapshotAgentConfig get(String name, Output<String> id, @Nullable RaftSnapshotAgentConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RaftSnapshotAgentConfig(name, id, state, options);
    }
}
