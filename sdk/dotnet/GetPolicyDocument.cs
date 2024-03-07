// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static class GetPolicyDocument
    {
        /// <summary>
        /// This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `vault.Policy` resource.
        /// 
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
        ///     var examplePolicyDocument = Vault.GetPolicyDocument.Invoke(new()
        ///     {
        ///         Rules = new[]
        ///         {
        ///             new Vault.Inputs.GetPolicyDocumentRuleInputArgs
        ///             {
        ///                 Path = "secret/*",
        ///                 Capabilities = new[]
        ///                 {
        ///                     "create",
        ///                     "read",
        ///                     "update",
        ///                     "delete",
        ///                     "list",
        ///                 },
        ///                 Description = "allow all on secrets",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var examplePolicy = new Vault.Policy("examplePolicy", new()
        ///     {
        ///         PolicyContents = examplePolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Hcl),
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPolicyDocumentResult> InvokeAsync(GetPolicyDocumentArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPolicyDocumentResult>("vault:index/getPolicyDocument:getPolicyDocument", args ?? new GetPolicyDocumentArgs(), options.WithDefaults());

        /// <summary>
        /// This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `vault.Policy` resource.
        /// 
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
        ///     var examplePolicyDocument = Vault.GetPolicyDocument.Invoke(new()
        ///     {
        ///         Rules = new[]
        ///         {
        ///             new Vault.Inputs.GetPolicyDocumentRuleInputArgs
        ///             {
        ///                 Path = "secret/*",
        ///                 Capabilities = new[]
        ///                 {
        ///                     "create",
        ///                     "read",
        ///                     "update",
        ///                     "delete",
        ///                     "list",
        ///                 },
        ///                 Description = "allow all on secrets",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var examplePolicy = new Vault.Policy("examplePolicy", new()
        ///     {
        ///         PolicyContents = examplePolicyDocument.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Hcl),
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPolicyDocumentResult> Invoke(GetPolicyDocumentInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPolicyDocumentResult>("vault:index/getPolicyDocument:getPolicyDocument", args ?? new GetPolicyDocumentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPolicyDocumentArgs : global::Pulumi.InvokeArgs
    {
        [Input("namespace")]
        public string? Namespace { get; set; }

        [Input("rules")]
        private List<Inputs.GetPolicyDocumentRuleArgs>? _rules;
        public List<Inputs.GetPolicyDocumentRuleArgs> Rules
        {
            get => _rules ?? (_rules = new List<Inputs.GetPolicyDocumentRuleArgs>());
            set => _rules = value;
        }

        public GetPolicyDocumentArgs()
        {
        }
        public static new GetPolicyDocumentArgs Empty => new GetPolicyDocumentArgs();
    }

    public sealed class GetPolicyDocumentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("rules")]
        private InputList<Inputs.GetPolicyDocumentRuleInputArgs>? _rules;
        public InputList<Inputs.GetPolicyDocumentRuleInputArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.GetPolicyDocumentRuleInputArgs>());
            set => _rules = value;
        }

        public GetPolicyDocumentInvokeArgs()
        {
        }
        public static new GetPolicyDocumentInvokeArgs Empty => new GetPolicyDocumentInvokeArgs();
    }


    [OutputType]
    public sealed class GetPolicyDocumentResult
    {
        /// <summary>
        /// The above arguments serialized as a standard Vault HCL policy document.
        /// </summary>
        public readonly string Hcl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        public readonly ImmutableArray<Outputs.GetPolicyDocumentRuleResult> Rules;

        [OutputConstructor]
        private GetPolicyDocumentResult(
            string hcl,

            string id,

            string? @namespace,

            ImmutableArray<Outputs.GetPolicyDocumentRuleResult> rules)
        {
            Hcl = hcl;
            Id = id;
            Namespace = @namespace;
            Rules = rules;
        }
    }
}
