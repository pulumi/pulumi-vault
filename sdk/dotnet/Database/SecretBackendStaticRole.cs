// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database
{
    /// <summary>
    /// Creates a Database Secret Backend static role in Vault. Database secret backend
    /// static roles can be used to manage 1-to-1 mapping of a Vault Role to a user in a
    /// database for the database.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var db = new Vault.Mount("db", new Vault.MountArgs
    ///         {
    ///             Path = "postgres",
    ///             Type = "database",
    ///         });
    ///         var postgres = new Vault.Database.SecretBackendConnection("postgres", new Vault.Database.SecretBackendConnectionArgs
    ///         {
    ///             AllowedRoles = 
    ///             {
    ///                 "*",
    ///             },
    ///             Backend = db.Path,
    ///             Postgresql = new Vault.Database.Inputs.SecretBackendConnectionPostgresqlArgs
    ///             {
    ///                 ConnectionUrl = "postgres://username:password@host:port/database",
    ///             },
    ///         });
    ///         var staticRole = new Vault.Database.SecretBackendStaticRole("staticRole", new Vault.Database.SecretBackendStaticRoleArgs
    ///         {
    ///             Backend = db.Path,
    ///             DbName = postgres.Name,
    ///             RotationPeriod = 3600,
    ///             RotationStatements = 
    ///             {
    ///                 "ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';",
    ///             },
    ///             Username = "example",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SecretBackendStaticRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The unique name of the Vault mount to configure.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The unique name of the database connection to use for the static role.
        /// </summary>
        [Output("dbName")]
        public Output<string> DbName { get; private set; } = null!;

        /// <summary>
        /// A unique name to give the static role.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The amount of time Vault should wait before rotating the password, in seconds.
        /// </summary>
        [Output("rotationPeriod")]
        public Output<int> RotationPeriod { get; private set; } = null!;

        /// <summary>
        /// Database statements to execute to rotate the password for the configured database user.
        /// </summary>
        [Output("rotationStatements")]
        public Output<ImmutableArray<string>> RotationStatements { get; private set; } = null!;

        /// <summary>
        /// The database username that this static role corresponds to.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendStaticRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendStaticRole(string name, SecretBackendStaticRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, args ?? new SecretBackendStaticRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendStaticRole(string name, Input<string> id, SecretBackendStaticRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:database/secretBackendStaticRole:SecretBackendStaticRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendStaticRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendStaticRole Get(string name, Input<string> id, SecretBackendStaticRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendStaticRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendStaticRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the Vault mount to configure.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The unique name of the database connection to use for the static role.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// A unique name to give the static role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The amount of time Vault should wait before rotating the password, in seconds.
        /// </summary>
        [Input("rotationPeriod", required: true)]
        public Input<int> RotationPeriod { get; set; } = null!;

        [Input("rotationStatements")]
        private InputList<string>? _rotationStatements;

        /// <summary>
        /// Database statements to execute to rotate the password for the configured database user.
        /// </summary>
        public InputList<string> RotationStatements
        {
            get => _rotationStatements ?? (_rotationStatements = new InputList<string>());
            set => _rotationStatements = value;
        }

        /// <summary>
        /// The database username that this static role corresponds to.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SecretBackendStaticRoleArgs()
        {
        }
    }

    public sealed class SecretBackendStaticRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the Vault mount to configure.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The unique name of the database connection to use for the static role.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// A unique name to give the static role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The amount of time Vault should wait before rotating the password, in seconds.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        [Input("rotationStatements")]
        private InputList<string>? _rotationStatements;

        /// <summary>
        /// Database statements to execute to rotate the password for the configured database user.
        /// </summary>
        public InputList<string> RotationStatements
        {
            get => _rotationStatements ?? (_rotationStatements = new InputList<string>());
            set => _rotationStatements = value;
        }

        /// <summary>
        /// The database username that this static role corresponds to.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SecretBackendStaticRoleState()
        {
        }
    }
}
