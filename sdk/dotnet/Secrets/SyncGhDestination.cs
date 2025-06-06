// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    ///     var gh = new Vault.Secrets.SyncGhDestination("gh", new()
    ///     {
    ///         Name = "gh-dest",
    ///         AccessToken = accessToken,
    ///         RepositoryOwner = repoOwner,
    ///         RepositoryName = "repo-name-example",
    ///         SecretNameTemplate = "vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub Secrets sync destinations can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:secrets/syncGhDestination:SyncGhDestination gh gh-dest
    /// ```
    /// </summary>
    [VaultResourceType("vault:secrets/syncGhDestination:SyncGhDestination")]
    public partial class SyncGhDestination : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Fine-grained or personal access token.
        /// Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
        /// variable.
        /// </summary>
        [Output("accessToken")]
        public Output<string?> AccessToken { get; private set; } = null!;

        /// <summary>
        /// The user-defined name of the GitHub App configuration. This is a reference to the name used   
        /// on the new endpoint when configuring the GitHub app on the Vault Server. Can be modified.
        /// Takes precedence over the `access_token` field.
        /// </summary>
        [Output("appName")]
        public Output<string?> AppName { get; private set; } = null!;

        /// <summary>
        /// Determines what level of information is synced as a distinct resource
        /// at the destination. Supports `secret-path` and `secret-key`.
        /// </summary>
        [Output("granularity")]
        public Output<string?> Granularity { get; private set; } = null!;

        /// <summary>
        /// The ID of the installation generated by GitHub when the app referenced by the `app_name` 
        /// was installed in the user’s GitHub account. Can be modified. Necessary if the `app_name` field is also provided.
        /// </summary>
        [Output("installationId")]
        public Output<int?> InstallationId { get; private set; } = null!;

        /// <summary>
        /// Unique name of the GitHub destination.
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
        /// Name of the repository.
        /// Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
        /// variable.
        /// </summary>
        [Output("repositoryName")]
        public Output<string?> RepositoryName { get; private set; } = null!;

        /// <summary>
        /// GitHub organization or username that owns the repository.
        /// Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
        /// variable.
        /// </summary>
        [Output("repositoryOwner")]
        public Output<string?> RepositoryOwner { get; private set; } = null!;

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Output("secretNameTemplate")]
        public Output<string> SecretNameTemplate { get; private set; } = null!;

        /// <summary>
        /// The type of the secrets destination (`gh`).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SyncGhDestination resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncGhDestination(string name, SyncGhDestinationArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:secrets/syncGhDestination:SyncGhDestination", name, args ?? new SyncGhDestinationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncGhDestination(string name, Input<string> id, SyncGhDestinationState? state = null, CustomResourceOptions? options = null)
            : base("vault:secrets/syncGhDestination:SyncGhDestination", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accessToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SyncGhDestination resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncGhDestination Get(string name, Input<string> id, SyncGhDestinationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncGhDestination(name, id, state, options);
        }
    }

    public sealed class SyncGhDestinationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        private Input<string>? _accessToken;

        /// <summary>
        /// Fine-grained or personal access token.
        /// Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
        /// variable.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The user-defined name of the GitHub App configuration. This is a reference to the name used   
        /// on the new endpoint when configuring the GitHub app on the Vault Server. Can be modified.
        /// Takes precedence over the `access_token` field.
        /// </summary>
        [Input("appName")]
        public Input<string>? AppName { get; set; }

        /// <summary>
        /// Determines what level of information is synced as a distinct resource
        /// at the destination. Supports `secret-path` and `secret-key`.
        /// </summary>
        [Input("granularity")]
        public Input<string>? Granularity { get; set; }

        /// <summary>
        /// The ID of the installation generated by GitHub when the app referenced by the `app_name` 
        /// was installed in the user’s GitHub account. Can be modified. Necessary if the `app_name` field is also provided.
        /// </summary>
        [Input("installationId")]
        public Input<int>? InstallationId { get; set; }

        /// <summary>
        /// Unique name of the GitHub destination.
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
        /// Name of the repository.
        /// Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
        /// variable.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        /// <summary>
        /// GitHub organization or username that owns the repository.
        /// Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
        /// variable.
        /// </summary>
        [Input("repositoryOwner")]
        public Input<string>? RepositoryOwner { get; set; }

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Input("secretNameTemplate")]
        public Input<string>? SecretNameTemplate { get; set; }

        public SyncGhDestinationArgs()
        {
        }
        public static new SyncGhDestinationArgs Empty => new SyncGhDestinationArgs();
    }

    public sealed class SyncGhDestinationState : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        private Input<string>? _accessToken;

        /// <summary>
        /// Fine-grained or personal access token.
        /// Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
        /// variable.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The user-defined name of the GitHub App configuration. This is a reference to the name used   
        /// on the new endpoint when configuring the GitHub app on the Vault Server. Can be modified.
        /// Takes precedence over the `access_token` field.
        /// </summary>
        [Input("appName")]
        public Input<string>? AppName { get; set; }

        /// <summary>
        /// Determines what level of information is synced as a distinct resource
        /// at the destination. Supports `secret-path` and `secret-key`.
        /// </summary>
        [Input("granularity")]
        public Input<string>? Granularity { get; set; }

        /// <summary>
        /// The ID of the installation generated by GitHub when the app referenced by the `app_name` 
        /// was installed in the user’s GitHub account. Can be modified. Necessary if the `app_name` field is also provided.
        /// </summary>
        [Input("installationId")]
        public Input<int>? InstallationId { get; set; }

        /// <summary>
        /// Unique name of the GitHub destination.
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
        /// Name of the repository.
        /// Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
        /// variable.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        /// <summary>
        /// GitHub organization or username that owns the repository.
        /// Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
        /// variable.
        /// </summary>
        [Input("repositoryOwner")]
        public Input<string>? RepositoryOwner { get; set; }

        /// <summary>
        /// Template describing how to generate external secret names.
        /// Supports a subset of the Go Template syntax.
        /// </summary>
        [Input("secretNameTemplate")]
        public Input<string>? SecretNameTemplate { get; set; }

        /// <summary>
        /// The type of the secrets destination (`gh`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SyncGhDestinationState()
        {
        }
        public static new SyncGhDestinationState Empty => new SyncGhDestinationState();
    }
}
