// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendRoleArgs Empty = new SecretBackendRoleArgs();

    /**
     * The unique name of the Vault mount to configure.
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The unique name of the Vault mount to configure.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * The database statements to execute when
     * creating a user.
     * 
     */
    @Import(name="creationStatements", required=true)
    private Output<List<String>> creationStatements;

    /**
     * @return The database statements to execute when
     * creating a user.
     * 
     */
    public Output<List<String>> creationStatements() {
        return this.creationStatements;
    }

    /**
     * Specifies the configuration
     * for the given `credential_type`.
     * 
     * The following options are available for each `credential_type` value:
     * 
     */
    @Import(name="credentialConfig")
    private @Nullable Output<Map<String,Object>> credentialConfig;

    /**
     * @return Specifies the configuration
     * for the given `credential_type`.
     * 
     * The following options are available for each `credential_type` value:
     * 
     */
    public Optional<Output<Map<String,Object>>> credentialConfig() {
        return Optional.ofNullable(this.credentialConfig);
    }

    /**
     * Specifies the type of credential that
     * will be generated for the role. Options include: `password`, `rsa_private_key`, `client_certificate`.
     * See the plugin&#39;s API page for credential types supported by individual databases.
     * 
     */
    @Import(name="credentialType")
    private @Nullable Output<String> credentialType;

    /**
     * @return Specifies the type of credential that
     * will be generated for the role. Options include: `password`, `rsa_private_key`, `client_certificate`.
     * See the plugin&#39;s API page for credential types supported by individual databases.
     * 
     */
    public Optional<Output<String>> credentialType() {
        return Optional.ofNullable(this.credentialType);
    }

    /**
     * The unique name of the database connection to use for
     * the role.
     * 
     */
    @Import(name="dbName", required=true)
    private Output<String> dbName;

    /**
     * @return The unique name of the database connection to use for
     * the role.
     * 
     */
    public Output<String> dbName() {
        return this.dbName;
    }

    /**
     * The default number of seconds for leases for this
     * role.
     * 
     */
    @Import(name="defaultTtl")
    private @Nullable Output<Integer> defaultTtl;

    /**
     * @return The default number of seconds for leases for this
     * role.
     * 
     */
    public Optional<Output<Integer>> defaultTtl() {
        return Optional.ofNullable(this.defaultTtl);
    }

    /**
     * The maximum number of seconds for leases for this
     * role.
     * 
     */
    @Import(name="maxTtl")
    private @Nullable Output<Integer> maxTtl;

    /**
     * @return The maximum number of seconds for leases for this
     * role.
     * 
     */
    public Optional<Output<Integer>> maxTtl() {
        return Optional.ofNullable(this.maxTtl);
    }

    /**
     * A unique name to give the role.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name to give the role.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The database statements to execute when
     * renewing a user.
     * 
     */
    @Import(name="renewStatements")
    private @Nullable Output<List<String>> renewStatements;

    /**
     * @return The database statements to execute when
     * renewing a user.
     * 
     */
    public Optional<Output<List<String>>> renewStatements() {
        return Optional.ofNullable(this.renewStatements);
    }

    /**
     * The database statements to execute when
     * revoking a user.
     * 
     */
    @Import(name="revocationStatements")
    private @Nullable Output<List<String>> revocationStatements;

    /**
     * @return The database statements to execute when
     * revoking a user.
     * 
     */
    public Optional<Output<List<String>>> revocationStatements() {
        return Optional.ofNullable(this.revocationStatements);
    }

    /**
     * The database statements to execute when
     * rolling back creation due to an error.
     * 
     */
    @Import(name="rollbackStatements")
    private @Nullable Output<List<String>> rollbackStatements;

    /**
     * @return The database statements to execute when
     * rolling back creation due to an error.
     * 
     */
    public Optional<Output<List<String>>> rollbackStatements() {
        return Optional.ofNullable(this.rollbackStatements);
    }

    private SecretBackendRoleArgs() {}

    private SecretBackendRoleArgs(SecretBackendRoleArgs $) {
        this.backend = $.backend;
        this.creationStatements = $.creationStatements;
        this.credentialConfig = $.credentialConfig;
        this.credentialType = $.credentialType;
        this.dbName = $.dbName;
        this.defaultTtl = $.defaultTtl;
        this.maxTtl = $.maxTtl;
        this.name = $.name;
        this.namespace = $.namespace;
        this.renewStatements = $.renewStatements;
        this.revocationStatements = $.revocationStatements;
        this.rollbackStatements = $.rollbackStatements;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendRoleArgs $;

        public Builder() {
            $ = new SecretBackendRoleArgs();
        }

        public Builder(SecretBackendRoleArgs defaults) {
            $ = new SecretBackendRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The unique name of the Vault mount to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The unique name of the Vault mount to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param creationStatements The database statements to execute when
         * creating a user.
         * 
         * @return builder
         * 
         */
        public Builder creationStatements(Output<List<String>> creationStatements) {
            $.creationStatements = creationStatements;
            return this;
        }

        /**
         * @param creationStatements The database statements to execute when
         * creating a user.
         * 
         * @return builder
         * 
         */
        public Builder creationStatements(List<String> creationStatements) {
            return creationStatements(Output.of(creationStatements));
        }

        /**
         * @param creationStatements The database statements to execute when
         * creating a user.
         * 
         * @return builder
         * 
         */
        public Builder creationStatements(String... creationStatements) {
            return creationStatements(List.of(creationStatements));
        }

        /**
         * @param credentialConfig Specifies the configuration
         * for the given `credential_type`.
         * 
         * The following options are available for each `credential_type` value:
         * 
         * @return builder
         * 
         */
        public Builder credentialConfig(@Nullable Output<Map<String,Object>> credentialConfig) {
            $.credentialConfig = credentialConfig;
            return this;
        }

        /**
         * @param credentialConfig Specifies the configuration
         * for the given `credential_type`.
         * 
         * The following options are available for each `credential_type` value:
         * 
         * @return builder
         * 
         */
        public Builder credentialConfig(Map<String,Object> credentialConfig) {
            return credentialConfig(Output.of(credentialConfig));
        }

        /**
         * @param credentialType Specifies the type of credential that
         * will be generated for the role. Options include: `password`, `rsa_private_key`, `client_certificate`.
         * See the plugin&#39;s API page for credential types supported by individual databases.
         * 
         * @return builder
         * 
         */
        public Builder credentialType(@Nullable Output<String> credentialType) {
            $.credentialType = credentialType;
            return this;
        }

        /**
         * @param credentialType Specifies the type of credential that
         * will be generated for the role. Options include: `password`, `rsa_private_key`, `client_certificate`.
         * See the plugin&#39;s API page for credential types supported by individual databases.
         * 
         * @return builder
         * 
         */
        public Builder credentialType(String credentialType) {
            return credentialType(Output.of(credentialType));
        }

        /**
         * @param dbName The unique name of the database connection to use for
         * the role.
         * 
         * @return builder
         * 
         */
        public Builder dbName(Output<String> dbName) {
            $.dbName = dbName;
            return this;
        }

        /**
         * @param dbName The unique name of the database connection to use for
         * the role.
         * 
         * @return builder
         * 
         */
        public Builder dbName(String dbName) {
            return dbName(Output.of(dbName));
        }

        /**
         * @param defaultTtl The default number of seconds for leases for this
         * role.
         * 
         * @return builder
         * 
         */
        public Builder defaultTtl(@Nullable Output<Integer> defaultTtl) {
            $.defaultTtl = defaultTtl;
            return this;
        }

        /**
         * @param defaultTtl The default number of seconds for leases for this
         * role.
         * 
         * @return builder
         * 
         */
        public Builder defaultTtl(Integer defaultTtl) {
            return defaultTtl(Output.of(defaultTtl));
        }

        /**
         * @param maxTtl The maximum number of seconds for leases for this
         * role.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(@Nullable Output<Integer> maxTtl) {
            $.maxTtl = maxTtl;
            return this;
        }

        /**
         * @param maxTtl The maximum number of seconds for leases for this
         * role.
         * 
         * @return builder
         * 
         */
        public Builder maxTtl(Integer maxTtl) {
            return maxTtl(Output.of(maxTtl));
        }

        /**
         * @param name A unique name to give the role.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name to give the role.
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
         * The `namespace` is always relative to the provider&#39;s configured namespace.
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
         * The `namespace` is always relative to the provider&#39;s configured namespace.
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param renewStatements The database statements to execute when
         * renewing a user.
         * 
         * @return builder
         * 
         */
        public Builder renewStatements(@Nullable Output<List<String>> renewStatements) {
            $.renewStatements = renewStatements;
            return this;
        }

        /**
         * @param renewStatements The database statements to execute when
         * renewing a user.
         * 
         * @return builder
         * 
         */
        public Builder renewStatements(List<String> renewStatements) {
            return renewStatements(Output.of(renewStatements));
        }

        /**
         * @param renewStatements The database statements to execute when
         * renewing a user.
         * 
         * @return builder
         * 
         */
        public Builder renewStatements(String... renewStatements) {
            return renewStatements(List.of(renewStatements));
        }

        /**
         * @param revocationStatements The database statements to execute when
         * revoking a user.
         * 
         * @return builder
         * 
         */
        public Builder revocationStatements(@Nullable Output<List<String>> revocationStatements) {
            $.revocationStatements = revocationStatements;
            return this;
        }

        /**
         * @param revocationStatements The database statements to execute when
         * revoking a user.
         * 
         * @return builder
         * 
         */
        public Builder revocationStatements(List<String> revocationStatements) {
            return revocationStatements(Output.of(revocationStatements));
        }

        /**
         * @param revocationStatements The database statements to execute when
         * revoking a user.
         * 
         * @return builder
         * 
         */
        public Builder revocationStatements(String... revocationStatements) {
            return revocationStatements(List.of(revocationStatements));
        }

        /**
         * @param rollbackStatements The database statements to execute when
         * rolling back creation due to an error.
         * 
         * @return builder
         * 
         */
        public Builder rollbackStatements(@Nullable Output<List<String>> rollbackStatements) {
            $.rollbackStatements = rollbackStatements;
            return this;
        }

        /**
         * @param rollbackStatements The database statements to execute when
         * rolling back creation due to an error.
         * 
         * @return builder
         * 
         */
        public Builder rollbackStatements(List<String> rollbackStatements) {
            return rollbackStatements(Output.of(rollbackStatements));
        }

        /**
         * @param rollbackStatements The database statements to execute when
         * rolling back creation due to an error.
         * 
         * @return builder
         * 
         */
        public Builder rollbackStatements(String... rollbackStatements) {
            return rollbackStatements(List.of(rollbackStatements));
        }

        public SecretBackendRoleArgs build() {
            $.backend = Objects.requireNonNull($.backend, "expected parameter 'backend' to be non-null");
            $.creationStatements = Objects.requireNonNull($.creationStatements, "expected parameter 'creationStatements' to be non-null");
            $.dbName = Objects.requireNonNull($.dbName, "expected parameter 'dbName' to be non-null");
            return $;
        }
    }

}
