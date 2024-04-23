// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Secrets
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var aws = new Vault.Secrets.SyncAwsDestination("aws", new()
    ///     {
    ///         Name = "aws-dest",
    ///         AccessKeyId = accessKeyId,
    ///         SecretAccessKey = secretAccessKey,
    ///         Region = "us-east-1",
    ///         RoleArn = "role-arn",
    ///         ExternalId = "external-id",
    ///         SecretNameTemplate = "vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
    ///         CustomTags = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AWS Secrets sync destinations can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:secrets/syncAwsDestination:SyncAwsDestination aws aws-dest
    /// ```
    /// </summary>
    [VaultResourceType("vault:secrets/syncAwsDestination:SyncAwsDestination")]
    public partial class SyncAwsDestination : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access key id to authenticate against the AWS secrets manager.
        /// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
        /// variable.
        /// </summary>
        [Output("accessKeyId")]
        public Output<string?> AccessKeyId { get; private set; } = null!;

        /// <summary>
        /// Custom tags to set on the secret managed at the destination.
        /// </summary>
        [Output("customTags")]
        public Output<ImmutableDictionary<string, object>?> CustomTags { get; private set; } = null!;

        /// <summary>
        /// Optional extra protection that must match the trust policy granting access to the
        /// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
        /// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
        /// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
        /// denied errors. Ignored if the `role_arn` field is empty.
        /// </summary>
        [Output("externalId")]
        public Output<string?> ExternalId { get; private set; } = null!;

        /// <summary>
        /// Determines what level of information is synced as a distinct resource 
        /// at the destination. Supports `secret-path` and `secret-key`.
        /// </summary>
        [Output("granularity")]
        public Output<string?> Granularity { get; private set; } = null!;

        /// <summary>
        /// Unique name of the AWS destination.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Region where to manage the secrets manager entries.
        /// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
        /// variable.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// Specifies a role to assume when connecting to AWS. When assuming a role, 
        /// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
        /// exist for Vault to be able to assume this role. The role can be in a different account.
        /// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
        /// It is possible to provide both an access key pair and a role to assume.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Secret access key to authenticate against the AWS secrets manager.
        /// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
        /// variable.
        /// </summary>
        [Output("secretAccessKey")]
        public Output<string?> SecretAccessKey { get; private set; } = null!;

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Output("secretNameTemplate")]
        public Output<string> SecretNameTemplate { get; private set; } = null!;

        /// <summary>
        /// The type of the secrets destination (`aws-sm`).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SyncAwsDestination resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncAwsDestination(string name, SyncAwsDestinationArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:secrets/syncAwsDestination:SyncAwsDestination", name, args ?? new SyncAwsDestinationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncAwsDestination(string name, Input<string> id, SyncAwsDestinationState? state = null, CustomResourceOptions? options = null)
            : base("vault:secrets/syncAwsDestination:SyncAwsDestination", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secretAccessKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SyncAwsDestination resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncAwsDestination Get(string name, Input<string> id, SyncAwsDestinationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncAwsDestination(name, id, state, options);
        }
    }

    public sealed class SyncAwsDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access key id to authenticate against the AWS secrets manager.
        /// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
        /// variable.
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        [Input("customTags")]
        private InputMap<object>? _customTags;

        /// <summary>
        /// Custom tags to set on the secret managed at the destination.
        /// </summary>
        public InputMap<object> CustomTags
        {
            get => _customTags ?? (_customTags = new InputMap<object>());
            set => _customTags = value;
        }

        /// <summary>
        /// Optional extra protection that must match the trust policy granting access to the
        /// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
        /// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
        /// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
        /// denied errors. Ignored if the `role_arn` field is empty.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Determines what level of information is synced as a distinct resource 
        /// at the destination. Supports `secret-path` and `secret-key`.
        /// </summary>
        [Input("granularity")]
        public Input<string>? Granularity { get; set; }

        /// <summary>
        /// Unique name of the AWS destination.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Region where to manage the secrets manager entries.
        /// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
        /// variable.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies a role to assume when connecting to AWS. When assuming a role, 
        /// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
        /// exist for Vault to be able to assume this role. The role can be in a different account.
        /// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
        /// It is possible to provide both an access key pair and a role to assume.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("secretAccessKey")]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// Secret access key to authenticate against the AWS secrets manager.
        /// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
        /// variable.
        /// </summary>
        public Input<string>? SecretAccessKey
        {
            get => _secretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Input("secretNameTemplate")]
        public Input<string>? SecretNameTemplate { get; set; }

        public SyncAwsDestinationArgs()
        {
        }
        public static new SyncAwsDestinationArgs Empty => new SyncAwsDestinationArgs();
    }

    public sealed class SyncAwsDestinationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access key id to authenticate against the AWS secrets manager.
        /// Can be omitted and directly provided to Vault using the `AWS_ACCESS_KEY_ID` environment
        /// variable.
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        [Input("customTags")]
        private InputMap<object>? _customTags;

        /// <summary>
        /// Custom tags to set on the secret managed at the destination.
        /// </summary>
        public InputMap<object> CustomTags
        {
            get => _customTags ?? (_customTags = new InputMap<object>());
            set => _customTags = value;
        }

        /// <summary>
        /// Optional extra protection that must match the trust policy granting access to the
        /// AWS IAM role ARN. We recommend using a different random UUID per destination. The value is generated by users.
        /// The field is mutable with no special condition, but users must be careful that the new value fits with the trust
        /// relationship condition they set on AWS otherwise sync operations will start to fail due to client-side access
        /// denied errors. Ignored if the `role_arn` field is empty.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Determines what level of information is synced as a distinct resource 
        /// at the destination. Supports `secret-path` and `secret-key`.
        /// </summary>
        [Input("granularity")]
        public Input<string>? Granularity { get; set; }

        /// <summary>
        /// Unique name of the AWS destination.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Region where to manage the secrets manager entries.
        /// Can be omitted and directly provided to Vault using the `AWS_REGION` environment
        /// variable.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies a role to assume when connecting to AWS. When assuming a role, 
        /// Vault uses temporary STS credentials to authenticate. An initial session with the proper trust relationship must
        /// exist for Vault to be able to assume this role. The role can be in a different account.
        /// The value is mutable as long as the new role targets the same AWS account ID. If not, the BE will return an error.
        /// It is possible to provide both an access key pair and a role to assume.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("secretAccessKey")]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// Secret access key to authenticate against the AWS secrets manager.
        /// Can be omitted and directly provided to Vault using the `AWS_SECRET_ACCESS_KEY` environment
        /// variable.
        /// </summary>
        public Input<string>? SecretAccessKey
        {
            get => _secretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Input("secretNameTemplate")]
        public Input<string>? SecretNameTemplate { get; set; }

        /// <summary>
        /// The type of the secrets destination (`aws-sm`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SyncAwsDestinationState()
        {
        }
        public static new SyncAwsDestinationState Empty => new SyncAwsDestinationState();
    }
}
