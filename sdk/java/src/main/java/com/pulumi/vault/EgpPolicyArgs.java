// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EgpPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final EgpPolicyArgs Empty = new EgpPolicyArgs();

    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     * 
     */
    @Import(name="enforcementLevel", required=true)
    private Output<String> enforcementLevel;

    /**
     * @return Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     * 
     */
    public Output<String> enforcementLevel() {
        return this.enforcementLevel;
    }

    /**
     * The name of the policy
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the policy
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
     * List of paths to which the policy will be applied to
     * 
     */
    @Import(name="paths", required=true)
    private Output<List<String>> paths;

    /**
     * @return List of paths to which the policy will be applied to
     * 
     */
    public Output<List<String>> paths() {
        return this.paths;
    }

    /**
     * String containing a Sentinel policy
     * 
     */
    @Import(name="policy", required=true)
    private Output<String> policy;

    /**
     * @return String containing a Sentinel policy
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    private EgpPolicyArgs() {}

    private EgpPolicyArgs(EgpPolicyArgs $) {
        this.enforcementLevel = $.enforcementLevel;
        this.name = $.name;
        this.namespace = $.namespace;
        this.paths = $.paths;
        this.policy = $.policy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EgpPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EgpPolicyArgs $;

        public Builder() {
            $ = new EgpPolicyArgs();
        }

        public Builder(EgpPolicyArgs defaults) {
            $ = new EgpPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enforcementLevel Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
         * 
         * @return builder
         * 
         */
        public Builder enforcementLevel(Output<String> enforcementLevel) {
            $.enforcementLevel = enforcementLevel;
            return this;
        }

        /**
         * @param enforcementLevel Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
         * 
         * @return builder
         * 
         */
        public Builder enforcementLevel(String enforcementLevel) {
            return enforcementLevel(Output.of(enforcementLevel));
        }

        /**
         * @param name The name of the policy
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the policy
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
         * @param paths List of paths to which the policy will be applied to
         * 
         * @return builder
         * 
         */
        public Builder paths(Output<List<String>> paths) {
            $.paths = paths;
            return this;
        }

        /**
         * @param paths List of paths to which the policy will be applied to
         * 
         * @return builder
         * 
         */
        public Builder paths(List<String> paths) {
            return paths(Output.of(paths));
        }

        /**
         * @param paths List of paths to which the policy will be applied to
         * 
         * @return builder
         * 
         */
        public Builder paths(String... paths) {
            return paths(List.of(paths));
        }

        /**
         * @param policy String containing a Sentinel policy
         * 
         * @return builder
         * 
         */
        public Builder policy(Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy String containing a Sentinel policy
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public EgpPolicyArgs build() {
            if ($.enforcementLevel == null) {
                throw new MissingRequiredPropertyException("EgpPolicyArgs", "enforcementLevel");
            }
            if ($.paths == null) {
                throw new MissingRequiredPropertyException("EgpPolicyArgs", "paths");
            }
            if ($.policy == null) {
                throw new MissingRequiredPropertyException("EgpPolicyArgs", "policy");
            }
            return $;
        }
    }

}
