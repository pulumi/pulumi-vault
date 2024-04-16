// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.AD
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Vault.AD.SecretBackend("config", new()
    ///     {
    ///         Backend = "ad",
    ///         Binddn = "CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
    ///         Bindpass = "SuperSecretPassw0rd",
    ///         Url = "ldaps://ad",
    ///         InsecureTls = true,
    ///         Userdn = "CN=Users,DC=corp,DC=example,DC=net",
    ///     });
    /// 
    ///     var qa = new Vault.AD.SecretLibrary("qa", new()
    ///     {
    ///         Backend = config.Backend,
    ///         Name = "qa",
    ///         ServiceAccountNames = new[]
    ///         {
    ///             "Bob",
    ///             "Mary",
    ///         },
    ///         Ttl = 60,
    ///         DisableCheckInEnforcement = true,
    ///         MaxTtl = 120,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// AD secret backend libraries can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:ad/secretLibrary:SecretLibrary role ad/library/bob
    /// ```
    /// </summary>
    [VaultResourceType("vault:ad/secretLibrary:SecretLibrary")]
    public partial class SecretLibrary : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path the AD secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
        /// </summary>
        [Output("disableCheckInEnforcement")]
        public Output<bool?> DisableCheckInEnforcement { get; private set; } = null!;

        /// <summary>
        /// The maximum password time-to-live in seconds. Defaults to the configuration
        /// max_ttl if not provided.
        /// </summary>
        [Output("maxTtl")]
        public Output<int> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// The name to identify this set of service accounts.
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
        /// Specifies the slice of service accounts mapped to this set.
        /// </summary>
        [Output("serviceAccountNames")]
        public Output<ImmutableArray<string>> ServiceAccountNames { get; private set; } = null!;

        /// <summary>
        /// The password time-to-live in seconds. Defaults to the configuration
        /// ttl if not provided.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a SecretLibrary resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretLibrary(string name, SecretLibraryArgs args, CustomResourceOptions? options = null)
            : base("vault:ad/secretLibrary:SecretLibrary", name, args ?? new SecretLibraryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretLibrary(string name, Input<string> id, SecretLibraryState? state = null, CustomResourceOptions? options = null)
            : base("vault:ad/secretLibrary:SecretLibrary", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretLibrary resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretLibrary Get(string name, Input<string> id, SecretLibraryState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretLibrary(name, id, state, options);
        }
    }

    public sealed class SecretLibraryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AD secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
        /// </summary>
        [Input("disableCheckInEnforcement")]
        public Input<bool>? DisableCheckInEnforcement { get; set; }

        /// <summary>
        /// The maximum password time-to-live in seconds. Defaults to the configuration
        /// max_ttl if not provided.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The name to identify this set of service accounts.
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

        [Input("serviceAccountNames", required: true)]
        private InputList<string>? _serviceAccountNames;

        /// <summary>
        /// Specifies the slice of service accounts mapped to this set.
        /// </summary>
        public InputList<string> ServiceAccountNames
        {
            get => _serviceAccountNames ?? (_serviceAccountNames = new InputList<string>());
            set => _serviceAccountNames = value;
        }

        /// <summary>
        /// The password time-to-live in seconds. Defaults to the configuration
        /// ttl if not provided.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretLibraryArgs()
        {
        }
        public static new SecretLibraryArgs Empty => new SecretLibraryArgs();
    }

    public sealed class SecretLibraryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AD secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
        /// </summary>
        [Input("disableCheckInEnforcement")]
        public Input<bool>? DisableCheckInEnforcement { get; set; }

        /// <summary>
        /// The maximum password time-to-live in seconds. Defaults to the configuration
        /// max_ttl if not provided.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// The name to identify this set of service accounts.
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

        [Input("serviceAccountNames")]
        private InputList<string>? _serviceAccountNames;

        /// <summary>
        /// Specifies the slice of service accounts mapped to this set.
        /// </summary>
        public InputList<string> ServiceAccountNames
        {
            get => _serviceAccountNames ?? (_serviceAccountNames = new InputList<string>());
            set => _serviceAccountNames = value;
        }

        /// <summary>
        /// The password time-to-live in seconds. Defaults to the configuration
        /// ttl if not provided.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretLibraryState()
        {
        }
        public static new SecretLibraryState Empty => new SecretLibraryState();
    }
}
