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
    /// Manages additional request headers that appear in audited requests.
    /// 
    /// &gt; **Note**
    /// Because of the way the [sys/config/auditing/request-headers API](https://www.vaultproject.io/api-docs/system/config-auditing)
    /// is implemented in Vault, this resource will manage existing audited headers with
    /// matching names without requiring import.
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
    ///     var xForwardedFor = new Vault.AuditRequestHeader("x_forwarded_for", new()
    ///     {
    ///         Name = "X-Forwarded-For",
    ///         Hmac = false,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:index/auditRequestHeader:AuditRequestHeader")]
    public partial class AuditRequestHeader : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether this header's value should be HMAC'd in the audit logs.
        /// </summary>
        [Output("hmac")]
        public Output<bool?> Hmac { get; private set; } = null!;

        /// <summary>
        /// The name of the request header to audit.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a AuditRequestHeader resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuditRequestHeader(string name, AuditRequestHeaderArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:index/auditRequestHeader:AuditRequestHeader", name, args ?? new AuditRequestHeaderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuditRequestHeader(string name, Input<string> id, AuditRequestHeaderState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/auditRequestHeader:AuditRequestHeader", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuditRequestHeader resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuditRequestHeader Get(string name, Input<string> id, AuditRequestHeaderState? state = null, CustomResourceOptions? options = null)
        {
            return new AuditRequestHeader(name, id, state, options);
        }
    }

    public sealed class AuditRequestHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this header's value should be HMAC'd in the audit logs.
        /// </summary>
        [Input("hmac")]
        public Input<bool>? Hmac { get; set; }

        /// <summary>
        /// The name of the request header to audit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public AuditRequestHeaderArgs()
        {
        }
        public static new AuditRequestHeaderArgs Empty => new AuditRequestHeaderArgs();
    }

    public sealed class AuditRequestHeaderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this header's value should be HMAC'd in the audit logs.
        /// </summary>
        [Input("hmac")]
        public Input<bool>? Hmac { get; set; }

        /// <summary>
        /// The name of the request header to audit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public AuditRequestHeaderState()
        {
        }
        public static new AuditRequestHeaderState Empty => new AuditRequestHeaderState();
    }
}
