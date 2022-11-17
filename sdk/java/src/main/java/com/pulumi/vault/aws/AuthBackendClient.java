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
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
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
 *         var exampleAuthBackend = new AuthBackend(&#34;exampleAuthBackend&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;aws&#34;)
 *             .build());
 * 
 *         var exampleAuthBackendClient = new AuthBackendClient(&#34;exampleAuthBackendClient&#34;, AuthBackendClientArgs.builder()        
 *             .backend(exampleAuthBackend.path())
 *             .accessKey(&#34;INSERT_AWS_ACCESS_KEY&#34;)
 *             .secretKey(&#34;INSERT_AWS_SECRET_KEY&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AWS auth backend clients can be imported using `auth/`, the `backend` path, and `/config/client` e.g.
 * 
 * ```sh
 *  $ pulumi import vault:aws/authBackendClient:AuthBackendClient example auth/aws/config/client
 * ```
 * 
 */
@ResourceType(type="vault:aws/authBackendClient:AuthBackendClient")
public class AuthBackendClient extends com.pulumi.resources.CustomResource {
    /**
     * The AWS access key that Vault should use for the
     * auth backend.
     * 
     */
    @Export(name="accessKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> accessKey;

    /**
     * @return The AWS access key that Vault should use for the
     * auth backend.
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
    @Export(name="backend", type=String.class, parameters={})
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
    @Export(name="ec2Endpoint", type=String.class, parameters={})
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
    @Export(name="iamEndpoint", type=String.class, parameters={})
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
    @Export(name="iamServerIdHeaderValue", type=String.class, parameters={})
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
     * The AWS secret key that Vault should use for the
     * auth backend.
     * 
     */
    @Export(name="secretKey", type=String.class, parameters={})
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
    @Export(name="stsEndpoint", type=String.class, parameters={})
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
    @Export(name="stsRegion", type=String.class, parameters={})
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendClient(String name) {
        this(name, AuthBackendClientArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendClient(String name, @Nullable AuthBackendClientArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendClient(String name, @Nullable AuthBackendClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendClient:AuthBackendClient", name, args == null ? AuthBackendClientArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendClient(String name, Output<String> id, @Nullable AuthBackendClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendClient:AuthBackendClient", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static AuthBackendClient get(String name, Output<String> id, @Nullable AuthBackendClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendClient(name, id, state, options);
    }
}