// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap;

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


public final class SecretBackendLibrarySetArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendLibrarySetArgs Empty = new SecretBackendLibrarySetArgs();

    /**
     * Disable enforcing that service
     * accounts must be checked in by the entity or client token that checked them
     * out. Defaults to false.
     * 
     */
    @Import(name="disableCheckInEnforcement")
    private @Nullable Output<Boolean> disableCheckInEnforcement;

    /**
     * @return Disable enforcing that service
     * accounts must be checked in by the entity or client token that checked them
     * out. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> disableCheckInEnforcement() {
        return Optional.ofNullable(this.disableCheckInEnforcement);
    }

    /**
     * The maximum password time-to-live in seconds. Defaults
     * to the configuration max_ttl if not provided.
     * 
     */
    @Import(name="maxTtl")
    private @Nullable Output<Integer> maxTtl;

    /**
     * @return The maximum password time-to-live in seconds. Defaults
     * to the configuration max_ttl if not provided.
     * 
     */
    public Optional<Output<Integer>> maxTtl() {
        return Optional.ofNullable(this.maxTtl);
    }

    /**
     * The path where the LDAP secrets backend is mounted.
     * 
     */
    @Import(name="mount")
    private @Nullable Output<String> mount;

    /**
     * @return The path where the LDAP secrets backend is mounted.
     * 
     */
    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
    }

    /**
     * The name to identify this set of service accounts.
     * Must be unique within the backend.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name to identify this set of service accounts.
     * Must be unique within the backend.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Specifies the slice of service accounts mapped to this set.
     * 
     */
    @Import(name="serviceAccountNames", required=true)
    private Output<List<String>> serviceAccountNames;

    /**
     * @return Specifies the slice of service accounts mapped to this set.
     * 
     */
    public Output<List<String>> serviceAccountNames() {
        return this.serviceAccountNames;
    }

    /**
     * The password time-to-live in seconds. Defaults to the configuration
     * ttl if not provided.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The password time-to-live in seconds. Defaults to the configuration
     * ttl if not provided.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private SecretBackendLibrarySetArgs() {}

    private SecretBackendLibrarySetArgs(SecretBackendLibrarySetArgs $) {
        this.disableCheckInEnforcement = $.disableCheckInEnforcement;
        this.maxTtl = $.maxTtl;
        this.mount = $.mount;
        this.name = $.name;
        this.namespace = $.namespace;
        this.serviceAccountNames = $.serviceAccountNames;
        this.ttl = $.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendLibrarySetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendLibrarySetArgs $;

        public Builder() {
            $ = new SecretBackendLibrarySetArgs();
        }

        public Builder(SecretBackendLibrarySetArgs defaults) {
            $ = new SecretBackendLibrarySetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disableCheckInEnforcement Disable enforcing that service
         * accounts must be checked in by the entity or client token that checked them
         * out. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder disableCheckInEnforcement(@Nullable Output<Boolean> disableCheckInEnforcement) {
            $.disableCheckInEnforcement = disableCheckInEnforcement;
            return this;
        }

        /**
         * @param disableCheckInEnforcement Disable enforcing that service
         * accounts must be checked in by the entity or client token that checked them
         * out. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder disableCheckInEnforcement(Boolean disableCheckInEnforcement) {
            return disableCheckInEnforcement(Output.of(disableCheckInEnforcement));
        }

        /**
         * @param maxTtl The maximum password time-to-live in seconds. Defaults
         * to the configuration max_ttl if not provided.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(@Nullable Output<Integer> maxTtl) {
            $.maxTtl = maxTtl;
            return this;
        }

        /**
         * @param maxTtl The maximum password time-to-live in seconds. Defaults
         * to the configuration max_ttl if not provided.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(Integer maxTtl) {
            return maxTtl(Output.of(maxTtl));
        }

        /**
         * @param mount The path where the LDAP secrets backend is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount The path where the LDAP secrets backend is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
        }

        /**
         * @param name The name to identify this set of service accounts.
         * Must be unique within the backend.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name to identify this set of service accounts.
         * Must be unique within the backend.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespace The namespace to provision the resource in.
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
         * @param namespace The namespace to provision the resource in.
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

        /**
         * @param serviceAccountNames Specifies the slice of service accounts mapped to this set.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountNames(Output<List<String>> serviceAccountNames) {
            $.serviceAccountNames = serviceAccountNames;
            return this;
        }

        /**
         * @param serviceAccountNames Specifies the slice of service accounts mapped to this set.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountNames(List<String> serviceAccountNames) {
            return serviceAccountNames(Output.of(serviceAccountNames));
        }

        /**
         * @param serviceAccountNames Specifies the slice of service accounts mapped to this set.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountNames(String... serviceAccountNames) {
            return serviceAccountNames(List.of(serviceAccountNames));
        }

        /**
         * @param ttl The password time-to-live in seconds. Defaults to the configuration
         * ttl if not provided.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The password time-to-live in seconds. Defaults to the configuration
         * ttl if not provided.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        public SecretBackendLibrarySetArgs build() {
            if ($.serviceAccountNames == null) {
                throw new MissingRequiredPropertyException("SecretBackendLibrarySetArgs", "serviceAccountNames");
            }
            return $;
        }
    }

}
