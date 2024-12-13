// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vault.pkiSecret.inputs.BackendConfigEstAuthenticatorsArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackendConfigEstArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackendConfigEstArgs Empty = new BackendConfigEstArgs();

    /**
     * Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
     * 
     * &lt;a id=&#34;nestedatt--authenticators&#34;&gt;&lt;/a&gt;
     * 
     */
    @Import(name="auditFields")
    private @Nullable Output<List<String>> auditFields;

    /**
     * @return Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
     * 
     * &lt;a id=&#34;nestedatt--authenticators&#34;&gt;&lt;/a&gt;
     * 
     */
    public Optional<Output<List<String>>> auditFields() {
        return Optional.ofNullable(this.auditFields);
    }

    /**
     * Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
     * 
     */
    @Import(name="authenticators")
    private @Nullable Output<BackendConfigEstAuthenticatorsArgs> authenticators;

    /**
     * @return Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
     * 
     */
    public Optional<Output<BackendConfigEstAuthenticatorsArgs>> authenticators() {
        return Optional.ofNullable(this.authenticators);
    }

    /**
     * The path to the PKI secret backend to
     * read the EST configuration from, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
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
    @Import(name="defaultMount")
    private @Nullable Output<Boolean> defaultMount;

    /**
     * @return If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
     * 
     */
    public Optional<Output<Boolean>> defaultMount() {
        return Optional.ofNullable(this.defaultMount);
    }

    /**
     * Required to be set if default_mount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:&lt;role_name&gt;.
     * 
     */
    @Import(name="defaultPathPolicy")
    private @Nullable Output<String> defaultPathPolicy;

    /**
     * @return Required to be set if default_mount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:&lt;role_name&gt;.
     * 
     */
    public Optional<Output<String>> defaultPathPolicy() {
        return Optional.ofNullable(this.defaultPathPolicy);
    }

    /**
     * If set, parse out fields from the provided CSR making them available for Sentinel policies.
     * 
     */
    @Import(name="enableSentinelParsing")
    private @Nullable Output<Boolean> enableSentinelParsing;

    /**
     * @return If set, parse out fields from the provided CSR making them available for Sentinel policies.
     * 
     */
    public Optional<Output<Boolean>> enableSentinelParsing() {
        return Optional.ofNullable(this.enableSentinelParsing);
    }

    /**
     * Specifies whether EST is enabled.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Specifies whether EST is enabled.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:&lt;role_name&gt;. Labels must be unique across Vault cluster, and will register .well-known/est/&lt;label&gt; URL paths.
     * 
     */
    @Import(name="labelToPathPolicy")
    private @Nullable Output<Map<String,String>> labelToPathPolicy;

    /**
     * @return Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:&lt;role_name&gt;. Labels must be unique across Vault cluster, and will register .well-known/est/&lt;label&gt; URL paths.
     * 
     */
    public Optional<Output<Map<String,String>>> labelToPathPolicy() {
        return Optional.ofNullable(this.labelToPathPolicy);
    }

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private BackendConfigEstArgs() {}

    private BackendConfigEstArgs(BackendConfigEstArgs $) {
        this.auditFields = $.auditFields;
        this.authenticators = $.authenticators;
        this.backend = $.backend;
        this.defaultMount = $.defaultMount;
        this.defaultPathPolicy = $.defaultPathPolicy;
        this.enableSentinelParsing = $.enableSentinelParsing;
        this.enabled = $.enabled;
        this.labelToPathPolicy = $.labelToPathPolicy;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendConfigEstArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendConfigEstArgs $;

        public Builder() {
            $ = new BackendConfigEstArgs();
        }

        public Builder(BackendConfigEstArgs defaults) {
            $ = new BackendConfigEstArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param auditFields Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
         * 
         * &lt;a id=&#34;nestedatt--authenticators&#34;&gt;&lt;/a&gt;
         * 
         * @return builder
         * 
         */
        public Builder auditFields(@Nullable Output<List<String>> auditFields) {
            $.auditFields = auditFields;
            return this;
        }

        /**
         * @param auditFields Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
         * 
         * &lt;a id=&#34;nestedatt--authenticators&#34;&gt;&lt;/a&gt;
         * 
         * @return builder
         * 
         */
        public Builder auditFields(List<String> auditFields) {
            return auditFields(Output.of(auditFields));
        }

        /**
         * @param auditFields Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
         * 
         * &lt;a id=&#34;nestedatt--authenticators&#34;&gt;&lt;/a&gt;
         * 
         * @return builder
         * 
         */
        public Builder auditFields(String... auditFields) {
            return auditFields(List.of(auditFields));
        }

        /**
         * @param authenticators Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
         * 
         * @return builder
         * 
         */
        public Builder authenticators(@Nullable Output<BackendConfigEstAuthenticatorsArgs> authenticators) {
            $.authenticators = authenticators;
            return this;
        }

        /**
         * @param authenticators Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
         * 
         * @return builder
         * 
         */
        public Builder authenticators(BackendConfigEstAuthenticatorsArgs authenticators) {
            return authenticators(Output.of(authenticators));
        }

        /**
         * @param backend The path to the PKI secret backend to
         * read the EST configuration from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path to the PKI secret backend to
         * read the EST configuration from, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param defaultMount If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
         * 
         * @return builder
         * 
         */
        public Builder defaultMount(@Nullable Output<Boolean> defaultMount) {
            $.defaultMount = defaultMount;
            return this;
        }

        /**
         * @param defaultMount If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
         * 
         * @return builder
         * 
         */
        public Builder defaultMount(Boolean defaultMount) {
            return defaultMount(Output.of(defaultMount));
        }

        /**
         * @param defaultPathPolicy Required to be set if default_mount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:&lt;role_name&gt;.
         * 
         * @return builder
         * 
         */
        public Builder defaultPathPolicy(@Nullable Output<String> defaultPathPolicy) {
            $.defaultPathPolicy = defaultPathPolicy;
            return this;
        }

        /**
         * @param defaultPathPolicy Required to be set if default_mount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:&lt;role_name&gt;.
         * 
         * @return builder
         * 
         */
        public Builder defaultPathPolicy(String defaultPathPolicy) {
            return defaultPathPolicy(Output.of(defaultPathPolicy));
        }

        /**
         * @param enableSentinelParsing If set, parse out fields from the provided CSR making them available for Sentinel policies.
         * 
         * @return builder
         * 
         */
        public Builder enableSentinelParsing(@Nullable Output<Boolean> enableSentinelParsing) {
            $.enableSentinelParsing = enableSentinelParsing;
            return this;
        }

        /**
         * @param enableSentinelParsing If set, parse out fields from the provided CSR making them available for Sentinel policies.
         * 
         * @return builder
         * 
         */
        public Builder enableSentinelParsing(Boolean enableSentinelParsing) {
            return enableSentinelParsing(Output.of(enableSentinelParsing));
        }

        /**
         * @param enabled Specifies whether EST is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether EST is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param labelToPathPolicy Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:&lt;role_name&gt;. Labels must be unique across Vault cluster, and will register .well-known/est/&lt;label&gt; URL paths.
         * 
         * @return builder
         * 
         */
        public Builder labelToPathPolicy(@Nullable Output<Map<String,String>> labelToPathPolicy) {
            $.labelToPathPolicy = labelToPathPolicy;
            return this;
        }

        /**
         * @param labelToPathPolicy Configures a pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:&lt;role_name&gt;. Labels must be unique across Vault cluster, and will register .well-known/est/&lt;label&gt; URL paths.
         * 
         * @return builder
         * 
         */
        public Builder labelToPathPolicy(Map<String,String> labelToPathPolicy) {
            return labelToPathPolicy(Output.of(labelToPathPolicy));
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public BackendConfigEstArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("BackendConfigEstArgs", "backend");
            }
            return $;
        }
    }

}