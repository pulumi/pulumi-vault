// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
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
    ///     var aws = new Vault.Aws.SecretBackend("aws", new()
    ///     {
    ///         AccessKey = "AKIA.....",
    ///         SecretKey = "AWS secret key",
    ///     });
    /// 
    ///     var role = new Vault.Aws.SecretBackendRole("role", new()
    ///     {
    ///         Backend = aws.Path,
    ///         Name = "deploy",
    ///         CredentialType = "iam_user",
    ///         PolicyDocument = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Effect"": ""Allow"",
    ///       ""Action"": ""iam:*"",
    ///       ""Resource"": ""*""
    ///     }
    ///   ]
    /// }
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AWS secret backend roles can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:aws/secretBackendRole:SecretBackendRole role aws/roles/deploy
    /// ```
    /// </summary>
    [VaultResourceType("vault:aws/secretBackendRole:SecretBackendRole")]
    public partial class SecretBackendRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path the AWS secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of credential to be used when
        /// retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
        /// `federation_token`.
        /// </summary>
        [Output("credentialType")]
        public Output<string> CredentialType { get; private set; } = null!;

        /// <summary>
        /// The default TTL in seconds for STS credentials.
        /// When a TTL is not specified when STS credentials are requested,
        /// and a default TTL is specified on the role,
        /// then this default TTL will be used. Valid only when `credential_type` is one of
        /// `assumed_role` or `federation_token`.
        /// </summary>
        [Output("defaultStsTtl")]
        public Output<int> DefaultStsTtl { get; private set; } = null!;

        /// <summary>
        /// External ID to set for assume role creds. 
        /// Valid only when `credential_type` is set to `assumed_role`.
        /// </summary>
        [Output("externalId")]
        public Output<string?> ExternalId { get; private set; } = null!;

        /// <summary>
        /// A list of IAM group names. IAM users generated
        /// against this vault role will be added to these IAM Groups. For a credential
        /// type of `assumed_role` or `federation_token`, the policies sent to the
        /// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
        /// policies from each group in `iam_groups` combined with the `policy_document`
        /// and `policy_arns` parameters.
        /// </summary>
        [Output("iamGroups")]
        public Output<ImmutableArray<string>> IamGroups { get; private set; } = null!;

        /// <summary>
        /// A map of strings representing key/value pairs
        /// to be used as tags for any IAM user that is created by this role.
        /// </summary>
        [Output("iamTags")]
        public Output<ImmutableDictionary<string, string>?> IamTags { get; private set; } = null!;

        /// <summary>
        /// The max allowed TTL in seconds for STS credentials
        /// (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
        /// one of `assumed_role` or `federation_token`.
        /// </summary>
        [Output("maxStsTtl")]
        public Output<int> MaxStsTtl { get; private set; } = null!;

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
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
        /// The ARN of the AWS Permissions 
        /// Boundary to attach to IAM users created in the role. Valid only when
        /// `credential_type` is `iam_user`. If not specified, then no permissions boundary
        /// policy will be attached.
        /// </summary>
        [Output("permissionsBoundaryArn")]
        public Output<string?> PermissionsBoundaryArn { get; private set; } = null!;

        /// <summary>
        /// Specifies a list of AWS managed policy ARNs. The
        /// behavior depends on the credential type. With `iam_user`, the policies will be
        /// attached to IAM users when they are requested. With `assumed_role` and
        /// `federation_token`, the policy ARNs will act as a filter on what the credentials
        /// can do, similar to `policy_document`. When `credential_type` is `iam_user` or
        /// `federation_token`, at least one of `policy_document` or `policy_arns` must
        /// be specified.
        /// </summary>
        [Output("policyArns")]
        public Output<ImmutableArray<string>> PolicyArns { get; private set; } = null!;

        /// <summary>
        /// The IAM policy document for the role. The
        /// behavior depends on the credential type. With `iam_user`, the policy document
        /// will be attached to the IAM user generated and augment the permissions the IAM
        /// user has. With `assumed_role` and `federation_token`, the policy document will
        /// act as a filter on what the credentials can do, similar to `policy_arns`.
        /// </summary>
        [Output("policyDocument")]
        public Output<string?> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// Specifies the ARNs of the AWS roles this Vault role
        /// is allowed to assume. Required when `credential_type` is `assumed_role` and
        /// prohibited otherwise.
        /// </summary>
        [Output("roleArns")]
        public Output<ImmutableArray<string>> RoleArns { get; private set; } = null!;

        /// <summary>
        /// A map of strings representing key/value pairs to be set
        /// during assume role creds creation. Valid only when `credential_type` is set to
        /// `assumed_role`.
        /// </summary>
        [Output("sessionTags")]
        public Output<ImmutableDictionary<string, string>?> SessionTags { get; private set; } = null!;

        /// <summary>
        /// The path for the user name. Valid only when 
        /// `credential_type` is `iam_user`. Default is `/`.
        /// </summary>
        [Output("userPath")]
        public Output<string?> UserPath { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:aws/secretBackendRole:SecretBackendRole", name, args ?? new SecretBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRole(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/secretBackendRole:SecretBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRole Get(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AWS secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Specifies the type of credential to be used when
        /// retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
        /// `federation_token`.
        /// </summary>
        [Input("credentialType", required: true)]
        public Input<string> CredentialType { get; set; } = null!;

        /// <summary>
        /// The default TTL in seconds for STS credentials.
        /// When a TTL is not specified when STS credentials are requested,
        /// and a default TTL is specified on the role,
        /// then this default TTL will be used. Valid only when `credential_type` is one of
        /// `assumed_role` or `federation_token`.
        /// </summary>
        [Input("defaultStsTtl")]
        public Input<int>? DefaultStsTtl { get; set; }

        /// <summary>
        /// External ID to set for assume role creds. 
        /// Valid only when `credential_type` is set to `assumed_role`.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("iamGroups")]
        private InputList<string>? _iamGroups;

        /// <summary>
        /// A list of IAM group names. IAM users generated
        /// against this vault role will be added to these IAM Groups. For a credential
        /// type of `assumed_role` or `federation_token`, the policies sent to the
        /// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
        /// policies from each group in `iam_groups` combined with the `policy_document`
        /// and `policy_arns` parameters.
        /// </summary>
        public InputList<string> IamGroups
        {
            get => _iamGroups ?? (_iamGroups = new InputList<string>());
            set => _iamGroups = value;
        }

        [Input("iamTags")]
        private InputMap<string>? _iamTags;

        /// <summary>
        /// A map of strings representing key/value pairs
        /// to be used as tags for any IAM user that is created by this role.
        /// </summary>
        public InputMap<string> IamTags
        {
            get => _iamTags ?? (_iamTags = new InputMap<string>());
            set => _iamTags = value;
        }

        /// <summary>
        /// The max allowed TTL in seconds for STS credentials
        /// (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
        /// one of `assumed_role` or `federation_token`.
        /// </summary>
        [Input("maxStsTtl")]
        public Input<int>? MaxStsTtl { get; set; }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
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
        /// The ARN of the AWS Permissions 
        /// Boundary to attach to IAM users created in the role. Valid only when
        /// `credential_type` is `iam_user`. If not specified, then no permissions boundary
        /// policy will be attached.
        /// </summary>
        [Input("permissionsBoundaryArn")]
        public Input<string>? PermissionsBoundaryArn { get; set; }

        [Input("policyArns")]
        private InputList<string>? _policyArns;

        /// <summary>
        /// Specifies a list of AWS managed policy ARNs. The
        /// behavior depends on the credential type. With `iam_user`, the policies will be
        /// attached to IAM users when they are requested. With `assumed_role` and
        /// `federation_token`, the policy ARNs will act as a filter on what the credentials
        /// can do, similar to `policy_document`. When `credential_type` is `iam_user` or
        /// `federation_token`, at least one of `policy_document` or `policy_arns` must
        /// be specified.
        /// </summary>
        public InputList<string> PolicyArns
        {
            get => _policyArns ?? (_policyArns = new InputList<string>());
            set => _policyArns = value;
        }

        /// <summary>
        /// The IAM policy document for the role. The
        /// behavior depends on the credential type. With `iam_user`, the policy document
        /// will be attached to the IAM user generated and augment the permissions the IAM
        /// user has. With `assumed_role` and `federation_token`, the policy document will
        /// act as a filter on what the credentials can do, similar to `policy_arns`.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        [Input("roleArns")]
        private InputList<string>? _roleArns;

        /// <summary>
        /// Specifies the ARNs of the AWS roles this Vault role
        /// is allowed to assume. Required when `credential_type` is `assumed_role` and
        /// prohibited otherwise.
        /// </summary>
        public InputList<string> RoleArns
        {
            get => _roleArns ?? (_roleArns = new InputList<string>());
            set => _roleArns = value;
        }

        [Input("sessionTags")]
        private InputMap<string>? _sessionTags;

        /// <summary>
        /// A map of strings representing key/value pairs to be set
        /// during assume role creds creation. Valid only when `credential_type` is set to
        /// `assumed_role`.
        /// </summary>
        public InputMap<string> SessionTags
        {
            get => _sessionTags ?? (_sessionTags = new InputMap<string>());
            set => _sessionTags = value;
        }

        /// <summary>
        /// The path for the user name. Valid only when 
        /// `credential_type` is `iam_user`. Default is `/`.
        /// </summary>
        [Input("userPath")]
        public Input<string>? UserPath { get; set; }

        public SecretBackendRoleArgs()
        {
        }
        public static new SecretBackendRoleArgs Empty => new SecretBackendRoleArgs();
    }

    public sealed class SecretBackendRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AWS secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Specifies the type of credential to be used when
        /// retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
        /// `federation_token`.
        /// </summary>
        [Input("credentialType")]
        public Input<string>? CredentialType { get; set; }

        /// <summary>
        /// The default TTL in seconds for STS credentials.
        /// When a TTL is not specified when STS credentials are requested,
        /// and a default TTL is specified on the role,
        /// then this default TTL will be used. Valid only when `credential_type` is one of
        /// `assumed_role` or `federation_token`.
        /// </summary>
        [Input("defaultStsTtl")]
        public Input<int>? DefaultStsTtl { get; set; }

        /// <summary>
        /// External ID to set for assume role creds. 
        /// Valid only when `credential_type` is set to `assumed_role`.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("iamGroups")]
        private InputList<string>? _iamGroups;

        /// <summary>
        /// A list of IAM group names. IAM users generated
        /// against this vault role will be added to these IAM Groups. For a credential
        /// type of `assumed_role` or `federation_token`, the policies sent to the
        /// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
        /// policies from each group in `iam_groups` combined with the `policy_document`
        /// and `policy_arns` parameters.
        /// </summary>
        public InputList<string> IamGroups
        {
            get => _iamGroups ?? (_iamGroups = new InputList<string>());
            set => _iamGroups = value;
        }

        [Input("iamTags")]
        private InputMap<string>? _iamTags;

        /// <summary>
        /// A map of strings representing key/value pairs
        /// to be used as tags for any IAM user that is created by this role.
        /// </summary>
        public InputMap<string> IamTags
        {
            get => _iamTags ?? (_iamTags = new InputMap<string>());
            set => _iamTags = value;
        }

        /// <summary>
        /// The max allowed TTL in seconds for STS credentials
        /// (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
        /// one of `assumed_role` or `federation_token`.
        /// </summary>
        [Input("maxStsTtl")]
        public Input<int>? MaxStsTtl { get; set; }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
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
        /// The ARN of the AWS Permissions 
        /// Boundary to attach to IAM users created in the role. Valid only when
        /// `credential_type` is `iam_user`. If not specified, then no permissions boundary
        /// policy will be attached.
        /// </summary>
        [Input("permissionsBoundaryArn")]
        public Input<string>? PermissionsBoundaryArn { get; set; }

        [Input("policyArns")]
        private InputList<string>? _policyArns;

        /// <summary>
        /// Specifies a list of AWS managed policy ARNs. The
        /// behavior depends on the credential type. With `iam_user`, the policies will be
        /// attached to IAM users when they are requested. With `assumed_role` and
        /// `federation_token`, the policy ARNs will act as a filter on what the credentials
        /// can do, similar to `policy_document`. When `credential_type` is `iam_user` or
        /// `federation_token`, at least one of `policy_document` or `policy_arns` must
        /// be specified.
        /// </summary>
        public InputList<string> PolicyArns
        {
            get => _policyArns ?? (_policyArns = new InputList<string>());
            set => _policyArns = value;
        }

        /// <summary>
        /// The IAM policy document for the role. The
        /// behavior depends on the credential type. With `iam_user`, the policy document
        /// will be attached to the IAM user generated and augment the permissions the IAM
        /// user has. With `assumed_role` and `federation_token`, the policy document will
        /// act as a filter on what the credentials can do, similar to `policy_arns`.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        [Input("roleArns")]
        private InputList<string>? _roleArns;

        /// <summary>
        /// Specifies the ARNs of the AWS roles this Vault role
        /// is allowed to assume. Required when `credential_type` is `assumed_role` and
        /// prohibited otherwise.
        /// </summary>
        public InputList<string> RoleArns
        {
            get => _roleArns ?? (_roleArns = new InputList<string>());
            set => _roleArns = value;
        }

        [Input("sessionTags")]
        private InputMap<string>? _sessionTags;

        /// <summary>
        /// A map of strings representing key/value pairs to be set
        /// during assume role creds creation. Valid only when `credential_type` is set to
        /// `assumed_role`.
        /// </summary>
        public InputMap<string> SessionTags
        {
            get => _sessionTags ?? (_sessionTags = new InputMap<string>());
            set => _sessionTags = value;
        }

        /// <summary>
        /// The path for the user name. Valid only when 
        /// `credential_type` is `iam_user`. Default is `/`.
        /// </summary>
        [Input("userPath")]
        public Input<string>? UserPath { get; set; }

        public SecretBackendRoleState()
        {
        }
        public static new SecretBackendRoleState Empty => new SecretBackendRoleState();
    }
}
