// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.gcp.AuthBackendArgs;
import com.pulumi.vault.gcp.inputs.AuthBackendState;
import com.pulumi.vault.gcp.outputs.AuthBackendCustomEndpoint;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to configure the [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.gcp.AuthBackend;
 * import com.pulumi.vault.gcp.AuthBackendArgs;
 * import com.pulumi.vault.gcp.inputs.AuthBackendCustomEndpointArgs;
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
 *         var gcp = new AuthBackend(&#34;gcp&#34;, AuthBackendArgs.builder()        
 *             .credentials(Files.readString(Paths.get(&#34;vault-gcp-credentials.json&#34;)))
 *             .customEndpoint(AuthBackendCustomEndpointArgs.builder()
 *                 .api(&#34;www.googleapis.com&#34;)
 *                 .iam(&#34;iam.googleapis.com&#34;)
 *                 .crm(&#34;cloudresourcemanager.googleapis.com&#34;)
 *                 .compute(&#34;compute.googleapis.com&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GCP authentication backends can be imported using the backend name, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:gcp/authBackend:AuthBackend gcp gcp
 * ```
 * 
 */
@ResourceType(type="vault:gcp/authBackend:AuthBackend")
public class AuthBackend extends com.pulumi.resources.CustomResource {
    /**
     * The clients email associated with the credentials
     * 
     */
    @Export(name="clientEmail", type=String.class, parameters={})
    private Output<String> clientEmail;

    /**
     * @return The clients email associated with the credentials
     * 
     */
    public Output<String> clientEmail() {
        return this.clientEmail;
    }
    /**
     * The Client ID of the credentials
     * 
     */
    @Export(name="clientId", type=String.class, parameters={})
    private Output<String> clientId;

    /**
     * @return The Client ID of the credentials
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     * 
     */
    @Export(name="credentials", type=String.class, parameters={})
    private Output</* @Nullable */ String> credentials;

    /**
     * @return A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     * 
     */
    public Output<Optional<String>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     * 
     */
    @Export(name="customEndpoint", type=AuthBackendCustomEndpoint.class, parameters={})
    private Output</* @Nullable */ AuthBackendCustomEndpoint> customEndpoint;

    /**
     * @return Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     * 
     */
    public Output<Optional<AuthBackendCustomEndpoint>> customEndpoint() {
        return Codegen.optional(this.customEndpoint);
    }
    /**
     * A description of the auth method.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the auth method.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Export(name="disableRemount", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Output<Optional<Boolean>> disableRemount() {
        return Codegen.optional(this.disableRemount);
    }
    /**
     * Specifies if the auth method is local only.
     * 
     */
    @Export(name="local", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> local;

    /**
     * @return Specifies if the auth method is local only.
     * 
     */
    public Output<Optional<Boolean>> local() {
        return Codegen.optional(this.local);
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
     * The path to mount the auth method — this defaults to &#39;gcp&#39;.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output</* @Nullable */ String> path;

    /**
     * @return The path to mount the auth method — this defaults to &#39;gcp&#39;.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The ID of the private key from the credentials
     * 
     */
    @Export(name="privateKeyId", type=String.class, parameters={})
    private Output<String> privateKeyId;

    /**
     * @return The ID of the private key from the credentials
     * 
     */
    public Output<String> privateKeyId() {
        return this.privateKeyId;
    }
    /**
     * The GCP Project ID
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The GCP Project ID
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackend(String name) {
        this(name, AuthBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackend(String name, @Nullable AuthBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackend(String name, @Nullable AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:gcp/authBackend:AuthBackend", name, args == null ? AuthBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackend(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:gcp/authBackend:AuthBackend", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "credentials"
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
    public static AuthBackend get(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackend(name, id, state, options);
    }
}