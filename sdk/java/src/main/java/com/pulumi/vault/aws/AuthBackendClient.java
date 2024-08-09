// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.aws.AuthBackendClientArgs;
import com.pulumi.vault.aws.inputs.AuthBackendClientState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * You can setup the AWS auth engine with Workload Identity Federation (WIF) for a secret-less configuration:
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.aws.AuthBackendClient;
 * import com.pulumi.vault.aws.AuthBackendClientArgs;
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
 *         var example = new AuthBackend("example", AuthBackendArgs.builder()
 *             .type("aws")
 *             .build());
 * 
 *         var exampleAuthBackendClient = new AuthBackendClient("exampleAuthBackendClient", AuthBackendClientArgs.builder()
 *             .identityTokenAudience("<TOKEN_AUDIENCE>")
 *             .identityTokenTtl("<TOKEN_TTL>")
 *             .roleArn("<AWS_ROLE_ARN>")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.aws.AuthBackendClient;
 * import com.pulumi.vault.aws.AuthBackendClientArgs;
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
 *         var example = new AuthBackend("example", AuthBackendArgs.builder()
 *             .type("aws")
 *             .build());
 * 
 *         var exampleAuthBackendClient = new AuthBackendClient("exampleAuthBackendClient", AuthBackendClientArgs.builder()
 *             .backend(example.path())
 *             .accessKey("INSERT_AWS_ACCESS_KEY")
 *             .secretKey("INSERT_AWS_SECRET_KEY")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * AWS auth backend clients can be imported using `auth/`, the `backend` path, and `/config/client` e.g.
 * 
 * ```sh
 * $ pulumi import vault:aws/authBackendClient:AuthBackendClient example auth/aws/config/client
 * ```
 * 
 */
@ResourceType(type="vault:aws/authBackendClient:AuthBackendClient")
public class AuthBackendClient extends com.pulumi.resources.CustomResource {
    /**
     * The AWS access key that Vault should use for the
     * auth backend. Mutually exclusive with `identity_token_audience`.
     * 
     */
    @Export(name="accessKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessKey;

    /**
     * @return The AWS access key that Vault should use for the
     * auth backend. Mutually exclusive with `identity_token_audience`.
     * 
     */
    public Output<Optional<String>> accessKey() {
        return Codegen.optional(this.accessKey);
    }
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * Override the URL Vault uses when making EC2 API
     * calls.
     * 
     */
    @Export(name="ec2Endpoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ec2Endpoint;

    /**
     * @return Override the URL Vault uses when making EC2 API
     * calls.
     * 
     */
    public Output<Optional<String>> ec2Endpoint() {
        return Codegen.optional(this.ec2Endpoint);
    }
    /**
     * Override the URL Vault uses when making IAM API
     * calls.
     * 
     */
    @Export(name="iamEndpoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> iamEndpoint;

    /**
     * @return Override the URL Vault uses when making IAM API
     * calls.
     * 
     */
    public Output<Optional<String>> iamEndpoint() {
        return Codegen.optional(this.iamEndpoint);
    }
    /**
     * The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     * 
     */
    @Export(name="iamServerIdHeaderValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> iamServerIdHeaderValue;

    /**
     * @return The value to require in the
     * `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
     * that are used in the IAM auth method.
     * 
     */
    public Output<Optional<String>> iamServerIdHeaderValue() {
        return Codegen.optional(this.iamServerIdHeaderValue);
    }
    /**
     * The audience claim value. Mutually exclusive with `access_key`.
     * Requires Vault 1.17+. *Available only for Vault Enterprise*
     * 
     */
    @Export(name="identityTokenAudience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenAudience;

    /**
     * @return The audience claim value. Mutually exclusive with `access_key`.
     * Requires Vault 1.17+. *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<String>> identityTokenAudience() {
        return Codegen.optional(this.identityTokenAudience);
    }
    /**
     * The TTL of generated identity tokens in seconds. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="identityTokenTtl", refs={Integer.class}, tree="[0]")
    private Output<Integer> identityTokenTtl;

    /**
     * @return The TTL of generated identity tokens in seconds. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Integer> identityTokenTtl() {
        return this.identityTokenTtl;
    }
    /**
     * Number of max retries the client should use for recoverable errors.
     * The default `-1` falls back to the AWS SDK&#39;s default behavior.
     * 
     */
    @Export(name="maxRetries", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxRetries;

    /**
     * @return Number of max retries the client should use for recoverable errors.
     * The default `-1` falls back to the AWS SDK&#39;s default behavior.
     * 
     */
    public Output<Optional<Integer>> maxRetries() {
        return Codegen.optional(this.maxRetries);
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Role ARN to assume for plugin identity token federation. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleArn;

    /**
     * @return Role ARN to assume for plugin identity token federation. Requires Vault 1.17+.
     * *Available only for Vault Enterprise*
     * 
     */
    public Output<Optional<String>> roleArn() {
        return Codegen.optional(this.roleArn);
    }
    /**
     * The AWS secret key that Vault should use for the
     * auth backend.
     * 
     */
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secretKey;

    /**
     * @return The AWS secret key that Vault should use for the
     * auth backend.
     * 
     */
    public Output<Optional<String>> secretKey() {
        return Codegen.optional(this.secretKey);
    }
    /**
     * Override the URL Vault uses when making STS API
     * calls.
     * 
     */
    @Export(name="stsEndpoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stsEndpoint;

    /**
     * @return Override the URL Vault uses when making STS API
     * calls.
     * 
     */
    public Output<Optional<String>> stsEndpoint() {
        return Codegen.optional(this.stsEndpoint);
    }
    /**
     * Override the default region when making STS API
     * calls. The `sts_endpoint` argument must be set when using `sts_region`.
     * 
     */
    @Export(name="stsRegion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stsRegion;

    /**
     * @return Override the default region when making STS API
     * calls. The `sts_endpoint` argument must be set when using `sts_region`.
     * 
     */
    public Output<Optional<String>> stsRegion() {
        return Codegen.optional(this.stsRegion);
    }
    /**
     * Available in Vault v1.15+. If set,
     * overrides both `sts_endpoint` and `sts_region` to instead use the region
     * specified in the client request headers for IAM-based authentication.
     * This can be useful when you have client requests coming from different
     * regions and want flexibility in which regional STS API is used.
     * 
     */
    @Export(name="useStsRegionFromClient", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useStsRegionFromClient;

    /**
     * @return Available in Vault v1.15+. If set,
     * overrides both `sts_endpoint` and `sts_region` to instead use the region
     * specified in the client request headers for IAM-based authentication.
     * This can be useful when you have client requests coming from different
     * regions and want flexibility in which regional STS API is used.
     * 
     */
    public Output<Boolean> useStsRegionFromClient() {
        return this.useStsRegionFromClient;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendClient(java.lang.String name) {
        this(name, AuthBackendClientArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendClient(java.lang.String name, @Nullable AuthBackendClientArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendClient(java.lang.String name, @Nullable AuthBackendClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendClient:AuthBackendClient", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthBackendClient(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendClient:AuthBackendClient", name, state, makeResourceOptions(options, id), false);
    }

    private static AuthBackendClientArgs makeArgs(@Nullable AuthBackendClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthBackendClientArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessKey",
                "secretKey"
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
    public static AuthBackendClient get(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendClient(name, id, state, options);
    }
}
