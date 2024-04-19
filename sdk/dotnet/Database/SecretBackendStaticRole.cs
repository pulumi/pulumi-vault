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
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var db = new Vault.Mount("db", new()
    ///     {
    ///         Path = "postgres",
    ///         Type = "database",
    ///     });
    /// 
    ///     var postgres = new Vault.Database.SecretBackendConnection("postgres", new()
    ///     {
    ///         Backend = db.Path,
    ///         AllowedRoles = new[]
    ///         {
    ///             "*",
    ///         },
    ///         Postgresql = new Vault.Database.Inputs.SecretBackendConnectionPostgresqlArgs
    ///         {
    ///             ConnectionUrl = "postgres://username:password@host:port/database",
    ///         },
    ///     });
    /// 
    ///     // configure a static role with period-based rotations
    ///     var periodRole = new Vault.Database.SecretBackendStaticRole("periodRole", new()
    ///     {
    ///         Backend = db.Path,
    ///         DbName = postgres.Name,
    ///         Username = "example",
    ///         RotationPeriod = 3600,
    ///         RotationStatements = new[]
    ///         {
    ///             "ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';",
    ///         },
    ///     });
    /// 
    ///     // configure a static role with schedule-based rotations
    ///     var scheduleRole = new Vault.Database.SecretBackendStaticRole("scheduleRole", new()
    ///     {
    ///         Backend = db.Path,
    ///         DbName = postgres.Name,
    ///         Username = "example",
    ///         RotationSchedule = "0 0 * * SAT",
    ///         RotationWindow = 172800,
    ///         RotationStatements = new[]
    ///         {
    ///             "ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database secret backend static roles can be imported using the `backend`, `/static-roles/`, and the `name` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:database/secretBackendStaticRole:SecretBackendStaticRole example postgres/static-roles/my-role
    /// ```
    /// </summary>
    [VaultResourceType("vault:database/secretBackendStaticRole:SecretBackendStaticRole")]
    public partial class SecretBackendStaticRole : global::Pulumi.CustomResource
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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The amount of time Vault should wait before rotating the password, in seconds.
        /// Mutually exclusive with `rotation_schedule`.
        /// </summary>
        [Output("rotationPeriod")]
        public Output<int?> RotationPeriod { get; private set; } = null!;

        /// <summary>
        /// A cron-style string that will define the schedule on which rotations should occur.
        /// Mutually exclusive with `rotation_period`.
        /// 
        /// **Warning**: The `rotation_period` and `rotation_schedule` fields are
        /// mutually exclusive. One of them must be set but not both.
        /// </summary>
        [Output("rotationSchedule")]
        public Output<string?> RotationSchedule { get; private set; } = null!;

        /// <summary>
        /// Database statements to execute to rotate the password for the configured database user.
        /// </summary>
        [Output("rotationStatements")]
        public Output<ImmutableArray<string>> RotationStatements { get; private set; } = null!;

        /// <summary>
        /// The amount of time, in seconds, in which rotations are allowed to occur starting
        /// from a given `rotation_schedule`.
        /// </summary>
        [Output("rotationWindow")]
        public Output<int?> RotationWindow { get; private set; } = null!;

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

    public sealed class SecretBackendStaticRoleArgs : global::Pulumi.ResourceArgs
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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The amount of time Vault should wait before rotating the password, in seconds.
        /// Mutually exclusive with `rotation_schedule`.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// A cron-style string that will define the schedule on which rotations should occur.
        /// Mutually exclusive with `rotation_period`.
        /// 
        /// **Warning**: The `rotation_period` and `rotation_schedule` fields are
        /// mutually exclusive. One of them must be set but not both.
        /// </summary>
        [Input("rotationSchedule")]
        public Input<string>? RotationSchedule { get; set; }

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
        /// The amount of time, in seconds, in which rotations are allowed to occur starting
        /// from a given `rotation_schedule`.
        /// </summary>
        [Input("rotationWindow")]
        public Input<int>? RotationWindow { get; set; }

        /// <summary>
        /// The database username that this static role corresponds to.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SecretBackendStaticRoleArgs()
        {
        }
        public static new SecretBackendStaticRoleArgs Empty => new SecretBackendStaticRoleArgs();
    }

    public sealed class SecretBackendStaticRoleState : global::Pulumi.ResourceArgs
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
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The amount of time Vault should wait before rotating the password, in seconds.
        /// Mutually exclusive with `rotation_schedule`.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// A cron-style string that will define the schedule on which rotations should occur.
        /// Mutually exclusive with `rotation_period`.
        /// 
        /// **Warning**: The `rotation_period` and `rotation_schedule` fields are
        /// mutually exclusive. One of them must be set but not both.
        /// </summary>
        [Input("rotationSchedule")]
        public Input<string>? RotationSchedule { get; set; }

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
        /// The amount of time, in seconds, in which rotations are allowed to occur starting
        /// from a given `rotation_schedule`.
        /// </summary>
        [Input("rotationWindow")]
        public Input<int>? RotationWindow { get; set; }

        /// <summary>
        /// The database username that this static role corresponds to.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SecretBackendStaticRoleState()
        {
        }
        public static new SecretBackendStaticRoleState Empty => new SecretBackendStaticRoleState();
    }
}
