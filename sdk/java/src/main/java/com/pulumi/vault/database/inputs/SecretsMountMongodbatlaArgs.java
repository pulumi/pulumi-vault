// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

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


public final class SecretsMountMongodbatlaArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretsMountMongodbatlaArgs Empty = new SecretsMountMongodbatlaArgs();

    /**
     * A list of roles that are allowed to use this
     * connection.
     * 
     */
    @Import(name="allowedRoles")
    private @Nullable Output<List<String>> allowedRoles;

    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public Optional<Output<List<String>>> allowedRoles() {
        return Optional.ofNullable(this.allowedRoles);
    }

    /**
     * A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    @Import(name="data")
    private @Nullable Output<Map<String,String>> data;

    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    public Optional<Output<Map<String,String>>> data() {
        return Optional.ofNullable(this.data);
    }

    /**
     * Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    @Import(name="disableAutomatedRotation")
    private @Nullable Output<Boolean> disableAutomatedRotation;

    /**
     * @return Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    public Optional<Output<Boolean>> disableAutomatedRotation() {
        return Optional.ofNullable(this.disableAutomatedRotation);
    }

    /**
     * Name of the database connection.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the database connection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Specifies the name of the plugin to use.
     * 
     */
    @Import(name="pluginName")
    private @Nullable Output<String> pluginName;

    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Optional<Output<String>> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }

    /**
     * The Private Programmatic API Key used to connect with MongoDB Atlas API.
     * 
     */
    @Import(name="privateKey", required=true)
    private Output<String> privateKey;

    /**
     * @return The Private Programmatic API Key used to connect with MongoDB Atlas API.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }

    /**
     * The Project ID the Database User should be created within.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The Project ID the Database User should be created within.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    @Import(name="publicKey", required=true)
    private Output<String> publicKey;

    /**
     * @return The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    /**
     * A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    @Import(name="rootRotationStatements")
    private @Nullable Output<List<String>> rootRotationStatements;

    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    public Optional<Output<List<String>>> rootRotationStatements() {
        return Optional.ofNullable(this.rootRotationStatements);
    }

    /**
     * The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationPeriod")
    private @Nullable Output<Integer> rotationPeriod;

    /**
     * @return The amount of time in seconds Vault should wait before rotating the root credential.
     * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<Integer>> rotationPeriod() {
        return Optional.ofNullable(this.rotationPeriod);
    }

    /**
     * The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationSchedule")
    private @Nullable Output<String> rotationSchedule;

    /**
     * @return The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
     * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<String>> rotationSchedule() {
        return Optional.ofNullable(this.rotationSchedule);
    }

    /**
     * The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    @Import(name="rotationWindow")
    private @Nullable Output<Integer> rotationWindow;

    /**
     * @return The maximum amount of time in seconds allowed to complete
     * a rotation when a scheduled token rotation occurs. The default rotation window is
     * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
     * 
     */
    public Optional<Output<Integer>> rotationWindow() {
        return Optional.ofNullable(this.rotationWindow);
    }

    /**
     * Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    @Import(name="verifyConnection")
    private @Nullable Output<Boolean> verifyConnection;

    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    public Optional<Output<Boolean>> verifyConnection() {
        return Optional.ofNullable(this.verifyConnection);
    }

    private SecretsMountMongodbatlaArgs() {}

    private SecretsMountMongodbatlaArgs(SecretsMountMongodbatlaArgs $) {
        this.allowedRoles = $.allowedRoles;
        this.data = $.data;
        this.disableAutomatedRotation = $.disableAutomatedRotation;
        this.name = $.name;
        this.pluginName = $.pluginName;
        this.privateKey = $.privateKey;
        this.projectId = $.projectId;
        this.publicKey = $.publicKey;
        this.rootRotationStatements = $.rootRotationStatements;
        this.rotationPeriod = $.rotationPeriod;
        this.rotationSchedule = $.rotationSchedule;
        this.rotationWindow = $.rotationWindow;
        this.verifyConnection = $.verifyConnection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretsMountMongodbatlaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretsMountMongodbatlaArgs $;

        public Builder() {
            $ = new SecretsMountMongodbatlaArgs();
        }

        public Builder(SecretsMountMongodbatlaArgs defaults) {
            $ = new SecretsMountMongodbatlaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(@Nullable Output<List<String>> allowedRoles) {
            $.allowedRoles = allowedRoles;
            return this;
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(List<String> allowedRoles) {
            return allowedRoles(Output.of(allowedRoles));
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(String... allowedRoles) {
            return allowedRoles(List.of(allowedRoles));
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * @return builder
         * 
         */
        public Builder data(@Nullable Output<Map<String,String>> data) {
            $.data = data;
            return this;
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * @return builder
         * 
         */
        public Builder data(Map<String,String> data) {
            return data(Output.of(data));
        }

        /**
         * @param disableAutomatedRotation Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder disableAutomatedRotation(@Nullable Output<Boolean> disableAutomatedRotation) {
            $.disableAutomatedRotation = disableAutomatedRotation;
            return this;
        }

        /**
         * @param disableAutomatedRotation Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder disableAutomatedRotation(Boolean disableAutomatedRotation) {
            return disableAutomatedRotation(Output.of(disableAutomatedRotation));
        }

        /**
         * @param name Name of the database connection.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the database connection.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param pluginName Specifies the name of the plugin to use.
         * 
         * @return builder
         * 
         */
        public Builder pluginName(@Nullable Output<String> pluginName) {
            $.pluginName = pluginName;
            return this;
        }

        /**
         * @param pluginName Specifies the name of the plugin to use.
         * 
         * @return builder
         * 
         */
        public Builder pluginName(String pluginName) {
            return pluginName(Output.of(pluginName));
        }

        /**
         * @param privateKey The Private Programmatic API Key used to connect with MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey The Private Programmatic API Key used to connect with MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param projectId The Project ID the Database User should be created within.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The Project ID the Database User should be created within.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param publicKey The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(@Nullable Output<List<String>> rootRotationStatements) {
            $.rootRotationStatements = rootRotationStatements;
            return this;
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(List<String> rootRotationStatements) {
            return rootRotationStatements(Output.of(rootRotationStatements));
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(String... rootRotationStatements) {
            return rootRotationStatements(List.of(rootRotationStatements));
        }

        /**
         * @param rotationPeriod The amount of time in seconds Vault should wait before rotating the root credential.
         * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(@Nullable Output<Integer> rotationPeriod) {
            $.rotationPeriod = rotationPeriod;
            return this;
        }

        /**
         * @param rotationPeriod The amount of time in seconds Vault should wait before rotating the root credential.
         * A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationPeriod(Integer rotationPeriod) {
            return rotationPeriod(Output.of(rotationPeriod));
        }

        /**
         * @param rotationSchedule The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
         * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(@Nullable Output<String> rotationSchedule) {
            $.rotationSchedule = rotationSchedule;
            return this;
        }

        /**
         * @param rotationSchedule The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
         * defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationSchedule(String rotationSchedule) {
            return rotationSchedule(Output.of(rotationSchedule));
        }

        /**
         * @param rotationWindow The maximum amount of time in seconds allowed to complete
         * a rotation when a scheduled token rotation occurs. The default rotation window is
         * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(@Nullable Output<Integer> rotationWindow) {
            $.rotationWindow = rotationWindow;
            return this;
        }

        /**
         * @param rotationWindow The maximum amount of time in seconds allowed to complete
         * a rotation when a scheduled token rotation occurs. The default rotation window is
         * unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
         * 
         * @return builder
         * 
         */
        public Builder rotationWindow(Integer rotationWindow) {
            return rotationWindow(Output.of(rotationWindow));
        }

        /**
         * @param verifyConnection Whether the connection should be verified on
         * initial configuration or not.
         * 
         * @return builder
         * 
         */
        public Builder verifyConnection(@Nullable Output<Boolean> verifyConnection) {
            $.verifyConnection = verifyConnection;
            return this;
        }

        /**
         * @param verifyConnection Whether the connection should be verified on
         * initial configuration or not.
         * 
         * @return builder
         * 
         */
        public Builder verifyConnection(Boolean verifyConnection) {
            return verifyConnection(Output.of(verifyConnection));
        }

        public SecretsMountMongodbatlaArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("SecretsMountMongodbatlaArgs", "name");
            }
            if ($.privateKey == null) {
                throw new MissingRequiredPropertyException("SecretsMountMongodbatlaArgs", "privateKey");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("SecretsMountMongodbatlaArgs", "projectId");
            }
            if ($.publicKey == null) {
                throw new MissingRequiredPropertyException("SecretsMountMongodbatlaArgs", "publicKey");
            }
            return $;
        }
    }

}
