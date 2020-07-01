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
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Vault.Mount("example", new Vault.MountArgs
    ///         {
    ///             Description = "This is an example mount",
    ///             Path = "dummy",
    ///             Type = "generic",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Mount : Pulumi.CustomResource
    {
        /// <summary>
        /// The accessor for this mount.
        /// </summary>
        [Output("accessor")]
        public Output<string> Accessor { get; private set; } = null!;

        /// <summary>
        /// Default lease duration for tokens and secrets in seconds
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description of the mount
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        /// <summary>
        /// Maximum possible lease duration for tokens and secrets in seconds
        /// </summary>
        [Output("maxLeaseTtlSeconds")]
        public Output<int> MaxLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// Specifies mount type specific options that are passed to the backend
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, object>?> Options { get; private set; } = null!;

        /// <summary>
        /// Where the secret backend will be mounted
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        /// </summary>
        [Output("sealWrap")]
        public Output<bool> SealWrap { get; private set; } = null!;

        /// <summary>
        /// Type of the backend, such as "aws"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Mount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Mount(string name, MountArgs args, CustomResourceOptions? options = null)
            : base("vault:index/mount:Mount", name, args ?? new MountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Mount(string name, Input<string> id, MountState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/mount:Mount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Mount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Mount Get(string name, Input<string> id, MountState? state = null, CustomResourceOptions? options = null)
        {
            return new Mount(name, id, state, options);
        }
    }

    public sealed class MountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default lease duration for tokens and secrets in seconds
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Human-friendly description of the mount
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum possible lease duration for tokens and secrets in seconds
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        [Input("options")]
        private InputMap<object>? _options;

        /// <summary>
        /// Specifies mount type specific options that are passed to the backend
        /// </summary>
        public InputMap<object> Options
        {
            get => _options ?? (_options = new InputMap<object>());
            set => _options = value;
        }

        /// <summary>
        /// Where the secret backend will be mounted
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        /// </summary>
        [Input("sealWrap")]
        public Input<bool>? SealWrap { get; set; }

        /// <summary>
        /// Type of the backend, such as "aws"
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MountArgs()
        {
        }
    }

    public sealed class MountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accessor for this mount.
        /// </summary>
        [Input("accessor")]
        public Input<string>? Accessor { get; set; }

        /// <summary>
        /// Default lease duration for tokens and secrets in seconds
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Human-friendly description of the mount
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum possible lease duration for tokens and secrets in seconds
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        [Input("options")]
        private InputMap<object>? _options;

        /// <summary>
        /// Specifies mount type specific options that are passed to the backend
        /// </summary>
        public InputMap<object> Options
        {
            get => _options ?? (_options = new InputMap<object>());
            set => _options = value;
        }

        /// <summary>
        /// Where the secret backend will be mounted
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        /// </summary>
        [Input("sealWrap")]
        public Input<bool>? SealWrap { get; set; }

        /// <summary>
        /// Type of the backend, such as "aws"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MountState()
        {
        }
    }
}
