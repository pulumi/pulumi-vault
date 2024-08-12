// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.AuthBackendArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.AuthBackendState;
import com.pulumi.vault.outputs.AuthBackendTune;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Auth methods can be imported using the `path`, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/authBackend:AuthBackend example github
 * ```
 * 
 */
@ResourceType(type="vault:index/authBackend:AuthBackend")
public class AuthBackend extends com.pulumi.resources.CustomResource {
    /**
     * The accessor for this auth method
     * 
     */
    @Export(name="accessor", refs={String.class}, tree="[0]")
    private Output<String> accessor;

    /**
     * @return The accessor for this auth method
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * A description of the auth method.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
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
    @Export(name="disableRemount", refs={Boolean.class}, tree="[0]")
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
     * The key to use for signing identity tokens.
     * 
     */
    @Export(name="identityTokenKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityTokenKey;

    /**
     * @return The key to use for signing identity tokens.
     * 
     */
    public Output<Optional<String>> identityTokenKey() {
        return Codegen.optional(this.identityTokenKey);
    }
    /**
     * Specifies if the auth method is local only.
     * 
     */
    @Export(name="local", refs={Boolean.class}, tree="[0]")
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
     * The path to mount the auth method — this defaults to the name of the type.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The path to mount the auth method — this defaults to the name of the type.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    @Export(name="tune", refs={AuthBackendTune.class}, tree="[0]")
    private Output<AuthBackendTune> tune;

    /**
     * @return Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    public Output<AuthBackendTune> tune() {
        return this.tune;
    }
    /**
     * The name of the auth method type.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The name of the auth method type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackend(java.lang.String name) {
        this(name, AuthBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackend(java.lang.String name, AuthBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackend(java.lang.String name, AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/authBackend:AuthBackend", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthBackend(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/authBackend:AuthBackend", name, state, makeResourceOptions(options, id), false);
    }

    private static AuthBackendArgs makeArgs(AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthBackendArgs.Empty : args;
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
    public static AuthBackend get(java.lang.String name, Output<java.lang.String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackend(name, id, state, options);
    }
}
