// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.aws.AuthBackendLoginArgs;
import com.pulumi.vault.aws.inputs.AuthBackendLoginState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Logs into a Vault server using an AWS auth backend. Login can be
 * accomplished using a signed identity request from IAM or using ec2
 * instance metadata. For more information, see the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/aws.html).
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.aws.AuthBackendClient;
 * import com.pulumi.vault.aws.AuthBackendClientArgs;
 * import com.pulumi.vault.aws.AuthBackendRole;
 * import com.pulumi.vault.aws.AuthBackendRoleArgs;
 * import com.pulumi.vault.aws.AuthBackendLogin;
 * import com.pulumi.vault.aws.AuthBackendLoginArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var aws = new AuthBackend(&#34;aws&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;aws&#34;)
 *             .path(&#34;aws&#34;)
 *             .build());
 * 
 *         var exampleAuthBackendClient = new AuthBackendClient(&#34;exampleAuthBackendClient&#34;, AuthBackendClientArgs.builder()        
 *             .backend(aws.path())
 *             .accessKey(&#34;123456789012&#34;)
 *             .secretKey(&#34;AWSSECRETKEYGOESHERE&#34;)
 *             .build());
 * 
 *         var exampleAuthBackendRole = new AuthBackendRole(&#34;exampleAuthBackendRole&#34;, AuthBackendRoleArgs.builder()        
 *             .backend(aws.path())
 *             .role(&#34;test-role&#34;)
 *             .authType(&#34;ec2&#34;)
 *             .boundAmiId(&#34;ami-8c1be5f6&#34;)
 *             .boundAccountId(&#34;123456789012&#34;)
 *             .boundVpcId(&#34;vpc-b61106d4&#34;)
 *             .boundSubnetId(&#34;vpc-133128f1&#34;)
 *             .boundIamInstanceProfileArns(&#34;arn:aws:iam::123456789012:instance-profile/MyProfile&#34;)
 *             .ttl(60)
 *             .maxTtl(120)
 *             .tokenPolicies(            
 *                 &#34;default&#34;,
 *                 &#34;dev&#34;,
 *                 &#34;prod&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(&#34;vault_aws_auth_backend_client.example&#34;)
 *                 .build());
 * 
 *         var exampleAuthBackendLogin = new AuthBackendLogin(&#34;exampleAuthBackendLogin&#34;, AuthBackendLoginArgs.builder()        
 *             .backend(vault_auth_backend.example().path())
 *             .role(exampleAuthBackendRole.role())
 *             .identity(&#34;BASE64ENCODEDIDENTITYDOCUMENT&#34;)
 *             .signature(&#34;BASE64ENCODEDSHA256IDENTITYDOCUMENTSIGNATURE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vault:aws/authBackendLogin:AuthBackendLogin")
public class AuthBackendLogin extends com.pulumi.resources.CustomResource {
    /**
     * The token&#39;s accessor.
     * 
     */
    @Export(name="accessor", type=String.class, parameters={})
    private Output<String> accessor;

    /**
     * @return The token&#39;s accessor.
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * The authentication type used to generate this token.
     * 
     */
    @Export(name="authType", type=String.class, parameters={})
    private Output<String> authType;

    /**
     * @return The authentication type used to generate this token.
     * 
     */
    public Output<String> authType() {
        return this.authType;
    }
    /**
     * The unique name of the AWS auth backend. Defaults to
     * &#39;aws&#39;.
     * 
     */
    @Export(name="backend", type=String.class, parameters={})
    private Output</* @Nullable */ String> backend;

    /**
     * @return The unique name of the AWS auth backend. Defaults to
     * &#39;aws&#39;.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * The token returned by Vault.
     * 
     */
    @Export(name="clientToken", type=String.class, parameters={})
    private Output<String> clientToken;

    /**
     * @return The token returned by Vault.
     * 
     */
    public Output<String> clientToken() {
        return this.clientToken;
    }
    /**
     * The HTTP method used in the signed IAM
     * request.
     * 
     */
    @Export(name="iamHttpRequestMethod", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamHttpRequestMethod;

    /**
     * @return The HTTP method used in the signed IAM
     * request.
     * 
     */
    public Output<Optional<String>> iamHttpRequestMethod() {
        return Codegen.optional(this.iamHttpRequestMethod);
    }
    /**
     * The base64-encoded body of the signed
     * request.
     * 
     */
    @Export(name="iamRequestBody", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamRequestBody;

    /**
     * @return The base64-encoded body of the signed
     * request.
     * 
     */
    public Output<Optional<String>> iamRequestBody() {
        return Codegen.optional(this.iamRequestBody);
    }
    /**
     * The base64-encoded, JSON serialized
     * representation of the GetCallerIdentity HTTP request headers.
     * 
     */
    @Export(name="iamRequestHeaders", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamRequestHeaders;

    /**
     * @return The base64-encoded, JSON serialized
     * representation of the GetCallerIdentity HTTP request headers.
     * 
     */
    public Output<Optional<String>> iamRequestHeaders() {
        return Codegen.optional(this.iamRequestHeaders);
    }
    /**
     * The base64-encoded HTTP URL used in the signed
     * request.
     * 
     */
    @Export(name="iamRequestUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamRequestUrl;

    /**
     * @return The base64-encoded HTTP URL used in the signed
     * request.
     * 
     */
    public Output<Optional<String>> iamRequestUrl() {
        return Codegen.optional(this.iamRequestUrl);
    }
    /**
     * The base64-encoded EC2 instance identity document to
     * authenticate with. Can be retrieved from the EC2 metadata server.
     * 
     */
    @Export(name="identity", type=String.class, parameters={})
    private Output</* @Nullable */ String> identity;

    /**
     * @return The base64-encoded EC2 instance identity document to
     * authenticate with. Can be retrieved from the EC2 metadata server.
     * 
     */
    public Output<Optional<String>> identity() {
        return Codegen.optional(this.identity);
    }
    /**
     * The duration in seconds the token will be valid, relative
     * to the time in `lease_start_time`.
     * 
     */
    @Export(name="leaseDuration", type=Integer.class, parameters={})
    private Output<Integer> leaseDuration;

    /**
     * @return The duration in seconds the token will be valid, relative
     * to the time in `lease_start_time`.
     * 
     */
    public Output<Integer> leaseDuration() {
        return this.leaseDuration;
    }
    /**
     * Time at which the lease was read, using the clock of the system where Terraform was running
     * 
     */
    @Export(name="leaseStartTime", type=String.class, parameters={})
    private Output<String> leaseStartTime;

    /**
     * @return Time at which the lease was read, using the clock of the system where Terraform was running
     * 
     */
    public Output<String> leaseStartTime() {
        return this.leaseStartTime;
    }
    /**
     * A map of information returned by the Vault server about the
     * authentication used to generate this token.
     * 
     */
    @Export(name="metadata", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> metadata;

    /**
     * @return A map of information returned by the Vault server about the
     * authentication used to generate this token.
     * 
     */
    public Output<Map<String,Object>> metadata() {
        return this.metadata;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", type=String.class, parameters={})
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * The unique nonce to be used for login requests. Can be
     * set to a user-specified value, or will contain the server-generated value
     * once a token is issued. EC2 instances can only acquire a single token until
     * the whitelist is tidied again unless they keep track of this nonce.
     * 
     */
    @Export(name="nonce", type=String.class, parameters={})
    private Output<String> nonce;

    /**
     * @return The unique nonce to be used for login requests. Can be
     * set to a user-specified value, or will contain the server-generated value
     * once a token is issued. EC2 instances can only acquire a single token until
     * the whitelist is tidied again unless they keep track of this nonce.
     * 
     */
    public Output<String> nonce() {
        return this.nonce;
    }
    /**
     * The PKCS#7 signature of the identity document to
     * authenticate with, with all newline characters removed. Can be retrieved from
     * the EC2 metadata server.
     * 
     */
    @Export(name="pkcs7", type=String.class, parameters={})
    private Output</* @Nullable */ String> pkcs7;

    /**
     * @return The PKCS#7 signature of the identity document to
     * authenticate with, with all newline characters removed. Can be retrieved from
     * the EC2 metadata server.
     * 
     */
    public Output<Optional<String>> pkcs7() {
        return Codegen.optional(this.pkcs7);
    }
    /**
     * The Vault policies assigned to this token.
     * 
     */
    @Export(name="policies", type=List.class, parameters={String.class})
    private Output<List<String>> policies;

    /**
     * @return The Vault policies assigned to this token.
     * 
     */
    public Output<List<String>> policies() {
        return this.policies;
    }
    /**
     * Set to true if the token can be extended through renewal.
     * 
     */
    @Export(name="renewable", type=Boolean.class, parameters={})
    private Output<Boolean> renewable;

    /**
     * @return Set to true if the token can be extended through renewal.
     * 
     */
    public Output<Boolean> renewable() {
        return this.renewable;
    }
    /**
     * The name of the AWS auth backend role to create tokens
     * against.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The name of the AWS auth backend role to create tokens
     * against.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * The base64-encoded SHA256 RSA signature of the
     * instance identity document to authenticate with, with all newline characters
     * removed. Can be retrieved from the EC2 metadata server.
     * 
     */
    @Export(name="signature", type=String.class, parameters={})
    private Output</* @Nullable */ String> signature;

    /**
     * @return The base64-encoded SHA256 RSA signature of the
     * instance identity document to authenticate with, with all newline characters
     * removed. Can be retrieved from the EC2 metadata server.
     * 
     */
    public Output<Optional<String>> signature() {
        return Codegen.optional(this.signature);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendLogin(String name) {
        this(name, AuthBackendLoginArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendLogin(String name, @Nullable AuthBackendLoginArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendLogin(String name, @Nullable AuthBackendLoginArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendLogin:AuthBackendLogin", name, args == null ? AuthBackendLoginArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendLogin(String name, Output<String> id, @Nullable AuthBackendLoginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendLogin:AuthBackendLogin", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientToken"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AuthBackendLogin get(String name, Output<String> id, @Nullable AuthBackendLoginState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendLogin(name, id, state, options);
    }
}