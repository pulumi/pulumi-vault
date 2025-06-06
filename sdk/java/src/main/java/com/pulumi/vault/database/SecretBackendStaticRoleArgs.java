// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendStaticRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendStaticRoleArgs Empty = new SecretBackendStaticRoleArgs();

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

    @Import(name="credentialConfig")
    private @Nullable Output<Map<String,String>> credentialConfig;

    public Optional<Output<Map<String,String>>> credentialConfig() {
        return Optional.ofNullable(this.credentialConfig);
    }

    /**
     * The credential type for the user, can be one of &#34;password&#34;, &#34;rsa_private_key&#34; or &#34;client_certificate&#34;.The configuration
     * can be done in `credential_config`.
     * 
     */
    @Import(name="credentialType")
    private @Nullable Output<String> credentialType;

    /**
     * @return The credential type for the user, can be one of &#34;password&#34;, &#34;rsa_private_key&#34; or &#34;client_certificate&#34;.The configuration
     * can be done in `credential_config`.
     * 
     */
    public Optional<Output<String>> credentialType() {
        return Optional.ofNullable(this.credentialType);
    }

    /**
     * The unique name of the database connection to use for the static role.
     * 
     */
    @Import(name="dbName", required=true)
    private Output<String> dbName;

    /**
     * @return The unique name of the database connection to use for the static role.
     * 
     */
    public Output<String> dbName() {
        return this.dbName;
    }

    /**
     * A unique name to give the static role.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name to give the static role.
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
     * The amount of time Vault should wait before rotating the password, in seconds.
     * Mutually exclusive with `rotation_schedule`.
     * 
     */
    @Import(name="rotationPeriod")
    private @Nullable Output<Integer> rotationPeriod;

    /**
     * @return The amount of time Vault should wait before rotating the password, in seconds.
     * Mutually exclusive with `rotation_schedule`.
     * 
     */
    public Optional<Output<Integer>> rotationPeriod() {
        return Optional.ofNullable(this.rotationPeriod);
    }

    /**
     * A cron-style string that will define the schedule on which rotations should occur.
     * Mutually exclusive with `rotation_period`.
     * 
     * **Warning**: The `rotation_period` and `rotation_schedule` fields are
     * mutually exclusive. One of them must be set but not both.
     * 
     */
    @Import(name="rotationSchedule")
    private @Nullable Output<String> rotationSchedule;

    /**
     * @return A cron-style string that will define the schedule on which rotations should occur.
     * Mutually exclusive with `rotation_period`.
     * 
     * **Warning**: The `rotation_period` and `rotation_schedule` fields are
     * mutually exclusive. One of them must be set but not both.
     * 
     */
    public Optional<Output<String>> rotationSchedule() {
        return Optional.ofNullable(this.rotationSchedule);
    }

    /**
     * Database statements to execute to rotate the password for the configured database user.
     * 
     */
    @Import(name="rotationStatements")
    private @Nullable Output<List<String>> rotationStatements;

    /**
     * @return Database statements to execute to rotate the password for the configured database user.
     * 
     */
    public Optional<Output<List<String>>> rotationStatements() {
        return Optional.ofNullable(this.rotationStatements);
    }

    /**
     * The amount of time, in seconds, in which rotations are allowed to occur starting
     * from a given `rotation_schedule`.
     * 
     */
    @Import(name="rotationWindow")
    private @Nullable Output<Integer> rotationWindow;

    /**
     * @return The amount of time, in seconds, in which rotations are allowed to occur starting
     * from a given `rotation_schedule`.
     * 
     */
    public Optional<Output<Integer>> rotationWindow() {
        return Optional.ofNullable(this.rotationWindow);
    }

    /**
     * The password corresponding to the username in the database.
     * Required when using the Rootless Password Rotation workflow for static roles. Only enabled for
     * select DB engines (Postgres). Requires Vault 1.18+ Enterprise.
     * 
     */
    @Import(name="selfManagedPassword")
    private @Nullable Output<String> selfManagedPassword;

    /**
     * @return The password corresponding to the username in the database.
     * Required when using the Rootless Password Rotation workflow for static roles. Only enabled for
     * select DB engines (Postgres). Requires Vault 1.18+ Enterprise.
     * 
     */
    public Optional<Output<String>> selfManagedPassword() {
        return Optional.ofNullable(this.selfManagedPassword);
    }

    /**
     * If set to true, Vault will skip the
     * initial secret rotation on import. Requires Vault 1.18+ Enterprise.
     * 
     */
    @Import(name="skipImportRotation")
    private @Nullable Output<Boolean> skipImportRotation;

    /**
     * @return If set to true, Vault will skip the
     * initial secret rotation on import. Requires Vault 1.18+ Enterprise.
     * 
     */
    public Optional<Output<Boolean>> skipImportRotation() {
        return Optional.ofNullable(this.skipImportRotation);
    }

    /**
     * The database username that this static role corresponds to.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The database username that this static role corresponds to.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private SecretBackendStaticRoleArgs() {}

    private SecretBackendStaticRoleArgs(SecretBackendStaticRoleArgs $) {
        this.backend = $.backend;
        this.credentialConfig = $.credentialConfig;
        this.credentialType = $.credentialType;
        this.dbName = $.dbName;
        this.name = $.name;
        this.namespace = $.namespace;
        this.rotationPeriod = $.rotationPeriod;
        this.rotationSchedule = $.rotationSchedule;
        this.rotationStatements = $.rotationStatements;
        this.rotationWindow = $.rotationWindow;
        this.selfManagedPassword = $.selfManagedPassword;
        this.skipImportRotation = $.skipImportRotation;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendStaticRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendStaticRoleArgs $;

        public Builder() {
            $ = new SecretBackendStaticRoleArgs();
        }

        public Builder(SecretBackendStaticRoleArgs defaults) {
            $ = new SecretBackendStaticRoleArgs(Objects.requireNonNull(defaults));
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

        public Builder credentialConfig(@Nullable Output<Map<String,String>> credentialConfig) {
            $.credentialConfig = credentialConfig;
            return this;
        }

        public Builder credentialConfig(Map<String,String> credentialConfig) {
            return credentialConfig(Output.of(credentialConfig));
        }

        /**
         * @param credentialType The credential type for the user, can be one of &#34;password&#34;, &#34;rsa_private_key&#34; or &#34;client_certificate&#34;.The configuration
         * can be done in `credential_config`.
         * 
         * @return builder
         * 
         */
        public Builder credentialType(@Nullable Output<String> credentialType) {
            $.credentialType = credentialType;
            return this;
        }

        /**
         * @param credentialType The credential type for the user, can be one of &#34;password&#34;, &#34;rsa_private_key&#34; or &#34;client_certificate&#34;.The configuration
         * can be done in `credential_config`.
         * 
         * @return builder
         * 
         */
        public Builder credentialType(String credentialType) {
            return credentialType(Output.of(credentialType));
        }

        /**
         * @param dbName The unique name of the database connection to use for the static role.
         * 
         * @return builder
         * 
         */
        public Builder dbName(Output<String> dbName) {
            $.dbName = dbName;
            return this;
        }

        /**
         * @param dbName The unique name of the database connection to use for the static role.
         * 
         * @return builder
         * 
         */
        public Builder dbName(String dbName) {
            return dbName(Output.of(dbName));
        }

        /**
         * @param name A unique name to give the static role.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name to give the static role.
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
         * @param rotationPeriod The amount of time Vault should wait before rotating the password, in seconds.
         * Mutually exclusive with `rotation_schedule`.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(@Nullable Output<Integer> rotationPeriod) {
            $.rotationPeriod = rotationPeriod;
            return this;
        }

        /**
         * @param rotationPeriod The amount of time Vault should wait before rotating the password, in seconds.
         * Mutually exclusive with `rotation_schedule`.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Integer rotationPeriod) {
            return rotationPeriod(Output.of(rotationPeriod));
        }

        /**
         * @param rotationSchedule A cron-style string that will define the schedule on which rotations should occur.
         * Mutually exclusive with `rotation_period`.
         * 
         * **Warning**: The `rotation_period` and `rotation_schedule` fields are
         * mutually exclusive. One of them must be set but not both.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(@Nullable Output<String> rotationSchedule) {
            $.rotationSchedule = rotationSchedule;
            return this;
        }

        /**
         * @param rotationSchedule A cron-style string that will define the schedule on which rotations should occur.
         * Mutually exclusive with `rotation_period`.
         * 
         * **Warning**: The `rotation_period` and `rotation_schedule` fields are
         * mutually exclusive. One of them must be set but not both.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(String rotationSchedule) {
            return rotationSchedule(Output.of(rotationSchedule));
        }

        /**
         * @param rotationStatements Database statements to execute to rotate the password for the configured database user.
         * 
         * @return builder
         * 
         */
        public Builder rotationStatements(@Nullable Output<List<String>> rotationStatements) {
            $.rotationStatements = rotationStatements;
            return this;
        }

        /**
         * @param rotationStatements Database statements to execute to rotate the password for the configured database user.
         * 
         * @return builder
         * 
         */
        public Builder rotationStatements(List<String> rotationStatements) {
            return rotationStatements(Output.of(rotationStatements));
        }

        /**
         * @param rotationStatements Database statements to execute to rotate the password for the configured database user.
         * 
         * @return builder
         * 
         */
        public Builder rotationStatements(String... rotationStatements) {
            return rotationStatements(List.of(rotationStatements));
        }

        /**
         * @param rotationWindow The amount of time, in seconds, in which rotations are allowed to occur starting
         * from a given `rotation_schedule`.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(@Nullable Output<Integer> rotationWindow) {
            $.rotationWindow = rotationWindow;
            return this;
        }

        /**
         * @param rotationWindow The amount of time, in seconds, in which rotations are allowed to occur starting
         * from a given `rotation_schedule`.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(Integer rotationWindow) {
            return rotationWindow(Output.of(rotationWindow));
        }

        /**
         * @param selfManagedPassword The password corresponding to the username in the database.
         * Required when using the Rootless Password Rotation workflow for static roles. Only enabled for
         * select DB engines (Postgres). Requires Vault 1.18+ Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder selfManagedPassword(@Nullable Output<String> selfManagedPassword) {
            $.selfManagedPassword = selfManagedPassword;
            return this;
        }

        /**
         * @param selfManagedPassword The password corresponding to the username in the database.
         * Required when using the Rootless Password Rotation workflow for static roles. Only enabled for
         * select DB engines (Postgres). Requires Vault 1.18+ Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder selfManagedPassword(String selfManagedPassword) {
            return selfManagedPassword(Output.of(selfManagedPassword));
        }

        /**
         * @param skipImportRotation If set to true, Vault will skip the
         * initial secret rotation on import. Requires Vault 1.18+ Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder skipImportRotation(@Nullable Output<Boolean> skipImportRotation) {
            $.skipImportRotation = skipImportRotation;
            return this;
        }

        /**
         * @param skipImportRotation If set to true, Vault will skip the
         * initial secret rotation on import. Requires Vault 1.18+ Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder skipImportRotation(Boolean skipImportRotation) {
            return skipImportRotation(Output.of(skipImportRotation));
        }

        /**
         * @param username The database username that this static role corresponds to.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The database username that this static role corresponds to.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public SecretBackendStaticRoleArgs build() {
            if ($.backend == null) {
                throw new MissingRequiredPropertyException("SecretBackendStaticRoleArgs", "backend");
            }
            if ($.dbName == null) {
                throw new MissingRequiredPropertyException("SecretBackendStaticRoleArgs", "dbName");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("SecretBackendStaticRoleArgs", "username");
            }
            return $;
        }
    }

}
