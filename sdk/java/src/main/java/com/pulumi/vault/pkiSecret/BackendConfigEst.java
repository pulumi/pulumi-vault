// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.BackendConfigEstArgs;
import com.pulumi.vault.pkiSecret.inputs.BackendConfigEstState;
import com.pulumi.vault.pkiSecret.outputs.BackendConfigEstAuthenticators;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows setting the EST configuration on a PKI Secret Backend
 * 
 * ## Import
 * 
 * The PKI config cluster can be imported using the resource&#39;s `id`.
 * In the case of the example above the `id` would be `pki-root/config/est`,
 * where the `pki-root` component is the resource&#39;s `backend`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:pkiSecret/backendConfigEst:BackendConfigEst example pki-root/config/est
 * ```
 * 
 */
@ResourceType(type="vault:pkiSecret/backendConfigEst:BackendConfigEst")
public class BackendConfigEst extends com.pulumi.resources.CustomResource {
    /**
     * Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
     * 
     * &lt;a id=&#34;nestedatt--authenticators&#34;&gt;&lt;/a&gt;
     * 
     */
    @Export(name="auditFields", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> auditFields;

    /**
     * @return Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
     * 
     * &lt;a id=&#34;nestedatt--authenticators&#34;&gt;&lt;/a&gt;
     * 
     */
    public Output<List<String>> auditFields() {
        return this.auditFields;
    }
    /**
     * Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
     * 
     */
    @Export(name="authenticators", refs={BackendConfigEstAuthenticators.class}, tree="[0]")
    private Output<BackendConfigEstAuthenticators> authenticators;

    /**
     * @return Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
     * 
     */
    public Output<BackendConfigEstAuthenticators> authenticators() {
        return this.authenticators;
    }
    /**
     * The path to the PKI secret backend to
     * read the EST configuration from, with no leading or trailing `/`s.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output<String> backend;

    /**
     * @return The path to the PKI secret backend to
     * read the EST configuration from, with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
     * 
     */
    @Export(name="defaultMount", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> defaultMount;

    /**
     * @return If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
     * 
     */
    public Output<Optional<Boolean>> defaultMount() {
        return Codegen.optional(this.defaultMount);
    }
    /**
     * Required to be set if default_mount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:&lt;role_name&gt;.
     * 
     */
    @Export(name="defaultPathPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultPathPolicy;

    /**
     * @return Required to be set if default_mount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:&lt;role_name&gt;.
     * 
     */
    public Output<Optional<String>> defaultPathPolicy() {
        return Codegen.optional(this.defaultPathPolicy);
    }
    /**
     * If set, parse out fields from the provided CSR making them available for Sentinel policies.
     * 
     */
    @Export(name="enableSentinelParsing", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableSentinelParsing;

    /**
     * @return If set, parse out fields from the provided CSR making them available for Sentinel policies.
     * 
     */
    public Output<Optional<Boolean>> enableSentinelParsing() {
        return Codegen.optional(this.enableSentinelParsing);
    }
    /**
     * Specifies whether EST is enabled.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Specifies whether EST is enabled.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:&lt;role_name&gt;. Labels must be unique across Vault cluster, and will register .well-known/est/&lt;label&gt; URL paths.
     * 
     */
    @Export(name="labelToPathPolicy", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> labelToPathPolicy;

    /**
     * @return Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:&lt;role_name&gt;. Labels must be unique across Vault cluster, and will register .well-known/est/&lt;label&gt; URL paths.
     * 
     */
    public Output<Optional<Map<String,Object>>> labelToPathPolicy() {
        return Codegen.optional(this.labelToPathPolicy);
    }
    /**
     * A read-only timestamp representing the last time the configuration was updated.
     * 
     */
    @Export(name="lastUpdated", refs={String.class}, tree="[0]")
    private Output<String> lastUpdated;

    /**
     * @return A read-only timestamp representing the last time the configuration was updated.
     * 
     */
    public Output<String> lastUpdated() {
        return this.lastUpdated;
    }
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackendConfigEst(java.lang.String name) {
        this(name, BackendConfigEstArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackendConfigEst(java.lang.String name, BackendConfigEstArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackendConfigEst(java.lang.String name, BackendConfigEstArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/backendConfigEst:BackendConfigEst", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BackendConfigEst(java.lang.String name, Output<java.lang.String> id, @Nullable BackendConfigEstState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/backendConfigEst:BackendConfigEst", name, state, makeResourceOptions(options, id), false);
    }

    private static BackendConfigEstArgs makeArgs(BackendConfigEstArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BackendConfigEstArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static BackendConfigEst get(java.lang.String name, Output<java.lang.String> id, @Nullable BackendConfigEstState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackendConfigEst(name, id, state, options);
    }
}
