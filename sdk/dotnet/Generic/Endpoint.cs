// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Generic
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
    ///     var userpass = new Vault.AuthBackend("userpass", new()
    ///     {
    ///         Type = "userpass",
    ///     });
    /// 
    ///     var u1 = new Vault.Generic.Endpoint("u1", new()
    ///     {
    ///         Path = "auth/userpass/users/u1",
    ///         IgnoreAbsentFields = true,
    ///         DataJson = @"{
    ///   ""policies"": [""p1""],
    ///   ""password"": ""changeme""
    /// }
    /// ",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             userpass, 
    ///         },
    ///     });
    /// 
    ///     var u1Token = new Vault.Generic.Endpoint("u1Token", new()
    ///     {
    ///         Path = "auth/userpass/login/u1",
    ///         DisableRead = true,
    ///         DisableDelete = true,
    ///         DataJson = @"{
    ///   ""password"": ""changeme""
    /// }
    /// ",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             u1, 
    ///         },
    ///     });
    /// 
    ///     var u1Entity = new Vault.Generic.Endpoint("u1Entity", new()
    ///     {
    ///         DisableRead = true,
    ///         DisableDelete = true,
    ///         Path = "identity/lookup/entity",
    ///         IgnoreAbsentFields = true,
    ///         WriteFields = new[]
    ///         {
    ///             "id",
    ///         },
    ///         DataJson = @"{
    ///   ""alias_name"": ""u1"",
    ///   ""alias_mount_accessor"": vault_auth_backend.userpass.accessor
    /// }
    /// ",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             u1Token, 
    ///         },
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["u1Id"] = u1Entity.WriteData.Apply(writeData =&gt; writeData.Id),
    ///     };
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Required Vault Capabilities
    /// 
    /// Use of this resource requires the `create` or `update` capability
    /// (depending on whether the resource already exists) on the given path. If
    /// `disable_delete` is false, the `delete` capability is also required. If
    /// `disable_read` is false, the `read` capability is required.
    /// 
    /// ## Import
    /// 
    /// Import is not supported for this resource.
    /// </summary>
    [VaultResourceType("vault:generic/endpoint:Endpoint")]
    public partial class Endpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written to the given path as the secret data.
        /// </summary>
        [Output("dataJson")]
        public Output<string> DataJson { get; private set; } = null!;

        /// <summary>
        /// - (Optional) True/false. Set this to true if your
        /// vault authentication is not able to delete the data or if the endpoint
        /// does not support the `DELETE` method. Defaults to false.
        /// </summary>
        [Output("disableDelete")]
        public Output<bool?> DisableDelete { get; private set; } = null!;

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data or if the endpoint does
        /// not support the `GET` method. Setting this to `true` will break drift
        /// detection. You should set this to `true` for endpoints that are
        /// write-only. Defaults to false.
        /// </summary>
        [Output("disableRead")]
        public Output<bool?> DisableRead { get; private set; } = null!;

        /// <summary>
        /// - (Optional) True/false. If set to true,
        /// ignore any fields present when the endpoint is read but that were not
        /// in `data_json`. Also, if a field that was written is not returned when
        /// the endpoint is read, treat that field as being up to date. You should
        /// set this to `true` when writing to endpoint that, when read, returns a
        /// different set of fields from the ones you wrote, as is common with
        /// many configuration endpoints. Defaults to false.
        /// </summary>
        [Output("ignoreAbsentFields")]
        public Output<bool?> IgnoreAbsentFields { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The full logical path at which to write the given
        /// data. Consult each backend's documentation to see which endpoints
        /// support the `PUT` methods and to determine whether they also support
        /// `DELETE` and `GET`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// - A map whose keys are the top-level data keys
        /// returned from Vault by the write operation and whose values are the
        /// corresponding values. This map can only represent string data, so
        /// any non-string values returned from Vault are serialized as JSON.
        /// Only fields set in `write_fields` are present in the JSON data.
        /// </summary>
        [Output("writeData")]
        public Output<ImmutableDictionary<string, string>> WriteData { get; private set; } = null!;

        /// <summary>
        /// - The JSON data returned by the write operation.
        /// Only fields set in `write_fields` are present in the JSON data.
        /// </summary>
        [Output("writeDataJson")]
        public Output<string> WriteDataJson { get; private set; } = null!;

        /// <summary>
        /// - (Optional). A list of fields that should be returned
        /// in `write_data_json` and `write_data`. If omitted, data returned by
        /// the write operation is not available to the resource or included in
        /// state. This helps to avoid accidental storage of sensitive values in
        /// state. Some endpoints, such as many dynamic secrets endpoints, return
        /// data from writing to an endpoint rather than reading it. You should
        /// use `write_fields` if you need information returned in this way.
        /// </summary>
        [Output("writeFields")]
        public Output<ImmutableArray<string>> WriteFields { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("vault:generic/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("vault:generic/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "dataJson",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataJson", required: true)]
        private Input<string>? _dataJson;

        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written to the given path as the secret data.
        /// </summary>
        public Input<string>? DataJson
        {
            get => _dataJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _dataJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// - (Optional) True/false. Set this to true if your
        /// vault authentication is not able to delete the data or if the endpoint
        /// does not support the `DELETE` method. Defaults to false.
        /// </summary>
        [Input("disableDelete")]
        public Input<bool>? DisableDelete { get; set; }

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data or if the endpoint does
        /// not support the `GET` method. Setting this to `true` will break drift
        /// detection. You should set this to `true` for endpoints that are
        /// write-only. Defaults to false.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        /// <summary>
        /// - (Optional) True/false. If set to true,
        /// ignore any fields present when the endpoint is read but that were not
        /// in `data_json`. Also, if a field that was written is not returned when
        /// the endpoint is read, treat that field as being up to date. You should
        /// set this to `true` when writing to endpoint that, when read, returns a
        /// different set of fields from the ones you wrote, as is common with
        /// many configuration endpoints. Defaults to false.
        /// </summary>
        [Input("ignoreAbsentFields")]
        public Input<bool>? IgnoreAbsentFields { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The full logical path at which to write the given
        /// data. Consult each backend's documentation to see which endpoints
        /// support the `PUT` methods and to determine whether they also support
        /// `DELETE` and `GET`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("writeFields")]
        private InputList<string>? _writeFields;

        /// <summary>
        /// - (Optional). A list of fields that should be returned
        /// in `write_data_json` and `write_data`. If omitted, data returned by
        /// the write operation is not available to the resource or included in
        /// state. This helps to avoid accidental storage of sensitive values in
        /// state. Some endpoints, such as many dynamic secrets endpoints, return
        /// data from writing to an endpoint rather than reading it. You should
        /// use `write_fields` if you need information returned in this way.
        /// </summary>
        public InputList<string> WriteFields
        {
            get => _writeFields ?? (_writeFields = new InputList<string>());
            set => _writeFields = value;
        }

        public EndpointArgs()
        {
        }
        public static new EndpointArgs Empty => new EndpointArgs();
    }

    public sealed class EndpointState : global::Pulumi.ResourceArgs
    {
        [Input("dataJson")]
        private Input<string>? _dataJson;

        /// <summary>
        /// String containing a JSON-encoded object that will be
        /// written to the given path as the secret data.
        /// </summary>
        public Input<string>? DataJson
        {
            get => _dataJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _dataJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// - (Optional) True/false. Set this to true if your
        /// vault authentication is not able to delete the data or if the endpoint
        /// does not support the `DELETE` method. Defaults to false.
        /// </summary>
        [Input("disableDelete")]
        public Input<bool>? DisableDelete { get; set; }

        /// <summary>
        /// True/false. Set this to true if your vault
        /// authentication is not able to read the data or if the endpoint does
        /// not support the `GET` method. Setting this to `true` will break drift
        /// detection. You should set this to `true` for endpoints that are
        /// write-only. Defaults to false.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        /// <summary>
        /// - (Optional) True/false. If set to true,
        /// ignore any fields present when the endpoint is read but that were not
        /// in `data_json`. Also, if a field that was written is not returned when
        /// the endpoint is read, treat that field as being up to date. You should
        /// set this to `true` when writing to endpoint that, when read, returns a
        /// different set of fields from the ones you wrote, as is common with
        /// many configuration endpoints. Defaults to false.
        /// </summary>
        [Input("ignoreAbsentFields")]
        public Input<bool>? IgnoreAbsentFields { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The full logical path at which to write the given
        /// data. Consult each backend's documentation to see which endpoints
        /// support the `PUT` methods and to determine whether they also support
        /// `DELETE` and `GET`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("writeData")]
        private InputMap<string>? _writeData;

        /// <summary>
        /// - A map whose keys are the top-level data keys
        /// returned from Vault by the write operation and whose values are the
        /// corresponding values. This map can only represent string data, so
        /// any non-string values returned from Vault are serialized as JSON.
        /// Only fields set in `write_fields` are present in the JSON data.
        /// </summary>
        public InputMap<string> WriteData
        {
            get => _writeData ?? (_writeData = new InputMap<string>());
            set => _writeData = value;
        }

        /// <summary>
        /// - The JSON data returned by the write operation.
        /// Only fields set in `write_fields` are present in the JSON data.
        /// </summary>
        [Input("writeDataJson")]
        public Input<string>? WriteDataJson { get; set; }

        [Input("writeFields")]
        private InputList<string>? _writeFields;

        /// <summary>
        /// - (Optional). A list of fields that should be returned
        /// in `write_data_json` and `write_data`. If omitted, data returned by
        /// the write operation is not available to the resource or included in
        /// state. This helps to avoid accidental storage of sensitive values in
        /// state. Some endpoints, such as many dynamic secrets endpoints, return
        /// data from writing to an endpoint rather than reading it. You should
        /// use `write_fields` if you need information returned in this way.
        /// </summary>
        public InputList<string> WriteFields
        {
            get => _writeFields ?? (_writeFields = new InputList<string>());
            set => _writeFields = value;
        }

        public EndpointState()
        {
        }
        public static new EndpointState Empty => new EndpointState();
    }
}
