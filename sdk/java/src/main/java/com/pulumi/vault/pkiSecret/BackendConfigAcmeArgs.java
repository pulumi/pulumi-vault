// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackendConfigAcmeArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackendConfigAcmeArgs Empty = new BackendConfigAcmeArgs();

    /**
     * Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
     * 
     */
    @Import(name="allowRoleExtKeyUsage")
    private @Nullable Output<Boolean> allowRoleExtKeyUsage;

    /**
     * @return Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
     * 
     */
    public Optional<Output<Boolean>> allowRoleExtKeyUsage() {
        return Optional.ofNullable(this.allowRoleExtKeyUsage);
    }

    /**
     * Specifies which issuers are allowed for use with ACME.
     * 
     */
    @Import(name="allowedIssuers")
    private @Nullable Output<List<String>> allowedIssuers;

    /**
     * @return Specifies which issuers are allowed for use with ACME.
     * 
     */
    public Optional<Output<List<String>>> allowedIssuers() {
        return Optional.ofNullable(this.allowedIssuers);
    }

    /**
     * Specifies which roles are allowed for use with ACME.
     * 
     */
    @Import(name="allowedRoles")
    private @Nullable Output<List<String>> allowedRoles;

    /**
     * @return Specifies which roles are allowed for use with ACME.
     * 
     */
    public Optional<Output<List<String>>> allowedRoles() {
        return Optional.ofNullable(this.allowedRoles);
    }

    /**
     * The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * Specifies the policy to be used for non-role-qualified ACME requests.
     * Allowed values are `forbid`, `sign-verbatim`, `role:&lt;role_name&gt;`, `external-policy` or `external-policy:&lt;policy&gt;`.
     * 
     */
    @Import(name="defaultDirectoryPolicy")
    private @Nullable Output<String> defaultDirectoryPolicy;

    /**
     * @return Specifies the policy to be used for non-role-qualified ACME requests.
     * Allowed values are `forbid`, `sign-verbatim`, `role:&lt;role_name&gt;`, `external-policy` or `external-policy:&lt;policy&gt;`.
     * 
     */
    public Optional<Output<String>> defaultDirectoryPolicy() {
        return Optional.ofNullable(this.defaultDirectoryPolicy);
    }

    /**
     * DNS resolver to use for domain resolution on this mount.
     * Must be in the format `&lt;host&gt;:&lt;port&gt;`, with both parts mandatory.
     * 
     */
    @Import(name="dnsResolver")
    private @Nullable Output<String> dnsResolver;

    /**
     * @return DNS resolver to use for domain resolution on this mount.
     * Must be in the format `&lt;host&gt;:&lt;port&gt;`, with both parts mandatory.
     * 
     */
    public Optional<Output<String>> dnsResolver() {
        return Optional.ofNullable(this.dnsResolver);
    }

    /**
     * Specifies the policy to use for external account binding behaviour.
     * Allowed values are `not-required`, `new-account-required` or `always-required`.
     * 
     */
    @Import(name="eabPolicy")
    private @Nullable Output<String> eabPolicy;

    /**
     * @return Specifies the policy to use for external account binding behaviour.
     * Allowed values are `not-required`, `new-account-required` or `always-required`.
     * 
     */
    public Optional<Output<String>> eabPolicy() {
        return Optional.ofNullable(this.eabPolicy);
    }

    /**
     * Specifies whether ACME is enabled.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Specifies whether ACME is enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * The maximum TTL in seconds for certificates issued by ACME. **Vault 1.17.0+**
     * 
     */
    @Import(name="maxTtl")
    private @Nullable Output<Integer> maxTtl;

    /**
     * @return The maximum TTL in seconds for certificates issued by ACME. **Vault 1.17.0+**
     * 
     */
    public Optional<Output<Integer>> maxTtl() {
        return Optional.ofNullable(this.maxTtl);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    private BackendConfigAcmeArgs() {}

    private BackendConfigAcmeArgs(BackendConfigAcmeArgs $) {
        this.allowRoleExtKeyUsage = $.allowRoleExtKeyUsage;
        this.allowedIssuers = $.allowedIssuers;
        this.allowedRoles = $.allowedRoles;
        this.backend = $.backend;
        this.defaultDirectoryPolicy = $.defaultDirectoryPolicy;
        this.dnsResolver = $.dnsResolver;
        this.eabPolicy = $.eabPolicy;
        this.enabled = $.enabled;
        this.maxTtl = $.maxTtl;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendConfigAcmeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendConfigAcmeArgs $;

        public Builder() {
            $ = new BackendConfigAcmeArgs();
        }

        public Builder(BackendConfigAcmeArgs defaults) {
            $ = new BackendConfigAcmeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowRoleExtKeyUsage Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
         * 
         * @return builder
         * 
         */
        public Builder allowRoleExtKeyUsage(@Nullable Output<Boolean> allowRoleExtKeyUsage) {
            $.allowRoleExtKeyUsage = allowRoleExtKeyUsage;
            return this;
        }

        /**
         * @param allowRoleExtKeyUsage Specifies whether the ExtKeyUsage field from a role is used. **Vault 1.14.1+**
         * 
         * @return builder
         * 
         */
        public Builder allowRoleExtKeyUsage(Boolean allowRoleExtKeyUsage) {
            return allowRoleExtKeyUsage(Output.of(allowRoleExtKeyUsage));
        }

        /**
         * @param allowedIssuers Specifies which issuers are allowed for use with ACME.
         * 
         * @return builder
         * 
         */
        public Builder allowedIssuers(@Nullable Output<List<String>> allowedIssuers) {
            $.allowedIssuers = allowedIssuers;
            return this;
        }

        /**
         * @param allowedIssuers Specifies which issuers are allowed for use with ACME.
         * 
         * @return builder
         * 
         */
        public Builder allowedIssuers(List<String> allowedIssuers) {
            return allowedIssuers(Output.of(allowedIssuers));
        }

        /**
         * @param allowedIssuers Specifies which issuers are allowed for use with ACME.
         * 
         * @return builder
         * 
         */
        public Builder allowedIssuers(String... allowedIssuers) {
            return allowedIssuers(List.of(allowedIssuers));
        }

        /**
         * @param allowedRoles Specifies which roles are allowed for use with ACME.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(@Nullable Output<List<String>> allowedRoles) {
            $.allowedRoles = allowedRoles;
            return this;
        }

        /**
         * @param allowedRoles Specifies which roles are allowed for use with ACME.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(List<String> allowedRoles) {
            return allowedRoles(Output.of(allowedRoles));
        }

        /**
         * @param allowedRoles Specifies which roles are allowed for use with ACME.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(String... allowedRoles) {
            return allowedRoles(List.of(allowedRoles));
        }

        /**
         * @param backend The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param defaultDirectoryPolicy Specifies the policy to be used for non-role-qualified ACME requests.
         * Allowed values are `forbid`, `sign-verbatim`, `role:&lt;role_name&gt;`, `external-policy` or `external-policy:&lt;policy&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder defaultDirectoryPolicy(@Nullable Output<String> defaultDirectoryPolicy) {
            $.defaultDirectoryPolicy = defaultDirectoryPolicy;
            return this;
        }

        /**
         * @param defaultDirectoryPolicy Specifies the policy to be used for non-role-qualified ACME requests.
         * Allowed values are `forbid`, `sign-verbatim`, `role:&lt;role_name&gt;`, `external-policy` or `external-policy:&lt;policy&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder defaultDirectoryPolicy(String defaultDirectoryPolicy) {
            return defaultDirectoryPolicy(Output.of(defaultDirectoryPolicy));
        }

        /**
         * @param dnsResolver DNS resolver to use for domain resolution on this mount.
         * Must be in the format `&lt;host&gt;:&lt;port&gt;`, with both parts mandatory.
         * 
         * @return builder
         * 
         */
        public Builder dnsResolver(@Nullable Output<String> dnsResolver) {
            $.dnsResolver = dnsResolver;
            return this;
        }

        /**
         * @param dnsResolver DNS resolver to use for domain resolution on this mount.
         * Must be in the format `&lt;host&gt;:&lt;port&gt;`, with both parts mandatory.
         * 
         * @return builder
         * 
         */
        public Builder dnsResolver(String dnsResolver) {
            return dnsResolver(Output.of(dnsResolver));
        }

        /**
         * @param eabPolicy Specifies the policy to use for external account binding behaviour.
         * Allowed values are `not-required`, `new-account-required` or `always-required`.
         * 
         * @return builder
         * 
         */
        public Builder eabPolicy(@Nullable Output<String> eabPolicy) {
            $.eabPolicy = eabPolicy;
            return this;
        }

        /**
         * @param eabPolicy Specifies the policy to use for external account binding behaviour.
         * Allowed values are `not-required`, `new-account-required` or `always-required`.
         * 
         * @return builder
         * 
         */
        public Builder eabPolicy(String eabPolicy) {
            return eabPolicy(Output.of(eabPolicy));
        }

        /**
         * @param enabled Specifies whether ACME is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether ACME is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param maxTtl The maximum TTL in seconds for certificates issued by ACME. **Vault 1.17.0+**
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(@Nullable Output<Integer> maxTtl) {
            $.maxTtl = maxTtl;
            return this;
        }

        /**
         * @param maxTtl The maximum TTL in seconds for certificates issued by ACME. **Vault 1.17.0+**
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(Integer maxTtl) {
            return maxTtl(Output.of(maxTtl));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public BackendConfigAcmeArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("BackendConfigAcmeArgs", "backend");
            }
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("BackendConfigAcmeArgs", "enabled");
            }
            return $;
        }
    }

}
