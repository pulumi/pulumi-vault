// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendStaticRoleState extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendStaticRoleState Empty = new SecretBackendStaticRoleState();

    /**
     * The unique name of the Vault mount to configure.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The unique name of the Vault mount to configure.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * The unique name of the database connection to use for the static role.
     * 
     */
    @Import(name="dbName")
    private @Nullable Output<String> dbName;

    /**
     * @return The unique name of the database connection to use for the static role.
     * 
     */
    public Optional<Output<String>> dbName() {
        return Optional.ofNullable(this.dbName);
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
     */
    @Import(name="rotationSchedule")
    private @Nullable Output<String> rotationSchedule;

    /**
     * @return A cron-style string that will define the schedule on which rotations should occur.
     * Mutually exclusive with `rotation_period`.
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
     * The database username that this static role corresponds to.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The database username that this static role corresponds to.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private SecretBackendStaticRoleState() {}

    private SecretBackendStaticRoleState(SecretBackendStaticRoleState $) {
        this.backend = $.backend;
        this.dbName = $.dbName;
        this.name = $.name;
        this.namespace = $.namespace;
        this.rotationPeriod = $.rotationPeriod;
        this.rotationSchedule = $.rotationSchedule;
        this.rotationStatements = $.rotationStatements;
        this.rotationWindow = $.rotationWindow;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendStaticRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendStaticRoleState $;

        public Builder() {
            $ = new SecretBackendStaticRoleState();
        }

        public Builder(SecretBackendStaticRoleState defaults) {
            $ = new SecretBackendStaticRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The unique name of the Vault mount to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
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
         * @param dbName The unique name of the database connection to use for the static role.
         * 
         * @return builder
         * 
         */
        public Builder dbName(@Nullable Output<String> dbName) {
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
         * @param username The database username that this static role corresponds to.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
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

        public SecretBackendStaticRoleState build() {
            return $;
        }
    }

}
